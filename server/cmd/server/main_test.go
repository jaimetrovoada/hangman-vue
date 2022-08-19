package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
