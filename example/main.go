package main

import (
	"fmt"

	"github.com/Grandbusta/regexlite"
)

func main() {
	res := regexlite.Value("://g.com").Contains("x").Validate()
	fmt.Println(res)
}
