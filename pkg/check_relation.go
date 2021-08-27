package pkg

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
	"github.com/ory/keto-client-go/client"
	"github.com/ory/keto-client-go/client/read"
	"os"
	"time"
)

var cfg = &client.TransportConfig{
	Host:     "localhost:4466",
	BasePath: "/",
	Schemes:  []string{"http"},
}

func RegisterCheckFuncs() {
	if ketoUrl, exists := os.LookupEnv("KETO_URL"); exists {
		cfg.Host = ketoUrl
	}
	rego.RegisterBuiltin4(
		&rego.Function{
			Name:    "ketoCheck",
			Decl:    types.NewFunction(types.Args(types.S, types.S, types.S, types.S), types.Any{}),
			Memoize: true,
		},
		func(bctx rego.BuiltinContext, a, b, c, d *ast.Term) (*ast.Term, error) {
			var namespace, relation, object, subject string

			if err := ast.As(a.Value, &subject); err != nil {
				return nil, err
			} else if err = ast.As(b.Value, &relation); err != nil {
				return nil, err
			} else if err = ast.As(c.Value, &namespace); err != nil {
				return nil, err
			} else if err = ast.As(d.Value, &object); err != nil {
				return nil, err
			}

			ketoClient := client.NewHTTPClientWithConfig(strfmt.Default, cfg)
			getCheckParams := read.NewGetCheckParamsWithTimeout(time.Second)
			getCheckParams.Subject = &subject
			getCheckParams.Relation = relation
			getCheckParams.Namespace = namespace
			getCheckParams.Object = object
			checkPayload, err := ketoClient.Read.GetCheck(getCheckParams)

			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
			fmt.Println(checkPayload.Payload.Allowed)
			resultTupleTerm := ast.Item(ast.StringTerm("Result"), ast.BooleanTerm(*checkPayload.Payload.Allowed))
			explanationTupleTerm := ast.Item(ast.StringTerm("Explanation"), ast.StringTerm(
				fmt.Sprintf("User %s has a %s relationship with object %s in namespace %s", subject,
					relation, object, namespace)))
			return ast.ObjectTerm(resultTupleTerm, explanationTupleTerm), nil
		},
	)
}
