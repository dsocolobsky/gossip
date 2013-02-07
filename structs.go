package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

type User struct {
	Nick     string
	Username string
}

type Server struct {
	Address string
	Port    string
	Channel string
}

type Connection struct {
	Srv    Server
	Usr    User
	Socket net.Conn
}

type UserDefined struct {
	Server  string
	Port    string
	Nick    string
	User    string
	Channel string
}

func decodeJSON() *UserDefined {
	var uds UserDefined

	file, err := ioutil.ReadFile("./pr1.json")
	if err != nil {
		fmt.Println("error loading json file")
	}

	json.Unmarshal(file, &uds)
	return &uds
}

func CreateConnection() *Connection {
	var conn Connection
	uds := decodeJSON()

	conn.Srv.Address = uds.Server
	conn.Srv.Port = uds.Port
	conn.Usr.Nick = uds.Nick
	conn.Usr.Username = uds.User
	conn.Srv.Channel = uds.Channel

	conn.Socket = nil

	return &conn
}

func (conn *Connection) Connect() (sock net.Conn, err error) {
	tcpcon, err := net.Dial("tcp", conn.Srv.Address+":"+conn.Srv.Port)
	if err != nil {
		log.Fatal("unable to Connect to IRC Server ", err)
	}

	conn.Socket = tcpcon
	log.Printf("Connected to IRC Server %s (%s)\n", conn.Srv.Address, conn.Socket.RemoteAddr())
	return conn.Socket, nil
}

func (conn *Connection) Start(sock net.Conn) {
	sock.Write([]byte("USER " + conn.Usr.Username + " 8 * :" + conn.Usr.Username + "\n"))
	sock.Write([]byte("NICK " + conn.Usr.Nick + "\n"))
	sock.Write([]byte("JOIN " + conn.Srv.Channel + "\n"))
}

func (conn *Connection) SendMessage(sock net.Conn, text string) {
	sock.Write([]byte("PRIVMSG " + conn.Srv.Channel + " :" + text + "\n"))
}
