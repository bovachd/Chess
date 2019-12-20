// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-20 16:19:05.0367574 +0700 +07 m=+5.960030401

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is list api for chess project",
        "title": "Swagger Chess project API",
        "contact": {},
        "license": {},
        "version": "2.0"
    },
    "host": "chess-apis.herokuapp.com",
    "basePath": "/api/v1/be",
    "paths": {
        "/access/login": {
            "post": {
                "description": "login with username, password. return token string jwt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "access-apis"
                ],
                "summary": "Login system",
                "parameters": [
                    {
                        "description": "username and password",
                        "name": "accountInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.LoginResponse"
                        }
                    }
                }
            }
        },
        "/access/login/token": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "login with token string. return new token string jwt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "access-apis"
                ],
                "summary": "Login system by token string",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.LoginResponse"
                        }
                    }
                }
            }
        },
        "/account/accounts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "filter list user and paging filtered",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account-manager-apis"
                ],
                "summary": "Filter users and paging",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username for user",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "rank of user",
                        "name": "rank",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "nickname of user",
                        "name": "nickname",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.MetaDataResponse"
                        }
                    }
                }
            }
        },
        "/account/accounts/top10": {
            "get": {
                "description": "Get top 10",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account-manager-apis"
                ],
                "summary": "Get top 10",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/account/create": {
            "post": {
                "description": "Create new account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account-manager-apis"
                ],
                "summary": "Create new Account",
                "parameters": [
                    {
                        "description": "User information",
                        "name": "UserInformation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.UserM"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/account/password": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update password for exists user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account-manager-apis"
                ],
                "summary": "Update password",
                "parameters": [
                    {
                        "description": "cuple value id and reset password",
                        "name": "Update_Password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.UpdatePassPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/account/{accountId}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update user by field:name, avatar, status, role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account-manager-apis"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of user account",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "include old user and new update user",
                        "name": "update_model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.UpdateAccountPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/account/{userId}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "find user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account-manager-apis"
                ],
                "summary": "Get user by user id",
                "parameters": [
                    {
                        "type": "string",
                        "default": "1",
                        "description": "user id is wanted find",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Soft Delete user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account-manager-apis"
                ],
                "summary": "Remove User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id is wanted remove",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/friend/friends/all/{userId}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all friends",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friend-manager-apis"
                ],
                "summary": "get all friends",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of user account",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/friend/friends/new": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new Friend",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friend-manager-apis"
                ],
                "summary": "Create new Friend",
                "parameters": [
                    {
                        "description": "Friend information",
                        "name": "FriendInformation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.FriendPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/friend/friends/{userId}/{friendId}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete friend by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friend-manager-apis"
                ],
                "summary": "Delete friend by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of user account",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id of friend account",
                        "name": "friendId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/report/reports": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Send report",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "report-manager-apis"
                ],
                "summary": "Send report",
                "parameters": [
                    {
                        "description": "Report information",
                        "name": "ReportInformation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.ReportPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/report/reports/all": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all report",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "report-manager-apis"
                ],
                "summary": "Get all report",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/report/reports/filter/{reporterId}/{reportedAccountId}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Filter report by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "report-manager-apis"
                ],
                "summary": "Filter report by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of report",
                        "name": "reporterId",
                        "in": "path"
                    },
                    {
                        "type": "integer",
                        "description": "id of report",
                        "name": "reportedAccountId",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/report/reports/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete report by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "report-manager-apis"
                ],
                "summary": "Delete report by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of report",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/room/rooms": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room-manager-apis"
                ],
                "summary": "Create room",
                "parameters": [
                    {
                        "description": "Room information",
                        "name": "RoomInformation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.RoomPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/room/rooms/all": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room-manager-apis"
                ],
                "summary": "Get all room",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/room/rooms/{roomId}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete room by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room-manager-apis"
                ],
                "summary": "Delete room by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of room",
                        "name": "roomId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.FriendPayload": {
            "type": "object",
            "properties": {
                "friendId": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "controllers.LoginRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controllers.LoginResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "controllers.ReportPayload": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "reportedAccountId": {
                    "type": "integer"
                },
                "reporterId": {
                    "type": "integer"
                }
            }
        },
        "controllers.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controllers.RoomPayload": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.UpdateAccountPayload": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                },
                "point": {
                    "type": "integer"
                },
                "rank": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "controllers.UpdatePassPayload": {
            "type": "object",
            "properties": {
                "newPass": {
                    "type": "string"
                },
                "oldPass": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "controllers.UserM": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.MetaData": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.MetaDataResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "metaData": {
                    "type": "object",
                    "$ref": "#/definitions/model.MetaData"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
