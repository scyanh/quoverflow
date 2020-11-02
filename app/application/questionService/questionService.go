package questionService

import (
	"encoding/json"
	"fmt"
	"github.com/scyanh/quoverflow/app/domain/models"
	"github.com/scyanh/quoverflow/app/domain/repositories/questionRepository"
	"github.com/scyanh/quoverflow/app/infrastructure/melodySocket"
)

func GetQuestions(page int) []model.Question {
	questions := questionRepository.Find(page)
	return questions
}

// Deprecated
func GetQuestionsByIds(ids []int) []model.Question {
	questions := questionRepository.FindByIDs(ids)
	return questions
}

func GetQuestion(productID int64) (*model.Question, error) {
	product, err := questionRepository.FindByID(productID)
	if err!=nil{
		return nil, err
	}
	return product, nil
}

func CreateQuestion(validator model.QuestionValidator) error {
	question := model.Question{
		Text: validator.Text,
	}

	if _, err := questionRepository.Create(question); err != nil {
		return err
	}
	callSocket()
	return nil
}

func UpdateAnswer(validator model.AnswerValidator) error {
	question, err := questionRepository.FindByID(validator.ID)
	if err != nil {
		return err
	}

	question.Answer = validator.Answer

	if _, err := questionRepository.Update(*question); err != nil {
		return err
	}
	callSocket()
	return nil
}
func UpvoteQuestion(validator model.QuestionUpvoteValidator) error {
	question, err := questionRepository.FindByID(validator.ID)
	if err != nil {
		return err
	}

	question.ID = validator.ID
	question.Upvotes++

	if _, err := questionRepository.Upvote(*question); err != nil {
		return err
	}
	callSocket()
	return nil
}

func callSocket(){
	fmt.Println("in callSocket()")
	questions := GetQuestions(1)
	b, err := json.Marshal(questions)
	if err != nil {
		fmt.Println("err=",err)
		return
	}
	melodySocket.GetMelody().Broadcast(b)
}
