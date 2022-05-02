package classes

import (
	"mygin/src/goft"
	"mygin/src/models"

	"github.com/gin-gonic/gin"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) UserList(ctx *gin.Context) string {
	return "abc"
}

func (this *UserClass) UserDetail(ctx *gin.Context) goft.Model {
	return &models.UserModel{UserId: 101, UserName: "Hsezr"}
}

func (this *UserClass) Build(g *goft.Goft) {
	g.Handle("GET", "/user1", this.UserList)
	g.Handle("GET", "/user2", this.UserDetail)
}


