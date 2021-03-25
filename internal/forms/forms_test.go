package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "c")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("whatever")

	if has {
		t.Error("form shows has field when it does not")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")

	if !has {
		t.Error("shows form does not have field when it should")
	}

}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	minLength := form.MinLength("whatever", 1)

	if minLength {
		t.Error("form shows min length for field that doesnt exist")
	}

	postedData = url.Values{}
	postedData.Add("a", "abcd")
	form = New(postedData)

	minLength = form.MinLength("a", 100)

	if form.Valid() {
		t.Error("form valid for field below MinLength requirements")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("Should have an error but didn't get one")
	}

	postedData = url.Values{}
	postedData.Add("a", "abcd")
	form = New(postedData)

	form.MinLength("a", 1)

	if !form.Valid() {
		t.Error("form not valid for field when it is")
	}

	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("Should not have an error but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.IsEmail("email")

	if form.Valid() {
		t.Error("form shows email for non-existent field")
	}

	postedData = url.Values{}
	postedData.Add("email", "abcd")
	form = New(postedData)

	form.IsEmail("email")

	if form.Valid() {
		t.Error("form shows email valid if not valid")
	}

	postedData = url.Values{}
	postedData.Add("email", "abcd@gmail.com")
	form = New(postedData)

	form.IsEmail("email")

	if !form.Valid() {
		t.Error("form shows email invalid if valid")
	}

}
