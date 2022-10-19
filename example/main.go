package main

import (
	"errors"
	"fmt"
	eh "gin-error-handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

type User struct {
	Name string `json:"name,omitempty"`
}

type UserLoginResp struct {
	Seq  int64  `json:"seq" eh:"-1"`                     // will be -1
	Code int64  `json:"code" eh:"pre:CodeUncaughtError"` // will be filled with Preset.CodeUncaughtError (551)
	Msg  string `json:"msg" eh:"err"`                    // will eventually be err.Error()
	User *User  `json:"user,omitempty" eh:""`            // will omit this field
}

func handler(c *gin.Context) {
	// create a new handler with *gin.Context and your JSON interface
	errHandler := eh.JSONHandler{
		C: c,
		V: UserLoginResp{},
	}
	err := errors.New("something went wrong") // some error occurred
	errHandler.Handle(err)                    // handle the error
}

func main() {
	req, _ := http.NewRequest(http.MethodGet, "/err", nil) // make a mock request
	rec := httptest.NewRecorder()                          // record the mock request
	// use Gin to handle the request
	r := gin.Default()
	r.GET("/err", handler)
	r.ServeHTTP(rec, req)
	fmt.Println(rec.Body.String())
}
