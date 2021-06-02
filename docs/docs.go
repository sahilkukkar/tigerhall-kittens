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
        "contact": {
            "name": "API Support",
            "email": "sahil.kukkar99@gmail.com"
        },
        "license": {
            "name": "SAHIL"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/createtiger": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Createtiger create tiger in the wild",
                "parameters": [
                    {
                        "description": "create tiger in wild",
                        "name": "requesttiger",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ReqCreateTiger"
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
        "/v1/listtigers": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "ListTigers list tigers with provided query feature",
                "parameters": [
                    {
                        "type": "string",
                        "description": "q",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResListTigers"
                        }
                    }
                }
            }
        },
        "/v1/listtigersights": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "ListTigerSight list tigers previous sights",
                "parameters": [
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResListTigers"
                        }
                    }
                }
            }
        },
        "/v1/sighttiger": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Createtiger create tiger in the wild",
                "parameters": [
                    {
                        "description": "create sight of tiger in wild",
                        "name": "requesttigersight",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ReqSightATiger"
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
        "/v1/uploadimage": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "upload image add images in variations depending on requirement",
                "parameters": [
                    {
                        "description": "upload a image",
                        "name": "uploadimage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ReqUploadImage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.ResUploadImage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Coordinates": {
            "type": "object",
            "properties": {
                "lat": {
                    "type": "number"
                },
                "long": {
                    "type": "number"
                }
            }
        },
        "controllers.ImageInfo": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "format": {
                    "type": "string"
                }
            }
        },
        "controllers.ReqCreateTiger": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "$ref": "#/definitions/controllers.Coordinates"
                },
                "dob": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.ReqSightATiger": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "$ref": "#/definitions/controllers.Coordinates"
                },
                "imagePath": {
                    "type": "string"
                },
                "seenAt": {
                    "type": "string"
                },
                "tigerID": {
                    "type": "string"
                }
            }
        },
        "controllers.ReqUploadImage": {
            "type": "object",
            "properties": {
                "image": {
                    "$ref": "#/definitions/controllers.ImageInfo"
                },
                "name": {
                    "type": "string"
                },
                "resource": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "controllers.ResListTigers": {
            "type": "object",
            "properties": {
                "tigers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.TigerData"
                    }
                },
                "totalCount": {
                    "type": "integer"
                }
            }
        },
        "controllers.ResUploadImage": {
            "type": "object",
            "properties": {
                "imagePath": {
                    "type": "string"
                }
            }
        },
        "controllers.TigerData": {
            "type": "object",
            "properties": {
                "dob": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "seenAt": {
                    "type": "string"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "/tigerhall/",
	Schemes:     []string{},
	Title:       "Merchant Swagger API",
	Description: "Swagger API for tigerhall Project.",
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