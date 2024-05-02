package tests

import (
	"letsgo/internal/app"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := app.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Expected HTTP status code to be 200")

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err, "Error decoding JSON")
	
	assert.Equal(t, "pong", response["message"], "Expected message to match 'pong'")
}