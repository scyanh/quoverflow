package router

import (
	"github.com/gin-gonic/gin"
	"github.com/scyanh/quoverflow/app/infrastructure/controllers/questionController"
)

func QuestionRoutes(v1 *gin.RouterGroup, httpRouter Router) {
	questionR := v1.Group("questions")
	httpRouter.GET(questionR,"", questionController.GetQuestions)
	httpRouter.GET(questionR,"/id/:questionID", questionController.GetQuestion)
	httpRouter.POST(questionR,"", questionController.CreateQuestion)
	httpRouter.PUT(questionR,"", questionController.UpdateAnswer)
	httpRouter.PUT(questionR,"/upvote", questionController.UpvoteQuestion)

}
