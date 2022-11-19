package server

import "github.com/gin-gonic/gin"

func (s *Server) Healthz(c *gin.Context) {
	c.String(200, "ok")
	return
}
