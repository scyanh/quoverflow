package questionController

import (
	"github.com/scyanh/quoverflow/app/domain/models"
)

type GetQuestionsSwagger struct {
	model.ServiceResponse
	Data []model.Question `json:"data"`
}