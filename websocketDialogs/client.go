package main

import (
	"diploma/domain"
	"github.com/gorilla/websocket"
	"io"
)

type Client struct {
	id     string
	ws     *websocket.Conn
	msgCh  chan domain.Message
	doneCh chan bool
}

func (c Client) Listen() {
	go c.listen()
	c.write()
}

//TODO: доделать метод https://github.com/gorilla/websocket/blob/master/examples/chat/client.go
func (c Client) listen() {
	for {
		select {
		case <-c.doneCh:
			return
		default:
			var msg domain.Message
			err := websocket.ReadJSON(c.ws, &msg)
			if err == io.EOF {
				c.doneCh <- true
				return
			}
			if err != nil {
				return
			}
		}
	}
}

//TODO: доделать метод https://github.com/gorilla/websocket/blob/master/examples/chat/client.go
func (c Client) write() {

}

func newClient(ws *websocket.Conn, id string) Client {
	msgCh := make(chan domain.Message)
	doneCh := make(chan bool)

	return Client{
		ws:     ws,
		msgCh:  msgCh,
		doneCh: doneCh,
	}
}
