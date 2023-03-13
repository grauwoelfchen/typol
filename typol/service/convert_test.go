//go:build !integration

package service

import (
	"reflect"
	"testing"

	"git.sr.ht/grauwoelfchen/typol/typol"
	"github.com/google/go-cmp/cmp"
)

func TestNewConvertCommand(t *testing.T) {
	t.Run("convert command type", func(t *testing.T) {
		got := NewConvertCommand()

		typ := reflect.TypeOf(got).String()
		if typ != "*service.ConvertCommand" {
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
			errMsg: "failed to parse args: flag provided but not defined: -unknown",
		},
		"unknown layout": {
			args:   []string{"-in", "Colemak"},
			errMsg: typol.UnknownLayoutErr.Error(),
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

func TestConvert(t *testing.T) {
	// reverse
	dataDQ := make(map[rune]rune, len(typol.DataQD))
	for k, v := range typol.DataQD {
		dataDQ[v] = k
	}

	tests := map[string]struct {
		text string
		data map[rune]rune
		want string
	}{
		"loadkeys dvorak - qwerty to dvorak": {
			text: "loadkeys dvorak",
			data: typol.DataQD,
			want: "nraet.fo ekrpat",
		},
		"loadkeys dvorak - dvorak to qwerty": {
			text: "loadkeys /usr/share/keymaps/i386/qwerty/us.map.gz",
			data: dataDQ,
			want: "psahvdt; [f;o[;jaod[vdtmar;[g386[x,dokt[f;emareu\\",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cmd := NewConvertCommand()
			cmd.txt = tt.text

			got := cmd.convert(tt.data)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("err mismatch (-want +got)\n%s", diff)
			}
		})
	}
}
