package regexlite_test

import (
	"testing"

	"github.com/Grandbusta/regexlite"
)

func TestStringValidation(t *testing.T) {
	// Some value to test:
	myURL := "https://github.com/Grandbusta/regexlite"
	myEmail := "amatsagu@github.com"
	myValue := "Hello world!"

	t.Run("success/url", func(t *testing.T) {
		err := regexlite.ValidateString(myURL).IsURL().Check()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("success/email", func(t *testing.T) {
		err := regexlite.ValidateString(myEmail).Min(2).Check()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("success/length", func(t *testing.T) {
		err := regexlite.ValidateString(myValue).Min(1).Max(20).Check()
		if err != nil {
			t.Error(err)
		}
	})
}
