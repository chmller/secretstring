package secretstring

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSecretString_New(t *testing.T) {
	s := New("hello world")

	if s.secret != "hello world" {
		t.Error("secret is not set correctly")
	}
	if s.mask != "********" {
		t.Error("mask is not set correctly")
	}
	if s.marshallMasked != false {
		t.Error("marshallMasked is not set correctly")
	}
}

func TestSecretString_NewWithOptions(t *testing.T) {
	o := Options{
		MarshallMasked: true,
		Mask:           "???",
	}
	s := NewWithOptions("hello world", o)

	if s.secret != "hello world" {
		t.Error("secret is not set correctly")
	}
	if s.mask != "???" {
		t.Error("mask is not set correctly")
	}
	if s.marshallMasked != true {
		t.Error("marshallMasked is not set correctly")
	}
}

func TestSecretString_String(t *testing.T) {
	expected := "********"

	s := New("this_is_a_secret")

	// when using the String() method directly
	direct := s.String()
	if direct != expected {
		t.Errorf("secret should be %s but is %s", expected, direct)
	}

	// when using fmt.Sprint or alike, we also expect the secret to be masked
	indirect := fmt.Sprint(s)
	if indirect != expected {
		t.Errorf("secret should be %s but is %s", expected, indirect)
	}
}

func TestSecretString_InnerStringMasked(t *testing.T) {
	expected := "********"

	s := New("this_is_a_secret")

	if s.secret.String() != expected {
		t.Errorf("inner secret should always be masked")
	}
}

func TestSecretString_GetSecret(t *testing.T) {
	expected := "this_is_a_secret"

	s := New("this_is_a_secret")

	actual := s.GetSecret()
	if actual != expected {
		t.Errorf("secret should be %s but is %s", expected, actual)
	}
}

func TestSecretString_MarshalJSON(t *testing.T) {
	s := New("hello world")

	type testStruct struct {
		Password *SecretString `json:"password"`
	}

	testInstance := testStruct{
		Password: s,
	}

	bytes, err := json.Marshal(testInstance)
	if err != nil {
		return
	}

	expected := "{\"password\":\"hello world\"}"
	actual := string(bytes)
	if actual != expected {
		t.Errorf("secret should be %s but is %s", expected, actual)
	}
}

func TestSecretString_MarshalJSON_Masked(t *testing.T) {
	o := Options{
		MarshallMasked: true,
		Mask:           "???",
	}
	s := NewWithOptions("hello world", o)

	type testStruct struct {
		Password *SecretString `json:"password"`
	}

	testInstance := testStruct{
		Password: s,
	}

	bytes, err := json.Marshal(testInstance)
	if err != nil {
		return
	}

	expected := "{\"password\":\"???\"}"
	actual := string(bytes)
	if actual != expected {
		t.Errorf("secret should be %s but is %s", expected, actual)
	}
}

func TestSecretString_UnmarshalJSON(t *testing.T) {
	bytes := []byte("{\"password\":\"supersecret!!\"}")

	type testStruct struct {
		Password *SecretString `json:"password"`
	}

	var testInstance testStruct
	_ = json.Unmarshal(bytes, &testInstance)

	if testInstance.Password.String() == "supersecret" {
		t.Errorf("field is not masked!")
	}
}
