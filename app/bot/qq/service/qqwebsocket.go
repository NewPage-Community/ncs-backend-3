package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/miRemid/amy/websocket/model"
	"log"
	"net/http"
)

type baseClient struct {
	URL   string
	Port  int
	conn  *websocket.Conn
	token string
}

// Connect to the cqhttp
func (c *baseClient) Connect(url string, header http.Header) {
	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Fatal(err)
	}
	c.conn = conn
}

// MsgClient is the struct of cq websocket client
type MsgClient struct {
	baseClient
	message WebEventHandler
}

// WebEventHandler is the handler function of receive message
type WebEventHandler func(event model.CQEvent)

// NewClient return a new websocket server ptr
func NewWSClient(url string, port int) *MsgClient {
	var client MsgClient
	client.URL = url
	client.Port = port
	client.token = ""
	return &client
}

// Run will open a websocket client
func (c *MsgClient) Run() {
	c.Connect()
	for {
		_, body, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("read error:%v", err)
			continue
		}
		evt := model.NewCQEvent(body)
		go c.message(evt)
	}
}

// Connect to the ws server
func (c *MsgClient) Connect() {
	url := fmt.Sprintf("ws://%s:%d/?access_token=%s", c.URL, c.Port, c.token)
	log.Printf("Event Connect:%s", url)
	c.baseClient.Connect(url, nil)
}

// OnMessage will set the message parse function
func (c *MsgClient) OnMessage(handler WebEventHandler) {
	c.message = handler
}

// SetToken will set the api client's token
func (c *baseClient) SetToken(token string) {
	c.token = token
}
