package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/shiningflint/go-gym/websocket/models"
)

var addr = flag.String("addr", ":8888", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	chat, err := models.GetChat(1)
	if err != nil {
		log.Fatal(err)
	}
	messages, err := models.AllChatMessages(chat.Id)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range messages {
		fmt.Println(v.Message)
	}

	http.ServeFile(w, r, "index.html")
	return
}

func main() {
	models.DbConnect()
	fmt.Println("Server starting on port :8888")
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
