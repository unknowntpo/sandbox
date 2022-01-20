package main

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
)

// NewOpenAPI3 instantiates the OpenAPI specification for this service.
// Ref: https://github.com/MarioCarrion/todo-api-microservice-example/blob/074bbb9f4d0f79e5bced943c10c56013705969a9/internal/rest/open_api.go
func NewOpenAPI3() openapi3.Swagger {
	swagger := openapi3.Swagger{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "TODOs API",
			Description: "This is the api documentation of TODOS service",
			Version:     "0.1.0",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				URL: "https://github.com/unknowntpo/todos",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://localhost:4000",
			},
			&openapi3.Server{
				Description: "Production server",
				URL:         "https:todos.unknowntpo.net",
			},
		},
	}
	// TODO: add more things
	// Schemas
	swagger.Components.Schemas = openapi3.Schemas{
		/*
			"Priority": openapi3.NewSchemaRef("",
				openapi3.NewStringSchema().
					WithEnum("none", "low", "medium", "high").
					WithDefault("none")),
			"Dates": openapi3.NewSchemaRef("",
				openapi3.NewObjectSchema().
					WithProperty("start", openapi3.NewStringSchema().
						WithFormat("date-time").
						WithNullable()).
					WithProperty("due", openapi3.NewStringSchema().
						WithFormat("date-time").
						WithNullable())),
			"Task": openapi3.NewSchemaRef("",
				openapi3.NewObjectSchema().
					WithProperty("id", openapi3.NewUUIDSchema()).
					WithProperty("description", openapi3.NewStringSchema()).
					WithProperty("is_done", openapi3.NewBoolSchema()).
					WithPropertyRef("priority", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Priority",
					}).
					WithPropertyRef("dates", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Dates",
					})),
		*/
		"api.AuthenticationRequestBody": openapi3.NewSchemaRef(
			"",
			openapi3.NewObjectSchema().
				WithProperty("email", openapi3.NewStringSchema()).
				WithProperty("password", openapi3.NewStringSchema())),
		"api.AuthenticationResponse": openapi3.NewSchemaRef(
			"",
			openapi3.NewObjectSchema().
				WithProperty("token", &openapi.SchemaRef{
					Ref: "#/components/schemas/domain.Token",
				})),
	}
}