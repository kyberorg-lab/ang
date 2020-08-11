package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	app := initApp()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r := httptest.NewRecorder()

	app.router.ServeHTTP(r, req)

	checkStatusCode(r.Code, http.StatusOK, t)
	checkContentType(r, t)

	var pingJson Ping
	_ = json.Unmarshal(r.Body.Bytes(), &pingJson)
	if pingJson.Ping != "pong" {
		t.Errorf("Ping response should match: got %v want %v", pingJson.Ping, "pong")
	}
}

func checkStatusCode(code int, want int, t *testing.T) {
	if code != want {
		t.Errorf("Wrong status code: got %v want %v", code, want)
	}
}

func checkContentType(r *httptest.ResponseRecorder, t *testing.T) {
	ct := r.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("Wrong Content Type: got %v want application/json", ct)
	}
}

func initApp() App {
	app := App{}
	app.router = setupRouter(app)
	return app
}

func setupRouter(app App) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", app.getPing)
	r.NoRoute(app.serveStatic)
	return r
}

func main() {
	a := initApp()
	log.Fatal(a.router.Run(":8080"))
}
