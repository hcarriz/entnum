package entnum

import (
	"io/fs"
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestNew(t *testing.T) {

	t.Run("checking", func(t *testing.T) {

		data, err := fs.ReadFile(os.DirFS("./internal/ent"), "entnum.go")
		if err != nil {
			t.Errorf("unable to open directory: %v", err)
		}

		snaps.MatchSnapshot(t, string(data))

	})

}
