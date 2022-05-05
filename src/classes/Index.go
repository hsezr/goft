package classes

import (
	"mygin/goft"

	"github.com/gin-gonic/gin"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

func (this *IndexClass) GetIndex(ctx *gin.Context) goft.View {
	ctx.Set("name", "zhangsan")
	return "index"
}

func (this *IndexClass) Build(g *goft.Goft) {
	g.Handle("GET", "/", this.GetIndex)
}
