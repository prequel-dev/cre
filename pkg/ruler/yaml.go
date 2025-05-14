package ruler

import (
	"github.com/prequel-dev/prequel-compiler/pkg/parser"
)

type RuleIncludeT struct {
	Metadata   parser.ParseRuleMetadataT `yaml:"metadata"`
	Tags       []TagT                    `yaml:"tags,omitempty"`
	Categories []TagT                    `yaml:"categories,omitempty"`
	Rules      []parser.ParseRuleT       `yaml:"rule"`
}

type TagT struct {
	Name        string `yaml:"name" binding:"required"`
	DisplayName string `yaml:"displayName" binding:"required"`
	Description string `yaml:"description" binding:"required"`
	Hash        string `yaml:"hash,omitempty"`
	Kind        string `yaml:"kind,omitempty"`
}
