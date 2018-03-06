# go-socket-server
Implements a socket server in GOLANG using gorilla library of websocket. This socket listens to events published by Gcloud pub-sub
Client :- 1 socket connection
Hub :- Collection of many clients which have same property eg:- (All clients which wants to listen to events from store-1 etc)
	Hub is identified by name like store-1, store-24 etc
HubMap :- It is basically list of all hubs. There will be 1 hubMap for an application. This hubMap listens to all events from gcloud pubsub.
	Whenever an event comes, this hub send the broadcast message to respective hub and that hub in turn sends it to all of its client.

In frontend, one can connect to sever through javascript
	var conn = new WebSocket("ws://localhost:8083/ws/store?storeId=1");

As of now, this istens to only order-status event of cloud pubsub. but it can be easily extended for all other events as well.



Setup
-----------
1) install go and set GOPATH to a place where u have placed ur repo.. Note that GOPATH point to a folder which have the structure like it contains 3 folders. - bin, pkg and src

2) go to $GOPATH/src and clone the repo

3) Installl deependent packages
	go get -u cloud.google.com/go/pubsub
	go get -u golang.org/x/net/context

Compiling and Running
----------------------

1) Goto $GOPATH/src/go-socket-server 
	and run go build. It will create a binary file called go-socket-server
		go install. It will install this executable file in $GOPATH/bin
2) to run ./go-socket-server


Running test cases
-----------------------
1) Goto $GOPATH/src/go-socket-serve
	and run go build
		go test -v
