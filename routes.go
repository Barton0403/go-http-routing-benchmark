package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	gozerorouter "github.com/zeromicro/go-zero/rest/router"
	"net/http"
)

type route struct {
	method string
	path   string
}

func init() {
	// beego sets it to runtime.NumCPU()
	// Currently none of the contesters does concurrent routing
	//runtime.GOMAXPROCS(1)

	initBeego()
	initGin()
}

// beego
func beegoHandler(ctx *context.Context) {
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
	ctx.WriteString("success")
}

func initBeego() {
	beego.BConfig.RunMode = beego.PROD
	beego.BeeLogger.Close()
}

func loadBeego(routes []route) http.Handler {
	h := beegoHandler

	app := beego.NewControllerRegister()
	for _, route := range routes {
		switch route.method {
		case "GET":
			app.Get(route.path, h)
		case "POST":
			app.Post(route.path, h)
		case "PUT":
			app.Put(route.path, h)
		case "PATCH":
			app.Patch(route.path, h)
		case "DELETE":
			app.Delete(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return app
}

func loadBeegoSingle(method, path string, handler beego.FilterFunc) http.Handler {
	app := beego.NewControllerRegister()
	switch method {
	case "GET":
		app.Get(path, handler)
	case "POST":
		app.Post(path, handler)
	case "PUT":
		app.Put(path, handler)
	case "PATCH":
		app.Patch(path, handler)
	case "DELETE":
		app.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return app
}

// Echo
func echoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "success")
}

func loadEcho(routes []route) http.Handler {
	var h echo.HandlerFunc = echoHandler

	e := echo.New()
	for _, r := range routes {
		switch r.method {
		case "GET":
			e.GET(r.path, h)
		case "POST":
			e.POST(r.path, h)
		case "PUT":
			e.PUT(r.path, h)
		case "PATCH":
			e.PATCH(r.path, h)
		case "DELETE":
			e.DELETE(r.path, h)
		default:
			panic("Unknow HTTP method: " + r.method)
		}
	}
	return e
}

func loadEchoSingle(method, path string, h echo.HandlerFunc) http.Handler {
	e := echo.New()
	switch method {
	case "GET":
		e.GET(path, h)
	case "POST":
		e.POST(path, h)
	case "PUT":
		e.PUT(path, h)
	case "PATCH":
		e.PATCH(path, h)
	case "DELETE":
		e.DELETE(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return e
}

// Gin
func ginHandle(c *gin.Context) {
	c.String(http.StatusOK, "success")
	c.JSON()
}

func initGin() {
	gin.SetMode(gin.ReleaseMode)
}

func loadGin(routes []route) http.Handler {
	h := ginHandle

	router := gin.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.GET(route.path, h)
		case "POST":
			router.POST(route.path, h)
		case "PUT":
			router.PUT(route.path, h)
		case "PATCH":
			router.PATCH(route.path, h)
		case "DELETE":
			router.DELETE(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return router
}

func loadGinSingle(method, path string, handle gin.HandlerFunc) http.Handler {
	router := gin.New()
	switch method {
	case "GET":
		router.GET(path, handle)
	case "POST":
		router.POST(path, handle)
	case "PUT":
		router.PUT(path, handle)
	case "PATCH":
		router.PATCH(path, handle)
	case "DELETE":
		router.DELETE(path, handle)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return router
}

// go-restful
func goRestfulHandler(r *restful.Request, w *restful.Response) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func loadGoRestful(routes []route) http.Handler {
	h := goRestfulHandler

	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)

	for _, r := range routes {
		path := r.path

		switch r.method {
		case "GET":
			ws.Route(ws.GET(path).To(h))
		case "POST":
			ws.Route(ws.POST(path).To(h))
		case "PUT":
			ws.Route(ws.PUT(path).To(h))
		case "PATCH":
			ws.Route(ws.PATCH(path).To(h))
		case "DELETE":
			ws.Route(ws.DELETE(path).To(h))
		default:
			panic("Unknow HTTP method: " + r.method)
		}
	}
	wsContainer.Add(ws)
	return wsContainer
}

func loadGoRestfulSingle(method, path string, handler restful.RouteFunction) http.Handler {
	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)
	switch method {
	case "GET":
		ws.Route(ws.GET(path).To(handler))
	case "POST":
		ws.Route(ws.POST(path).To(handler))
	case "PUT":
		ws.Route(ws.PUT(path).To(handler))
	case "PATCH":
		ws.Route(ws.PATCH(path).To(handler))
	case "DELETE":
		ws.Route(ws.DELETE(path).To(handler))
	default:
		panic("Unknow HTTP method: " + method)
	}
	wsContainer.Add(ws)
	return wsContainer
}

// HttpRouter
func httpRouterHandle(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func loadHttpRouter(routes []route) http.Handler {
	h := httpRouterHandle

	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadHttpRouterSingle(method, path string, handle httprouter.Handle) http.Handler {
	router := httprouter.New()
	router.Handle(method, path, handle)
	return router
}

func goZeroHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("success"))
}

func loadGoZero(routes []route) http.Handler {
	h := http.HandlerFunc(goZeroHandle)

	// 模拟真实初始化
	var c rest.RestConf
	c.Log = logx.LogConf{
		Path: "./log",
	}
	s := rest.MustNewServer(c)
	s.AddRoute()

	router := gozerorouter.NewRouter()
	for _, r := range routes {
		router.Handle(r.method, r.path, h)
	}

	return router
}

func loadGoZeroSingle(method, path string, handle http.HandlerFunc) http.Handler {
	router := gozerorouter.NewRouter()
	router.Handle(method, path, handle)
	return router
}
