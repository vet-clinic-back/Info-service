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
        "/info/v1/pets": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all pets details",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pets"
                ],
                "summary": "Get all pets",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Pet ID",
                        "name": "pet_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Veterinarian ID",
                        "name": "vet_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Owner ID",
                        "name": "owner_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved pets",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.OutputPetDTO"
                            }
                        }
                    },
                    "404": {
                        "description": "Not found in db",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new pet in the system. Age \u0026 weight should be \u003e 0 \u0026 Gender should be 'Male' or 'Female'",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pets"
                ],
                "summary": "Create Pet",
                "parameters": [
                    {
                        "description": "Pet details",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.createPetDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created pet",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "400": {
                        "description": "Invalid input body",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    }
                }
            }
        },
        "/info/v1/pets/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get pet details by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pets"
                ],
                "summary": "Get Pet",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Pet ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved pet",
                        "schema": {
                            "$ref": "#/definitions/models.Pet"
                        }
                    },
                    "404": {
                        "description": "Pet not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update pet details by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pets"
                ],
                "summary": "Update Pet",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Pet ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Pet details",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Pet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated pet",
                        "schema": {
                            "$ref": "#/definitions/models.Pet"
                        }
                    },
                    "400": {
                        "description": "Invalid input body or pet ID",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "404": {
                        "description": "Pet not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
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
                "description": "Delete pet details by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pets"
                ],
                "summary": "Delete Pet",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Pet ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted pet",
                        "schema": {
                            "$ref": "#/definitions/models.Pet"
                        }
                    },
                    "404": {
                        "description": "Pet not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    }
                }
            }
        },
        "/info/v1/record/entries": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new med entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MedEntry"
                ],
                "summary": "getEntries",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Entry ID",
                        "name": "entry_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Pet ID",
                        "name": "pet_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created утекн",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.MedicalEntry"
                            }
                        }
                    },
                    "400": {
                        "description": "failed to parse filters",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new med entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MedEntry"
                ],
                "summary": "Create med entry",
                "parameters": [
                    {
                        "description": "entry data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MedicalEntry"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created утекн",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "400": {
                        "description": "Invalid input body",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.createPetDTO": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "animal_type": {
                    "type": "string"
                },
                "behavior": {
                    "type": "string"
                },
                "condition": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "integer"
                },
                "research_status": {
                    "type": "string"
                },
                "vet_id": {
                    "type": "integer"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "models.ErrorDTO": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.MedicalEntry": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "device_number": {
                    "type": "integer"
                },
                "disease": {
                    "type": "string"
                },
                "entry_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "medical_record_id": {
                    "type": "integer"
                },
                "recommendation": {
                    "type": "string"
                },
                "vaccinations": {
                    "type": "string"
                },
                "vet_id": {
                    "type": "integer"
                }
            }
        },
        "models.OutputPetDTO": {
            "type": "object",
            "properties": {
                "owner_id": {
                    "type": "integer"
                },
                "pet_info": {
                    "$ref": "#/definitions/models.Pet"
                },
                "vet_id": {
                    "type": "integer"
                }
            }
        },
        "models.Pet": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "animal_type": {
                    "type": "string"
                },
                "behavior": {
                    "type": "string"
                },
                "condition": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "research_status": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Vet clinic auth service",
	Description:      "auth service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
