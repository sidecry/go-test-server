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
	s.engine.GET("/healthz", s.Healthz)
	s.engine.GET("/readyz", s.Readyz)
	s.engine.GET("/tests", s.ListTests)

	return s.engine
}

func (s *Server) Healthz(c *gin.Context) {
	c.String(200, "ok")
	return
}

func (s *Server) Readyz(c *gin.Context) {
	mes := ""
	if !isDataLoaded() {
		mes += "Initial data is not loaded."
	}

	if len(mes) > 0 {
		c.String(503, mes)
	} else {
		c.String(200, "ok")
	}
	return
}

func isDataLoaded() bool {
	return true
}

func (s *Server) ListTests(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	return
}
