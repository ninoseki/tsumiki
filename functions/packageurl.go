package functions

import (
	"strings"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
	"github.com/package-url/packageurl-go"
)

type PackageURL struct {
	*packageurl.PackageURL
}

func (a PackageURL) Compare(b PackageURL) int {
	return strings.Compare(a.String(), b.String())
}

func init() {
	rego.RegisterBuiltin1(
		&rego.Function{
			Name:        "purl.parse",
			Description: "Parses the input as a package URL.",
			Decl: types.NewFunction(
				types.Args(
					types.Named("purl", types.S),
				),
				types.Named("result", types.A),
			),
		},
		func(_ rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
			var v string

			if err := ast.As(a.Value, &v); err != nil {
				return nil, err
			}

			purl, err := packageurl.FromString(v)
			if err != nil {
				return nil, err
			}

			qualifiers := ast.NewObject()
			for _, q := range purl.Qualifiers {
				qualifiers.Insert(ast.StringTerm("key"), ast.StringTerm(q.Key))
				qualifiers.Insert(ast.StringTerm("value"), ast.StringTerm(q.Value))
			}

			obj := ast.NewObject(
				ast.Item(ast.StringTerm("name"), ast.StringTerm(purl.Name)),
				ast.Item(ast.StringTerm("namespace"), ast.StringTerm(purl.Namespace)),
				ast.Item(ast.StringTerm("subpath"), ast.StringTerm(purl.Subpath)),
				ast.Item(ast.StringTerm("type"), ast.StringTerm(purl.Type)),
				ast.Item(ast.StringTerm("version"), ast.StringTerm(purl.Version)),
				ast.Item(ast.StringTerm("qualifiers"), ast.NewTerm(qualifiers)),
			)

			return &ast.Term{Value: obj}, nil
		},
	)
}
