package service

import (
	"errors"
	"flag"
	"fmt"
	"regexp"
)

// Executor is an interface for each subcommand
type Executor interface {
	Name() string
	Init([]string) error
	Exec()
	Output() string
}

var helpArg = regexp.MustCompile(`^--?(h|help)$`)
var helpMsg = `usage: typol <subcommand> [OPTION]...

subcommands
  convert    Convert input texts`

// Run executes a given subcommand.
func Run(args []string) (string, error) {
	if len(args) < 1 {
		err := errors.New("subcommand is required")
		return "", err
	}
	name := args[0]

	if helpArg.MatchString(name) {
		return helpMsg, nil
	}

	// NOTE:
	// Each subcommand has buffered output (state), so we need to generate them
	// at runtime for now.
	var subcommands = []Executor{
		NewConvertCommand(),
	}

	for _, cmd := range subcommands {
		var err error
		if cmd.Name() == name {
			err = cmd.Init(args[1:])
			// We don't os.Exit(1) for requested ErrHelp
			if err == flag.ErrHelp {
				return cmd.Output(), nil
			}
			if err != nil {
				return "", err
			}

			cmd.Exec()
			return cmd.Output(), nil
		}
	}
	err := fmt.Errorf("unknown subcommand: %s", name)
	return "", err
}
