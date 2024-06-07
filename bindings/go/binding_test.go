package tree_sitter_prisma_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-prisma"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_prisma.Language())
	if language == nil {
		t.Errorf("Error loading Prisma grammar")
	}
}
