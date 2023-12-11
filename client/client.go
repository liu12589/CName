package client

import (
	"CName/model"
	"CName/service"
)

func NewClient() *service.Client {
	return &service.Client{
		Ratio:         model.Ratio,
		Range:         model.Range,
		LastNameRange: model.LastNameRange,
		SurnameRange:  model.SurNameRange,
		Surname:       model.SurName,
		LastName:      model.LastName,
	}
}
