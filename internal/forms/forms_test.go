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
		t.Error("got invalid when should have ben valid")
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
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

// Check for valid and invalid reponse from validator

func TestForm_Has(t *testing.T) {
	postData := url.Values{}
	postData.Add("FindMe", "GotMe")

	form := New(postData)
	if !form.Has("FindMe") {
		t.Error("Unable to find field in Form")
	}

	if form.Has("ImNotThere") {
		t.Error("Found a field that shouldn't be there.")
	}
}

func TestForm_MinLength(t *testing.T) {
	postData := url.Values{}
	postData.Add("AmIToShort", "NoIAmNot")

	form := New(postData)
	if !form.MinLength("AmIToShort", 5) {
		t.Error("Field NoIAmNot is longer then 5 and should have passed")
	}

	isError := form.Errors.Get("AmIToShort")
	if isError != "" {
		t.Error("error added when field does not violate min length requirement")
	}

	if form.MinLength("AmIToShort", 20) {
		t.Error("Field NoIAmNot is shorter then 20 and should have failed")
	}

	isError = form.Errors.Get("AmIToShort")
	if isError == "" {
		t.Error("no error added when field violates min length requirement")
	}
}

func TestForm_IsEmail(t *testing.T) {
	validEmailData := url.Values{}
	validEmailData.Add("email", "alex@gmail.com")
	validEmailForm := New(validEmailData)
	validEmailForm.IsEmail("email")
	if !validEmailForm.Valid() {
		t.Error("Valid email address provided but form said it was not valid")
	}

	invalidEmailData := url.Values{}
	invalidEmailData.Add("email", "notanemail")
	invalidEmailForm := New(invalidEmailData)
	invalidEmailForm.IsEmail("email")

	if invalidEmailForm.Valid() {
		t.Error("Invalid email in form but form is still valid")
	}

}
