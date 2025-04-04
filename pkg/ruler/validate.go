package ruler

import (
	"errors"

	"github.com/prequel-dev/prequel/pkg/schema"
	"github.com/rs/zerolog/log"
)

var (
	ErrMissingId             = errors.New("missing id field")
	ErrEmptyId               = errors.New("empty id field")
	ErrMissingName           = errors.New("missing name field")
	ErrEmptyName             = errors.New("empty name field")
	ErrMissingDisplayName    = errors.New("missing display name field")
	ErrEmptyDisplayName      = errors.New("empty display name field")
	ErrMissingDescription    = errors.New("missing description field")
	ErrEmptyDescription      = errors.New("empty description field")
	ErrInvalidTagsKind       = errors.New("invalid tags kind")
	ErrInvalidCategoriesKind = errors.New("invalid categories kind")
	ErrDuplicateName         = errors.New("duplicate name")
	ErrUnknownTag            = errors.New("unknown tag")
	ErrUnknownCategory       = errors.New("unknown category")
	ErrMissingCategory       = errors.New("missing category field")
)

type tagsT map[string]struct{}

func validateTagsFields(t RuleIncludeT, tags tagsT) error {

	if t.Metadata.Kind != schema.KindTags {
		return ErrInvalidTagsKind
	}

	return validateTags(t.Tags, t.Metadata.Kind, tags)
}

func validateCategoriesFields(t RuleIncludeT, tags tagsT) error {
	if t.Metadata.Kind != schema.KindCategories {
		return ErrInvalidCategoriesKind
	}

	return validateTags(t.Categories, t.Metadata.Kind, tags)
}

func validateTags(tags []TagT, kind string, dupes tagsT) error {
	for _, t := range tags {
		if t.Name == "" {
			return ErrMissingDisplayName
		}

		if _, ok := dupes[t.Name]; ok {
			log.Error().
				Str("name", t.Name).
				Str("kind", kind).
				Msg("Duplicate name")
			return ErrDuplicateName
		}

		dupes[t.Name] = struct{}{}

		if t.DisplayName == "" {
			return ErrMissingDisplayName
		}

		if t.Description == "" {
			return ErrMissingDescription
		}
	}

	return nil
}

func validateRules(rules RuleIncludeT, tags tagsT) error {

	if rules.Cre.Id == "" {
		log.Error().
			Any("rule", rules).
			Msg("Missing CRE id")
		return ErrMissingId
	}

	if rules.Metadata.Id == "" {
		log.Error().
			Any("rule", rules).
			Msg("Missing rule id")
		return ErrMissingId
	}

	if rules.Cre.Category == "" {
		log.Error().
			Any("rule", rules).
			Msg("Missing category")
		return ErrMissingCategory
	}

	if _, ok := tags[rules.Cre.Category]; !ok {
		log.Error().
			Str("category", rules.Cre.Category).
			Msg("Unknown category")
		return ErrUnknownCategory
	}

	for _, tag := range rules.Cre.Tags {
		if _, ok := tags[tag]; !ok {
			log.Error().
				Str("tag", tag).
				Msg("Uknown tag")
			return ErrUnknownTag
		}
	}

	return nil
}
