package main

import (
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
			log.Println("come a client:", client.ipAddress)
			m.Clients[client] = true
			go client.ReadAndStore()
		case client = <-m.Checkout:
			log.Println("leave a client:", client.ipAddress)
			delete(m.Clients, client)
		case message = <-m.BoardCast:
			log.Println("boardCast a msg:", message)
			m.BoardCastFunc(message)
		}
	}

}

func (m *ClientManager) BoardCastFunc(message *models.Chat) {
	for client, _ := range m.Clients {
		client.conn.WriteJSON(message)

	}
}

func consoleWsHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//conn, err := upgrade.Upgrade(w, r, nil)
	//if err != nil {
	//	log.Panicf("upgrade ws failed with err :%v\n", err)
	//	return
	//}
	//log.Printf("client ip is:%v\t %v\t\n", conn.RemoteAddr(), r.RemoteAddr)
	////conn.WriteMessage(1, utils.GetShellPrefix())
	////conn.WriteMessage(1,append([]byte("\n"),utils.GetShellPrefix()...))
	//session := exec.Command("/bin/bash")
	//f,_ := pty.Start(session)
	//go func() {
	//	command := bytes.Buffer{}
	//	for {
	//		buffer := make([]byte,1)
	//		_, wsReader, _ := conn.NextReader()
	//		length, _ := wsReader.Read(buffer)
	//		//fmt.Printf("%#v",string(buffer[:length]))
	//		if buffer[length-1] == '\r' {
	//			break
	//		}
	//		command.Write(buffer)
	//
	//	}
	//	execCommand := string(command.Bytes())
	//	fmt.Printf("command is:%#v\n",execCommand)
	//	//res, _ := exec.Command("/bin/bash", "-c",execCommand).Output()
	//	//res = bytes.ReplaceAll(res,[]byte("\n"),[]byte("\t"))
	//	f.Write(command.Bytes())
	//	_,webReader,_ := conn.NextReader()
	//	io.Copy(webReader,f)
	//	log.Printf("std out is : %#v\n", string(res))
	//	conn.WriteMessage(1, append([]byte("\r\n"),res...))
	//}()

}

func chatWsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Panicf("upgrade ws failed with err :%v\n", err)
		return
	}
	log.Printf("client ip is:%v\t %v\t\n", conn.RemoteAddr(), r.RemoteAddr)
	newClient := &Client{
		ipAddress: conn.RemoteAddr().String(),
		Online:    true,
		conn:      conn,
	}
	manager.LoginChan <- newClient

}
