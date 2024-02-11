package entnum

import (
	"embed"
	"errors"
	"strings"
	"text/template"
	"unicode"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

var (
	_ entc.Extension = (*Extension)(nil)

	//go:embed entnum
	templates     embed.FS
	TemplateFuncs = template.FuncMap{
		"title": func(in string) string {

			r := []rune(strings.ToLower(in))

			if len(r) > 0 {
				r[0] = unicode.ToUpper(r[0])
			}

			return string(r)
		},
		"lower": strings.ToLower,
		"isEnum": func(in *gen.Field) bool {
			return in != nil && in.Type.Type == field.TypeEnum

		},
	}
)

type Extension struct {
	entc.DefaultExtension
}

// type option func(*Extension) error

// func (o option) apply(c *Extension) error {
// 	return o(c)
// }

// Option for future use.
type Option interface {
	apply(*Extension) error
}

func (e *Extension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("entnum").Funcs(TemplateFuncs).ParseFS(templates, "entnum")),
	}
}

// New creates a new entnum.Extension to use with entc.
func New(opts ...Option) (*Extension, error) {

	var (
		e   = Extension{}
		err error
	)

	for _, opt := range opts {
		err = errors.Join(err, opt.apply(&e))
	}

	if err != nil {
		return nil, err
	}

	return &e, nil
}
