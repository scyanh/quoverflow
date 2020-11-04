package questionController

import (
	"github.com/scyanh/quoverflow/app/application/questionService"
	"github.com/scyanh/quoverflow/app/domain/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Get question detail response godoc
// @Summary question response
// @Tags question
// @Accept  json
// @Produce  json
// @Param questionID path int true "id"
// @Success 200 {object} questionController.GetQuestionSwagger
// @Router /v1/questions/id/{questionID} [get]
func GetQuestion(c *gin.Context) {
	responseContextData := model.ResponseContext{Ctx: c}
	questionID, _ := strconv.ParseInt(c.Param("questionID"), 10, 64)

	question, err := questionService.GetQuestion(questionID)

	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(model.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(model.StatusSuccess, "", question))
}

// Get questions response godoc
// @Summary questions response
// @Tags question
// @Accept  json
// @Produce  json
// @Success 200 {array} questionController.GetQuestionsSwagger
// @Router /v1/questions [get]
func GetQuestions(c *gin.Context) {
	responseContextData := model.ResponseContext{Ctx: c}
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	questions := questionService.GetQuestions(page)

	c.JSON(http.StatusOK, responseContextData.ResponseData(model.StatusSuccess, "", questions))
}

// Create question godoc
// @Summary Create question
// @Tags question
// @Accept  json
// @Produce  json
// @Param QuestionValidator body model.QuestionValidator true "QuestionValidator"
// @Success 200
// @Router /v1/questions [post]
func CreateQuestion(c *gin.Context) {
	responseContextData := model.ResponseContext{Ctx: c}
	var validator model.QuestionValidator

	err := c.ShouldBindJSON(&validator)
	if err != nil {
		fmt.Println("err=", err)
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(model.StatusFail, err.Error(), ""))
		return
	}

	if err := questionService.CreateQuestion(validator); err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(model.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(model.StatusSuccess, "Question created successfully", ""))
}

// Update answer godoc
// @Summary Update answer
// @Tags question
// @Accept  json
// @Produce  json
// @Param AnswerValidator body model.AnswerValidator true "AnswerValidator"
// @Success 200
// @Router /v1/questions [put]
func UpdateAnswer(c *gin.Context) {
	responseContextData := model.ResponseContext{Ctx: c}
	var validator model.AnswerValidator

	err := c.ShouldBindJSON(&validator)
	if err != nil {
		fmt.Println("err=", err)
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(model.StatusFail, err.Error(), ""))
		return
	}

	if err := questionService.UpdateAnswer(validator); err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(model.StatusFail, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, responseContextData.ResponseData(model.StatusSuccess, "Answer updated successfully", ""))
}

// Update question godoc
// @Summary Update question
// @Tags question
// @Accept  json
// @Produce  json
// @Param QuestionUpvoteValidator body model.QuestionUpvoteValidator true "QuestionUpvoteValidator"
// @Success 200
// @Router /v1/questions/upvote [put]
func UpvoteQuestion(c *gin.Context) {
	responseContextData := model.ResponseContext{Ctx: c}
	var validator model.QuestionUpvoteValidator

	err := c.ShouldBindJSON(&validator)
	if err != nil {
		fmt.Println("err=", err)
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(model.StatusFail, err.Error(), ""))
		return
	}

	if err := questionService.UpvoteQuestion(validator); err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(model.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(model.StatusSuccess, "", ""))
}
