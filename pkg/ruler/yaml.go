package ruler

import (
	"github.com/prequel-dev/prequel/pkg/parser"
)

type RuleIncludeT struct {
	Metadata   parser.ParseRuleMetadataT `yaml:"metadata"`
	Tags       []TagT                    `yaml:"tags,omitempty"`
	Categories []TagT                    `yaml:"categories,omitempty"`
	Cre        parser.ParseCreT          `yaml:"cre,omitempty"`
	Rule       parser.ParseRuleDataT     `yaml:"rule"`
}

type TagT struct {
	Name        string `yaml:"name" binding:"required"`
	DisplayName string `yaml:"displayName" binding:"required"`
	Description string `yaml:"description" binding:"required"`
}
