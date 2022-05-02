package classes

import (
	"mygin/src/goft"

	"github.com/gin-gonic/gin"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}


func (this *UserClass) UserList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"result":"user success",
		})
	}
}

func (this *UserClass) Build(g *goft.Goft) {
	g.Handle("GET", "/user", this.UserList())
}