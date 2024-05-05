package functions

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

func init() {
	rego.RegisterBuiltin1(
		&rego.Function{
			Name:        "validators.is_url",
			Description: "Validates if the input is a valid URL.",
			Decl: types.NewFunction(
				types.Args(
					types.Named("value", types.S),
				),
				types.Named("result", types.B),
			),
		},
		func(_ rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
			var v string

			if err := ast.As(a.Value, &v); err != nil {
				return nil, err
			}

			err := validation.Validate(v, validation.Required, is.URL)
			return ast.BooleanTerm(err == nil), nil
		},
	)

	rego.RegisterBuiltin1(
		&rego.Function{
			Name:        "validators.is_domain",
			Description: "Validates if the input is a valid domain.",
			Decl: types.NewFunction(
				types.Args(
					types.Named("value", types.S),
				),
				types.Named("result", types.B),
			),
		},
		func(_ rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
			var v string

			if err := ast.As(a.Value, &v); err != nil {
				return nil, err
			}

			err := validation.Validate(v, validation.Required, is.Domain)
			return ast.BooleanTerm(err == nil), nil
		},
	)

	rego.RegisterBuiltin1(
		&rego.Function{
			Name:        "validators.is_ipv4",
			Description: "Validates if the input is a valid IPv4.",
			Decl: types.NewFunction(
				types.Args(
					types.Named("value", types.S),
				),
				types.Named("result", types.B),
			),
		},
		func(_ rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
			var v string

			if err := ast.As(a.Value, &v); err != nil {
				return nil, err
			}

			err := validation.Validate(v, validation.Required, is.IPv4)
			return ast.BooleanTerm(err == nil), nil
		},
	)

	rego.RegisterBuiltin1(
		&rego.Function{
			Name:        "validators.is_ipv6",
			Description: "Validates if the input is a valid IPv6.",
			Decl: types.NewFunction(
				types.Args(
					types.Named("value", types.S),
				),
				types.Named("result", types.B),
			),
		},
		func(_ rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
			var v string

			if err := ast.As(a.Value, &v); err != nil {
				return nil, err
			}

			err := validation.Validate(v, validation.Required, is.IPv6)
			return ast.BooleanTerm(err == nil), nil
		},
	)

	rego.RegisterBuiltin1(
		&rego.Function{
			Name:        "validators.is_email",
			Description: "Validates if the input is a valid email.",
			Decl: types.NewFunction(
				types.Args(
					types.Named("value", types.S),
				),
				types.Named("result", types.B),
			),
		},
		func(_ rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
			var v string

			if err := ast.As(a.Value, &v); err != nil {
				return nil, err
			}

			err := validation.Validate(v, validation.Required, is.Email)
			return ast.BooleanTerm(err == nil), nil
		},
	)

}
