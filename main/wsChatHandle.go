package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"linuxPhoneTransition/models"
	"log"
	"net/http"
	"time"
)

type Client struct {
	ipAddress string
	Online    bool
	conn      *websocket.Conn
}

func (c *Client) ReadAndStore() {

	for {
		var CurMsg = models.Chat{}
		err := c.conn.ReadJSON(&CurMsg)
		//_, msg, err := c.conn.ReadMessage()
		if err != nil {
			manager.Checkout <- c
			return
		}

		CurMsg.CreateTime = time.Now().Unix()
		db.Create(&CurMsg)
		manager.BoardCast <- &CurMsg
	}

}

type ClientManager struct {
	LoginChan chan *Client
	Checkout  chan *Client
	BoardCast chan *models.Chat
	Clients   map[*Client]bool
}

func (m *ClientManager) Start() {
	var client *Client
	var message *models.Chat
	for {
		select {
		case client = <-m.LoginChan:
			fmt.Println("come a client:", client.ipAddress)
			m.Clients[client] = true
			go client.ReadAndStore()
		case client = <-m.Checkout:
			fmt.Println("leave a client:", client.ipAddress)
			delete(m.Clients, client)
		case message = <-m.BoardCast:
			fmt.Println("boardCast a msg:",message)
			m.BoardCastFunc(message)
		}
	}

}

func (m *ClientManager) BoardCastFunc(message *models.Chat) {
	for client, _ := range m.Clients {
		client.conn.WriteJSON(message)

	}
}

func chatWsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Panicf("upgrade ws failed with err :%v\n", err)
		return
	}
	fmt.Printf("client ip is:%v\t %v\t\n", conn.RemoteAddr(), r.RemoteAddr)
	newClient := &Client{
		ipAddress: conn.RemoteAddr().String(),
		Online:    true,
		conn:      conn,
	}
	manager.LoginChan <- newClient

}
