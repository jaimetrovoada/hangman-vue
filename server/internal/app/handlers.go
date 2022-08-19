package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

func (s *Server) WebSocketHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		wsupgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		}
		wsupgrader.CheckOrigin(c.Request)
		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Printf("Failed to set websocket upgrade: %s", err)
		}

		for {
			t, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
			conn.WriteMessage(t, []byte(string("pong")))
		}
	}
}
