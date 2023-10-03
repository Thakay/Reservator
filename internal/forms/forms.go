package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

// Valid return true if there are errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func New(data url.Values) *Form {

	return &Form{data,
		map[string][]string{}}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, fmt.Sprintf("%s field can not be blank", field))
		}
	}
}

func (f *Form) Has(field string, r *http.Request) (ok bool) {
	x := r.Form.Get(field)
	if x == "" {
		ok = false
	}
	ok = true
	return ok
}

// MinLength check for min char in a field
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("%s must be greater than %d", field, length))
		return false
	}
	return true
}
