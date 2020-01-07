package services

import (
	"github.com/NothingXiang/go-study/models"
	"github.com/NothingXiang/go-study/store"
	"github.com/NothingXiang/go-study/store/mongo"
)

type DemoServiceIml struct {
	store.DemoStore
}

func NewDemoServiceIml() *DemoServiceIml {
	return &DemoServiceIml{DemoStore: &mongo.DemoMgoStore{}}
}

func (d *DemoServiceIml) GetDemo(id string) (m *models.DemoModel, err error) {
	if m, err = d.DemoStore.GetDemo(id); err != nil {
		// todo: 打log
		return nil, err
	}
	return
}

func (d *DemoServiceIml) SetDemo(model models.DemoModel) (err error) {
	if err = d.DemoStore.SetDemo(model); err != nil {
		// todo: 打log
		return err
	}
	return nil
}
