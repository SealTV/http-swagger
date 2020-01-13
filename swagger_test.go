package httpSwagger

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/SealTV/http-swagger/example/go-chi/docs"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestWrapHandler(t *testing.T) {
	router := chi.NewRouter()

	router.Get("/*", WrapHandler)

	w1 := performRequest("GET", "/index.html", router)
	assert.Equal(t, 200, w1.Code)

	w2 := performRequest("GET", "/doc.json", router)
	assert.Equal(t, 200, w2.Code)

	w3 := performRequest("GET", "/favicon-16x16.png", router)
	assert.Equal(t, 200, w3.Code)

	w4 := performRequest("GET", "/notfound", router)
	assert.Equal(t, 404, w4.Code)

	w5 := performRequest("GET", "/", router)
	assert.Equal(t, 301, w5.Code)
}

func performRequest(method, target string, h http.Handler) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)
	return w
}
