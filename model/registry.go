package registry

import "{{ProjectName}}/internal/{{AppName}}/api/v1/handle"

func init() {
	v1 := Router.Group("/api/v1")
	{
		{{tableName}} := new(handle.{{TableName}})
		v1.GET("/{{tableName}}", {{tableName}}.GetAll)
		v1.GET("/{{tableName}}/:id", {{tableName}}.Get)
		v1.POST("/{{tableName}}", {{tableName}}.Post)
		v1.PUT("/{{tableName}}/:id", {{tableName}}.Put)
		v1.DELETE("/{{tableName}}/:id", {{tableName}}.Del)
	}
}