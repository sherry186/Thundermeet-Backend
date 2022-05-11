// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Wu, Chien Yin and Yeh, Hsiao Li"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/users": {
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "The body to create a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Update"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetUserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "The body to create a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Register"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users/login/": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "The body to login a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users/resetPassword": {
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "The body to create a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ForgotInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "GetUserResponse": {
            "type": "object",
            "required": [
                "password_answer",
                "status",
                "user_id",
                "username"
            ],
            "properties": {
                "password_answer": {
                    "type": "string",
                    "example": "NTU"
                },
                "status": {
                    "type": "integer",
                    "example": 0
                },
                "user_id": {
                    "type": "string",
                    "example": "christine891225"
                },
                "username": {
                    "type": "string",
                    "example": "Christine Wang"
                }
            }
        },
        "Login": {
            "type": "object",
            "required": [
                "password",
                "userId"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "userId": {
                    "type": "string",
                    "example": "christine891225"
                }
            }
        },
        "Register": {
            "type": "object",
            "required": [
                "password",
                "passwordAnswer",
                "userId"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "passwordAnswer": {
                    "type": "string",
                    "example": "NTU"
                },
                "userId": {
                    "type": "string",
                    "example": "christine891225"
                },
                "userName": {
                    "type": "string",
                    "example": "Christine Wang"
                }
            }
        },
        "Update": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "passwordAnswer": {
                    "type": "string",
                    "example": "NTU"
                },
                "userName": {
                    "type": "string",
                    "example": "Christine Wang"
                }
            }
        },
        "controller.ForgotInfo": {
            "type": "object",
            "required": [
                "password",
                "passwordAnswer",
                "userId"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "passwordAnswer": {
                    "type": "string",
                    "example": "NTU"
                },
                "userId": {
                    "type": "string",
                    "example": "christine891225"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080/",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "ThunderMeet APIs",
	Description:      "This is the backend server for the Thundermeet App.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
