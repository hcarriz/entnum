package entnum

import (
	"embed"
	"errors"
	"text/template"
	"unicode"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema"
)

var (
	_ entc.Extension = (*Extension)(nil)

	//go:embed entnum.tmpl
	templates embed.FS

	// TemplateFuncs holds the functions that are being used by entnum during generation.
	TemplateFuncs = template.FuncMap{
		"title":  title,
		"isEnum": isEnum,
	}
)

type Extension struct {
	name string
	entc.DefaultExtension
}

type option func(*Extension) error

func (o option) apply(c *Extension) error {
	return o(c)
}

// Option for future use.
type Option interface {
	apply(*Extension) error
}

type Retrieve struct {
	Verb string
}

func (r Retrieve) Name() string {
	return "EntnumVerb"
}

var _ schema.Annotation = (*Retrieve)(nil)

func Name(name string) Option {
	return option(func(e *Extension) error {

		if name == "" {
			return errors.New("entnum.Name() is missing a name")
		}

		if !unicode.IsLetter([]rune(name)[0]) {
			return errors.New("entnum.Name() required can only start with a letter")
		}

		e.name = title(name)
		return nil
	})
}

func (e *Extension) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Retrieve{Verb: e.name},
	}
}

func (e *Extension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("entnum").Funcs(TemplateFuncs).ParseFS(templates, "entnum.tmpl")),
	}
}

// New creates a new entnum.Extension to use with entc.
func New(opts ...Option) (*Extension, error) {

	var (
		e = Extension{
			name: "All",
		}
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
