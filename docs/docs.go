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
        "/v1/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth Module"
                ],
                "summary": "User Registration/Login",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/aphrodite-go_api_v1.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/aphrodite-go_api_v1.Response"
                        }
                    }
                }
            }
        },
        "/v1/logout": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth Module"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/aphrodite-go_api_v1.Response"
                        }
                    }
                }
            }
        },
        "/v1/send-code": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth Module"
                ],
                "summary": "Send Verification Vode",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/aphrodite-go_api_v1.SendVerifyCodeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/aphrodite-go_api_v1.Response"
                        }
                    }
                }
            }
        },
        "/v1/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Module"
                ],
                "summary": "Get User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/aphrodite-go_api_v1.GetUserResp"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Module"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/aphrodite-go_api_v1.UpdateUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/aphrodite-go_api_v1.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Module"
                ],
                "summary": "Delete User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/aphrodite-go_api_v1.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "aphrodite-go_api_v1.GetUserResp": {
            "description": "Complete response for user information",
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "description": "User information data",
                    "allOf": [
                        {
                            "$ref": "#/definitions/aphrodite-go_api_v1.GetUserRespData"
                        }
                    ]
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "aphrodite-go_api_v1.GetUserRespData": {
            "description": "Response data for user information",
            "type": "object",
            "required": [
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john@example.com"
                },
                "nickname": {
                    "type": "string",
                    "example": "john"
                },
                "phone": {
                    "type": "string",
                    "example": "13800138000"
                },
                "userCode": {
                    "type": "string"
                },
                "userNo": {
                    "type": "string"
                }
            }
        },
        "aphrodite-go_api_v1.LoginReq": {
            "description": "User login request data",
            "type": "object",
            "required": [
                "phone"
            ],
            "properties": {
                "phone": {
                    "type": "string",
                    "example": "13800138000"
                },
                "verifyCode": {
                    "type": "string",
                    "example": "1234"
                }
            }
        },
        "aphrodite-go_api_v1.Response": {
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
        "aphrodite-go_api_v1.SendVerifyCodeReq": {
            "description": "Request data for sending verification code",
            "type": "object",
            "required": [
                "phone"
            ],
            "properties": {
                "phone": {
                    "type": "string",
                    "example": "13800138000"
                }
            }
        },
        "aphrodite-go_api_v1.UpdateUserReq": {
            "description": "Request data for updating user information",
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john@example.com"
                },
                "nickname": {
                    "type": "string",
                    "example": "john"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Aphrodite API",
	Description:      "API Description",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
