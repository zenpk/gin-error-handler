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
	Seq  int64  `json:"seq" eh:"-1"`                      // the default seq is -1 when an uncaught error happened
	Code int64  `json:"code" eh:"pre: CodeUncaughtError"` // will be filled with Preset.CodeUncaughtError (551)
	Msg  string `json:"msg" eh:"err"`                     // will eventually be err.Error()
	User *User  `json:"user,omitempty" eh:"nil"`          // will omit this field
}

func handler(c *gin.Context) {
	// some error occurred
	err := errors.New("something went wrong")
	eh.Handle(c, UserLoginResp{}, err)
}

func main() {
	// make a mock request
	req, _ := http.NewRequest(http.MethodGet, "/err", nil)
	// record the mock request
	rec := httptest.NewRecorder()
	// use Gin to handle the request
	r := gin.Default()
	r.GET("/err", handler)
	r.ServeHTTP(rec, req)
	fmt.Println(rec.Body.String())
}
