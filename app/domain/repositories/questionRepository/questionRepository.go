package questionRepository

import (
	"errors"
	"github.com/scyanh/quoverflow/app/domain/models"
	"github.com/scyanh/quoverflow/app/infrastructure/db"
)

func Create(question model.Question) (*model.Question, error) {
	if err := db.GetDb().Create(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}
func Update(question model.Question) (*model.Question, error) {
	if err := db.GetDb().Save(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func Find() []model.Question {
	var questions []model.Question

	db.GetDb().
		Find(&questions)

	return questions
}

func FindByID(questionID int64) (*model.Question, error) {
	var question model.Question

	if db.GetDb().
		Where("id = ?", questionID).
		First(&question).RecordNotFound() {
		return nil, errors.New("Question not found")
	}

	return &question, nil
}
