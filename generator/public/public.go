package public

import (
	_ "embed"

	"gitlab.com/mnm/bud/budfs"
	"gitlab.com/mnm/bud/mod"

	"gitlab.com/mnm/bud/gen"
	"gitlab.com/mnm/bud/internal/gotemplate"
)

//go:embed public.gotext
var template string

var generator = gotemplate.MustParse("public", template)

func New(bfs budfs.FS, module *mod.Module) *Generator {
	return &Generator{
		BFS:    bfs,
		Module: module,
		// Embed  bool
		// Minify bool
	}
}

type Generator struct {
	BFS    budfs.FS
	Module *mod.Module
	Embed  bool
	Minify bool
}

func (g *Generator) GenerateFile(_ gen.F, file *gen.File) error {
	state, err := Load(g.BFS, g.Module, g.Embed, g.Minify)
	if err != nil {
		return err
	}
	code, err := generator.Generate(state)
	if err != nil {
		return err
	}
	file.Write(code)
	return nil
}