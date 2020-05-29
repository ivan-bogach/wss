package main

import (
	"flag"
	"log"
	"net/http"

	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
}

func write(msg chan *Message) {
	for {
		time.Sleep(5 * time.Second)
		msg <- &Message{Message: "AAAAAAAAAAAAAAAAAA"}
	}
}

func main() {
	msg := make(chan *Message)

	hub := newHub()
	go hub.run(msg)

	go func() {
		d := &Device{hub: hub, send: make(chan []byte), localAddr: "loc"}
		d.hub.reg <- d
		go write(msg)
	}()

	flag.Parse()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveW(w, r, msg)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
