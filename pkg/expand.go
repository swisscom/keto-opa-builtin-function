package opa_keto

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

/*
	RegisterExpand registers the ketoExpand(relation, namespace, object) function
*/
func RegisterExpand() {
	rego.RegisterBuiltin3(
		&rego.Function{
			Name:    KetoExpand,
			Decl:    types.NewFunction(types.Args(types.S, types.S, types.S), types.Any{}),
			Memoize: true,
		},
		func(bctx rego.BuiltinContext, a, b, c *ast.Term) (*ast.Term, error) {
			var namespace, relation, object string
			var err error

			if err := ast.As(a.Value, &relation); err != nil {
				return nil, err
			}
			if err = ast.As(b.Value, &namespace); err != nil {
				return nil, err
			}
			if err = ast.As(c.Value, &object); err != nil {
				return nil, err
			}

			ketoClient := client.NewHTTPClientWithConfig(strfmt.Default, cfg)
			getExpandParams := read.NewGetExpandParamsWithTimeout(time.Second)
			getExpandParams.Relation = relation
			getExpandParams.Namespace = namespace
			getExpandParams.Object = object
			getExpandParams.MaxDepth = 100
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
