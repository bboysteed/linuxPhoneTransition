package main

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"linuxPhoneTransition/models"
	"log"
	"net/http"
)

var (
	storePath = "./files"
	db *gorm.DB
	upgrade = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	manager = &ClientManager{
		LoginChan: make(chan *Client),
		Checkout:  make(chan *Client),
		BoardCast: make(chan *models.Chat),
		Clients:   make(map[*Client]bool),
	}
)



func main() {
	//initSome()
	http.Handle("/",http.FileServer(http.Dir("dist/")))
	http.HandleFunc("/upload", uploadFileHandler)
	http.HandleFunc("/download/", downloadFileHandler)
	http.HandleFunc("/getFiles", getFileHandler)
	http.HandleFunc("/chatSocket", chatWsHandler)
	http.HandleFunc("/console", consoleWsHandler)
	http.HandleFunc("/getMsgs", getMsgsHandler)
	//http.HandleFunc("/sendMsg", sendMsgHandler)
	log.Print("Server started on localhost:10000...")
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		log.Printf("listen at 10000 failed,err is: %v\n", err)
	}
}









func initSome() {
	go manager.Start()
	ConnectMysql()
}




