package regexlite

import (
	"errors"
	"fmt"
	"log"
	"regexp"
)

type regexPart struct {
	errorText string
	regex     string
}

type lengthData struct {
	length    int
	errorText string
}

type data struct {
	value         string
	regexparts    []regexPart
	minLengthData lengthData
	maxLengthData lengthData
}

func getRegexPart(regex string, errorText string) regexPart {
	return regexPart{errorText: errorText, regex: regex}
}

func Value(value string) *data {
	return &data{value: value}
}

func (d *data) HasText() *data {
	d.regexparts = append(
		d.regexparts,
		getRegexPart(".*[A-Za-z]+.*", "text not present in value"),
	)
	return d
}

func (d *data) HasNumbers() *data {
	d.regexparts = append(
		d.regexparts,
		getRegexPart(".*[0-9]", "number not present in value"),
	)
	return d
}

func (d *data) HasSpecialCharacter() *data {
	d.regexparts = append(
		d.regexparts,
		getRegexPart(".*[!@#$%^&*]", "special character not present in value"),
	)
	return d
}
func (d *data) HasUpperCase() *data {
	d.regexparts = append(d.regexparts, getRegexPart(".*[A-Z]", "uppercase not present in value"))
	return d
}
func (d *data) HasLowerCase() *data {
	d.regexparts = append(
		d.regexparts,
		getRegexPart(".*[a-z]", "lower case not present in value"),
	)
	return d
}

func (d *data) IsEmail() *data {
	d.regexparts = append(
		d.regexparts,
		getRegexPart("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", "invalid email"),
	)
	return d
}

func (d *data) IsUrl() *data {
	d.regexparts = append(
		d.regexparts,
		getRegexPart("^(https?|ftp):\\/\\/[^\\s/$.?#].[^\\s]*$", "invalid Url"),
	)
	return d
}

func (d *data) Contains(substring string) *data {
	d.regexparts = append(
		d.regexparts,
		getRegexPart(fmt.Sprintf(".*%v", substring), fmt.Sprintf("substring '%v' not present in value", substring)),
	)
	return d
}

func (d *data) Min(length int) *data {
	d.minLengthData = lengthData{
		length:    length,
		errorText: fmt.Sprintf("value should have a minimum of %v characters", length),
	}
	return d
}

func (d *data) Max(length int) *data {
	d.maxLengthData = lengthData{
		length:    length,
		errorText: fmt.Sprintf("value should have a maximum of %v characters", length),
	}
	return d
}

func (d *data) Validate() (valid bool, err error) {
	if d.minLengthData.length != 0 {
		if len(d.value) < d.minLengthData.length {
			return false, errors.New(d.minLengthData.errorText)
		}
	}
	if d.maxLengthData.length != 0 {
		if len(d.value) > d.maxLengthData.length {
			return false, errors.New(d.maxLengthData.errorText)
		}
	}
	for _, regexpart := range d.regexparts {
		r, err := regexp.Compile(regexpart.regex)
		if err != nil {
			log.Fatal(err)
		}
		if !r.MatchString(d.value) {
			return false, errors.New(regexpart.errorText)
		}
	}
	return true, nil
}
