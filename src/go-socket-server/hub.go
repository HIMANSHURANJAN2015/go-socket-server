package main

import "log"

type Hub struct {
	// Registered clients
	clients map[*Client]bool
	// Messages to be broadcasted in this hub
	broadcast chan string
	//Register requests from clients
	register chan *Client
	//Unregister requests from clients
	unregister chan *Client 

}

func newHub() *Hub {
	hub := &Hub{
		clients: make(map[*Client]bool),
		broadcast: make(chan string),
		register: make(chan *Client),
		unregister: make(chan *Client),
	}
	go hub.run()
	log.Println("Creating new Hub")
	return hub
}

func (h *Hub) run() {
	log.Println("Hub is waiting for messages")
	for {
		select {
		case client := <-h.register:
			log.Println("Client registered with hub")
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <- h.broadcast:
			log.Println("Sending message to all of its clients")
			for client := range h.clients {
				select {
					case client.send <- message:
					default:
						delete(h.clients, client) 
						close(client.send)
				}
			}
		}
	}
}
