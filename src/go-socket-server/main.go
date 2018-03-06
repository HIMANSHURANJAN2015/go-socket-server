package main

import (
	"log"
	"net/http"
)


var hubMap *HubMap

func init () {
	log.Println("init called")
	hubMap = &HubMap{hubs : make(map[string]*Hub)}
	hubMap.addSubscriber("order-status", "order-notifications-desktop", "zopsmart-176211")		
	go hubMap.subscribe()
}

func main() {
	log.Println("Golang websocket Server Running.....")	
	// This will have all the routes
	http.HandleFunc("/ws/store", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}