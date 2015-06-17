package slugify

import (
	"bytes"
	"strings"
)

func validCharacter(c uint8) bool {

	if c >= 'a' && c <= 'z' {
		return true
	}

	if c >= '0' && c <= '9' {
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
