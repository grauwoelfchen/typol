//go:generate stringer -type=Layout -output=layout_string.go
package command

type Layout int

const (
	Unknown Layout = iota
	Dvorak
	Qwerty
)
