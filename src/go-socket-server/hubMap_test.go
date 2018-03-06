package main 

import(
	"testing"
	"reflect"
)	


func TestHubMapGet(t *testing.T) {
	hubMap := &HubMap{hubs : make(map[string]*Hub)}
	// Calling hubmap Get with store-1
	hub1 := hubMap.Get("store-1")
	t.Log(hub1)
	// Calling hubmap Get with store-2
	hub2 := hubMap.Get("store-2")
	t.Log(hub2)
	// Calling hubmap Get again with store-1
	hub3 := hubMap.Get("store-1")
	t.Log(hub3)

	if !reflect.DeepEqual(hub1, hub3) {
		t.Error("Same hubName gave different hub")
	}

	if reflect.DeepEqual(hub1, hub2) {
		t.Error("Different hubName gave same hub")
	}
	// Calling hubmap Get without name -> Build itself wont work

}

func TestHubSubscribe(t *testing.T) {
	hubMap := &HubMap{hubs : make(map[string]*Hub)}
	go hubMap.subscribe()
	hub1 := hubMap.Get("store-1")
	t.Log(hub1)
	hub2 := hubMap.Get("store-2")
	t.Log(hub2)
	hub3 := hubMap.Get("store-1")
	t.Log(hub3)
	// Passing value to hubmap channel
	t.Log(reflect.TypeOf(channel))
	channel <- "{\"referenceNumber\":12972,\"oldStatus\":null,\"newStatus\":\"PENDING\",\"organizationId\":\"1\",\"storeId\":1}"
	t.Log(hub1.broadcast)
	t.Log(hub2.broadcast)
	t.Log(hub3.broadcast)
}