# GOLang Tutorial

## To install please visit <https://go.dev/dl/>

## Basic

To initialize the package run `go mod init <dependency name>.`

To import any module we need to use `import` keyword

```golang
  import (
    "fmt"
  )
```

To create any new function `func` keyword. `main()` func is used as entry point for executable code.

```golang
func main() {
  fmt.Println("Hello ಜಗತ್ತು");
}
```

To Run the go file run `go run .` or `go run <fileName>`

## Data Types

| Type  | Data Type | Description |
| --- | --- | --- |
| signed Number | int (32 or 64), int8, int16, int32, int64 | Signed Integer numbers |
| unsigned Number | uint(32 or 64), uint8, uint16, uint32, uint64| Un-signed integer numbers |
| unicode  | rune | Unicode characters internally stored as int32 |
| Byte | byte | A byte (8 bits) of non-negative integers. |
| boolean | bool | true or false |
| String | string | sequence of characters |
| Decimal number | float | decimal numbers like 9.21 |
 