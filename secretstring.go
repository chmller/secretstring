// Package secretstring is a simple package that provides a type that can be used to store secrets.
// The secret is masked when printed or marshalled, but can be retrieved using the GetSecret() method.
package secretstring

import (
	"reflect"
	"strings"
)

const (
	defaultMask           string = "********"
	defaultMarshallMasked bool   = false
)

// secretString is the (private) underlying type of SecretString that holds the actual secret.
// It is used to prevent fmt.Sprint or alike from printing the secret,
// since they don't call the String() method, but instead read the underlying value from memory.
type secretString string

// SecretString is a string that is masked when printed.
// It implements the Stringer interface.
type SecretString struct {
	secret         secretString
	mask           string
	marshallMasked bool
}

// Options are used to configure the SecretString.
type Options struct {
	// MarshallMasked determines whether the secret is masked or not when marshalled as JSON.
	MarshallMasked bool
	// Mask is the string that is used to mask the secret.
	Mask string
}

// New creates a new instance of SecretString with default options.
func New(secret string) *SecretString {
	return &SecretString{
		secret:         secretString(secret),
		mask:           defaultMask,
		marshallMasked: defaultMarshallMasked,
	}
}

// NewWithOptions creates a new instance of SecretString with custom options.
func NewWithOptions(secret string, options Options) *SecretString {
	return &SecretString{
		secret:         secretString(secret),
		mask:           options.Mask,
		marshallMasked: options.MarshallMasked,
	}
}

// String returns the masked secret.
func (s *secretString) String() string {
	return "********"
}

// GetSecret returns the actual secret.
func (s *secretString) GetSecret() string {
	return reflect.ValueOf(s).Elem().String()
}

// String returns the masked secret.
//
// Example:
//
//		s := New("my magic secret")
// 		s.String() // returns "********"
func (s *SecretString) String() string {
	return s.mask
}

// GetSecret returns the actual secret.
//
// Example:
//
// 		s := New("my magic secret")
// 		s.GetSecret() // returns "my magic secret"
func (s *SecretString) GetSecret() string {
	return s.secret.GetSecret()
}

// MarshalJSON returns the actual secret as JSON if Options.MarshallMasked is true
// or the masked secret if Options.MarshallMasked is false.
//
// Example of a masked secret:
//
// 		o := Options{
//			MarshallMasked: true,
//			Mask:           "???",
//		}
//		s := NewWithOptions("hello world", o)
//		s.MarshalJSON() // returns "???"
//
// Example of an unmasked secret:
//
// 		o := Options{
//			MarshallMasked: false,
//			Mask:           "********",
//		}
//		s := NewWithOptions("hello world", o)
//		s.MarshalJSON() // returns "hello world"
func (s *SecretString) MarshalJSON() ([]byte, error) {
	var marshalled []byte

	if s.marshallMasked {
		marshalled = []byte("\"" + s.String() + "\"")
	} else {
		marshalled = []byte("\"" + s.GetSecret() + "\"")
	}

	return marshalled, nil
}

// UnmarshalJSON sets the secret from JSON. Any value is treated as a string.
//
// Example:
//
// 		body := []byte(`{"pin_number":"0000"}`)
// 		type testStruct struct {
// 			PinNumber *SecretString `json:"pin_number"`
// 		}
// 		var t testStruct
// 		json.Unmarshal(body, &t)
// 		t.PinNumber.String() // returns "********"
// 		t.PinNumber.GetSecret() // returns "0000"
//
func (s *SecretString) UnmarshalJSON(b []byte) error {
	encodedValue := string(b)
	encodedValue = strings.Trim(encodedValue, "\"")

	s.secret = secretString(encodedValue)
	s.mask = defaultMask
	s.marshallMasked = defaultMarshallMasked

	return nil
}
