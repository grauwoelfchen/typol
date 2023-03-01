package command

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// ConvertCommand converts Dvorak/Qverty input.
type ConvertCommand struct {
	flags *flag.FlagSet

	in string
	to string
}

var layoutTypes = []string{
	"dvorak",
	"qwerty",
}

var _ Executor = &ConvertCommand{}

// NewConvertCommand returns ConvertCommand instance.
func NewConvertCommand() *ConvertCommand {
	cc := &ConvertCommand{
		flags: flag.NewFlagSet("convert", flag.ExitOnError),
	}
	cc.flags.StringVar(&cc.in, "in", "", "Input value needs to be converted")
	cc.flags.StringVar(&cc.to, "to", "Dvorak",
		"Layout type ([Dd]vorak|[Qq]werty)")

	cc.flags.Usage = func() {
		fmt.Fprintln(os.Stdout, "usage: convert [input]")
		cc.flags.PrintDefaults()
	}
	return cc
}

// Name returns this subcommand's name.
func (c *ConvertCommand) Name() string {
	return c.flags.Name()
}

// Init parses arguments.
func (c *ConvertCommand) Init(args []string) error {
	err := c.flags.Parse(args)
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
	var nArgs = c.flags.Args()
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
	return fmt.Errorf("unknown layout: %s", layout)
}

// Exec is actual command operations invoked from main function.
func (c *ConvertCommand) Exec() (string, error) {
	var value string
	switch c.toLayout() {
	case "dvorak":
		value = c.toDvorak()
	case "qverty":
		value = c.toQwerty()
	default:
		value = c.toDvorak()
	}
	if value != "" {
		// TODO
		return "TODO", nil
	}
	return "", nil
}

func (c *ConvertCommand) toLayout() string {
	return strings.ToLower(c.to)
}

func (c *ConvertCommand) toDvorak() string {
	// TODO
	return c.in
}

func (c *ConvertCommand) toQwerty() string {
	// TODO
	return c.in
}
