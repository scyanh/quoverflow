package questionController

import (
	"github.com/scyanh/quoverflow/app/application/questionService"
	"github.com/scyanh/quoverflow/app/domain/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetQuestion(c *gin.Context) {
	responseContextData := model.ResponseContext{Ctx: c}
	questionID, _ := strconv.ParseInt(c.Param("questionID"), 10, 64)

	product, err := questionService.GetQuestion(questionID)

	if err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(model.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(model.StatusSuccess, "", product))
}

// questions response godoc
// @Summary questions response
// @Description get response questions
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

	c.JSON(http.StatusOK, responseContextData.ResponseData(model.StatusSuccess, "Answer created successfully", ""))
}

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
