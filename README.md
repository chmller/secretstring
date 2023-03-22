# SecretString

SecretString is a string, that prevents you from accidentally printing secret values to the console, a log file etc.
You have to call it´s GetSecret() method to get the actual value.

It is highly inspired by Pydantic´s SecretStr (https://docs.pydantic.dev/usage/types/#secret-types)

Example usage:
```go
package main

import (
	"fmt"
	"github.com/chmller/secretstring"
)

func main() {
	// create a new secret string with default options
	s := secretstring.New("this_is_a_secret")

	fmt.Println(s)
	fmt.Println(s.GetSecret())
	
	// output:
	// **********
	// this_is_a_secret
}
```

Or with custom options:

```go
package main

import (
	"fmt"
	"github.com/chmller/secretstring"
)

func main() {
	// create a new secret string with custom options
	o := secretstring.Options{
		MarshallMasked: false,
		Mask: "???",
    }
	s := secretstring.NewWithOptions("this_is_a_secret", o)

	fmt.Println(s)
	fmt.Println(s.GetSecret())

	// output:
	// ???
	// this_is_a_secret
}
```

It has build in marshaller and unmarshaller for JSON.
