package server

import "github.com/gin-gonic/gin"

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
