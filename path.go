package akko

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/nanozuki/akko/typ"
)

type PathBuilder struct {
	path string
	item *openapi3.PathItem
	ms   []Middleware
}

func Path(path string) *PathBuilder {
	return &PathBuilder{
		path: path,
		item: &openapi3.PathItem{},
	}
}

func (b *PathBuilder) Path(path *PathBuilder) *PathBuilder {
	path.AttachToPath(b)
	return b
}

func (b *PathBuilder) Use(middlewares ...Middleware) *PathBuilder {
	b.ms = append(b.ms, middlewares...)
	return b
}

func (b *PathBuilder) Summary(summary string) *PathBuilder {
	b.item.Summary = summary
	return b
}

func (b *PathBuilder) Description(description string) *PathBuilder {
	b.item.Description = description
	return b
}

func (b *PathBuilder) Servers(servers ...*ServerBuilder) *PathBuilder {
	for _, server := range servers {
		b.item.Servers = append(b.item.Servers, server.server)
	}
	return b
}

func buildParameter(pb typ.ParameterPropBuilder) *openapi3.Parameter {
	panic("not implemented")
}

func (b *PathBuilder) Parameters(parameters ...typ.ParameterPropBuilder) *PathBuilder {
	for _, pb := range parameters {
		b.item.Parameters = append(b.item.Parameters, &openapi3.ParameterRef{Value: buildParameter(pb)})
	}
	return b
}

func (b *PathBuilder) CONNECT(op *OperationBuilder) *PathBuilder {
	b.item.Connect = op.operation
	return b
}

func (b *PathBuilder) DELETE(op *OperationBuilder) *PathBuilder {
	b.item.Delete = op.operation
	return b
}

func (b *PathBuilder) GET(op *OperationBuilder) *PathBuilder {
	b.item.Get = op.operation
	return b
}

func (b *PathBuilder) HEAD(op *OperationBuilder) *PathBuilder {
	b.item.Head = op.operation
	return b
}

func (b *PathBuilder) OPTIONS(op *OperationBuilder) *PathBuilder {
	b.item.Options = op.operation
	return b
}

func (b *PathBuilder) PATCH(op *OperationBuilder) *PathBuilder {
	b.item.Patch = op.operation
	return b
}

func (b *PathBuilder) POST(op *OperationBuilder) *PathBuilder {
	b.item.Post = op.operation
	return b
}

func (b *PathBuilder) PUT(op *OperationBuilder) *PathBuilder {
	b.item.Put = op.operation
	return b
}

func (b *PathBuilder) TRACE(op *OperationBuilder) *PathBuilder {
	b.item.Trace = op.operation
	return b
}

func (b *PathBuilder) AttachToPath(path *PathBuilder) *PathBuilder {
	panic("not implemented")
}

func (b *PathBuilder) AttachToAPI(api *OpenAPIBuilder) *PathBuilder {
	panic("not implemented")
}
