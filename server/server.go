package server

import "github.com/gin-gonic/gin"

type Server interface {
	Start()
}

type server struct {
	gin *gin.Engine
}

func New() Server {
	return &server{}
}

func (s *server) Start() {
	s.gin = gin.Default()

	s.gin.Run()
}
