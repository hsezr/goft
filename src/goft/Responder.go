package goft

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

func init() {
	ResponderList = []Responder{new(StringResponder), new(ModelResponder)}
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

var ResponderList []Responder

type StringResponder func(ctx *gin.Context) string
type ModelResponder func(ctx *gin.Context) Model

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(200, this(ctx))
	}
}

func (this ModelResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, this(ctx))
	}
}

func Convert(handler interface{}) gin.HandlerFunc {
	h_ref := reflect.ValueOf(handler)
	for _, r := range ResponderList {
		r_ref := reflect.ValueOf(r).Elem()
		if h_ref.Type().ConvertibleTo(r_ref.Type()) {
			r_ref.Set(h_ref)
			return r_ref.Interface().(Responder).RespondTo()
		}
	}

	return nil
}
