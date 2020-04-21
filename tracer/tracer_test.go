package tracer

import (
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/aereal/gqlgen-tracer-opencensus/testdata"
	"github.com/aereal/gqlgen-tracer-opencensus/testdata/generated"
)

func TestCompile(t *testing.T) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &testdata.Resolver{}}))
	srv.Use(Tracer{})
}
