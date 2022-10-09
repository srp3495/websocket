package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func setupConn() {
	http.HandleFunc("/", httpFunc)
	http.HandleFunc("/ws", wsFunc)
}

func httpFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Home page")
}
func reader(wc *websocket.Conn) {
	for {
		messageType, p, err := wc.ReadMessage()
		if err != nil {
			log.Println("Error while reading message", err)
			return
		}
		log.Println(string(p))
		if err := wc.WriteMessage(messageType, p); err != nil {
			log.Println("Error while writing ", err)
			return
		}

	}
}

func wsFunc(w http.ResponseWriter, req *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Error in connection", err)
	}
	log.Println("Connection established")
	reader(ws)

}

func main() {
	fmt.Println("Go Websockets")
	setupConn()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
