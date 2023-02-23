package command

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

// Runner is an interface for each subcommand
type Runner interface {
	Name() string
	PrintHelp()
	Init([]string) error
	Run() error
}

var subcommands = []Runner{
	NewConvertCommand(),
}

// PrintHelp shows global help message.
func PrintHelp() {
	fmt.Fprintf(os.Stdout, `usage: typol <subcommand> [OPTION]...

subcommands
  convert    Convert input texts
`)

}

// Run executes a given subcommand.
func Run(args []string) error {
	if len(args) < 1 {
		err := errors.New("subcommand is required")
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err
	}
	cmd := args[0]

	var helpArgs = regexp.MustCompile(`^--?(h|help)$`)
	if helpArgs.MatchString(cmd) {
		PrintHelp()
		return nil
	}

	for _, s := range subcommands {
		if s.Name() == cmd {
			err := s.Init(args[1:])
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				s.PrintHelp()
				return err
			}
			return s.Run()
		}
	}
	err := fmt.Errorf("unknown subcommand: %s", cmd)
	fmt.Fprintf(os.Stderr, "%s\n", err)
	PrintHelp()
	return err
}
