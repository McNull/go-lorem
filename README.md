# go-lorem

Seeded or nonseeded lorem ipsum generator for Go

Forked from https://github.com/drhodes/golorem

This fork adds the ability to generated seeded lorem ipsum text and uses the upgraded random/v2 package for random number generation.

## Example

```go
package main

import (
    "fmt"
    "github.com/mcnull/go-lorem"
)

func main() {
    // Generate a random paragraph
    fmt.Println(lorem.Paragraph(5, 10))

    // Generate a random sentence
    fmt.Println(lorem.Sentence(5, 10))

    // Generate a random word
    fmt.Println(lorem.Word(5, 10))

    // Generate a random email
    fmt.Println(lorem.Email())

    // Generate a random URL
    fmt.Println(lorem.URL())

    // Generate a random host
    fmt.Println(lorem.Host())

    // Generate a random UUID
    fmt.Println(lorem.UUID())
}
```

## Seeded

A seeded lorem generator is also available. This allows you to generate the same random text each time you run your program. This is useful for testing.

```go
package main

import (
    "fmt"
    "github.com/mcnull/go-lorem"
)

func main() {
    // Generate a seeded lorem ipsum generator
    l := lorem.NewSeed(1,2)

    // Generate a random paragraph
    fmt.Println(l.Paragraph(5, 10))

    // Same function as above can be called on the instance.
}
```
