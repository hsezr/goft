package goft

import (
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	RG *gin.RouterGroup
}

func Ignite() *Goft {
	return &Goft{Engine: gin.New()}
}

func (this *Goft) Launch() {
	this.Run(":8080")
}

func (this *Goft) Mount(group string, classes ...IClass) (*Goft) {
	this.RG = this.Group(group)
	for _, class := range classes {
		class.Build(this)
	}

	return this
}

func (this *Goft) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Goft {
	this.RG.Handle(httpMethod, relativePath, handlers...)

	return this
}

func (this *Goft) Attach(f gin.HandlerFunc) *Goft {
	this.Use(f)
	return this
}
