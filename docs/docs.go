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
        "/BlockChain": {
            "post": {
                "description": "用于同步本地区块链数据",
                "tags": [
                    "服务端"
                ],
                "summary": "当前区块链数据",
                "responses": {
                    "200": {
                        "description": "statuc\":\"ok\", \"data\":\"bytesdata\"}",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/balance": {
            "get": {
                "description": "返回指定用户的余额信息",
                "tags": [
                    "前端"
                ],
                "summary": "余额",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ivan",
                        "name": "address",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/blockhttp.Response"
                        }
                    }
                }
            }
        },
        "/balancedetailed": {
            "get": {
                "description": "返回指定地址的交易明细",
                "tags": [
                    "前端"
                ],
                "summary": "余额明细",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ivan",
                        "name": "address",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/entry": {
            "post": {
                "description": "生产茶叶数据， 用于交易",
                "tags": [
                    "前端"
                ],
                "summary": "数据录入",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ivan",
                        "name": "address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "5000",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "json",
                        "name": "data",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "statuc\":\"ok\", \"data\":\"\"}",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "statuc\":\"error\", \"message\":\"失败原因\"}",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/registerinfo": {
            "get": {
                "description": "返回指当前所有的注册信息",
                "tags": [
                    "服务端"
                ],
                "summary": "返回注册信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/teadata": {
            "post": {
                "description": "获取指定地址的茶叶数据",
                "tags": [
                    "前端"
                ],
                "summary": "茶叶数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ASHASDSABDKJQWFKJBASFKAF",
                        "name": "address",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "statuc\":\"ok\", \"data\":\"\"}",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "statuc\":\"error\", \"message\":\"失败原因\"}",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/transaction": {
            "post": {
                "description": "用于两个不同地址之间的数据交易",
                "tags": [
                    "前端"
                ],
                "summary": "茶叶交易",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ivan",
                        "name": "from",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ble",
                        "name": "to",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "300",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "statuc\":\"ok\"}",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "statuc\":\"error\", \"data\":\"失败原因\"}",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "返回当前区块链长度",
                "tags": [
                    "服务端"
                ],
                "summary": "区块链版本",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "blockhttp.Response": {
            "type": "object"
        },
        "gin.H": {
            "type": "object",
            "additionalProperties": true
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
	Title:       "",
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
