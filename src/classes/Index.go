package classes

import (
	"mygin/src/goft"

	"github.com/gin-gonic/gin"
)

type IndexClass struct {
	
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

func (this *IndexClass) GetIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"result":"index success",
		})
	}
}

func (this *IndexClass) Build(g *goft.Goft) {
	g.Handle("GET", "/", this.GetIndex())
}