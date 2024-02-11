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

	t.Run("generate with entc", func(t *testing.T) {

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

		en, err := entnum.New()
		if err != nil {
			t.Fatalf("unable to start entenum: %v", err)
		}

		opts := []entc.Option{
			entc.Extensions(en),
		}

		if err := entc.Generate(schmaPath, config, opts...); err != nil {
			t.Fatalf("running ent codegen: %v", err)
		}

		f := os.DirFS(dir)

		data, err := fs.ReadFile(f, filepath.Join(pkg, result))
		if err != nil {
			t.Errorf("unable to open directory: %v", err)
		}

		snaps.MatchSnapshot(t, string(data))

	})

}
