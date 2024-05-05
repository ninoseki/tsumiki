package functions

import (
	"strings"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
	"golang.org/x/mod/semver"
)

func init() {
	rego.RegisterBuiltin1(
		&rego.Function{
			Name:        "semver.coerce",
			Description: "Coerces the input as a SemVer string.",
			Decl: types.NewFunction(
				types.Args(
					types.Named("vsn", types.S),
				),
				types.Named("result", types.S).Description("'' if `vsn` is a invalid SemVer."),
			),
		},
		func(_ rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
			var v string

			if err := ast.As(a.Value, &v); err != nil {
				return nil, err
			}

			return ast.StringTerm(Coerce(v)), nil
		},
	)
}

func Coerce(v string) string {
	if semver.IsValid((v)) {
		return v
	}

	np, _ := strings.CutPrefix(v, "v")
	cv := semver.Canonical("v" + np)

	coerced, _ := strings.CutPrefix(cv, "v")
	return coerced
}
