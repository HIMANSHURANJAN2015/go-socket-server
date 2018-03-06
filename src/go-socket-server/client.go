package main 

import (
	"log"
	"net/http"
	"time"
	"strings"
	"github.com/gorilla/websocket"
)

const (
	// Maximum message size allowed from peer.
	maxMessageSize = 512
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize : 1024,
	WriteBufferSize : 1024,
	CheckOrigin: func(r *http.Request) bool {
			return true //By default gorilla/websocket returns false if the request is cross-origin.
	},
}

type Client struct {
	hub *Hub
	// websocket connection
	conn *websocket.Conn
	// Buffered Channel of outboud messages (Hub writes data to this channel during broadcasting)
	send chan string
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <- c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			str := []byte(message)
			//Compare this method with other methof to wrte 
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(str)
			if err := w.Close(); err != nil {
				return
			}
		case <- ticker.C:
			// Checking if connection is alive
			log.Println("closing connection of client1")
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("closing connection of client")
				return
			}	
		}
	}
}


func serveWs(w http.ResponseWriter, r *http.Request) {	
	log.Println("Reached serveWs function")
	requestUrl := r.URL.RequestURI()
	parts := strings.Split(requestUrl, "/")
	socketType := strings.Split(parts[2], "?")[0] // "/ws/store?storeId=21"
	hubName := ""
	if (socketType == "store") {
		mandatoryParam := "storeId" // Can become an array
		keys, ok := r.URL.Query()[mandatoryParam]
    	if !ok || len(keys) < 1 {
    			errorMessage := "Url Param 'key' is missing :" +mandatoryParam
        		log.Println(errorMessage)
        		return
    	}
		hubName = "store-"+keys[0]	
	}
	log.Println("Hub Name :", hubName)
	hub := hubMap.Get(hubName)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		hub: hub, 
		conn: conn, 
		send: make(chan string),
	}
	client.hub.register <- client
	log.Println("Client added to hub-"+hubName)
	go client.writePump()

}







