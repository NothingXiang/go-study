/*
 *实体类
*/
package models

type DemoModel struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
