package regexlite_test

import (
	"testing"

	"github.com/Grandbusta/regexlite"
)

func TestNumberComparison(t *testing.T) {
	value := 5.24 // Some value to test

	t.Run("success", func(t *testing.T) {
		err := regexlite.ValidateNumber(value).Min(4).Max(12.78).Check()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("failure", func(t *testing.T) {
		err := regexlite.ValidateNumber(value).Min(6).Check()
		if err == nil {
			t.Error(err)
		}
	})
}
