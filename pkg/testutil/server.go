package testutil

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)


func CreateServer() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func MakeRequest(method, url, body string)( *http.Request, *httptest.ResponseRecorder){
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Contet-Type", "application/json")
	return req, httptest.NewRecorder()
}