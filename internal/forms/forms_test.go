package forms

import (
	"net/http"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r, err := http.NewRequest("POST", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}

	form := New(r.PostForm)
	isValid := form.Valid()
	if !isValid {
		t.Error("got an invalid form but wanted valid")
	}
}

func TestForm_Fail_Valid(t *testing.T) {
	r, err := http.NewRequest("POST", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}
	
	form := New(r.PostForm)
	form.Errors.Add("first_name", "This field cannot be blank")
	isValid := form.Valid()
	if isValid {
		t.Error("got an valid form but wanted invalid")
	}
}

func TestForm_New(t *testing.T){
	data := url.Values{}
	data.Add("first_name", "Flint")
	form := New(data)
	if form.Get("first_name") != "Flint" {
		t.Error("Expected first_name to be equal to Flint")
	}
	if len(form.Errors) != 0 {
		t.Error("Expected form.Errors to be empty")
	}
}

func TestForm_Required(t *testing.T) {
	data := url.Values{}
	data.Add("first_name", "Flint")
	form := New(data)
	form.Required("first_name")
	if len(form.Errors) != 0 {
		t.Error("Expected no errors but got one")
	}
}

func TestForm_Fail_Required(t *testing.T) {
	data := url.Values{}
	form := New(data)
	form.Required("first_name")
	if len(form.Errors) == 0 {
		t.Error("Expected errors but didn't get one")
	}
}

func TestForm_Has(t *testing.T) {
	data := url.Values{}
	data.Add("first_name", "Flint")
	form := New(data)
	form.Has("first_name", &http.Request{Form: data})
	if len(form.Errors) != 0 {
		t.Error("Expected no errors but got one")
	}
}

func TestForm_Fail_Has(t *testing.T) {
	data := url.Values{}
	form := New(data)
	
	if form.Has("first_name", &http.Request{Form: data}) == true {
		t.Error("Expected errors but didn't get one")
	}
}

func TestForm_MinLength(t *testing.T) {
	data := url.Values{}
	data.Add("first_name", "Flint")
	form := New(data)
	form.MinLength("first_name", 3, &http.Request{Form: data})
	if len(form.Errors) != 0 {
		t.Error("Expected no errors but got one")
	}
}

func TestForm_Fail_MinLength(t *testing.T) {
	data := url.Values{}
	data.Add("first_name", "Fl")
	form := New(data)
	if form.MinLength("first_name", 3, &http.Request{Form: data}) == true {
		t.Error("Expected errors but didn't get one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	data := url.Values{}
	data.Add("email", "flintory@gmail.com")
	form := New(data)
	if !form.IsEmail("email") {
		t.Error("Expected valid email but got invalid")
	}
}

func TestForm_Fail_IsEmail(t *testing.T) {
	data := url.Values{}
	data.Add("email", "flintory")
	form := New(data)
	if form.IsEmail("email") {
		t.Error("Expected invalid email but got valid")
	}
}