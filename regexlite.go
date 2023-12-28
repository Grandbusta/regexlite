package regexlite

import (
	"fmt"
	"log"
	"regexp"
)

type data struct {
	value      string
	regexparts []string
}

func Value(value string) *data {
	return &data{value: value}
}

func (d *data) HasText() *data {
	d.regexparts = append(d.regexparts, ".*[A-Za-z]+.*")
	return d
}

func (d *data) HasNumbers() *data {
	d.regexparts = append(d.regexparts, ".*[0-9]")
	return d
}

func (d *data) HasSpecialCharacter() *data {
	d.regexparts = append(d.regexparts, ".*[!@#$%^&*]")
	return d
}
func (d *data) HasUpperCase() *data {
	d.regexparts = append(d.regexparts, ".*[A-Z]")
	return d
}
func (d *data) HasLowerCase() *data {
	d.regexparts = append(d.regexparts, ".*[a-z]")
	return d
}

func (d *data) IsEmail() *data {
	d.regexparts = append(d.regexparts, "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
	return d
}

func (d *data) IsUrl() *data {
	d.regexparts = append(d.regexparts, "^(https?|ftp):\\/\\/[^\\s/$.?#].[^\\s]*$")
	return d
}

func (d *data) Contains(substring string) *data {
	d.regexparts = append(d.regexparts, fmt.Sprintf(".*%v", substring))
	return d
}

func (d *data) Validate() bool {
	for _, regex := range d.regexparts {
		r, err := regexp.Compile(regex)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(r, r.MatchString(d.value))
		if !r.MatchString(d.value) {
			return false
		}
	}
	return true
}
