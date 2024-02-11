package entnum_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/hcarriz/entnum"
)

func TestEntc(t *testing.T) {

	type test struct {
		name    string
		args    []entnum.Option
		wantErr bool
	}

	tests := []test{
		{
			name:    "generate with entc",
			args:    []entnum.Option{},
			wantErr: false,
		},
		{
			name: "custom verb",
			args: []entnum.Option{
				entnum.Name("return"),
			},
			wantErr: false,
		},
		{
			name: "missing verb",
			args: []entnum.Option{
				entnum.Name(""),
			},
			wantErr: true,
		},
		{
			name: "verb that starts with number",
			args: []entnum.Option{
				entnum.Name("4you"),
			},
			wantErr: true,
		},
		{
			name: "verb that starts with punctuation",
			args: []entnum.Option{
				entnum.Name("!you"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			dir := t.TempDir()

			pkg := "ent"
			sch := "schema"
			result := "entnum.go"
			schmaPath := "./test/schema"

			config := &gen.Config{
				Target:  filepath.Join(dir, pkg),
				Package: pkg,
				Schema:  sch,
			}

			got, err := entnum.New(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr {
				return
			}

			opts := []entc.Option{
				entc.Extensions(got),
			}

			if err := entc.Generate(schmaPath, config, opts...); err != nil {
				t.Fatalf("running ent codegen: %v", err)
				return
			}

			f := os.DirFS(dir)

			data, err := fs.ReadFile(f, filepath.Join(pkg, result))
			if err != nil {
				t.Errorf("unable to open directory: %v", err)
				return
			}

			snaps.MatchSnapshot(t, string(data))

		})
	}

}

// func TestNew(t *testing.T) {
// 	type args struct {
// 		opts []Option
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *Extension
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := New(tt.args.opts...)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("New() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
