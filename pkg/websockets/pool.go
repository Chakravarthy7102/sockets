package websokets

type Pool struct {
	Register   chan *Clinet
	Unregister chan *Client
	Clients    map[*Client]bool
	Brodcast   chan Message
}

func NewPool() {}
