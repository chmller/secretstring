# SecretString

SecretString is a string, that prevents you from accidentally printing secret values to the console, a log file etc.
You have to call it´s GetSecret() method to get the actual value.

It is highly inspired by Pydantic´s SecretStr (https://docs.pydantic.dev/usage/types/#secret-types)

```go
package main

import (
	"fmt"
	"github.com/chmller/secretstring"
)

func main() {
	s := secretstring.SecretString("this_is_a_secret")
	// or
	// var s SecretString = "this_is_a_secret"

	fmt.Println(s)
	fmt.Println(s.GetSecret())
}
```

this should output:

```
********
this_is_a_secret
```

It has build in marshaller and unmarshaller for JSON.
