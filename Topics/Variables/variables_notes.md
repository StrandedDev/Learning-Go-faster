
# Learning Go Lang Faster

## Go : Variables

Variables are containers that store a data to help us manually write the same data over and over again. Variables come in very handly when the data is changing or we don't know the data yet. We reference the container then the program does takes care of the job when the data is received, changed or assigned. 

**Note: The examples below does not have proper imports as well as print functions which may result into errors if you just copy and paste the code without the print functions.**

## View output

```go
package main
import "fmt"

func main(){

  var a int = 5
  var b int = 6  // variables

  fmt.Println(a, b) // output

}


```


## Keep in mind

- Every declared variables must be used, else it'll throw and error in compile time
- Variables are case sensitive in Go
- Variables declared with constants cannot be altered later in the program
- Go only stores the same type of data in same type of variable
- Shorthands do not require any keywords or explicit type declaration
- Defining types explicitly so debugging becomes easier, it's a good practice

## Declaring (Creating) variables:

### 1. `var` keyword:

```go
var name type = value;
```

### 2. `const` keyword:

```go
const name type = value;
```


`Note:`
- `const` variables cannot be reassigned or updated. 
- `const` variables must be initialized and assigned in the same line

### 3. `:=` shorthand:

```go
name := value;
```

> `Note:` 
- This is a shorthand for creating variables 
- `var` `const` keywords are not required, also not allowed
- The variable type gets inferred from the type of data is stored in the variable.

## Some points to remember

| `var` | `const` | `:=` |
|:-:|:-:|:-:|
| `Global scoped` | `Global scoped`| `Function scoped`|
| `Can overwrite` | `Cannot overwrite`| `Can overwrite`|
| `Loose type supported` | `Strict type`| `Loose type`|
| `Multiline` | `Single line`| `Single line`|


## Variable Initialization, Assignment

- Initialization is when the variable gets created (with/without value);
- Assignment is when data is stored into the variable.

### Initialization

```go
func main(){
    var name string;     // Initialized without value
    var age int = 24;    // Initialized and assigned value
    var num = 24;        // Initialized and assigned value without type
}

```

### Assignment

```go
func main(){
    var name string = "John doe";   // Initialized and assigned value
    car := audi;                    // Initialized and assigned value

    var age int;                    // Initialized
    age = 24;                       // Assigned value
}

```

## Declaring variables without initial values 

In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:

```go
func main(){
    var a int;      // Output will be 0
    var b bool;     // Output will be false
    var c string;   // Output will be "" or empty string
}

```


## Declaring multiple variables

- Go has this amazing feature of declaring multiple variables at once like below: 

```go
func main(){
    var a,b,c int = 1,2,3           // Same type in one line
	var d,e,f = 10," hello ",true   // Different types in one line

	m,n := " a ",false              // Different types in single line shorthand
}
```

- Go also supports multi-line declaration like below: 
- Wrap declaration inside `var()`

```go
func main(){
	var (
		o int  = 11     //
		p bool = true   // Multi type and multi line variable declaration
		q      = " hi " // 
		r int           //
	)

	r = 12 // Initialized inside brackets, Assigned outside
}
```

## Next section : Data types


| [&larr; Back to hello world]("https://github.com/StrandedDev/Learning-Go-faster/blob/main/Topics/Installation/installation_notes.md") | [Data types &rarr;]("https://github.com/StrandedDev/Learning-Go-faster/blob/main/Topics/Installation/installation_notes.md") |
|:-|-:|