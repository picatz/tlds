# tlds
> Top-level domains made easy.

A simple [top-level domain](https://en.wikipedia.org/wiki/Top-level_domain) library for gophers.

## Installation

```
go get github.com/picatz/tlds
```

## Example Usage

```go
package main

import (
  "fmt"
  "github.com/picatz/tlds"
)

func main() {
  for domain := tlds.StreamAllDomains() {
    fmt.Println(domain)
  }
}
```
