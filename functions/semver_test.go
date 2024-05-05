package functions_test

import (
	"testing"

	"github.com/ninoseki/tsumiki/functions"
)

var params = []struct {
	in  string
	out string
}{
	{"0", "0.0.0"},
	{"1.0", "1.0.0"},
	{"1.0.0", "1.0.0"},
	{"1.2.3.4", ""},
}

func TestCoerce(t *testing.T) {
	for _, tt := range params {
		t.Run(tt.in, func(t *testing.T) {
			s := functions.Coerce(tt.in)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
