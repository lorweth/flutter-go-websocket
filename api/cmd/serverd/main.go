package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			if err.Error() == "EOF" {
				log.Println("Disconnect")
				break
			}

			log.Printf("Can't receive: %v \n", err.Error())
			break
		}

		log.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		log.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			log.Printf("Can't send: %v \n", err.Error())
			break
		}
	}
}

func main() {
	http.HandleFunc("/chat", func(w http.ResponseWriter, req *http.Request) {
		s := websocket.Server{Handler: websocket.Handler(Echo)}
		s.ServeHTTP(w, req)
	})

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
