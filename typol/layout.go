//go:generate stringer -type=Layout -output=layout_string.go
package typol

import (
	"errors"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Layout int

const (
	Unknown Layout = iota
	Dvorak
	Qwerty
)

var UnknownLayoutErr = errors.New("unknown layout")

var layoutTypes = []Layout{
	Dvorak,
	Qwerty,
}

func FindLayoutType(s string) Layout {
	c := cases.Title(language.English)
	name := c.String(s)
	for _, t := range layoutTypes {
		if name == t.String() {
			return t
		}
	}
	return Unknown
}

// Qwerty => Dvorak
var DataQD = map[rune]rune{
	// line 0
	'q': '\'',
	'w': ',',
	'e': '.',
	'r': 'p',
	't': 'y',
	'y': 'f',
	'u': 'g',
	'i': 'c',
	'o': 'r',
	'p': 'l',
	// line 1
	'a': 'a',
	's': 'o',
	'd': 'e',
	'f': 'u',
	'g': 'i',
	'h': 'd',
	'j': 'h',
	'k': 't',
	'l': 'n',
	';': 's',
	// line 2
	'z':  ';',
	'x':  'q',
	'c':  'j',
	'v':  'k',
	'b':  'x',
	'n':  'b',
	'm':  'm',
	',':  'w',
	'.':  'v',
	'\\': 'z',
	// others
	'\'': '-',
	'-':  '\\',
	'=':  ']',
	'[':  '/',
	']':  '=',
	// line 0 + SHIFT
	'Q': '"',
	'W': '<',
	'E': '>',
	'R': 'P',
	'T': 'Y',
	'Y': 'F',
	'U': 'G',
	'I': 'C',
	'O': 'R',
	'P': 'L',
	// line 1 + SHIFT
	'A': 'A',
	'S': 'O',
	'D': 'E',
	'F': 'U',
	'G': 'I',
	'H': 'D',
	'J': 'H',
	'K': 'T',
	'L': 'N',
	':': 'S',
	// line 2 + SHIFT
	'Z': ':',
	'X': 'Q',
	'C': 'J',
	'V': 'K',
	'B': 'X',
	'N': 'B',
	'M': 'M',
	'<': 'W',
	'>': 'V',
	'?': 'Z',
	// others + SHIFT
	'"': '_',
	'+': '}',
	'_': '{',
	'{': '?',
	'}': '+',
}
