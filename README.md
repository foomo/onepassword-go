# onepassword-go

[![Build Status](https://github.com/foomo/onepassword-go/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/foomo/onepassword-go/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/onepassword-go)](https://goreportcard.com/report/github.com/foomo/onepassword-go)
[![godoc](https://godoc.org/github.com/foomo/onepassword-go?status.svg)](https://godoc.org/github.com/foomo/onepassword-go)
[![goreleaser](https://github.com/foomo/onepassword-go/actions/workflows/release.yml/badge.svg)](https://github.com/foomo/onepassword-go/actions)

> Unified way to handle op, service accounts & connect


## Usage

```go
package main

import (
  "context"
  "fmt"
  "testing"

  "github.com/foomo/onepassword-go"
)

func TestSecret(t *testing.T) {
  secret, err := onepassword.Secret(context.TODO(), "your-name", "vault", "item", "field")
  if err != nil {
    panic(err)
  }
  fmt.Println(secret)
}
```

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
