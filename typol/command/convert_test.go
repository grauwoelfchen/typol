//go:build !integration

package command

import (
	"reflect"
	"testing"
)

func TestNewConvertCommand(t *testing.T) {
	t.Run("convert command type", func(t *testing.T) {
		got := NewConvertCommand()

		typ := reflect.TypeOf(got).String()
		if typ != "*command.ConvertCommand" {
			t.Errorf("failed to create ConvertCommand")
		}
	})
}

func TestName(t *testing.T) {
	t.Run("convert command name", func(t *testing.T) {
		got := NewConvertCommand()

		name := got.Name()
		if name != "convert" {
			t.Errorf("Name() mismatch: %s", name)
		}
	})
}

func TestInit(t *testing.T) {
	tests := map[string]struct {
		args   []string
		errMsg string
	}{
		"no args": {
			args:   []string{},
			errMsg: "",
		},
		"invalid arg": {
			args:   []string{"-in", "hello", "--unknown"},
			errMsg: "flag provided but not defined: -unknown",
		},
		"unknown layout": {
			args:   []string{"-to", "Colemak"},
			errMsg: UnknownLayoutErr.Error(),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cmd := NewConvertCommand()

			err := cmd.Init(tt.args)
			if err != nil {
				if err.Error() != tt.errMsg {
					t.Errorf("err: %v", err)
				}
			}
		})
	}
}
