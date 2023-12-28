package regexlite

import (
	"errors"
	"fmt"
	"log"
	"reflect"
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

type AnyType interface {
	string | int | uint
}

type data[K AnyType] struct {
	Value         K
	regexparts    []regexPart
	minLengthData lengthData
	maxLengthData lengthData
}

func (d *data[K]) String() string {
	return reflect.ValueOf(d.Value).String()
}

func getRegexPart(regex string, errorText string) regexPart {
	return regexPart{errorText: errorText, regex: regex}
}

func Value[K AnyType](value K) *data[K] {
	return &data[K]{Value: value}
}

func (d *data[K]) HasText() *data[K] {
	d.regexparts = append(
		d.regexparts,
		getRegexPart(".*[A-Za-z]+.*", "text not present in value"),
	)
	return d
}

func (d *data[K]) HasNumbers() *data[K] {
	d.regexparts = append(
		d.regexparts,
		getRegexPart(".*[0-9]", "number not present in value"),
	)
	return d
}

func (d *data[K]) HasSpecialCharacter() *data[K] {
	d.regexparts = append(
		d.regexparts,
		getRegexPart(".*[!@#$%^&*]", "special character not present in value"),
	)
	return d
}
func (d *data[K]) HasUpperCase() *data[K] {
	d.regexparts = append(d.regexparts, getRegexPart(".*[A-Z]", "uppercase not present in value"))
	return d
}
func (d *data[K]) HasLowerCase() *data[K] {
	d.regexparts = append(
		d.regexparts,
		getRegexPart(".*[a-z]", "lower case not present in value"),
	)
	return d
}

func (d *data[K]) IsEmail() *data[K] {
	d.regexparts = append(
		d.regexparts,
		getRegexPart("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", "invalid email"),
	)
	return d
}

func (d *data[K]) IsUrl() *data[K] {
	d.regexparts = append(
		d.regexparts,
		getRegexPart("^(https?|ftp):\\/\\/[^\\s/$.?#].[^\\s]*$", "invalid Url"),
	)
	return d
}

func (d *data[K]) Contains(substring string) *data[K] {
	d.regexparts = append(
		d.regexparts,
		getRegexPart(fmt.Sprintf(".*%v", substring), fmt.Sprintf("substring '%v' not present in value", substring)),
	)
	return d
}

func (d *data[K]) Min(length int) *data[K] {
	d.minLengthData = lengthData{
		length:    length,
		errorText: fmt.Sprintf("value should have a minimum of %v characters", length),
	}
	return d
}

func (d *data[K]) Max(length int) *data[K] {
	d.maxLengthData = lengthData{
		length:    length,
		errorText: fmt.Sprintf("value should have a maximum of %v characters", length),
	}
	return d
}

func (d *data[K]) Validate() (valid bool, err error) {
	strValue := d.String()

	if d.minLengthData.length != 0 {
		if len(strValue) < d.minLengthData.length {
			return false, errors.New(d.minLengthData.errorText)
		}
	}
	if d.maxLengthData.length != 0 {
		if len(strValue) > d.maxLengthData.length {
			return false, errors.New(d.maxLengthData.errorText)
		}
	}
	for _, regexpart := range d.regexparts {
		r, err := regexp.Compile(regexpart.regex)
		if err != nil {
			log.Fatal(err)

		}
		if !r.MatchString(strValue) {
			return false, errors.New(regexpart.errorText)
		}
	}
	return true, nil
}
