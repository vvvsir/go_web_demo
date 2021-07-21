// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/adminUser": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户"
                ],
                "summary": "获取用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "姓名|手机",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示多少条",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AdminUserPage"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户"
                ],
                "summary": "添加用户",
                "parameters": [
                    {
                        "description": "用户",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserAdd"
                        }
                    }
                ]
            }
        },
        "/adminUser/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户"
                ],
                "summary": "获取用户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AdminUserList"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户"
                ],
                "summary": "编辑用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "用户",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserEdit"
                        }
                    }
                ]
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/auth/role": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户"
                ],
                "summary": "获取角色树",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AdminUserPage"
                        }
                    }
                }
            }
        },
        "/auth/tree": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "获取权限树",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RolePage"
                        }
                    }
                }
            }
        },
        "/captcha": {
            "get": {
                "tags": [
                    "登录操作"
                ],
                "summary": "验证码",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录操作"
                ],
                "summary": "登录操作",
                "parameters": [
                    {
                        "description": "用户",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/role": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "获取所有角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示多少条",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RolePage"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "添加角色",
                "parameters": [
                    {
                        "description": "角色",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RoleAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/role/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "获取角色详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RoleList"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "编辑角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "角色",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RoleEdit"
                        }
                    }
                ]
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "删除角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "request.LoginUser": {
            "type": "object",
            "required": [
                "password",
                "username",
                "verify_code"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "admin"
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "admin"
                },
                "verify_code": {
                    "description": "验证码",
                    "type": "string",
                    "example": "9527"
                }
            }
        },
        "request.RoleAdd": {
            "type": "object",
            "required": [
                "auth",
                "name"
            ],
            "properties": {
                "auth": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2,
                        3,
                        4,
                        5,
                        6,
                        7,
                        8,
                        9,
                        10,
                        11,
                        12,
                        13,
                        14,
                        15,
                        16,
                        17,
                        18,
                        19
                    ]
                },
                "name": {
                    "description": "角色名",
                    "type": "string",
                    "example": "程序员"
                },
                "pid": {
                    "description": "上级ID",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "request.RoleEdit": {
            "type": "object",
            "required": [
                "auth",
                "name"
            ],
            "properties": {
                "auth": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2,
                        3,
                        4,
                        5,
                        6,
                        7,
                        8,
                        9,
                        10,
                        11,
                        12,
                        13,
                        14,
                        15,
                        16,
                        17,
                        18,
                        19
                    ]
                },
                "name": {
                    "description": "角色名",
                    "type": "string",
                    "example": "程序员"
                },
                "pid": {
                    "description": "上级ID",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "request.UserAdd": {
            "type": "object",
            "required": [
                "password",
                "real_name",
                "roles",
                "tel",
                "user_name"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "test"
                },
                "real_name": {
                    "description": "真实姓名",
                    "type": "string",
                    "example": "test"
                },
                "roles": {
                    "description": "所属角色",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2
                    ]
                },
                "status": {
                    "description": "状态",
                    "type": "integer",
                    "example": 1
                },
                "tel": {
                    "description": "电话号码",
                    "type": "string",
                    "example": "13054174174"
                },
                "user_name": {
                    "description": "账号",
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "request.UserEdit": {
            "type": "object",
            "required": [
                "real_name",
                "roles",
                "tel",
                "user_name"
            ],
            "properties": {
                "password": {
                    "description": "密码（非必填）",
                    "type": "string",
                    "example": "test"
                },
                "real_name": {
                    "description": "真实姓名",
                    "type": "string",
                    "example": "test"
                },
                "roles": {
                    "description": "所属角色",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2
                    ]
                },
                "status": {
                    "description": "状态",
                    "type": "integer",
                    "example": 1
                },
                "tel": {
                    "description": "电话号码",
                    "type": "string",
                    "example": "13054174174"
                },
                "user_name": {
                    "description": "账号",
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "response.AdminUserList": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "real_name": {
                    "description": "真实姓名",
                    "type": "string"
                },
                "roles": {
                    "description": "角色信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.CasRole"
                    }
                },
                "status": {
                    "description": "用户状态",
                    "type": "integer"
                },
                "tel": {
                    "description": "手机号码",
                    "type": "string"
                },
                "user_name": {
                    "description": "登录名",
                    "type": "string"
                }
            }
        },
        "response.AdminUserPage": {
            "type": "object",
            "properties": {
                "current_page": {
                    "description": "每页显示多少条",
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.AdminUserList"
                    }
                },
                "per_page": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "total": {
                    "description": "总共多少页",
                    "type": "integer"
                }
            }
        },
        "response.CasRole": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "角色ID",
                    "type": "integer"
                },
                "name": {
                    "description": "角色名",
                    "type": "string"
                }
            }
        },
        "response.RoleList": {
            "type": "object",
            "properties": {
                "auth": {
                    "description": "所有权限arr",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "base_auth": {
                    "description": "api权限",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "角色名",
                    "type": "string"
                },
                "parent_name": {
                    "description": "上级角色名",
                    "type": "string"
                },
                "pid": {
                    "description": "上级ID",
                    "type": "integer"
                }
            }
        },
        "response.RolePage": {
            "type": "object",
            "properties": {
                "current_page": {
                    "description": "每页显示多少条",
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.RoleList"
                    }
                },
                "per_page": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "total": {
                    "description": "总共多少页",
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "go_web_demo",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
