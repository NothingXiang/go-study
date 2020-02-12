package mongo

import (
	"github.com/NothingXiang/go-study/common/dbutil"
	"github.com/NothingXiang/go-study/models"
)

const Col_Demo = "demo"

type DemoMgoStore struct {
}

func (d *DemoMgoStore) GetDemo(id string) (m *models.DemoModel, e error) {
	// connect to mongo db
	db := dbutil.NewMgoDB()
	defer db.Session.Close()

	if e = db.C(Col_Demo).FindId(id).One(&m); e != nil {
		return nil, e
	}

	return
}

func (d *DemoMgoStore) SetDemo(model models.DemoModel) error {
	// connect to mongo db
	db := dbutil.NewMgoDB()
	defer db.Session.Close()

	if err := db.C(Col_Demo).Insert(model); err != nil {
		return err
	}
	return nil
}
