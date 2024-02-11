//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hcarriz/entnum"
)

func main() {

	config := &gen.Config{
		Target:  "./ent",
		Package: "github.com/hcarriz/entnum/internal/ent",
		Schema:  "github.com/hcarriz/entnum/internal/schema",
	}

	en, err := entnum.New()
	if err != nil {
		log.Fatalf("unable to start entenum: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(en),
	}

	if err := entc.Generate("./schema", config, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

}
