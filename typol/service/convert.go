package service

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
	"unicode/utf8"

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
	if err := c.fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse args: %w", err)
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

	nArgs := c.fs.Args()
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
func (c *ConvertCommand) Exec() {
	if c.txt == "" {
		return
	}

	var out string

	if c.from == typol.Dvorak && c.to == typol.Qwerty {
		// reverse DataQD
		data := make(map[rune]rune, len(typol.DataQD))
		for k, v := range typol.DataQD {
			data[v] = k
		}

		out = c.convert(data)
	} else if c.from == typol.Qwerty && c.to == typol.Dvorak {
		out = c.convert(typol.DataQD)
	}

	c.buf.WriteString(out)
}

// Output returns combined outputs from the buffer.
func (c *ConvertCommand) Output() string {
	out := strings.TrimSuffix(c.buf.String(), "\n")
	c.buf.Reset()

	return out
}

func (c *ConvertCommand) convert(data map[rune]rune) string {
	out := ""

	for i, w := 0, 0; i < len(c.txt); i += w {
		qr, width := utf8.DecodeRuneInString(c.txt[i:])
		c := qr

		if dr, ok := data[qr]; ok {
			c = dr
		}

		out = strings.Join([]string{out, string(c)}, "")
		w = width
	}

	return out
}
