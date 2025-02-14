package httpsvr

import (
	"Feed/producer/handler"
	"Feed/producer/httpctx"
	"Feed/producer/service/elastic"
	"careyads/tools/logrus"
	"os"

	"github.com/kataras/iris"
)

type HttpServer struct {
	Es      *elastic.ElasticManager
	App     *iris.Application
	Context *httpctx.ServerContext
	Logger  *logrus.Logger
}

//HTTPserver包含几大部分功能
//1：路由
//2：保存session，因为session本质上也是存储在内存之中，golang也并没有原生支持session，所以可以直接将所有信息直接保存在系统里
//	也可以存储静态表，也可以存储用户信息
//3：保存service层及其他信息
func NewHttpServer(es *elastic.ElasticManager) *HttpServer {

	return &HttpServer{
		Es:  es,
		App: iris.New(),
		Context: &httpctx.ServerContext{
			EsManager: es,
		},
	}
}

func (this *HttpServer) Start() {

	//this.LoadData()

	this.App.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("Hello world!")
	})

	this.App.Post("/write", handler.WriteHandler)

	this.App.Post("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	err := this.App.Run(iris.Addr(":8090"), iris.WithoutServerError(iris.ErrServerClosed))

	if err != nil {

		this.Logger.Error("[http server] iris app run error " + err.Error())

		os.Exit(1)
	}

	//this.gin.GET("/ws", Controller.Upload)

}

func (this *HttpServer) LoadData() {
	this.Context.EsManager = this.Es
}

func (this *HttpServer) Wrap(handler func(ctx iris.Context, logger *logrus.Logger, context *httpctx.ServerContext)) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		handler(ctx, this.Logger, this.Context)
	}
}
