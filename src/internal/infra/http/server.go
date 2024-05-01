package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type TServer struct {
	host   string
	port   int
	router *gin.Engine
}

type IServer interface {
	Serve() error
}

func (s *TServer) Serve() error {
	return s.router.Run(fmt.Sprintf("%s:%d", s.host, s.port))
}

func NewServer(host string, port int, router *gin.Engine) (*TServer, error) {
	return &TServer{host: host, port: port, router: router}, nil
}
