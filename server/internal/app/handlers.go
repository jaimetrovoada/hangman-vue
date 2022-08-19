package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		// response := map[string]string{
		// 	"status": "UP",
		// }

		c.String(http.StatusOK, "pong")
	}
}
