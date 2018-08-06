package main

import (
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/hpcloud/tail"
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

func writeLog(filepath string, server *gosocketio.Server) {
	t, err := tail.TailFile(filepath, tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		server.BroadcastToAll("logevent", LogEventData{line.Text})
	}
}

func main() {
	hp := loadConfig()
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New Client Connected")
		c.Join("webtail")
	})

	for _, p := range hp.Program {
		go writeLog(p.Stdout, server)
		go writeLog(p.Stderr, server)
	}

	serveMux := http.NewServeMux()
	serveMux.Handle("/", server)
	log.Panic(http.ListenAndServe("0.0.0.0:8080", serveMux))
}
