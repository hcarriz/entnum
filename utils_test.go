package entnum

import (
	"testing"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func Test_title(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "blank",
			args: "",
			want: "",
		},
		{
			name: "title",
			args: "title",
			want: "Title",
		},
		{
			name: "Number",
			args: "4title",
			want: "4title",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := title(tt.args); got != tt.want {
				t.Errorf("title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEnum(t *testing.T) {
	tests := []struct {
		name string
		args *gen.Field
		want bool
	}{
		{
			name: "nil",
			args: nil,
			want: false,
		},
		{
			name: "nil field",
			args: &gen.Field{Name: "Nil Field"},
			want: false,
		},
		{
			name: "wrong type",
			args: &gen.Field{Name: "Nil Field", Type: &field.TypeInfo{Type: field.TypeBool}},
			want: false,
		},
		{
			name: "correct type",
			args: &gen.Field{Name: "Nil Field", Type: &field.TypeInfo{Type: field.TypeEnum}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEnum(tt.args); got != tt.want {
				t.Errorf("isEnum() = %v, want %v", got, tt.want)
			}
		})
	}
}
