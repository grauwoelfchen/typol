//go:build integration

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var binName = "typol"
var binPath = ""

var globalHelpMsg = `usage: typol <subcommand> [OPTION]...

subcommands
  convert    Convert input texts`

// The output from PrintDefaults() contains `\n    \t` :'(
// https://github.com/golang/go/blob/master/src/flag/flag.go#L575
var convertHelpMsg = `usage: convert [OPTION]... TEXT
  -in string
    	Input layout type ([Dd]vorak|[Qq]werty) (default "Dvorak")
  -out string
    	Output layout type ([Dd]vorak|[Qq]werty) (default "Qwerty")`

func execute(args []string) ([]byte, error) {
	cmd := exec.Command(binPath, args...)
	cmd.Env = os.Environ()
	return cmd.CombinedOutput()
}

func TestMain(m *testing.M) {
	err := os.Chdir(filepath.Join("..", "..", "dst"))
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	// Update the binPath variable at runtime
	binPath = filepath.Join(dir, binName)
	os.Exit(m.Run())
}

func TestRun(t *testing.T) {
	tests := map[string]struct {
		args   []string
		hasErr bool
		want   string // combined
	}{
		"no args": {
			args:   []string{},
			hasErr: true,
			want:   "subcommand is required\n",
		},
		"unknown subcommand": {
			args: []string{
				"invalid",
			},
			hasErr: true,
			want:   "unknown subcommand: invalid\n",
		},
		"with -h": {
			args: []string{
				"-h",
			},
			hasErr: false,
			want:   fmt.Sprintf("%s\n", globalHelpMsg),
		},
		"with -help": {
			args: []string{
				"-help",
			},
			hasErr: false,
			want:   fmt.Sprintf("%s\n", globalHelpMsg),
		},
		"with --help": {
			args: []string{
				"--help",
			},
			hasErr: false,
			want:   fmt.Sprintf("%s\n", globalHelpMsg),
		},
		"convert with no input": {
			args: []string{
				"convert",
			},
			hasErr: false,
			want:   "",
		},
		"convert with a direct input": {
			args: []string{
				"convert",
				"hello",
			},
			hasErr: false,
			want:   "jdpps\n",
		},
		"convert with a --help flag": {
			args: []string{
				"convert",
				"--help",
			},
			hasErr: false,
			want:   fmt.Sprintf("%s\n", convertHelpMsg),
		},
		"convert with a -in flag": {
			args: []string{
				"convert",
				"-in",
				"dvorak",
			},
			hasErr: false,
			want:   "",
		},
		"convert with unknown flag": {
			args: []string{
				"convert",
				"-foo",
			},
			hasErr: true,
			want:   "flag provided but not defined: -foo\n",
		},
		"convert with full args": {
			args: []string{
				"convert",
				"-in",
				"dvorak",
				"-out",
				"qwerty",
				"hello",
			},
			hasErr: false,
			want:   "jdpps\n",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := execute(tt.args)

			if tt.hasErr {
				if err == nil {
					t.Fatal("command should be failed")
				}
			} else {
				if err != nil {
					t.Fatalf("err: %v", err)
				}
			}

			got := string(out)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
