package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jaimetrovoada/hangman/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	r := gin.Default()
	server := app.NewServer(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/ping", nil)
	server.Routes().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestWebSocket(t *testing.T) {
	r := gin.Default()
	server := app.NewServer(r)

	d := websocket.Dialer{}
	s := httptest.NewServer(server.Routes().Handler())
	defer s.Close()
	_, resp, err := d.Dial("ws://"+s.Listener.Addr().String()+"/v1/api/ws", nil)
	if err != nil {
		t.Fatal(err)
	}

	got, want := resp.StatusCode, http.StatusSwitchingProtocols

	// msgT, _, err := c.ReadMessage()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// msg := []byte(string("ping"))
	// if err := c.WriteMessage(msgT, msg); err != nil {
	// 	t.Fatal(err)
	// }

	if ok := assert.Equal(t, want, got); !ok {
		t.Errorf("want:\t%d\ngot:\t%d", want, got)
	}
}
