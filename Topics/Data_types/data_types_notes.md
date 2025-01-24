
# Learning Go Lang Faster

## Go : Data types

Go supports multiples and specific memory size data types, which help to make safer and efficient code. 

| Type | Description |
|:-|:-|
| [Numeric &rarr;](#Numeric-data-types) | Numeric data types represent arithmatic types like integer, floating points, complex numbers and so on |
| [Boolean &rarr;](#Boolean-data-type) | Boolean data typs are either true of false |
| [String &rarr;](#String-data-type) | Represents the set of string values or sequence of bytes. Strings are immutable. |
| [Derived &rarr;](#Derived-data-types) | They include Pointers, Arrays, Structures, Unions, Functions, Slices, Interfaces, Maps, Channels |

## Cheatsheet 

**Signed integers** means it can store positive or negative numbers. ***ex: 35, -60, 4***
**Unsigned integers** means it can store positive numbers only. ***ex: 50, 32, 40***

| Category | Keyword | Size | Feature |  Range |
|:-:|:-:|:-:|:-:|:-:|
| Integer | `int` | 8 bits | Signed | **32-bit:** -2147483648 to 2147483647 <br> **64-bit:** -9223372036854775808 to 9223372036854775807 |
| Integer | `int8` <br> `int16` <br> `int32` <br> `int64` | 8 bits <br> 16 bits <br> 32 bits <br> 64 bits | Signed | -128 to 127 <br> -32768 to 32767 <br> -2147483648 to 2147483647 <br> -9223372036854775808 to 9223372036854775807 |
| Integer | `uint8` <br> `uint16` <br> `uint32` <br> `uint64` | 8 bits <br> 16 bits <br> 32 bits <br> 64 bits | Unsigned | 0 to 127 <br> 0 to 32767 <br> 0 to 2147483647 <br> 0 to 9223372036854775807 |
| Floating | `float32` <br> `float64` | 32 bits <br> 64 bits |Signed | -3.4e+38 to 3.4e+38 <br> -1.7e+308 to +1.7e+308 |
| Complex | `complex64` <br> `complex128` | 64 bits <br> 128 bits | Unsigned |-3.4e38 + -1.2e38i (approx.) <br> 3.4e38 + 1.2e38i (approx.) |
| Aliases| `byte` <br> `rune` | 8 bits <br> 32 bits | Unsigned <br> Signed | 0 to 255 <br> -2147483648 to 2147483647 |
| String | `string` | - | - | - | - |
| Boolean | `bool` | - | - | - | 

## Numeric data types

Numeric data can be positive or negative like: 50, -40, 133 etc. For example: `int8` can store numbers between `-127 to 127`, whereas `uint8` can store from `0 to 127`.

**Signed integers** means it can store positive or negative numbers. ***ex: 35, -60, 4***
**Unsigned integers** means it can store positive numbers only. ***ex: 50, 32, 40***

| Category | Keyword | Types | Description |
|:-|:-:|:-:|:-:|
| Integers | `int` | int8. int16, int32, int64 | Signed, Contains negatives |
| Integers | `int` | uint8. uint16, uint32, uint64 | Unsigned, No negatives |
| Floating | `float` | float32, float64 | Signed, Round fp to avoid fp overflow |
| Imaginary | complex64, complex128 | complex64, complex128 | Signed |

**Example:**

```go
var x int8 = 123
var y float32 = 12.3
var z complex = 12i
```


**Note: There are some special integer data types in Go as:** 

| Type | Alias of  | Description |
|:-:|:-:|:-:|
| `rune` | `int32` | Signed, Provides the storage size required for multiple bytes in UTF-8 encoding |
| `byte` | `uint8` | Unsigned, Commonly used for raw data and character representations |
| `uintptr` | none | Unsigned, Large enough to store the uninterpreted bits of a pointer value |

**Example:**
```go
var letter rune = 'A'
fmt.Println(ch) // Output: 65 (integer value of the rune 'A')

var b byte = 'a'
fmt.Println(b) // Output: 97 (integer value of the byte 'a')
```


## Boolean data type


| Type | Keyword | Values |
|:-:|:-:|:-:|
| Boolean | `bool` | true, false |

**Example:**
```go

var married bool = false
var alive bool = true
```


## String data type

Anything written inside quotes or backticks is interpreted as a string value.

| Type | Keyword | Values | Defined inside |
|:-:|:-:|:-:|:-:|
| String | `string` | "any text" | " ", ' ', \` ` |

**Example:**

```go 
var firstName = "John"      // Uses double quotes
var lastName = 'Smith'      // Uses single quotes

// Strings created  with backticks are called string literals
var fullName = `John smith`

```

[More about strings &rarr;]()

## Derived data types


[Pointers &rarr;]()
[Arrays &rarr;]()
[Structures &rarr;]()
[Unions &rarr;]()
[Functions &rarr;]()
[Slices &rarr;]()
[Interfaces &rarr;]()
[Maps &rarr;]()
[Channels &rarr;]()





## Examples based on data types

#### Boolean

```go
var isTrue bool = true
var isFalse bool = false
```

Booleans in Go can only be `true` or `false`.

#### Integer

Go provides several integer types:

```go
var int8Value int8 = -128
var int16Value int16 = 32767
var int32Value int32 = 2147483647
var int64Value int64 = 9223372036854775807
var intValue int = 2147483647
var uint8Value uint8 = 255
var uint16Value uint16 = 65535
var uint32Value uint32 = 4294967295
var uint64Value uint64 = 18446744073709551615
```

#### Floating Point

Go provides two floating-point types:

```go
var float32Value float32 = 3.14
var float64Value float64 = 3.14159265358979323846
```

#### Complex Numbers

Go also supports complex numbers:

```go
var complex128Num complex128 = complex(3, 4)
var complex64Num complex64 = complex(3, 4)
```

### String

Strings in Go are immutable sequences of bytes:

```go
var stringVar string = "Hello, World!"
var stringVar2 := "Another string"
```

### Derived Data Types

#### Array

Arrays in Go have fixed size:

```go
var intArray [6]int = [6]int{1, 2, 3, 4, 5}
var slice := []int{1, 2, 3, 4, 5}
```

#### Struct

Structs allow grouping related fields together:

```go
type Person struct {
    Name string
    Age  int
}

var person Person = Person{Name: "John Doe", Age: 30}
```

#### Pointer

Pointers store memory addresses:

```go
var ptr *int = new(int)
*ptr = 42
```

#### Slice

Slices are dynamic arrays:

```go
slice := []int{1, 2, 3, 4, 5}
```

#### Map

Maps store key-value pairs:

```go
var m map[string]int = make(map[string]int)
m["one"] = 1
```

#### Channel

Channels are used for communication between goroutines:

```go
ch := make(chan int)
close(ch)
```

#### Interface

Interfaces define a contract that types must satisfy:

```go
type Writer interface {
    Write([]byte) error
}
```


# Next section : Operators


| [&larr; Back to variables](https://github.com/StrandedDev/Learning-Go-faster/blob/main/Topics/Variables/variables_notes.md) | [Operators &rarr;](https://github.com/StrandedDev/Learning-Go-faster/blob/main/Topics/Operators/operators_notes.md) |
|:-|-:|
