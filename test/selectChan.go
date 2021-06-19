package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var signalChan chan string
var loginChan chan string
var registerChan chan string
var logoutChan chan string

func main() {

	loginChan = make(chan string, 1)
	registerChan = make(chan string, 1)
	logoutChan = make(chan string, 1)
	signalChan = make(chan string, 1)

	go receiveKey()
	go loopAndServer()
	wg.Add(2)
	wg.Wait()
}
func loopAndServer() {
	var start string
	for {

		select {
		case start = <-signalChan:
			fmt.Println("start...", start)
		case <-registerChan:
			fmt.Printf("someone register at:%v\n", time.Now().Format("15:04:05"))
		case <-logoutChan:
			fmt.Printf("someone logout at:%v\n", time.Now().Format("15:04:05"))
		case <-loginChan:
			fmt.Printf("someone login at:%v\n", time.Now().Format("15:04:05"))

		}
	}
}
func receiveKey() {

	var key string
	for {
		fmt.Scanln(&key)
		fmt.Printf("input key is :%v\n", key)
		switch key {
		case "1":
			signalChan <- "start"
		case "2":
			registerChan <- "start"
		case "3":
			loginChan <- "start"
		case "4":
			logoutChan <- "start"
		default:
			fmt.Println("wrong type!")

		}
		key=""

	}

}
