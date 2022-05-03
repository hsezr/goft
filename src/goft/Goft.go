package goft

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	RG  *gin.RouterGroup
	dba interface{}
}

func Ignite() *Goft {
	g := &Goft{Engine: gin.New()}
	g.Use(ErrorHandler())
	return g
}

func (this *Goft) Launch() {
	this.Run(":8080")
}

func (this *Goft) Mount(group string, classes ...IClass) *Goft {
	this.RG = this.Group(group)
	for _, class := range classes {
		class.Build(this)
		vClass := reflect.ValueOf(class).Elem()
		if vClass.NumField() > 0 {
			if this.dba != nil {
				vClass.Field(0).Set(reflect.New(vClass.Field(0).Type().Elem()))
				vClass.Field(0).Elem().Set(reflect.ValueOf(this.dba).Elem())
			}
		}
	}

	return this
}

func (this *Goft) Handle(httpMethod, relativePath string, handler interface{}) *Goft {
	if h := Convert(handler); h != nil {
		this.RG.Handle(httpMethod, relativePath, h)
	}

	return this
}

func (this *Goft) Attach(f Fairing) *Goft {
	this.Use(func(ctx *gin.Context) {
		err := f.OnRequest(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		} else {
			ctx.Next()
		}
	})
	return this
}

//设定数据库连接对象
func (this *Goft) DB(dba interface{}) *Goft {
	this.dba = dba
	return this
}
