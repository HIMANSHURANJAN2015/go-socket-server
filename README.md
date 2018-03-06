# go-socket-server
Implements a socket server in GOLANG using gorilla library of websocket. This socket listens to events published by Gcloud pub-sub
1) Client :- 1 socket connection
2) Hub :- Collection of many clients which have same property eg:- (All clients which wants to listen to events from store-1 etc)
	Hub is identified by name like store-1, store-24 etc
3) HubMap :- It is basically list of all hubs. There will be 1 hubMap for an application. This hubMap listens to all events from gcloud pubsub.
	Whenever an event comes, this hub send the broadcast message to respective hub and that hub in turn sends it to all of its client.

In frontend, one can connect to sever through javascript
	var conn = new WebSocket("ws://localhost:8083/ws/store?storeId=1");

As of now, this listens to only "_order-status_" event of cloud pubsub. but it can be easily extended for all other events as well.



Setup
-----------

1) clone the repository
2) install go and set GOPATH to point to the cloned repository (go-socket-server)
	* vim ~/.bashrc . Go to the end of the file and add/update the following line
	* export GOPATH=/home/<name>/Desktop/go-socket-server     (go-socket-server is our repo.)
	* source ~/bashrc

Note that GOPATH point to a folder which have the structure like it contains 3 folders. - bin, pkg and src

3) Install dependent packages.
	* cd $GOPATH
	* go get -u cloud.google.com/go/pubsub
	* go get -u golang.org/x/net/context

This will install the packages in pkg directory of our $GOPATH folder	
	

Compiling and Running
----------------------

1) Goto $GOPATH/src/go-socket-server and run
	* go build. It will create a binary file called go-socket-server
	* go install. It will install this executable file in $GOPATH/bin
2) to run ./go-socket-server


Running test cases
-----------------------
1) Goto $GOPATH/src/go-socket-serverand run
	* go build
	* go test -v
