package command

import (
	"errors"
	"fmt"
	"regexp"
)

// Executor is an interface for each subcommand
type Executor interface {
	Name() string
	Init([]string) error
	Exec() (string, error)
}

var subcommands = []Executor{
	NewConvertCommand(),
}

var helpMsg = `usage: typol <subcommand> [OPTION]...

subcommands
  convert    Convert input texts`

// Run executes a given subcommand.
func Run(args []string) (string, error) {
	if len(args) < 1 {
		err := errors.New("subcommand is required")
		return "", err
	}
	cmd := args[0]

	var helpArgs = regexp.MustCompile(`^--?(h|help)$`)
	if helpArgs.MatchString(cmd) {
		return helpMsg, nil
	}

	for _, s := range subcommands {
		if s.Name() == cmd {
			err := s.Init(args[1:])
			if err != nil {
				return helpMsg, err
			}
			return s.Exec()
		}
	}
	err := fmt.Errorf("unknown subcommand: %s", cmd)
	return "", err
}
