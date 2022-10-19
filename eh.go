package eh

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type JSONHandler struct {
	C *gin.Context
	V interface{}
}

// Handle return the given interface with tagged values and error messages
func (h *JSONHandler) Handle(err error) {
	ref := reflect.ValueOf(&h.V).Elem()
	refCopy := reflect.New(ref.Elem().Type()).Elem()
	for i := 0; i < ref.Elem().NumField(); i++ {
		tag := ref.Elem().Type().Field(i).Tag.Get("eh")
		if tag == "err" {
			refCopy.Field(i).SetString(err.Error())
			continue
		}
		if tag == "nil" {
			continue
		}
		if len(tag) > 5 && tag[0:5] == "pre: " {
			field := tag[5:]
			presetVal := reflect.ValueOf(Preset).FieldByName(field)
			refCopy.Field(i).Set(presetVal)
			continue
		}
		switch ref.Elem().Field(i).Interface().(type) {
		case string:
			refCopy.Field(i).SetString(tag)
		case int64, int32, int:
			val, _ := strconv.ParseInt(tag, 10, 64)
			refCopy.Field(i).SetInt(val)
		case uint64, uint32, uint:
			val, _ := strconv.ParseUint(tag, 10, 64)
			refCopy.Field(i).SetUint(val)
		case bool:
			val, _ := strconv.ParseBool(tag)
			refCopy.Field(i).SetBool(val)
		case float64, float32:
			val, _ := strconv.ParseFloat(tag, 64)
			refCopy.Field(i).SetFloat(val)
		default:
			// other implementations...
		}
	}
	log.Printf("uncaught error: %v, returned JSON\n", err) // do whatever you want here
	ref.Set(refCopy)
	h.C.JSON(http.StatusOK, h.V)
}
