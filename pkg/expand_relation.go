package pkg

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
	"github.com/ory/keto-client-go/client"
	"github.com/ory/keto-client-go/client/read"
	"time"
)

func RegisterExpandFuncs()  {
	rego.RegisterBuiltin3(
		&rego.Function{
			Name:    "ketoExpand",
			Decl:    types.NewFunction(types.Args(types.S, types.S, types.S), types.Any{}),
			Memoize: true,
		},
		func(bctx rego.BuiltinContext, a, b, c*ast.Term) (*ast.Term, error) {
			var namespace, relation, object string

			if err := ast.As(a.Value, &relation); err != nil {
				return nil, err
			} else if err = ast.As(b.Value, &namespace); err != nil {
				return nil, err
			} else if err = ast.As(c.Value, &object); err != nil {
				return nil, err
			}

			ketoClient := client.NewHTTPClientWithConfig(strfmt.Default, cfg)
			getExpandParams := read.NewGetExpandParamsWithTimeout(time.Second)
			getExpandParams.Relation = relation
			getExpandParams.Namespace = namespace
			getExpandParams.Object = object
			maxDepth := int64(100)
			getExpandParams.MaxDepth = &maxDepth
			expandPayload, err := ketoClient.Read.GetExpand(getExpandParams)
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
			fmt.Println(expandPayload.Payload)
			return ast.ObjectTerm(), nil
		},
	)
}
