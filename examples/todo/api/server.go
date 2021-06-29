package api

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct{}

func (s *Server) Run() {
	router := httprouter.New()
	router.GET("/", nil)
	router.GET("/hello/:name", nil)

	log.Fatal(http.ListenAndServe(":8080", router))
}
