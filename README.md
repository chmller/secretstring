# SecretString

SecretString is a handy wrapper around string, that prevents you from accidently printing secret values to the console, a log file etc.
It´s highly inspired by Pydantic´s SecrerStr (https://docs.pydantic.dev/usage/types/#secret-types)

## Usage
```go
secretString := New("this_is_a_secret")
fmt.Println(secretString)
```

this should output:

```
********
```

You can also have an alternative mask:
```go
secretString := NewWithMask("this_is_a_secret", "???")
fmt.Println(secretString)
```

which should output:

```
???
```