package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	http.NewRequest("GET", "/crawler", nil)
	w := httptest.NewRecorder()
	http.NewRequest("GET", "/animes/1", nil)
	assert.Equal(t, 200, w.Code)
}
