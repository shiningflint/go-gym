package main

import (
	"testing"
)

func TestHubRunRegister(t *testing.T) {
	hub := newHub()
	go hub.run()
	client := &Client{hub: hub, send: make(chan []byte, 256)}
	client.hub.register <- client
	if hub.clients[client] != true {
		t.Errorf("Expecting client to be true, got %v", hub.clients[client])
	}
}
