package main

import (
	// "fmt"
	"log"
	"net/http"
	// "os/exec"
)

// Websocket message for gui
type Message struct {
	// Winner device number
	Message string `json:"Message"`
	// Device local addresses list
	DevicesLocAddr []string `json:"devices"`
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func serveW(w http.ResponseWriter, r *http.Request, msg chan *Message) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Client servW: ", err)
		return
	}

	for {
		select {
		case message := <-msg:
			if err = conn.WriteJSON(message); err != nil {
				log.Println("Client serveW <- msg: ", err)
				return
			}

		}
	}
}
