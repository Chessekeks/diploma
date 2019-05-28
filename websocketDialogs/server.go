package main

import (
	"diploma/domain"
	"log"
	"net/http"
)

type Server struct {
	clients  map[string]Client
	messages []domain.Message
	addCh    chan Client
	deleteCh chan Client
	doneCh   chan bool
	errCh    chan error
}

func newServer() *Server {
	clients := make(map[string]Client)
	messages := []domain.Message{}
	addCh := make(chan Client)
	deleteCh := make(chan Client)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		clients,
		messages,
		addCh,
		deleteCh,
		doneCh,
		errCh,
	}
}

func (s *Server) listen(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	id := readCookie(r, "UIT")
	client := newClient(conn, id)
	s.add(client)
	client.Listen()

	for {
		select {
		case c := <-s.addCh:
			s.clients[c.id] = c
		case c := <-s.deleteCh:
			delete(s.clients, c.id)
		case err := <-s.errCh:
			log.Println(err.Error())
		case <-s.doneCh:
			return
		}

	}
}

func (s *Server) add(c Client) {
	s.addCh <- c
}

func (s *Server) del(c Client) {
	s.deleteCh <- c
}

func (s *Server) done() {
	s.doneCh <- true
}

func (s *Server) err(err error) {
	s.errCh <- err
}
