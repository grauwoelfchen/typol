//go:build !integration

package service

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExec(t *testing.T) {
	tests := map[string]struct {
		args []string
		out  string
		err  error
	}{
		"no subcommand": {
			args: []string{},
			out:  "",
			err:  errors.New("subcommand is required"),
		},
		"unknown subcommand": {
			args: []string{"invalid"},
			out:  "",
			err:  errors.New("unknown subcommand: invalid"),
		},
		"an arg -h": {
			args: []string{"-h"},
			out:  helpMsg,
			err:  nil,
		},
		"an arg -help": {
			args: []string{"-help"},
			out:  helpMsg,
			err:  nil,
		},
		"an arg --help": {
			args: []string{"--help"},
			out:  helpMsg,
			err:  nil,
		},
		"convert": {
			args: []string{"convert"},
			out:  "",
			err:  nil,
		},
		"convert with an arg": {
			args: []string{"convert", "hello"},
			out:  "TODO",
			err:  nil,
		},
		"convert with full args": {
			args: []string{"convert", "-in", "Dvorak", "-out", "Qwerty", "hello"},
			out:  "TODO",
			err:  nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(*testing.T) {
			var got, want string

			out, err := Run(tt.args)

			if tt.err != nil {
				got = err.Error()
				want = tt.err.Error()

				if diff := cmp.Diff(want, got); diff != "" {
					t.Fatalf("err mismatch (-want +got)\n%s", diff)
				}
			}

			got = out
			want = tt.out
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("Stdout mismatch (-want +got)\n%s", diff)
			}
		})
	}
}
