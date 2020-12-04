package registry

import "{{ProjectName}}/internal/{{appName}}/api/{{version}}/handle"

func init() {
	{{version}} := Router.Group("/api/{{version}}")
	{
		{{tableName}} := new(handle.{{TableName}})
		{{version}}.GET("/{{tableName}}", {{tableName}}.GetAll{{TableName}})
		{{version}}.GET("/{{tableName}}/:id", {{tableName}}.Get{{TableName}})
		{{version}}.POST("/{{tableName}}", {{tableName}}.Post{{TableName}})
		{{version}}.PUT("/{{tableName}}/:id", {{tableName}}.Put{{TableName}})
		{{version}}.DELETE("/{{tableName}}/:id", {{tableName}}.Del{{TableName}})
	}
}