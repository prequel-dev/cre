package tests

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/prequel-dev/cre/pkg/logs"
	"github.com/prequel-dev/cre/pkg/ruler"
	"github.com/prequel-dev/preq/pkg/eval"
	"github.com/prequel-dev/prequel-compiler/pkg/parser"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

const (
	creFolders    = "cre-*"
	creRules      = "*.yaml"
	testLogFile   = "test.log"
	testFPLogFile = "test-fp.log"
)

var (
	rulesPath        = os.Getenv("RULES_PATH")
	level            = os.Getenv("LOG_LEVEL")
	defaultRulesPath = "../rules"
	defaultLogLevel  = "info"
)

func initLogger() {
	logs.InitLogger(logs.WithPretty(), logs.WithLevel(strings.ToUpper(level)))
}

func TestMain(m *testing.M) {
	initLogger()

	if rulesPath == "" {
		rulesPath = defaultRulesPath
	}

	if level == "" {
		level = defaultLogLevel
	}

	log.Info().Str("rulesPath", rulesPath).Msg("Starting tests")
	code := m.Run()
	os.Exit(code)
}

func TestCres(t *testing.T) {

	ctx := context.Background()

	// Find all CRE directories
	cres, err := filepath.Glob(filepath.Join(rulesPath, creFolders))
	if err != nil {
		t.Fatalf("Error finding CRE test files: %v", err)
	}

	// Read each CRE directory and run the tests
	for _, cre := range cres {

		var (
			rulesT     *parser.RulesT
			ruleData   []byte
			testData   []byte
			testFpData []byte
			err        error
		)

		log.Info().Str("cre", cre).Msg("Reading CRE directory")

		rules, err := filepath.Glob(filepath.Join(cre, creRules))
		if err != nil {
			t.Fatalf("Error finding CRE rule files: %v", err)
		}

		if len(rules) != 1 {
			t.Fatalf("Expected 1 rule file, got %d", len(rules))
		}

		ruleData, err = os.ReadFile(rules[0])
		if err != nil {
			t.Fatalf("Error reading CRE rule file %s: %v", rules[0], err)
		}

		if rulesT, err = parser.Unmarshal(ruleData); err != nil {
			t.Fatalf("Error unmarshalling rule file %s: %v", rules[0], err)
		}

		if len(rulesT.Rules) != 1 {
			t.Fatalf("Expected 1 rule, got %d", len(rulesT.Rules))
		}

		var r = rulesT.Rules[0]
		rulesT.Rules[0].Metadata.Hash, err = ruler.HashRule(r)
		if err != nil {
			t.Fatalf("Error hashing rule %s: %v", rules[0], err)
		}

		if ruleData, err = yaml.Marshal(rulesT); err != nil {
			t.Fatalf("Error marshalling rule file %s: %v", rules[0], err)
		}

		testData, err = os.ReadFile(filepath.Join(cre, testLogFile))
		if err != nil {
			t.Fatalf("Error reading CRE test file: %v", err)
		}

		// Optional FP log file
		testFpData, _ = os.ReadFile(filepath.Join(cre, testFPLogFile))

		t.Run(filepath.Base(cre), func(t *testing.T) {
			_, stats, err := eval.Detect(ctx, nil, string(testData), string(ruleData))
			if err != nil {
				t.Fatalf("Error running detection: %v", err)
			}

			if stats["problems"] != 1 {
				t.Fatalf("Expected problems, got %d", stats["problems"])
			}

			if len(testFpData) > 0 {
				_, stats, err := eval.Detect(ctx, nil, string(testFpData), string(ruleData))
				if err != nil {
					t.Fatalf("Error running detection: %v", err)
				}

				if stats["problems"] != 0 {
					t.Fatalf("Expected no problems, got %d", stats["problems"])
				}

			}
		})
	}
}
