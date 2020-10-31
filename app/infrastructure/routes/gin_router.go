package router

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/scyanh/quoverflow/app/infrastructure/middlewares"
	"github.com/scyanh/quoverflow/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

type ginRouter struct{}

var (
	ginEngine *gin.Engine
)

func NewGinRouter() Router {
	return &ginRouter{}
}

func (*ginRouter) POST(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context)) {
	if groupPath != nil {
		groupPath.POST(relativePath, f)
		return
	}
	ginEngine.POST(relativePath, f)
}
func (*ginRouter) DELETE(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context)) {
	if groupPath != nil {
		groupPath.DELETE(relativePath, f)
		return
	}
	ginEngine.DELETE(relativePath, f)
}
func (*ginRouter) PUT(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context)) {
	if groupPath != nil {
		groupPath.PUT(relativePath, f)
		return
	}
	ginEngine.PUT(relativePath, f)
}

func (*ginRouter) GET(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context)) {
	if groupPath != nil {
		groupPath.GET(relativePath, f)
		return
	}
	ginEngine.GET(relativePath, f)
}

func (*ginRouter) SERVE(port string) {
	fmt.Printf("Gin server running om port: %v ", port)
	ginEngine.Run(":" + port)
}

func (*ginRouter) INIT() {

	ENV_NAME := os.Getenv("ENV_NAME")
	if ENV_NAME == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	ginEngine = gin.New()
	ginEngine.Use(gin.Recovery())

	//Documentation start - programatically set swagger info
	docs.SwaggerInfo.Title = "Quoverflow API"
	docs.SwaggerInfo.Description = "This is a Quoverflow server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// use ginSwagger middleware to serve the API docs
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (*ginRouter) SETUP_MIDDLEWARES() {
	ginEngine.Use(middleware.CORSMiddleware())
}

func (*ginRouter) SETUP_ROUTES() {
	logger, _ := zap.NewProduction()
	ginEngine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	ginEngine.Use(ginzap.RecoveryWithZap(logger, true))
	v1 := ginEngine.Group("/v1")

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response": "ok"})
	})

	QuestionRoutes(v1, NewGinRouter())
}
