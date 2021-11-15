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
	RegisterCheck registers the ketoCheck(subject, relation, namespace, object) function
	that returns:
	{
		"result": bool,
		"explanation": "string",
	}
*/

func RegisterCheck() {
	rego.RegisterBuiltin4(
		&rego.Function{
			Name:    KetoCheck,
			Decl:    types.NewFunction(types.Args(types.S, types.S, types.S, types.S), types.Any{}),
			Memoize: true,
		},
		func(bctx rego.BuiltinContext, a, b, c, d *ast.Term) (*ast.Term, error) {
			var namespace, relation, object, subject string
			var err error

			if err = ast.As(a.Value, &subject); err != nil {
				return nil, err
			}
			if err = ast.As(b.Value, &relation); err != nil {
				return nil, err
			}
			if err = ast.As(c.Value, &namespace); err != nil {
				return nil, err
			}
			if err = ast.As(d.Value, &object); err != nil {
				return nil, err
			}

			ketoClient := client.NewHTTPClientWithConfig(strfmt.Default, cfg)
			getCheckParams := read.NewGetCheckParamsWithTimeout(time.Second)
			getCheckParams.SubjectID = &subject
			getCheckParams.Relation = relation
			getCheckParams.Namespace = namespace
			getCheckParams.Object = object
			checkPayload, err := ketoClient.Read.GetCheck(getCheckParams)

			if err != nil {
				switch err.(type) {
				case *read.GetCheckForbidden:
					resultTupleTerm := ast.Item(ast.StringTerm("result"), ast.BooleanTerm(false))
					explanationTupleTerm := ast.Item(ast.StringTerm("explanation"), ast.StringTerm(
						fmt.Sprintf("Subject %s doesn't have %s relationship with object %s in namespace %s",
							subject, relation, object, namespace)))
					return ast.ObjectTerm(resultTupleTerm, explanationTupleTerm), nil
				default:
					fmt.Printf("getcheck failed: %v", err)
					return nil, err
				}
			}

			resultTupleTerm := ast.Item(ast.StringTerm("result"), ast.BooleanTerm(*checkPayload.Payload.Allowed))
			explanationTupleTerm := ast.Item(ast.StringTerm("explanation"), ast.StringTerm(
				fmt.Sprintf("User %s has a %s relationship with object %s in namespace %s", subject,
					relation, object, namespace)))
			return ast.ObjectTerm(resultTupleTerm, explanationTupleTerm), nil
		},
	)
}
