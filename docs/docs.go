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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/docs": {
            "post": {
                "description": "Create Docs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "doc"
                ],
                "summary": "Create Docs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Batch Delete Docs By Ids",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "doc"
                ],
                "summary": "Batch Delete Docs By Ids",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/docs/{id}": {
            "get": {
                "description": "Get Doc By Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "doc"
                ],
                "summary": "Get Doc By Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "doc id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete One Doc By Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "doc"
                ],
                "summary": "Delete One Doc By Id",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/documents/{document_id}": {
            "get": {
                "description": "Get key",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "document"
                ],
                "summary": "Get key",
                "parameters": [
                    {
                        "type": "string",
                        "description": "index name",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "put": {
                "description": "Update key,Use \"HSET\"",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "document"
                ],
                "summary": "Update key,Use \"HSET\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "index name",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "Delete key,Use \"Del\"",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "document"
                ],
                "summary": "Delete key,Use \"Del\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "index name",
                        "name": "document_id",
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
        },
        "/indexes": {
            "get": {
                "description": "List all indexes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "List all indexes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/indexes/{index}": {
            "get": {
                "description": "Get Index Info",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Get Index Info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "index name",
                        "name": "index",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create an index. ` + "`" + `schema.fields.type` + "`" + `: ` + "`" + `0` + "`" + `-` + "`" + `Text` + "`" + `;` + "`" + `1` + "`" + `-` + "`" + `Numeric` + "`" + `;` + "`" + `2` + "`" + `-` + "`" + `Geo` + "`" + `;` + "`" + `3` + "`" + `-` + "`" + `Tag` + "`" + `",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Create an index",
                "parameters": [
                    {
                        "type": "string",
                        "description": "index name",
                        "name": "index",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "Delete an index",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Delete an index",
                "parameters": [
                    {
                        "type": "string",
                        "description": "index name",
                        "name": "index",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "delete document",
                        "name": "deldocs",
                        "in": "query"
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/indexes/{index}/search": {
            "get": {
                "description": "Searches the index with a textual query, returning either documents or just count(when num=0 and offset=0).",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Search in an index with GET",
                "parameters": [
                    {
                        "type": "string",
                        "description": "index name",
                        "name": "index",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " the text query to search",
                        "name": "raw",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maximum": 1000000,
                        "minimum": 0,
                        "type": "integer",
                        "default": 10,
                        "description": "maximum number of documents returned. default is ` + "`" + `10` + "`" + `;max is ` + "`" + `1_000_000` + "`" + `. when num is ` + "`" + `0` + "`" + `, just return the count",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "default": 0,
                        "description": "number of documents to skip，default is ` + "`" + `0` + "`" + `",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "If set, we limit the result to a given set of keys specified in the list. ",
                        "name": "in_keys",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": " If set, filter the results to ones appearing only in specific fields of the document, like title or URL. num is the number of specified field arguments ",
                        "name": "in_fields",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Use this keyword to limit which fields from the document are returned.e.g: ` + "`" + `return_fields=id,name,age` + "`" + ` ",
                        "name": "return_fields",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "If specified, the results are ordered by the value of this field. This applies to both text and numeric fields. e.g: ` + "`" + `sort_by=name|asc` + "`" + `",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "If set, we use a stemmer for the supplied language during search for query expansion",
                        "name": "language",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/http.Response"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Searches the index with a textual query, returning either documents or just count(when num=0 and offset=0).",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Search in an index with POST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "index name",
                        "name": "index",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/http.Response"
                            }
                        }
                    }
                }
            }
        },
        "/search/{index}": {
            "get": {
                "description": "Searches the index with a textual query, returning either documents or just count(when num=0 and offset=0).",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Search in an index with GET",
                "parameters": [
                    {
                        "type": "string",
                        "description": "index name",
                        "name": "index",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " the text query to search",
                        "name": "raw",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maximum": 1000000,
                        "minimum": 0,
                        "type": "integer",
                        "default": 10,
                        "description": "maximum number of documents returned. default is ` + "`" + `10` + "`" + `;max is ` + "`" + `1_000_000` + "`" + `. when num is ` + "`" + `0` + "`" + `, just return the count",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "default": 0,
                        "description": "number of documents to skip，default is ` + "`" + `0` + "`" + `",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "If set, we limit the result to a given set of keys specified in the list. ",
                        "name": "in_keys",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": " If set, filter the results to ones appearing only in specific fields of the document, like title or URL. num is the number of specified field arguments ",
                        "name": "in_fields",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Use this keyword to limit which fields from the document are returned.e.g: ` + "`" + `return_fields=id,name,age` + "`" + ` ",
                        "name": "return_fields",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "If specified, the results are ordered by the value of this field. This applies to both text and numeric fields. e.g: ` + "`" + `sort_by=name|asc` + "`" + `",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "If set, we use a stemmer for the supplied language during search for query expansion",
                        "name": "language",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/http.Response"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Searches the index with a textual query, returning either documents or just count(when num=0 and offset=0).",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Search in an index with POST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "index name",
                        "name": "index",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/http.Response"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.HttpError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "http.Response": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "调用失败时，返回的出错信息",
                    "$ref": "#/definitions/http.HttpError"
                },
                "inventory": {
                    "description": "调用成功时，返回的数据清单",
                    "type": "object"
                },
                "success": {
                    "description": "是否调用成功：true表示调用成功，false表示调用失败",
                    "type": "boolean"
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
