package service

import (
	"bytes"
	"flag"
	"fmt"
	"strings"

	"git.sr.ht/grauwoelfchen/typol/typol"
)

// ConvertCommand converts Dvorak/Qverty input.
type ConvertCommand struct {
	fs  *flag.FlagSet
	buf *bytes.Buffer

	from typol.Layout
	to   typol.Layout

	// user inputs
	in  string
	out string
	txt string
}

var _ Executor = &ConvertCommand{}

// NewConvertCommand returns ConvertCommand instance.
func NewConvertCommand() *ConvertCommand {
	cc := &ConvertCommand{
		fs:  flag.NewFlagSet("convert", flag.ContinueOnError),
		buf: &bytes.Buffer{},
	}

	cc.fs.StringVar(&cc.in, "in", "Dvorak",
		"Input layout type ([Dd]vorak|[Qq]werty)")
	cc.fs.StringVar(&cc.out, "out", "Qwerty",
		"Output layout type ([Dd]vorak|[Qq]werty)")

	cc.fs.SetOutput(cc.buf)
	cc.fs.Usage = func() {
		fmt.Fprintln(cc.buf, "usage: convert [OPTION]... TEXT")
		cc.fs.PrintDefaults()
	}
	return cc
}

// Name returns this subcommand's name.
func (c *ConvertCommand) Name() string {
	return c.fs.Name()
}

// Init parses arguments.
func (c *ConvertCommand) Init(args []string) error {
	// Note:
	//   * We set ContinueOnError instead of ExitOnError or PanicOnError for this
	//     FlagSet
	//   * This returns ErrHelp if help message is requested
	err := c.fs.Parse(args)
	if err != nil {
		return err
	}

	// this allows an argument to be passed without -in as followings:
	//
	// ```zsh
	// % typol convert Hoi
	// % typol convert -in Dvorak Hoi
	// % typol convert -in Dvorak -out Qwerty Hoi
	// % typol convert -in Dvorak -out Qwerty --txt Hoi
	// # fyi, also these
	// % typol convert -- "-out"
	// % typol convert -in Dvorak -- "-out"
	// ```
	var nArgs = c.fs.Args()
	if len(nArgs) > 0 {
		if c.txt == "" && nArgs[0] != "" {
			c.txt = nArgs[0]
		}
	}

	// validations
	c.from = typol.FindLayoutType(c.in)
	if c.from == typol.Unknown {
		return typol.UnknownLayoutErr
	}
	c.to = typol.FindLayoutType(c.out)
	if c.to == typol.Unknown {
		return typol.UnknownLayoutErr
	}
	return nil
}

// Exec is actual command operations invoked from main function.
func (c *ConvertCommand) Exec() error {
	var out string
	var err error

	if c.txt == "" {
		return nil
	}

	switch c.from {
	case typol.Dvorak:
		out, err = c.toQwerty()
	case typol.Qwerty:
		out, err = c.toDvorak()
	default:
		out, err = c.toDvorak()
	}
	if err != nil {
		return err
	}
	c.buf.WriteString(out)
	return nil
}

// Output returns combined outputs from the buffer.
func (c *ConvertCommand) Output() string {
	out := strings.TrimSuffix(c.buf.String(), "\n")
	c.buf.Reset()
	return out
}

func (c *ConvertCommand) toDvorak() (string, error) {
	// FIXME: c.in
	return "TODO", nil
}

func (c *ConvertCommand) toQwerty() (string, error) {
	// FIXME: c.in
	return "TODO", nil
}
