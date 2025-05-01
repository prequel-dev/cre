package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/prequel-dev/cre/pkg/logs"
	"github.com/rs/zerolog/log"
)

const (
	creFolders = "cre-*"
	testFile   = "test.log"
)

var (
	rulesPath    = os.Getenv("RULES_PATH")
	defaultPath  = "../rules"
	level        = os.Getenv("LOG_LEVEL")
	defaultLevel = "info"
)

func initLogger() {
	logs.InitLogger(logs.WithPretty(), logs.WithLevel(strings.ToUpper(level)))
}

func TestMain(m *testing.M) {
	initLogger()

	if rulesPath == "" {
		rulesPath = defaultPath
	}

	if level == "" {
		level = defaultLevel
	}

	log.Info().Str("rulesPath", rulesPath).Msg("Starting tests")
	code := m.Run()
	os.Exit(code)
}

func TestJson(t *testing.T) {

	// Find all CRE directories
	cres, err := filepath.Glob(filepath.Join(rulesPath, creFolders))
	if err != nil {
		t.Fatalf("Error finding CRE test files: %v", err)
	}

	// Read each CRE directory and run the tests
	for _, cre := range cres {

		log.Info().Str("cre", cre).Msg("Reading CRE directory")

		// Read the test file
		_, err := os.ReadFile(filepath.Join(cre, testFile))
		if err != nil {
			t.Fatalf("Error reading CRE test file %s: %v", testFile, err)
		}

		t.Run(filepath.Base(testFile), func(t *testing.T) {
			log.Info().Str("cre", cre).Msg("Running test")
		})
	}
}
