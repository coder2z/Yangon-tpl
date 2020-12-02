package registry

import "{{ProjectName}}/internal/{{appName}}/api/v1/handle"

func init() {
	v1 := Router.Group("/api/v1")
	{
		{{tableName}} := new(handle.{{TableName}})
		v1.GET("/{{tableName}}", {{tableName}}.GetAll{{TableName}})
		v1.GET("/{{tableName}}/:id", {{tableName}}.Get{{TableName}})
		v1.POST("/{{tableName}}", {{tableName}}.Post{{TableName}})
		v1.PUT("/{{tableName}}/:id", {{tableName}}.Put{{TableName}})
		v1.DELETE("/{{tableName}}/:id", {{tableName}}.Del{{TableName}})
	}
}