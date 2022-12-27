package websokets

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID      string
	Connect *websocket.Conn
	Pool    *Pool
	mu      sync.Mutex
}
