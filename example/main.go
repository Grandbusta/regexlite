package main

import (
	"fmt"

	"github.com/Grandbusta/regexlite"
)

func main() {
	isValidPassword, err := regexlite.Value("Thisisapass&").Min(8).Max(30).HasUpperCase().HasSpecialCharacter().Validate()
	fmt.Println(isValidPassword, err) // returns true or false based on validation. err returns nil or error based on validation
}
