# GOLang Tutorial

## To install please visit <https://go.dev/dl/>

## [Basic](./basics/basic.go)

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

## [Data Types and Operators](./basics/dataTypes.go)

Syntax for declaring variables - `[var, const] variableNames(comma separated if multiple) optional_data_types = values(comma separated if multiple) (optional)`

`var` - Used for reassign the values for variable, Values can be mutated or changed.
`const` - Used where values should not be changed, Constants value cannot be changed.

Shorthand operator for assigning variables `numb := 1`

### Data Types

| Type  | Data Type | Description |
| --- | --- | --- |
| signed Number | int (32 or 64), int8, int16, int32, int64 | Signed Integer numbers |
| unsigned Number | uint(32 or 64), uint8, uint16, uint32, uint64| Un-signed integer numbers |
| unicode  | rune | Unicode characters internally stored as int32 |
| Byte | byte | A byte (8 bits) of non-negative integers. |
| boolean | bool | true or false |
| String | string | sequence of characters |
| Decimal number | float32 float64 | decimal numbers like 9.21 |

### Operators

| Type | Operator | Description |
| --- | --- | --- |
| Arithmetic | + - / * % ++ -- | A + B, A - B, A / B, A * B, A % B, A++, A-- |
| Logical | && \|\| ! | && -> And, || -> OR ! -> not |
| Assignment | =, +=, -=, *= | A+B=C, A+=B (A=A+B), A-=B (A=A-C), A*=B (A=A*B) |
| Relational | ==, != | returns true if Left operands value is equal to right operands value then returns true, vice versa for != |
| Miscellaneous operators | & * | & -> store the address of operators, * -> Refer Pointers of variable |

