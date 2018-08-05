package main

import (
	"github.com/hpcloud/tail"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
	"net/http"
)

type LogMessage struct {
	logtype string 
	name    string 
	message string 
}

type LogEventData struct {
	Data string
}

func writeLog(server *gosocketio.Server) {
	t, err := tail.TailFile("./nohup.out", tail.Config{Follow: true})
	if (err != nil) {
		panic(err)
	}
	for line := range t.Lines {
		server.BroadcastToAll("logevent", LogEventData{line.Text})
	}
}

func main() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New Client Connected")
		c.Join("webtail")
	})

	go writeLog(server)
	serveMux := http.NewServeMux()
	serveMux.Handle("/", server)
	log.Panic(http.ListenAndServe("0.0.0.0:8080", serveMux))
}
