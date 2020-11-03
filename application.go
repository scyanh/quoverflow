package main
/*
import (
	"github.com/scyanh/quoverflow/app/infrastructure/db"
	"github.com/scyanh/quoverflow/app/infrastructure/melodySocket"
	"github.com/scyanh/quoverflow/app/infrastructure/routes"
	"go.uber.org/zap"
	"os"
)

var (
	httpRouter = router.NewGinRouter()
)

func main() {
	//Db Connect and Close
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	db.Init()
	defer db.CloseDb()

	melodySocket.Init()
	defer melodySocket.CloseMelody()

	//Init Routes
	httpRouter.INIT()
	httpRouter.SETUP_MIDDLEWARES()
	httpRouter.SETUP_ROUTES()

	//db.GetDb().AutoMigrate(&model.Question{})

	//Run Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	httpRouter.SERVE(port)

}
*/