package ruler

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/mod/semver"
	"gopkg.in/yaml.v3"
)

var (
	ErrInvalidType = errors.New("invalid type")
)

var (
	packageName = "cre-rules"
	tagsDir     = "tags"
	tagsYaml    = "tags.yaml"
	catsYaml    = "categories.yaml"
)

type BuildCmd struct {
	InPath  string `name:"path" short:"p" help:"Path to read rules" default:"rules"`
	OutPath string `name:"out" short:"o" help:"Optional path to write files; default curdir"`
	Version string `name:"vers" short:"v" help:"Optional semantic version override"`
}

func (c *BuildCmd) Run() error {

	var (
		inPath  = c.InPath
		outPath = c.OutPath
		vers    = c.Version
	)

	if outPath == "" {
		var err error
		outPath, err = os.Getwd()
		if err != nil {
			log.Error().Err(err).Msg("Fail os.Getwd()")
			return err
		}
	}

	if vers == "" {
		vers = Semver()
	}

	if !strings.HasPrefix(vers, "v") {
		vers = "v" + vers
	}

	if !semver.IsValid(vers) {
		return fmt.Errorf("invalid semver: %s", vers)
	}

	if outPath != "" {
		if err := os.MkdirAll(outPath, 0755); err != nil {
			log.Error().Err(err).Msg("Fail mkdir all")
			return err
		}
	}

	if err := _build(vers, inPath, outPath, packageName); err != nil {
		return err
	}

	return nil
}

func processTags(inPath string) (tagsT, error) {
	var (
		tagsData          []byte
		categoriesData    []byte
		tagsSection       RuleIncludeT
		categoriesSection RuleIncludeT
		tags              = make(tagsT)
		err               error
	)

	tagsData, err = os.ReadFile(filepath.Join(inPath, tagsDir, tagsYaml))
	if err != nil {
		log.Error().Err(err).Msg("Fail read tags")
		return nil, err
	}

	if err := yaml.Unmarshal(tagsData, &tagsSection); err != nil {
		log.Error().Err(err).Msg("Fail unmarshal tags")
		return nil, err
	}

	categoriesData, err = os.ReadFile(filepath.Join(inPath, tagsDir, catsYaml))
	if err != nil {
		log.Error().Err(err).Msg("Fail read categories")
		return nil, err
	}

	if err := yaml.Unmarshal(categoriesData, &categoriesSection); err != nil {
		log.Error().Err(err).Msg("Fail unmarshal categories")
		return nil, err
	}

	if err := validateTagsFields(tagsSection, tags); err != nil {
		return nil, err
	}

	if err := validateCategoriesFields(categoriesSection, tags); err != nil {
		return nil, err
	}

	return tags, nil
}

func processRules(path string, tags tagsT) (*RuleIncludeT, error) {

	var (
		rulesData []byte
		rule      RuleIncludeT
		err       error
	)

	yamls, err := os.ReadDir(path)
	if err != nil {
		log.Error().Err(err).Msg("Fail read rules")
		return nil, err
	}

	log.Info().Int("count", len(yamls)).Msg("Processing rules")

	for _, y := range yamls {

		log.Info().Str("file", y.Name()).Msg("Processing rule")

		if !strings.HasSuffix(y.Name(), ".yaml") {
			continue
		}

		rulesData, err = os.ReadFile(filepath.Join(path, y.Name()))
		if err != nil {
			log.Error().Err(err).Msg("Fail read rules")
			return nil, err
		}

		if err := yaml.Unmarshal(rulesData, &rule); err != nil {
			log.Error().Err(err).Msg("Fail unmarshal rules")
			return nil, err
		}

		if err := validateRules(rule, tags); err != nil {
			return nil, err
		}

		rule.Metadata.Hash, err = hashRule(rule)
		if err != nil {
			return nil, err
		}

		log.Info().
			Str("hash", rule.Metadata.Hash).
			Msg("Rule")
	}

	return &rule, nil
}

func _build(vers, inPath, outPath, packageName string) error {

	var (
		rules = make(map[string]any)
		tags  tagsT
		err   error
	)

	log.Info().Str("vers", vers).Str("outPath", outPath).Msg("Building")

	if tags, err = processTags(inPath); err != nil {
		return err
	}

	log.Debug().Any("tags", tags).Msg("Tags")

	cres, err := os.ReadDir(inPath)
	if err != nil {
		log.Error().Err(err).Msg("Fail read rules dir")
		return err
	}

	for _, e := range cres {

		var (
			rule *RuleIncludeT
			err  error
		)

		if !e.IsDir() {
			log.Debug().Str("file", e.Name()).Msg("Skipping")
			continue
		}

		if !strings.HasPrefix(e.Name(), "cre-") {
			log.Debug().Str("file", e.Name()).Msg("Skipping")
			continue
		}

		log.Info().Str("file", e.Name()).Msg("Processing target")

		if rule, err = processRules(filepath.Join(inPath, e.Name()), tags); err != nil {
			return err
		}

		rules[rule.Metadata.Hash] = rule
	}

	doc, err := generateDocument(rules)
	if err != nil {
		return err
	}

	hash := _sha256(doc)

	fileName := makeFilename(packageName, vers, hash)
	fullPath := filepath.Join(outPath, fileName)

	if err = writeFile(fullPath, doc); err != nil {
		return err
	}

	hashPath := fmt.Sprintf("%s.sha256", fullPath)
	if err = writeFile(hashPath, []byte(hash)); err != nil {
		return err
	}

	fmt.Printf("Wrote file [sha256 %s]: %s\n", hash, fileName)
	fmt.Printf("Wrote hash file: %s\n", hashPath)

	return nil
}

func _sha256(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func writeFile(fn string, data []byte) error {
	fh, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	if _, err = fh.Write(data); err != nil {
		fh.Close()
		return err
	}

	return fh.Close()
}

func makeFilename(name, vers, hash string) string {
	name = strings.TrimSuffix(name, ".yaml")
	vers = strings.TrimPrefix(vers, "v")

	buildMeta := fmt.Sprintf(".%s", hash[:8])

	return fmt.Sprintf("%s.%s%s.yaml", name, vers, buildMeta)
}

// Convert to document per section

func generateDocument(rules map[string]any) ([]byte, error) {

	type docT struct {
		Rules []any `yaml:"rules,omitempty"`
	}

	// Gather keys to produce consistent order output
	keys := make([]string, 0, len(rules))
	for k := range rules {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer

	doc := docT{
		Rules: make([]any, 0),
	}

	for _, k := range keys {
		log.Debug().Any("rule", rules[k]).Msg("Adding rule")
		doc.Rules = append(doc.Rules, rules[k])
	}

	y, err := yaml.Marshal(&doc)
	if err != nil {
		return nil, err
	}

	buf.Write(y)

	return buf.Bytes(), nil
}
