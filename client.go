package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	User string
	Text string
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
	user string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	user := GenerateUsername()
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), user: user}
	hub.register <- client

	go client.writePump()
	client.readPump()
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.hub.broadcast <- Message{User: c.user, Text: string(msg)}
	}
}

func (c *Client) writePump() {
	for msg := range c.send {
		c.conn.WriteMessage(websocket.TextMessage, msg)
	}
}
