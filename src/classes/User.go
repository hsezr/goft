package classes

import (
	"mygin/goft"
	"mygin/models"

	"github.com/gin-gonic/gin"
)

type UserClass struct {
	*goft.GormAdapter
	Age *goft.Value `prefix:"user.age"`
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) UserTest(ctx *gin.Context) string {
	return "abc" + this.Age.String() 
}

func (this *UserClass) UserDetail(ctx *gin.Context) goft.Model {
	user := models.NewUserModel()
	err := ctx.BindUri(user)
	goft.Error(err, "ID参数不合法")
	this.Table("users").Where("user_id=?", user.UserId).Find(user)
	return user
}

func (this *UserClass) UserList(ctx *gin.Context) goft.Models {
	users := []*models.UserModel{
		&models.UserModel{UserId: 101, UserName: "Hsezr"},
		&models.UserModel{UserId: 102, UserName: "zhangsan"},
	}

	return goft.MakeModels(users)
}

func (this *UserClass) Build(g *goft.Goft) {
	g.Handle("GET", "/test", this.UserTest)
	g.Handle("GET", "/user/:id", this.UserDetail)
	g.Handle("GET", "/user3", this.UserList)
}

func (this *UserClass) Name() string {
	return "UserClass"
}
