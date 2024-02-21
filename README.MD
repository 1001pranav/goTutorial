# GOLang Tutorial

## To install please visit <https://go.dev/dl/>

### To initialize the package run `go mod init <dependency name>.`

## [Basic](./basics/basic.go)

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
| Logical | && \|\| ! | && -> And, \|\| -> OR ! -> not |
| Assignment | =, +=, -=, *= | A+B=C, A+=B (A=A+B), A-=B (A=A-C), A*=B (A=A*B) |
| Relational | ==, != | returns true if Left operands value is equal to right operands value then returns true, vice versa for != |
| Miscellaneous operators | & * | & -> store the address of operators, * -> Refer Pointers of variable |

## [Arrays and Slices](./3%20-%20arrays/array.go)

Syntax for Arrays - `<var> variable_name[size]<data_types> `

For pre defined array we can use the following syntax
```golang
var arr = [4]int{0,1,2,3};
// <OR>
arr1 := [...]int{25, 30, 122}; // Here compiler determines length
```

For multi-dimensional arrays, We can use following syntax: 
```golang
var multiDimensionalArray = [size][size]<DATA_TYPE> {
  {value1-1, value1-2},
  {value2-1, value2-2}
}
```

Pass By Value and Pass By Reference - 
* `Pass By Value` - Just passes the value to new variable, and does not change the original array.

```golang     
passByValue := arr;
```
* `Pass By Reference` - Points to original array, changes in this will reflect on the original array too.
```goLang
passByRef := &arr;
passByRef[0] = 12;
```

### Slices
Creating Slices using `make`
```goLang
  slice := make([]<DataType>, <Length>, <Capacity>);
  // OR

```
* `len()` or `length` - Length of slice is the number of element in the slice.
* `cap()` or `capacity` - The maximum number of elements that the slice can hold, Without reallocation of the array.  

## [For Loop](./4%20-%20conditions%20and%20loops/loops.go)
There are 4 ways to use for loops, in different ways we can use the for loops.

### Old School method.
```golang
arr := [8]int{3, 5, 9, 11, 26, 21, 223};

for i := 0; i < len(arr); i++ {
  if arr[i] == 0 {
    fmt.Println("Empty value in the array", i)
  } else {
    fmt.Println("Value of array is ", arr[i])
  }
}
```

### For with iteration condition.
```golang
arr := [8]int{3, 5, 9, 11, 26, 21, 223};

i := 0
for i < len(arr) {
  arr[i] += i
  i++
}
```

### Only FOR loops
```golang
arr := [8]int{3, 5, 9, 11, 26, 21, 223};
i = 0
for {
  fmt.Println("Array value is: ", arr[i])
  i++
  if i >= len(arr) {
    break
  }
}
```

### With range
```golang
arr := [8]int{3, 5, 9, 11, 26, 21, 223};
for index, arrVal := range arr {
  fmt.Println("Index is ", index, " Value of the array is", arrVal);
}
```

## [If Conditions](./4%20-%20conditions%20and%20loops/conditions.go)
In if condition there is no use of the condition
```golang
var x int = 10;
y := 20;
if x < y {
  fmt.printf("%d is greater than %d", x, y);
}
```