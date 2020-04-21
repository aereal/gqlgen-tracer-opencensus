# gqlgen-tracer-opencensus

OpenCensus Tracer for [gqlgen][]

## Synopsis

```go
package main

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/aereal/gqlgen-tracer-opencensus/tracer"
)

func main() {
	srv := handler.NewDefaultServer(...)
	srv.Use(tracer.Tracer{})
}
```

## See also

- [99designs/gqlgen-contrib][gqlgen-contrib]

[gqlgen]: https://gqlgen.com/
[gqlgen-contrib]: https://github.com/99designs/gqlgen-contrib
