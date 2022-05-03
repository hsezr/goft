package goft

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	RG  *gin.RouterGroup
	props []interface{}
}

func Ignite() *Goft {
	g := &Goft{Engine: gin.New(), props: make([]interface{}, 0)}
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
		this.setProp(class)
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
func (this *Goft) Beans(beans ...interface{}) *Goft {
	this.props = append(this.props, beans...) 
	return this
}

func (this *Goft) getProp(t reflect.Type) interface{} {
	for _, p := range this.props {
		if t == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}


func (this *Goft) setProp(class IClass) {
	vClass := reflect.ValueOf(class).Elem()
	for i := 0; i < vClass.NumField(); i++ {
		f := vClass.Field(i)
		if !f.IsNil() || f.Kind() != reflect.Ptr {
			continue
		}

		if p := this.getProp(f.Type()); p != nil {
			f.Set(reflect.New(vClass.Field(0).Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}
