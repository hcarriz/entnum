package entnum

import (
	"strings"
	"unicode"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func title(in string) string {

	r := []rune(strings.ToLower(in))

	if len(r) > 0 {
		r[0] = unicode.ToUpper(r[0])
	}

	return string(r)
}

func isEnum(in *gen.Field) bool {
	return in != nil && in.Type != nil && in.Type.Type == field.TypeEnum
}
