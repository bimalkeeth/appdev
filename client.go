package main

import "github.com/gorilla/websocket"


type clinet struct {
	socket *websocket.Conn
	send chan []byte
	room *room
}


func main() {


}
