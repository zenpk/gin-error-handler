# gin-error-handler

Automatically returns JSON when you don't want to bother with the err!=nil, with Gin framework.

## Usage

Copy the `eh.go` and `preset.go` file to wherever you want.

Define your Data Transfer Object with "eh" tags.

There are 3 meaningful tags, others stands for the default value:

1. `err` - the field with `err` tag will be returned as a string, its value equals to `err.Error()`
2. `nil` - the field with `nil` tag will be omitted
3. `pre: ` - set the field with value in `Preset` type, e.g. `pre: CodeOK` will set the field with `Preset.CodeOK` (200,
   int64). Be aware of the spelling: `p`,`r`,`e`,`colon`,`space`. The type of the original field and the Preset must be
   the same, otherwise it will panic.

## Example

```go
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

```

### Output

```text
{"seq":-1,"code":551,"msg":"something went wrong"}
```
