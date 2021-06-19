package models

import "github.com/gorilla/websocket"

type Client struct {
	ipAddress string
	Online    bool
	conn      *websocket.Conn
}

