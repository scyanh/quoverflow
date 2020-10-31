package questionService

import (
	"github.com/scyanh/quoverflow/app/domain/models"
	"github.com/scyanh/quoverflow/app/domain/repositories/questionRepository"
)

func GetQuestions() []model.Question {
	questions := questionRepository.Find()
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

	return nil
}

func UpdateAnswer(validator model.AnswerValidator) error {
	question, err := questionRepository.FindByID(validator.QuestionID)
	if err != nil {
		return err
	}

	question.Answer = validator.Answer

	if _, err := questionRepository.Update(*question); err != nil {
		return err
	}
	return nil
}
