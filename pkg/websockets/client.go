package websokets

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID         string
	Connection *websocket.Conn
	Pool       *Pool
	mu         sync.Mutex
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()

	for {
		messageType, p, err := c.Connection.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		message := Message{
			Type: messageType,
			Body: string(p),
		}

		c.Pool.Brodcast <- message
		fmt.Printf("message recived: %+V\n", message)

	}
}
