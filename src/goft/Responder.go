package goft

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

func init() {
	ResponderList = []Responder{new(StringResponder),
		new(ModelResponder),
		new(ModelsResponder),
		new(ViewResponder)}
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

var ResponderList []Responder

type StringResponder func(ctx *gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(200, this(ctx))
	}
}

type ModelResponder func(ctx *gin.Context) Model

func (this ModelResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, this(ctx))
	}
}

type ModelsResponder func(*gin.Context) Models

func (this ModelsResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-type", "application/json")
		ctx.Writer.WriteString(string(this(ctx)))
	}
}

type View string
type ViewResponder func(ctx *gin.Context) View

func (this ViewResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(200, string(this(ctx))+".html", ctx.Keys)
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
