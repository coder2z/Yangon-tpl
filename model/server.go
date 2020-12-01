package server

import (
	_map "{{ProjectName}}/internal/{{AppName}}/map"
	"{{ProjectName}}/internal/{{AppName}}/model/{{tableName}}"
)

func GetAllServer(page _map.PageList) (data []{{tableName}}.{{TableName}}, total int64, err error) {
	data = make([]{{tableName}}.{{TableName}}, 0)
	total, err = new({{tableName}}.{{TableName}}).Get(page.PageSize*(page.Page-1), page.PageSize, &data, nil)
	return
}

func AddServer(add _map.AddServer) (err error) {
	//todo 添加手动赋值
	data := &{{tableName}}.{{TableName}}{

	}
	err = data.Add()
	return
}

func GetByIdServer(idMap _map.IdMap) (data *{{tableName}}.{{TableName}}, err error) {
	data = new({{tableName}}.{{TableName}})
	data.{{Id}} = idMap.Id
	err = data.GetById()
	return
}

func PutByIdServer(put _map.PutServer) (err error) {
	//todo 修改手动赋值
	data := &{{tableName}}.{{TableName}}{

	}
	err = data.UpdateById()
	return
}

func DelServer(idMap _map.IdMap) (err error) {
	err = new({{tableName}}.{{TableName}}).Del(map[string]interface{}{"{{id}}=?": idMap.Id})
	return
}
