package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/shiningflint/go-gym/websocket/models"
)

var addr = flag.String("addr", ":8888", "http service address")

var templates = template.Must(template.ParseFiles("index.html"))

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

	err = templates.ExecuteTemplate(w, "index.html", messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func serveChatJs(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	http.ServeFile(w, r, "assets/javascripts/chat.js")
}

func serveChatCss(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	http.ServeFile(w, r, "assets/stylesheets/chat.css")
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
	http.HandleFunc("/assets/javascripts/chat.js", serveChatJs)
	http.HandleFunc("/assets/stylesheets/chat.css", serveChatCss)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
