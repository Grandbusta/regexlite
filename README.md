# Regex Lite🔥 [![Twitter Follow](https://img.shields.io/twitter/follow/iamgrandbusta?style=social)](https://twitter.com/iamgrandbusta)

## 📖Description

`regexlite` is a Golang utility library designed to simplify the creation and validation of regular expressions.

## ✨Features

- [x] Chainable Functions
- [x] Easy construction of complex regex patterns
- [x] Built-in validators for common use-cases like email, URL, etc.
- [x] Customizable and extendable

## Installation

```bash
go get github.com/Grandbusta/regexlite
```

## 🛠️Usage

Here's a quick example to get you started:

```go
package main

import (
	"fmt"

	"github.com/Grandbusta/regexlite"
)

func main() {
	isValidPassword, err := regexlite.Value("Thisisapass&").Min(8).Max(30).HasUpperCase().HasSpecialCharacter().Validate()
	fmt.Println(isValidPassword, err) // returns true or false based on validation. err returns nil or error based on validation
}
```

## 📮Available Methods

- `HasText()`: Ensures the string contains text (a-z, A-Z).
- `HasNumbers()`: Ensures the string contains numbers.
- `HasSpecialCharacter()`: Ensures the string contains special characters.
- `Min(length: int)`: Sets the minimum length of the string.
- `Max(length: int)`: Sets the maximum length of the string.
- `Validate()`: Executes the validation and returns a boolean and error result.
- `HasUpperCase()`: Ensures string contains at least one uppercase character.
- `HasLowerCase()`: Ensures string contains at least one lowercase character.
- `IsEmail()`: Ensures string is a valid email.
- `IsUrl()`: Ensures string is a valid url.
- `Contains(substring string)`: Ensures string contains a substring.

## ➕Contributing

If you have a suggestion that would make this project better, please fork the repo and create a pull request.
Don't forget to give the project a star! Thanks again!

## 🤓 Author(s)

**Olaifa Boluwatife Jonathan** [![Twitter Follow](https://img.shields.io/twitter/follow/iamgrandbusta?style=social)](https://twitter.com/iamgrandbusta)

## 🔖 License

Distributed under the MIT License. See `LICENSE` for more information.
