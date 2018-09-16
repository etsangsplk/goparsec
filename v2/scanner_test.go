package v2

import (
	"testing"

	goparsec "github.com/bricef/goparsec"
)

func TestScannerCanBeCreatedWithEmtpyString(t *testing.T) {
	s := goparsec.NewScanner([]byte(``))
	if !s.Endof() {
		t.Error("Empty Buffer for scanner should end immediately.")
	}
}
