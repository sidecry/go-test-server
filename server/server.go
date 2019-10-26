package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func New() *Server {
	return &Server{
		engine: gin.New(),
	}
}

func (s *Server) Handler() http.Handler {
	s.engine.GET("/tests", s.ListTests)

	return s.engine
}

func (s *Server) ListTests(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	return
}
