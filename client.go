package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct{
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
	CheckOrigin: func(r * http.Request) bool {return true},
}

func ServeWs()