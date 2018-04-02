# go-socket-server
Implements a socket server in GOLANG using gorilla library of websocket. This socket listens to events published by Gcloud pub-sub
1) Client :- Represents 1 socket connection. If the client goes offline/closes connection (eg:- closing tab/browser) then the connection will be closed after 54sec, through PING mechanism and subsequently removed from the hub. _Each tab is considered a different client_.
2) Hub :- Collection of many clients which have same property eg:- (All clients which wants to listen to events from store-1 etc)
	Hub is identified by name like store-1, store-24 etc
3) HubMap :- It is basically list of all hubs. There will be 1 hubMap for an application. This hubMap listens to all events from gcloud pubsub.
	Whenever an event comes, this hub send the broadcast message to respective hub and that hub in turn sends it to all of its client.

In frontend, one can connect to sever through javascript
	var conn = new WebSocket("ws://localhost:8083/ws/store?storeId=1");

As of now, this listens to only "_order-status_" event of cloud pubsub. but it can be easily extended for all other events as well.



Setup
-----------

1) Install Go and setup GOPATH environment variables. Also append $GOPATH/bin to $PATH variable 

   eg:- 
   * a sample gopath is /home/zopnow/Desktop/go-workspace
   * go-workspace folder has following directories - bin, pkg, projects and src. bin,pkg and src follows sematics of workspace as defined by GO.
   * __Projects folder has all our different repositories of go projects__ . 
2) Clone the repository inside /home/zopnow/Desktop/go-workspace/projects. (You may choose any other directory as well) The important point to remember is that go will fetch executables and libraries	from go path.


3) Install dependent packages.
	* cd $GOPATH
	* go get -u cloud.google.com/go/pubsub
	* go get -u golang.org/x/net/context
	* go get -u github.com/gorilla/websocket
	     
This will install the packages in pkg directory of our $GOPATH folder	
	

Compiling and Running
----------------------

1) Goto $GOPATH/src/go-socket-server and run
	* go build. It will create a binary file called go-socket-server
	* go install. It will install this executable file in $GOPATH/bin
2) to run ./go-socket-server

Setting Up GCLOUD EMULATOR
---------------------------
1) Setup emulator by following [https://cloud.google.com/pubsub/docs/emulator]
2) For testing, point your PUBSUB_EMULATOR_HOST locally. See subscriber.go file

Running test cases
-----------------------
1) Goto $GOPATH/src/go-socket-serverand run
	* go build
	* go test -v
