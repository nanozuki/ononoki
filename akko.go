package akko

import "github.com/getkin/kin-openapi/openapi3"

type Akko struct {
	Version     string
	Description string
	Author      Author
	ToS         string
	License     string
	LicenseURL  string

	// services []*Services
	api *openapi3.T
	ms  []Middleware
}

type Author struct {
	Name  string
	URL   string
	Email string
}

type Services struct {
	Name        string
	Description string
	Host        string
	BaseURL     string
}

func New(title string) *Akko {
	return &Akko{
		api: &openapi3.T{
			OpenAPI: "v3.0.3",
			Info: &openapi3.Info{
				Title:   title,
				Version: "v1.0.0",
			},
		},
	}
}

func (a *Akko) Path(path string) *Path {
	return &Path{
		url: path,
	}
}

func (a *Akko) Use(middlewares ...Middleware) {
	a.ms = middlewares
}
