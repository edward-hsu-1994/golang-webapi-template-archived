package core

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type GinRegexRoutes struct {
	*GinEngine
	basePath string
}

func (this *GinRegexRoutes) Use(handlers ...gin.HandlerFunc) *GinRegexRoutes {
	this.GinEngine.Use(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == this.basePath ||
			strings.HasPrefix(path, this.basePath+"/") {
			for _, handler := range handlers {
				handler(ctx)
			}
		}

		ctx.Next()
	})

	return this
}

func (this *GinRegexRoutes) Handle(httpMethod string, regexPattern string, handlers ...gin.HandlerFunc) *GinRegexRoutes {
	this.Any(regexPattern, func(ctx *gin.Context) {
		if strings.EqualFold(ctx.Request.Method, httpMethod) == false {
			ctx.Next()
			return
		}

		for _, handler := range handlers {
			handler(ctx)
		}
	})

	return this
}

func (this *GinRegexRoutes) Any(regexPattern string, handlers ...gin.HandlerFunc) *GinRegexRoutes {
	regex := regexp.MustCompile(regexPattern)

	this.Use(func(ctx *gin.Context) {
		subPath := ctx.Request.URL.Path[len(this.basePath):]
		if regex.MatchString(subPath) == false {
			ctx.Next()
			return
		}

		pathParamsValue := regex.FindStringSubmatch(subPath)
		for index, paramName := range regex.SubexpNames() {
			if len(paramName) == 0 {
				continue
			}
			ctx.AddParam(paramName, pathParamsValue[index])
		}

		for _, handler := range handlers {
			handler(ctx)
		}
	})

	return this
}

func (this *GinRegexRoutes) GET(regexPattern string, handlers ...gin.HandlerFunc) *GinRegexRoutes {
	this.Handle(http.MethodGet, regexPattern, handlers...)
	return this
}

func (this *GinRegexRoutes) POST(regexPattern string, handlers ...gin.HandlerFunc) *GinRegexRoutes {
	this.Handle(http.MethodPost, regexPattern, handlers...)
	return this
}

func (this *GinRegexRoutes) DELETE(regexPattern string, handlers ...gin.HandlerFunc) *GinRegexRoutes {
	this.Handle(http.MethodDelete, regexPattern, handlers...)
	return this
}

func (this *GinRegexRoutes) PATCH(regexPattern string, handlers ...gin.HandlerFunc) *GinRegexRoutes {
	this.Handle(http.MethodPatch, regexPattern, handlers...)
	return this
}

func (this *GinRegexRoutes) PUT(regexPattern string, handlers ...gin.HandlerFunc) *GinRegexRoutes {
	this.Handle(http.MethodPut, regexPattern, handlers...)
	return this
}

func (this *GinRegexRoutes) OPTIONS(regexPattern string, handlers ...gin.HandlerFunc) *GinRegexRoutes {
	this.Handle(http.MethodOptions, regexPattern, handlers...)
	return this
}

func (this *GinRegexRoutes) HEAD(regexPattern string, handlers ...gin.HandlerFunc) *GinRegexRoutes {
	this.Handle(http.MethodHead, regexPattern, handlers...)
	return this
}
