package server

import (
	"fmt"
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
	s.engine.GET("/healthz", s.Healthz)
	s.engine.GET("/readyz", s.Readyz)
	s.engine.GET("/tests", s.ListTests)
	s.engine.POST("/tests", s.CreateTest)

	return s.engine
}

func isDataLoaded() bool {
	return true
}

func (s *Server) ListTests(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	return
}

func (s *Server) CreateTest(c *gin.Context) {
	str, ok := c.Get("test")
	if ok {
		fmt.Println(str)
	} else {
		fmt.Println("test")
	}

	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	return
}
