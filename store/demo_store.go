/*
 *持久层
 */
package store

import (
	"github.com/NothingXiang/go-study/models"
)

type DemoStore interface {
	GetDemo(id string) (*models.DemoModel, error)
	SetDemo(model models.DemoModel) error
}
