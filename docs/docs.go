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
            "name": "changyoung"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/ingredients": {
            "get": {
                "description": "List all uploaded ingredients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List all uploaded ingredients",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name search by q",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.IngredientResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Upload single ingredient. The name must be unique",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Upload an ingredient",
                "parameters": [
                    {
                        "type": "file",
                        "description": "image of ingredient",
                        "name": "file",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "json structure of ingredient",
                        "name": "json",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.IngredientResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    }
                }
            }
        },
        "/api/recipecategories": {
            "get": {
                "description": "List all uploaded recpie-categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List all uploaded recipe-categories",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name search by q",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.RecipeCategoryResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Upload single recipe-category. The name must be unique",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Upload recipe-category",
                "parameters": [
                    {
                        "type": "file",
                        "description": "image of recipe-category",
                        "name": "file",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "json structure of recipe-category",
                        "name": "json",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.RecipeCategoryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    }
                }
            }
        },
        "/api/recipes": {
            "get": {
                "description": "List all uploaded recipes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List all uploaded recipes",
                "parameters": [
                    {
                        "enum": [
                            "weekly_views",
                            "created_at"
                        ],
                        "type": "string",
                        "description": "sort by field",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "numbers to fetch",
                        "name": "limits",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.RecipeThumbResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Upload single recipe. The name must be unique",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Upload the recipe",
                "parameters": [
                    {
                        "type": "file",
                        "description": "image of recipe",
                        "name": "file",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "image of step 1",
                        "name": "step_1",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "image of step 2",
                        "name": "step_2",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "image of step 3",
                        "name": "step_3",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "image of step 4",
                        "name": "step_4",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "image of step 5",
                        "name": "step_5",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "json structure of recipe",
                        "name": "json",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.RecipeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    }
                }
            }
        },
        "/api/recipes/challenge": {
            "get": {
                "description": "List all challenge recipes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List all challege recipes",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "numbers to fetch",
                        "name": "limits",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.RecipeThumbResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    }
                }
            }
        },
        "/api/recipes/recommend": {
            "get": {
                "description": "List recommended recipes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List recommended recipes",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "numbers to fetch",
                        "name": "limits",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.RecipeThumbResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    }
                }
            }
        },
        "/api/recipes/trending": {
            "get": {
                "description": "List all trending recipes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List trending recipes",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "numbers to fetch",
                        "name": "limits",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.RecipeThumbResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    }
                }
            }
        },
        "/api/recipes/{recipeID}": {
            "get": {
                "description": "Get the detail of recipe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get the detail of recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "recipeID",
                        "name": "recipeID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.RecipeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.RecipeStep": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "image_path": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "tip": {
                    "type": "string"
                }
            }
        },
        "service.ErrResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "application-specific error code",
                    "type": "integer"
                },
                "error": {
                    "description": "application-level error message, for debugging",
                    "type": "string"
                },
                "status": {
                    "description": "user-level status message",
                    "type": "string"
                }
            }
        },
        "service.IngredientQuantityResponse": {
            "type": "object",
            "properties": {
                "ingredient": {
                    "$ref": "#/definitions/service.IngredientResponse"
                },
                "quantity": {
                    "type": "object"
                }
            }
        },
        "service.IngredientResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "image_path": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "service.RecipeCategoryResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "image_path": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "service.RecipeResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "ease": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_path": {
                    "type": "string"
                },
                "ingredient_quantities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.IngredientQuantityResponse"
                    }
                },
                "is_clipped": {
                    "type": "boolean"
                },
                "preparation_time": {
                    "type": "integer"
                },
                "recipe_category": {
                    "$ref": "#/definitions/service.RecipeCategoryResponse"
                },
                "steps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RecipeStep"
                    }
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "writer": {
                    "$ref": "#/definitions/service.UserResponse"
                }
            }
        },
        "service.RecipeThumbResponse": {
            "type": "object",
            "properties": {
                "ease": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_path": {
                    "type": "string"
                },
                "is_clipped": {
                    "type": "boolean"
                },
                "preparation_time": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "writer": {
                    "$ref": "#/definitions/service.UserResponse"
                }
            }
        },
        "service.UserResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "image_path": {
                    "type": "string"
                },
                "name": {
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
	BasePath:    "",
	Schemes:     []string{},
	Title:       "cooker API",
	Description: "cookerserver",
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
