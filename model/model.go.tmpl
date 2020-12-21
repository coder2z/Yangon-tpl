package {{tableName}}

import (
	"{{ProjectName}}/internal/{{appName}}/model"
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
func (a *{{TableName}}) GetAll(data *[]{{TableName}}, wheres map[string][]interface{}) (err error) {
	db := model.MainDB.Table(a.TableName())
	for s, i := range wheres {
		db = db.Where(s, i...)
	}
	err = db.Find(&data).Error
	return
}

//偏移查询
func (a *{{TableName}}) Get(start int64, size int64, data *[]{{TableName}}, wheres map[string]interface{}, isDelete bool) (total int64, err error) {
	db := model.MainDB.Table(a.TableName())
	for k, v := range wheres {
		db = db.Where(k, v)
	}
	if isDelete {
		db = db.Unscoped().Where("{{tableName}}.deleted_at is not null")
	} else {
		db = db.Where(map[string]interface{}{"deleted_at": nil})
	}
	err = db.Limit(size).Offset(start).Find(&data).Error
	err = db.Count(&total).Error
	return
}

func (a *{{TableName}}) GetWhere(wheres map[string][]interface{}) (err error) {
	db := model.MainDB.Table(a.TableName())
	for s, i := range wheres {
		db = db.Where(s, i...)
	}
	err = db.Find(a).Error
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

//软删除恢复
func (a *{{TableName}}) DelRes() error {
	return model.MainDB.Table(a.TableName()).Where("id=?", a.{{ID}}).Update("deleted_at", nil).Error
}

