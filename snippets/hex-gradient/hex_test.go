package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRGBFromHex(t *testing.T) {
	tests := []struct {
		hex         string
		RGB         RGB
		checkForErr bool
	}{
		{
			hex:         "#F967DC",
			RGB:         RGB{249, 103, 220},
			checkForErr: false,
		},
		{
			hex:         "#6B50FF",
			RGB:         RGB{107, 80, 255},
			checkForErr: false,
		},
		{
			hex:         "#9c6c23",
			RGB:         RGB{156, 108, 35},
			checkForErr: false,
		},
		{
			hex:         "#e3d654",
			RGB:         RGB{227, 214, 84},
			checkForErr: false,
		},
		{
			hex:         "#zzzzzz",
			RGB:         RGB{},
			checkForErr: true,
		},
		{
			hex:         "",
			RGB:         RGB{},
			checkForErr: true,
		},
	}

	for _, test := range tests {
		got, err := RGBFromHex(test.hex)
		if test.checkForErr && err == nil {
			t.Error("err not nil")
		}
		if eq := cmp.Equal(got, test.RGB); !eq {
			t.Errorf("got: %+v, want: %+v", got, test.RGB)
		}
	}
}
