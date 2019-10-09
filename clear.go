package autocomplete

import (
	"bytes"
	"strings"
	"unicode"
)

func clear(f string) (r []byte) {
	f = strings.ToLower(f)
	var b bytes.Buffer
	for _, letter := range f {
		switch {
		case
			unicode.IsLetter(letter),
			unicode.IsSpace(letter),
			unicode.IsDigit(letter):
			b.WriteRune(letter)
		default:
			//b.WriteByte(' ')
		}
	}
	return b.Bytes()
}
