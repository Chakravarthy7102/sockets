package websokets

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Brodcast   chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Brodcast:   make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Print("Size of the connection pool:", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Connection.WriteJSON(Message{Type: 1, Body: "New user joined.."})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of the connection pool", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Connection.WriteJSON(Message{Type: 1, Body: "User Disconnected"})
			}
			break
		case message := <-pool.Brodcast:
			fmt.Println("Sending message to all clinets in the pool")
			for client, _ := range pool.Clients {
				if err := client.Connection.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}

	}
}

//
