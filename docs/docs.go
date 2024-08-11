// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "登录",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "登录"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "account",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取信息成功"
                    },
                    "400": {
                        "description": "获取信息失败"
                    }
                }
            }
        },
        "/publish": {
            "post": {
                "description": "发布",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "发布"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "文章内容",
                        "name": "artic",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取信息成功"
                    },
                    "400": {
                        "description": "获取信息失败"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
