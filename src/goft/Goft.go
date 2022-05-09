package goft

import (
	"fmt"
	"log"
	"mygin/funcs"

	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	RG          *gin.RouterGroup
	beanFactory *BeanFactory
	exprData    map[string]interface{}
}

func Ignite() *Goft {
	g := &Goft{Engine: gin.New(), beanFactory: NewBeanFactory(), exprData: map[string]interface{}{}}
	g.Use(ErrorHandler())
	config := InitConfig()
	g.beanFactory.setBean(config) //整个配置加载进bean中
	if config.Server.Html != "" {
		g.FuncMap = funcs.FuncMap
		g.LoadHTMLGlob(config.Server.Html)
	}

	return g
}

func (this *Goft) Launch() { //最终启动函数， 不用run，没有逼格
	//config:=InitConfig()
	var port int32 = 8080
	if config := this.beanFactory.GetBean(new(SysConfig)); config != nil {
		port = config.(*SysConfig).Server.Port
	}
	getCronTask().Start()
	this.Run(fmt.Sprintf(":%d", port))
}

func (this *Goft) Mount(group string, classes ...IClass) *Goft { // 这是挂载， 后面还需要加功能。
	this.RG = this.Group(group)
	for _, class := range classes {
		class.Build(this) //这一步是关键 。 这样在main里面 就不需要 调用了
		this.beanFactory.inject(class)
		this.Beans(class)
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
func (this *Goft) Beans(beans ...Bean) *Goft {
	for _, bean := range beans {
		this.exprData[bean.Name()] = bean
	}

	this.beanFactory.setBean(beans...)
	return this
}

//增加定时任务
func (this *Goft) Task(cron string, expr interface{}) *Goft {
	var err error
	if f, ok := expr.(func()); ok {
		_, err = getCronTask().AddFunc(cron, f)
	} else if exp, ok := expr.(Expr); ok {
		_, err = getCronTask().AddFunc(cron, func() {
			_, expErr := ExecExpr(exp, this.exprData)
			if expErr != nil {
				log.Println(expErr)
			}
		})
	}

	if err != nil {
		log.Println(err)
	}

	return this
}
