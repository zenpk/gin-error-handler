# gin-error-handler

Automatically returns JSON when you don't want to bother with ***every*** err!=nil, using Gin framework.

## Usage

Copy the `eh.go` and `preset.go` file to wherever you want.

Define your Data Transfer Object with "eh" tags.

There are 3 meaningful tags, others stand for default values:

1. `eh:"err"` - the field with `err` tag will be returned as a string, its value equals to `err.Error()`
2. `eh:""` - the field with empty tag will remain the default value of the type
3. `eh:"pre: "` - the field will be set the same as the variable in `Preset` struct with the same name, e.g. field
   with `pre: CodeOK` will be set with `Preset.CodeOK` (200,
   int64). Be aware of the spelling: `p`,`r`,`e`,`colon`,`space`. The type of the original field and the according "
   Preset variable" must be
   the same, otherwise will cause a panic.

## Example

```go
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
	User *User  `json:"user,omitempty" eh:""`             // will omit this field
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
```

### Output

```text
{"seq":-1,"code":551,"msg":"something went wrong"}
```
