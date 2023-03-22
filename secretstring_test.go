package secretstring

import (
	"fmt"
	"testing"
)

func TestMaskSecretString(t *testing.T) {
	expected := "********"
	secretString := New("this_is_a_secret")

	actual := fmt.Sprint(secretString)
	if actual != expected {
		t.Error("secret string is not masked")
	}
}

func TestNewWithMask(t *testing.T) {
	expected := "???"
	secretString := NewWithMask("this_is_a_secret", expected)

	actual := fmt.Sprint(secretString)
	if actual != expected {
		t.Error("secret string is not masked")
	}
}

func TestGetSecretString(t *testing.T) {
	expected := "this_is_a_secret"
	secretString := SecretString{secret: expected}

	actual := secretString.GetSecret()
	if actual != expected {
		t.Errorf("secret should be %s but is %s", expected, actual)
	}
}
