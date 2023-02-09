package main

import (
	"net/url"
	"strings"
)

type errors map[string][]string

func (e errors) Get(filed string) string {
	errorSlice := e[filed]
	if len(errorSlice) == 0 {
		return ""
	}

	return errorSlice[0]
}

func (e errors) Add(filed, message string) {
	e[filed] = append(e[filed], message)
}

type Form struct {
	Data   url.Values
	Errors errors
}

func NewForm(data url.Values) *Form {
	return &Form{
		Data:   data,
		Errors: map[string][]string{},
	}
}

func (f *Form) Has(filed string) bool {
	x := f.Data.Get(filed)
	if x == "" {
		return false
	}
	return true
}

func (f *Form) Required(fields ...string) {
	for _, filed := range fields {
		value := f.Data.Get(filed)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(filed, "This fields cannot be blank")
		}
	}
}

func (f *Form) Check(ok bool, key, message string) {
	if !ok {
		f.Errors.Add(key, message)
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
