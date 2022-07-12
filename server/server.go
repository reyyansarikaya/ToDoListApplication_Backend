package server

import (
	"bootcamp/handler"
	"bootcamp/repository"
	"bootcamp/service"
	"fmt"
	"net/http"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(port string) error {
	repo := repository.NewRepository()
	serv := service.NewService(repo)
	hand := handler.NewHandler(serv)
	http.HandleFunc("/api/v1/todo", hand.HandlerEndpoints)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	return err
}
