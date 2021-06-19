package main

import (
	"awesomeProject/models"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
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
	initSome()
	http.Handle("/",http.FileServer(http.Dir("dist/")))
	http.HandleFunc("/upload", uploadFileHandler)
	http.HandleFunc("/download/", downloadFileHandler)
	http.HandleFunc("/getFiles", getFileHandler)
	http.HandleFunc("/chatSocket", chatWsHandler)
	http.HandleFunc("/getMsgs", getMsgsHandler)
	//http.HandleFunc("/sendMsg", sendMsgHandler)
	log.Print("Server started on localhost:9000...")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Printf("listen at 9000 failed,err is: %v\n", err)
	}
}







func initSome() {
	go manager.Start()
	ConnectMysql()
}




