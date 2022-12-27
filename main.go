package main

import (
	"fmt"
	"net/http"

	websocket "github.com/Chakravarthy712/sockets/pkg/websockets"
)

func serverWS(pool *websocket.Pool, w *http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint reached")

	connection, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &websocket.Client{
		Connection: connection,
		Pool:       pool,
	}

	pool.Register <- client

	client.Read()

}

func setUpRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWS(pool, w, r)
	})
}

func main() {
	fmt.Println("Chakravarthy full stack websocketes lab.")
	setUpRoutes()

	http.ListenAndServe(":9000", nil)
}
