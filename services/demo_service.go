package services

import (
	"github.com/NothingXiang/go-study/models"
)

type DemoService interface {
	GetDemo(id string) (*models.DemoModel, error)
	SetDemo(model models.DemoModel) error
}
