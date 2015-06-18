package slugify

import (
	"bytes"
	"strings"
)

var (
	defaultSlugger = New(Configuration{})
)

// Slugify creates the slug for a given value
func Slugify(value string) string {
	return defaultSlugger.Slugify(value)
}

func validCharacter(c uint8) bool {

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
	isValidCharacter func(c uint8) bool
	replaceCharacter uint8
}

// Slugify creates a slug for a string
func (s Slugifier) Slugify(value string) string {

	value = strings.ToLower(value)
	var buffer bytes.Buffer
	lastCharacterWasInvalid := false

	for i := 0; i < len(value); i++ {
		c := value[i]
		if s.isValidCharacter(c) {
			buffer.WriteByte(c)
			lastCharacterWasInvalid = false
		} else if lastCharacterWasInvalid == false {
			buffer.WriteByte(s.replaceCharacter)
			lastCharacterWasInvalid = true
		}
	}

	return strings.Trim(buffer.String(), string(s.replaceCharacter))
}

// Configuration is the basic configuration for Slugifier
type Configuration struct {
	IsValidCharacterChecker func(uint8) bool
	ReplaceCharacter        uint8
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
