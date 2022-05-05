package goft

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	RG          *gin.RouterGroup
	beanFactory *BeanFactory
}

func Ignite() *Goft {
	g := &Goft{Engine: gin.New(), beanFactory: NewBeanFactory()}
	g.Use(ErrorHandler())
	config := InitConfig()
	g.beanFactory.setBean(config) //整个配置加载进bean中
	if config.Server.Html != "" {
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
	this.Run(fmt.Sprintf(":%d", port))
}

func (this *Goft) Mount(group string, classes ...IClass) *Goft { // 这是挂载， 后面还需要加功能。
	this.RG = this.Group(group)
	for _, class := range classes {
		class.Build(this) //这一步是关键 。 这样在main里面 就不需要 调用了
		this.beanFactory.inject(class)
	}
	return this
}

func (this *Goft) Handle(httpMethod, relativePath string, handler interface{}) *Goft {
	if h := Convert(handler); h != nil {
		fmt.Println(h)
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
	this.beanFactory.setBean(beans...)
	return this
}
