// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "An API to automate content creation on Notion.",
    "title": "Notion API",
    "contact": {},
    "version": "0.1"
  },
  "host": "api:3001",
  "basePath": "/v1",
  "paths": {
    "/tasks/campuses/{id}": {
      "post": {
        "description": "Create campus tasks like onboarding steps, etc..",
        "tags": [
          "Task"
        ],
        "summary": "Create campus tasks",
        "operationId": "campusCreate",
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Task"
              }
            }
          },
          "500": {
            "$ref": "#/responses/500"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Campus ID.",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/tasks/pools/{id}": {
      "post": {
        "description": "Create a task list linked to pool setup.",
        "tags": [
          "Task"
        ],
        "summary": "Create pool tasks",
        "operationId": "poolCreate",
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Task"
              }
            }
          },
          "500": {
            "$ref": "#/responses/500"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Pool event ID.",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 500
        },
        "message": {
          "type": "string",
          "example": "Error message"
        }
      }
    },
    "Task": {
      "type": "object",
      "properties": {
        "epicID": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "priority": {
          "type": "string"
        },
        "step": {
          "type": "integer"
        }
      }
    }
  },
  "responses": {
    "400": {
      "description": "Bad request",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "500": {
      "description": "Internal Server Error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "An API to automate content creation on Notion.",
    "title": "Notion API",
    "contact": {},
    "version": "0.1"
  },
  "host": "api:3001",
  "basePath": "/v1",
  "paths": {
    "/tasks/campuses/{id}": {
      "post": {
        "description": "Create campus tasks like onboarding steps, etc..",
        "tags": [
          "Task"
        ],
        "summary": "Create campus tasks",
        "operationId": "campusCreate",
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Task"
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Campus ID.",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/tasks/pools/{id}": {
      "post": {
        "description": "Create a task list linked to pool setup.",
        "tags": [
          "Task"
        ],
        "summary": "Create pool tasks",
        "operationId": "poolCreate",
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Task"
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Pool event ID.",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 500
        },
        "message": {
          "type": "string",
          "example": "Error message"
        }
      }
    },
    "Task": {
      "type": "object",
      "properties": {
        "epicID": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "priority": {
          "type": "string"
        },
        "step": {
          "type": "integer"
        }
      }
    }
  },
  "responses": {
    "400": {
      "description": "Bad request",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "500": {
      "description": "Internal Server Error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  }
}`))
}
