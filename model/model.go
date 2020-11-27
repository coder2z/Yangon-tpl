package {{tableName}}

import (
	"{{ProjectName}}/internal/{{AppName}}/model"
	"github.com/jinzhu/gorm"
	{{IsTime}}
)

func init() {
	model.MainDB.AutoMigrate(new({{TableName}}))
}

type {{TableName}} struct {
	{{TableFieldList}}
}

func (a *{{TableName}}) TableName() string {
	return "{{tableName}}"
}

//添加
func (a *{{TableName}}) Add() error {
	return model.MainDB.Table(a.TableName()).Create(a).Error
}

//删除where
func (a *{{TableName}}) Del(wheres map[string]interface{}) error {
	db := model.MainDB.Table(a.TableName())
	for k, v := range wheres {
		db = db.Where(k, v)
	}
	return db.Delete(a).Error
}

//查询所有
func (a *{{TableName}}) GetAll(data *[]{{TableName}}) (err error) {
	err = model.MainDB.Table(a.TableName()).Find(&data).Error
	return
}

//偏移查询
func (a *{{TableName}}) Get(start int64, size int64, data *[]{{TableName}}, wheres map[string]interface{}) (total int64, err error) {
	db := model.MainDB.Table(a.TableName())
	for k, v := range wheres {
		db = db.Where(k, v)
	}
	err = db.Limit(size).Offset(start).Find(&data).Error
	err = db.Count(&total).Error
	return
}

//根据id查询
func (a *{{TableName}}) GetById() error {
	return model.MainDB.Table(a.TableName()).Where("id=?", a.{{ID}}).First(a).Error
}

//修改ById
func (a *{{TableName}}) UpdateById() error {
	return model.MainDB.Table(a.TableName()).Where("id=?", a.{{ID}}).Update(a).Error
}