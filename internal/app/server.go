package app

import (
	"fmt"
	"graduatework/internal/config"
	"graduatework/internal/handler"
	dcollect "graduatework/internal/microservices"
	"graduatework/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	Server http.Server
	Router *mux.Router
	micro  *dcollect.MicroServiceStr
}

func NewServer() *server {
	return &server{
		Server: http.Server{},
		Router: mux.NewRouter(),
	}

}

func RunServer() {
	s := NewServer()
	microServ := s.micro.MicroService()
	service := service.NewServiceManage(microServ)
	handle := handler.NewHandler(service)
	handle.RegisterR(s.Router)
	fmt.Print("Server started\n")
	err := http.ListenAndServe(config.GlobalConfig.Addr, s)
	if err != nil {
		log.Fatal("Server didn't start: ", err)
	}

}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}