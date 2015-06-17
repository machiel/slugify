package slugify

import (
	"bytes"
	"strings"
)

const (
	characterLowercaseA = 97
	characterLowercaseZ = 122
	characterZero       = 48
	characterNine       = 57
)

func validCharacter(c uint8) bool {

	if c >= characterLowercaseA && c <= characterLowercaseZ {
		return true
	}

	if c >= characterZero && c <= characterNine {
		return true
	}

	return false
}

// Slugify creates a slug for a string
func Slugify(value string) string {

	value = strings.ToLower(value)
	var buffer bytes.Buffer
	lastCharacterWasInvalid := false

	for i := 0; i < len(value); i++ {
		c := value[i]
		if validCharacter(c) {
			buffer.WriteByte(c)
			lastCharacterWasInvalid = false
		} else if lastCharacterWasInvalid == false {
			buffer.WriteByte('-')
			lastCharacterWasInvalid = true
		}
	}

	return strings.Trim(buffer.String(), "-")
}
