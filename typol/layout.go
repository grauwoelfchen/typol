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
