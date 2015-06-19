package slugify

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

var (
	defaultSlugger = New(Configuration{})
)

// Slugify creates the slug for a given value
func Slugify(value string) string {
	return defaultSlugger.Slugify(value)
}

func validCharacter(c rune) bool {

	if c >= 'a' && c <= 'z' {
		return true
	}

	if c >= '0' && c <= '9' {
		return true
	}

	return false
}

// Slugifier based on settings
type Slugifier struct {
	isValidCharacter func(c rune) bool
	replaceCharacter rune
}

// Slugify creates a slug for a string
func (s Slugifier) Slugify(value string) string {

	value = strings.ToLower(value)
	var buffer bytes.Buffer
	lastCharacterWasInvalid := false

	for len(value) > 0 {

		c, size := utf8.DecodeRuneInString(value)

		if s.isValidCharacter(c) {
			buffer.WriteRune(c)
			lastCharacterWasInvalid = false
		} else if lastCharacterWasInvalid == false {
			buffer.WriteRune(s.replaceCharacter)
			lastCharacterWasInvalid = true
		}

		value = value[size:]

	}

	return strings.Trim(buffer.String(), string(s.replaceCharacter))
}

// Configuration is the basic configuration for Slugifier
type Configuration struct {
	IsValidCharacterChecker func(rune) bool
	ReplaceCharacter        rune
}

// New initialize a new slugifier
func New(config Configuration) *Slugifier {
	if config.IsValidCharacterChecker == nil {
		config.IsValidCharacterChecker = validCharacter
	}

	if config.ReplaceCharacter == 0 {
		config.ReplaceCharacter = '-'
	}

	return &Slugifier{isValidCharacter: config.IsValidCharacterChecker, replaceCharacter: config.ReplaceCharacter}
}
