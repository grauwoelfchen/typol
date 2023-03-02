package command

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"strings"
)

// ConvertCommand converts Dvorak/Qverty input.
type ConvertCommand struct {
	fs  *flag.FlagSet
	buf *bytes.Buffer

	in string
	to string
}

var layoutTypes = []string{
	"dvorak",
	"qwerty",
}

var _ Executor = &ConvertCommand{}

var UnknownLayoutErr = errors.New("unknown layout")

// NewConvertCommand returns ConvertCommand instance.
func NewConvertCommand() *ConvertCommand {
	cc := &ConvertCommand{
		fs:  flag.NewFlagSet("convert", flag.ContinueOnError),
		buf: &bytes.Buffer{},
	}

	cc.fs.StringVar(&cc.in, "in", "", "Input value needs to be converted")
	cc.fs.StringVar(&cc.to, "to", "Dvorak",
		"Layout type ([Dd]vorak|[Qq]werty)")

	cc.fs.SetOutput(cc.buf)
	cc.fs.Usage = func() {
		fmt.Fprintln(cc.buf, "usage: convert [input]")
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
	// % typol convert -to Dvorak Hoi
	// # fyi, also these
	// % typol convert -- "-to"
	// % typol convert -to Dvorak -- "-to"
	// ```
	var nArgs = c.fs.Args()
	if len(nArgs) > 0 {
		if c.in == "" && nArgs[0] != "" {
			c.in = nArgs[0]
		}
	}

	layout := c.toLayout()
	for _, t := range layoutTypes {
		if layout == t {
			return nil
		}
	}
	return UnknownLayoutErr
}

// Exec is actual command operations invoked from main function.
func (c *ConvertCommand) Exec() error {
	var out string
	var err error

	if c.in == "" {
		return nil
	}

	switch c.toLayout() {
	case "dvorak":
		out, err = c.toDvorak()
	case "qverty":
		out, err = c.toQwerty()
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

func (c *ConvertCommand) toLayout() string {
	return strings.ToLower(c.to)
}

func (c *ConvertCommand) toDvorak() (string, error) {
	// FIXME: c.in
	return "TODO", nil
}

func (c *ConvertCommand) toQwerty() (string, error) {
	// FIXME: c.in
	return "TODO", nil
}
