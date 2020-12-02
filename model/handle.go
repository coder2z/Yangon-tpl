package handle

import (
	"{{ProjectName}}/internal/{{appName}}/error/httpError"
	_map "{{ProjectName}}/internal/{{appName}}/map"
	"{{ProjectName}}/internal/{{appName}}/server/{{tableName}}"
	R "{{ProjectName}}/pkg/response"
	"{{ProjectName}}/pkg/validator"
	"github.com/gin-gonic/gin"
)

type {{TableName}} struct{}

func ({{TableName}}) GetAll{{TableName}}(ctx *gin.Context) {
	var page = _map.DefaultPageRequest
	if err := ctx.Bind(&page); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(page); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if data, total, err := server.GetAllServer(page); err != nil {
		R.Error(ctx, err.Error(), nil)
	} else {
		R.Ok(ctx, R.MSG_OK, R.Page(total, page.Page, page.PageSize, data))
	}
	return
}

func ({{TableName}}) Post{{TableName}}(ctx *gin.Context) {
	var addMap _map.{{AppName}}AddServer
	if err := ctx.Bind(&addMap); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(addMap); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if err := server.AddServer(addMap); err != nil {
		R.Error(ctx, err.Error(), nil)
		return
	}
	R.Ok(ctx, R.MSG_OK, nil)
	return
}

func ({{TableName}}) Get{{TableName}}(ctx *gin.Context) {
	var id _map.IdMap
	if err := ctx.BindUri(&id); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(id); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if data, err := server.GetByIdServer(id); err != nil {
		R.Error(ctx, err.Error(), nil)
	} else {
		R.Ok(ctx, R.MSG_OK, data)
	}
	return
}

func ({{TableName}}) Put{{TableName}}(ctx *gin.Context) {
	var put _map.{{AppName}}PutServer
	if err := ctx.Bind(&put); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(put); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if err := server.PutByIdServer(put); err != nil {
		R.Error(ctx, err.Error(), nil)
	} else {
		R.Ok(ctx, R.MSG_OK, nil)
	}
	return
}

func ({{TableName}}) Del{{TableName}}(ctx *gin.Context) {
	var del _map.IdMap
	if err := ctx.BindUri(&del); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(del); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if err := server.DelServer(del); err != nil {
		R.Error(ctx, err.Error(), nil)
	} else {
		R.Ok(ctx, R.MSG_OK, nil)
	}
	return
}
