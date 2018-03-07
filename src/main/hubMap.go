package main 

import (
	"log"
	"../subscriber"
	"sync"
	"encoding/json"
	"strconv"
)

type HubMap struct {
	// Store list of all hubs. It will be a key value pair eg:- store-1, store-2 
	hubs map[string]*Hub
	sync.RWMutex
}

var channel = make(chan string)

// Get by hubName
func (hm *HubMap) Get(hubName string) *Hub {
    hm.RLock()
    defer hm.RUnlock()
    // Does this hub exist yet
    hub, ok := hm.hubs[hubName]

    if !ok {
        // Create a new hub
        hub = newHub()
        hm.hubs[hubName] = hub
    }
    return hub
}

func (hm *HubMap) addSubscriber(topicName, subscriptionName, projectId string) {
	var sub subscriber.Subscriber
	sub.SubscriptionName = subscriptionName
	sub.TopicName = topicName
	sub.ProjectId = projectId
	go sub.PullMessages(channel)
	log.Println("Subscriber Added successfully for the topic - ", topicName)
}

func (hm *HubMap) subscribe() {
	log.Println("HubList is waiting for events")
	for {
		var parsedMessage map[string]interface{}
		message := <-channel
		log.Println("reached here")
		err := json.Unmarshal([]byte(message), &parsedMessage) 
		log.Println(err)
		log.Println("yahan")
		if err!= nil {
			log.Println(err)
			continue
		}
		storeId := int(0)
		log.Println(parsedMessage, "parsedMessage")		
		for key, value := range parsedMessage {
		 	if (key == "storeId") {
		 		storeId = int(value.(float64))
		 	}
		}
		log.Println("storeId", storeId)
		// log.Println("type of storeId"+reflect.TypeOf(storeId))
		log.Println("Store ID from pubsub : " + strconv.Itoa(storeId))
		if storeId != 0 {
			hubName := "store-" + strconv.Itoa(storeId)
			log.Println("will send data to this hub"+hubName)
			hub := hm.hubs[hubName]
			log.Println(hm.hubs)
			log.Println(hub)
			log.Println("done")
			if hub != nil {
				log.Println("Send broadcast message to hub :"+ hubName)
				hub.broadcast <- message
			}			
		}		
	}
}
