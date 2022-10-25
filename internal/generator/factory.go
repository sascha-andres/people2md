package generator

import (
	"github.com/sascha-andres/people2md/internal/generator/markdown"
	"github.com/sascha-andres/people2md/internal/types"
)

func GetGenerator() (types.DataBuilder, error) {
	return &markdown.MarkdownData{}, nil
}
