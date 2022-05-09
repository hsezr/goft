package classes

import (
	"log"
	"mygin/goft"
	"mygin/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ArticleClass struct {
	*goft.GormAdapter
}

func NewArticleClass() *ArticleClass {
	return &ArticleClass{}
}

func (this *ArticleClass) ArticleDetail(ctx *gin.Context) goft.Model {
	user := models.NewArticleModel()
	goft.Error(ctx.BindUri(user))
	goft.Error(this.Table("users").Where("user_id=?", user.UserId).Find(user).Error)

	goft.Task(this.UpdateViews, nil, user.UserId)
	return user
}

func (this *ArticleClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/article/:id", this.ArticleDetail)
}

func (this *ArticleClass) UpdateViews(params ...interface{}) {
	this.Table("users").Where("user_id=?", params[0]).Update("user_view", gorm.Expr("user_view+1"))
}

func (this *ArticleClass) Test() interface{} {
	log.Println("测试定时任务")
	return nil
}

func (this *ArticleClass) Name() string {
	return "ArticleClass"
}