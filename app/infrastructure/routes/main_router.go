package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	INIT()
	SETUP_MIDDLEWARES()
	SETUP_ROUTES()
	POST(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context))
	DELETE(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context))
	PUT(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context))
	GET(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context))
	SERVE(port string)
}
