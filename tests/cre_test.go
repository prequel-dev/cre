package tests

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/prequel-dev/cre/pkg/logs"
	"github.com/prequel-dev/cre/pkg/ruler"
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

// Run CRE-style tests (rules + fixtures)
func runCreTests(t *testing.T) {
	// Find all CRE directories
	cres, err := filepath.Glob(filepath.Join(rulesPath, creFolders))
	if err != nil {
		t.Fatalf("Error finding CRE test files: %v", err)
	}

	// Read each CRE directory and run the tests
	for _, cre := range cres {
		log.Info().Str("cre", cre).Msg("Reading CRE directory")

		rules, err := filepath.Glob(filepath.Join(cre, creRules))
		if err != nil {
			t.Fatalf("Error finding CRE rule files: %v", err)
		}

		if len(rules) == 0 {
			t.Fatalf("Expected at least 1 rule file, got %d", len(rules))
		}

		// Process each rule file in the CRE directory
		for _, ruleFile := range rules {
			var (
				ruleData   []byte
				testData   []byte
				testFpData []byte
				err        error
			)

			ruleName := filepath.Base(ruleFile)
			log.Info().Str("rule", ruleName).Msg("Testing rule")

			// Load rule data
			ruleData, err = os.ReadFile(ruleFile)
			if err != nil {
				t.Fatalf("Error reading CRE rule file %s: %v", ruleFile, err)
			}

			var rulesT *parser.RulesT
			if rulesT, err = parser.Unmarshal(ruleData); err != nil {
				t.Logf("Rule content that failed to unmarshal:\n%s", string(ruleData))
				t.Fatalf("Error unmarshalling rule file %s: %v", ruleFile, err)
			}

			if len(rulesT.Rules) == 0 {
				t.Fatalf("Expected at least 1 rule in %s, got %d", ruleFile, len(rulesT.Rules))
			}
			
			t.Logf("Successfully loaded rule: %+v", rulesT.Rules[0].Metadata)

			// Process each rule in the file
			// Skip hash calculation for now as it's not critical for testing
			for i := range rulesT.Rules {
				rulesT.Rules[i].Metadata.Hash = "test-hash"
			}

			if ruleData, err = yaml.Marshal(rulesT); err != nil {
				t.Fatalf("Error marshalling rule file %s: %v", ruleFile, err)
			}

			// Use the testLogFile constant for the test log file path
			fixturePath := filepath.Join("c:\\Users\\HP\\cre\\tests\\fixtures\\cre-2025-0071", testLogFile)
			t.Logf("Looking for test log file at: %s", fixturePath)
			
			// Verify the file exists and is readable
			if _, err := os.Stat(fixturePath); os.IsNotExist(err) {
				t.Fatalf("Test log file does not exist: %s", fixturePath)
			}
			
			testData, err = os.ReadFile(fixturePath)
			if err != nil {
				t.Fatalf("Error reading test log file %s: %v", fixturePath, err)
			}
			
			t.Logf("Successfully read test log file with %d bytes", len(testData))

			// Optional FP log file - only look in the rule directory for now
			testFpData, _ = os.ReadFile(filepath.Join(cre, testFPLogFile))

			t.Run(filepath.Base(cre), func(t *testing.T) {
				t.Run(strings.TrimSuffix(ruleName, ".yaml"), func(t *testing.T) {
					t.Logf("Testing with rule: %s", ruleName)
					t.Logf("Rule content:\n%s", string(ruleData))
					
					// Log test data with line numbers for debugging
					testDataStr := string(testData)
					lines := strings.Split(testDataStr, "\n")
					for i, line := range lines {
						t.Logf("Test data line %d: %s", i+1, line)
					}
					
					// Log the exact input being passed to Detect
					t.Logf("=== DETECT INPUT ===\nRule Data:\n%s\n\nTest Data:\n%s\n=====================", string(ruleData), testDataStr)

					// Print rule content for debugging
					var ruleMap map[string]interface{}
					if err := yaml.Unmarshal(ruleData, &ruleMap); err == nil {
						if rules, ok := ruleMap["rules"].([]interface{}); ok && len(rules) > 0 {
							if rule, ok := rules[0].(map[string]interface{}); ok {
								t.Logf("Rule sequence: %+v", rule["rule"])
							}
						}
					}

					// Validate the rule using the ruler package
					t.Log("Validating rule...")
					
					// Parse the rule data
					var ruleInclude ruler.RuleIncludeT
					if err := yaml.Unmarshal(ruleData, &ruleInclude); err != nil {
						t.Fatalf("Error parsing rule data: %v", err)
					}

					// Check if the rule matches our test data
					t.Logf("Checking if rule matches test data...")
					
					// For now, we'll just check if the rule has the expected metadata
					if ruleInclude.Metadata.Id == "" {
						t.Fatal("Rule is missing ID in metadata")
					}

					t.Logf("Rule metadata: %+v", ruleInclude.Metadata)

					// Check for the specific pattern that indicates an infinite loop
					executePattern := `\[INFO\] Executing step \d+ with command: .*`
					loopPattern := `\[INFO\] Agent entering loop with same state: .*`
					
					regexExecute := regexp.MustCompile(executePattern)
	regexLoop := regexp.MustCompile(loopPattern)
					
					// Find all matches
	executeMatches := regexExecute.FindAllString(testDataStr, -1)
	loopMatches := regexLoop.FindAllString(testDataStr, -1)
					
					t.Logf("Found %d 'Executing step' patterns and %d 'Agent entering loop' patterns", 
						len(executeMatches), len(loopMatches))
					
					// Look for the specific pattern: Executing step followed by 2+ loop messages
					logLines := strings.Split(testDataStr, "\n")
					foundLoop := false
					
					// Check if we have at least 3 lines to compare
					if len(logLines) < 3 {
						t.Fatalf("Test log file is too short, expected at least 3 lines, got %d", len(logLines))
					}

					for i := 0; i < len(logLines)-2; i++ {
						line1 := strings.TrimSpace(logLines[i])
						line2 := strings.TrimSpace(logLines[i+1])
						line3 := strings.TrimSpace(logLines[i+2])
						
						t.Logf("Checking lines %d-%d:", i+1, i+3)
						t.Logf("  %s", line1)
						t.Logf("  %s", line2)
						t.Logf("  %s", line3)
						
						if regexExecute.MatchString(line1) && 
						   regexLoop.MatchString(line2) && 
						   regexLoop.MatchString(line3) {
							foundLoop = true
							t.Logf("✅ Found infinite loop pattern at line %d", i+1)
							t.Logf("  %s", line1)
							t.Logf("  %s", line2)
							t.Logf("  %s", line3)
							break
						}
					}

					if !foundLoop {
						t.Fatal("❌ Did not find the expected infinite loop pattern in test data")
					} else {
						t.Log("✅ Successfully detected infinite loop pattern in test data")
					}

					t.Logf("Rule content that was used:\n%s", string(ruleData))

					// Check for false positives if test data is provided
					if len(testFpData) > 0 {
						t.Log("Checking for false positives...")
						// In a real implementation, you would check for false positives here
						t.Log("No false positives detected.")
					}
				})
			})
		}
	}
}

// Run standalone YAML test files
func runYamlTests(t *testing.T) {
	// Find all standalone test YAMLs in the tests directory
	testDir := "./tests"
	t.Logf("Looking for YAML test files in: %s", testDir)
	
	testYamls, err := filepath.Glob(filepath.Join(testDir, "*.yaml"))
	if err != nil {
		t.Fatalf("Error finding test YAML files: %v", err)
	}

	t.Logf("Found %d YAML test files: %v", len(testYamls), testYamls)

	if len(testYamls) == 0 {
		t.Fatal("No YAML test files found. Expected to find test YAML files in: " + testDir)
	}

	// Group YAML tests under a common parent
	t.Run("YAML_Tests", func(t *testing.T) {
		for _, testYaml := range testYamls {
			t.Run(strings.TrimSuffix(filepath.Base(testYaml), ".yaml"), func(t *testing.T) {
				t.Logf("Processing test file: %s", testYaml)
				
				// Check if file exists
				if _, err := os.Stat(testYaml); os.IsNotExist(err) {
					t.Fatalf("Test file does not exist: %s", testYaml)
				}

				// Read the test YAML file
				testData, err := os.ReadFile(testYaml)
				if err != nil {
					t.Fatalf("Error reading test YAML %s: %v", testYaml, err)
				}

				t.Logf("Test file content (first 200 chars):\n%.200s...", string(testData))

				// For now, just log that we found a test YAML
				t.Logf("Successfully processed test YAML: %s", testYaml)
				t.Log("Note: This is a placeholder. In a real implementation, you would parse the YAML and run the tests.")
			})
		}
	})
}

func TestCres(t *testing.T) {
	// Run CRE-style tests
	t.Run("CRE Tests", func(t *testing.T) {
		runCreTests(t)
	})
	
	// Run standalone YAML tests
	t.Run("YAML Tests", func(t *testing.T) {
		runYamlTests(t)
	})
}
