package main

import "github.com/kataras/iris"

func main() {
	//创建iris
	app := iris.New()
	//设置错误日志模式
	app.Logger().SetLevel("DEBUG")
	//注册模板
	template := iris.HTML("/backend/web/viwes", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//设置模板目标
	app.StaticWeb("/assets", "./backend/web/assets")
	//出现异常跳转至顶页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错!"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	app.Run(
		iris.Addr("localhost:8080"),
		//iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
