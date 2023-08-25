package main

import (
	"github.com/gin-gonic/gin"
)

type XEngine struct {
	*gin.Engine
}

type XContext struct {
	*gin.Context
}

func (x *XContext) WriteString(str string) (int, error) {
	return x.Writer.Write([]byte(str))
}

type XRouterGroup struct {
	*gin.RouterGroup
}

func NeXContext(ctx *gin.Context) *XContext {
	return &XContext{
		Context: ctx,
	}
}

type XHandlerFunc func(ctx *XContext) error

func wrapperContext(handlers ...XHandlerFunc) []gin.HandlerFunc {
	ret := make([]gin.HandlerFunc, 0, len(handlers))
	for _, xhandler := range handlers {
		hd := xhandler
		if nil != hd {
			r := func(ctx *gin.Context) {
				if err := hd(NeXContext(ctx)); err != nil {
					ctx.Error(err)
				}
			}
			ret = append(ret, r)
		}
	}
	return ret
}

func (x *XEngine) XGet(relativePath string, handlers ...XHandlerFunc) gin.IRoutes {
	return x.GET(relativePath, wrapperContext(handlers...)...)
}

func (x *XEngine) XPOST(relativePath string, handlers ...XHandlerFunc) gin.IRoutes {
	return x.POST(relativePath, wrapperContext(handlers...)...)
}

func (x *XEngine) XPATCH(relativePath string, handlers ...XHandlerFunc) gin.IRoutes {
	return x.PATCH(relativePath, wrapperContext(handlers...)...)
}

func (x *XEngine) XDELETE(relativePath string, handlers ...XHandlerFunc) gin.IRoutes {
	return x.DELETE(relativePath, wrapperContext(handlers...)...)
}

func (x *XEngine) XGroup(relativePath string, handlers ...XHandlerFunc) *XRouterGroup {
	return &XRouterGroup{x.Group(relativePath, wrapperContext(handlers...)...)}
}

func (x *XRouterGroup) XGET(relativePath string, handlers ...XHandlerFunc) gin.IRoutes {
	return x.GET(relativePath, wrapperContext(handlers...)...)
}

func (x *XRouterGroup) XPOST(relativePath string, handlers ...XHandlerFunc) gin.IRoutes {
	return x.POST(relativePath, wrapperContext(handlers...)...)
}

func (x *XRouterGroup) XPATCH(relativePath string, handlers ...XHandlerFunc) gin.IRoutes {
	return x.PATCH(relativePath, wrapperContext(handlers...)...)
}

func (x *XRouterGroup) XDELELTE(relativePath string, handlers ...XHandlerFunc) gin.IRoutes {
	return x.DELETE(relativePath, wrapperContext(handlers...)...)
}

func HelloHandler(ctx *XContext) error {
	_, err := ctx.WriteString("hello\n")
	return err
}

func main() {
	eg := &XEngine{gin.Default()}
	eg.XGet("/hello", HelloHandler)
	v1 := eg.XGroup("/v1")
	{
		v1.XGET("hellow", HelloHandler)
	}
	eg.Run()
}
