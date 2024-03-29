// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "miomiora",
            "url": "https://github.com/miomiora"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/post/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章相关接口"
                ],
                "summary": "通过获取文章视图列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生，需为管理员",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "新增的文章参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PostDTOAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/post/delete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章相关接口"
                ],
                "summary": "通过postId删除文章",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生，需为管理员",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "需要删除的postId",
                        "name": "postId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/post/get": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章相关接口"
                ],
                "summary": "管理员通过postId获取文章全部数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生，需为管理员",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "需要查询的postId",
                        "name": "postId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/post/get/vo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章相关接口"
                ],
                "summary": "通过postId获取文章视图",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "需要查找的文章id",
                        "name": "postId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/post/list/page": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章相关接口"
                ],
                "summary": "管理员获取全部文章详细信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生，需为管理员",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "分页查询所需要的参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/post/list/page/vo": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章相关接口"
                ],
                "summary": "通过获取文章视图列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "分页查询需要的参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/post/my": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章相关接口"
                ],
                "summary": "通过当前登录的用户所写的文章",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "分页查询需要的参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/post/new": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章相关接口"
                ],
                "summary": "新建文章",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "新建文章参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PostDTOInsert"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/post/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章相关接口"
                ],
                "summary": "管理员编辑文章",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生，需为管理员",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "需要更新的文章信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PostDTOUpdateByAdmin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/post/update/my": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章相关接口"
                ],
                "summary": "用户更新自己写的文章",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "修改后的数据",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PostDTOUpdateBySelf"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "管理员添加用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生，需为管理员",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "新用户的数据",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserDTOAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "管理员根据userId删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生，需为管理员",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "需要删除的userId",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/get": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "管理员根据userId查询用户完整信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生，需为管理员",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "需要查询的userId",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/get/login": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "获取当前登录用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/get/vo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "根据userId查找用户视图",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "需要查找的用户id",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/list/page": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "管理员根据查询用户完整信息列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生，需为管理员",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "分页查询所需要的参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/list/page/vo": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "获取用户视图列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "分页查询需要的参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "登录参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserDTOLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户登出",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "注册参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserDTORegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "管理员根据userId更新用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生，需为管理员",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "需要更新的用户信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserDTOUpdateByAdmin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/user/update/my": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "当前用户更新自己的信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌 Token 登录后产生",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "修改后的数据",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserDTOUpdateBySelf"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.ListParams": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "model.PostDTOAdd": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.PostDTOInsert": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.PostDTOUpdateByAdmin": {
            "type": "object",
            "required": [
                "content",
                "postId",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "postId": {
                    "type": "string",
                    "example": "0"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.PostDTOUpdateBySelf": {
            "type": "object",
            "required": [
                "content",
                "postId",
                "title",
                "userId"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "postId": {
                    "type": "string",
                    "example": "0"
                },
                "title": {
                    "type": "string"
                },
                "userId": {
                    "type": "string",
                    "example": "0"
                }
            }
        },
        "model.UserDTOAdd": {
            "type": "object",
            "required": [
                "account",
                "password",
                "rePassword"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "rePassword": {
                    "type": "string"
                }
            }
        },
        "model.UserDTOLogin": {
            "type": "object",
            "required": [
                "account",
                "password"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UserDTORegister": {
            "type": "object",
            "required": [
                "account",
                "password",
                "rePassword"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "rePassword": {
                    "type": "string"
                }
            }
        },
        "model.UserDTOUpdateByAdmin": {
            "type": "object",
            "required": [
                "account",
                "userId"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "rePassword": {
                    "type": "string"
                },
                "userId": {
                    "type": "string",
                    "example": "0"
                },
                "userRole": {
                    "type": "integer"
                }
            }
        },
        "model.UserDTOUpdateBySelf": {
            "type": "object",
            "required": [
                "account",
                "userId"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "rePassword": {
                    "type": "string"
                },
                "userId": {
                    "type": "string",
                    "example": "0"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8081",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "mio-init",
	Description:      "Go Web 开发脚手架",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
