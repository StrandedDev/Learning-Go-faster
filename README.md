
# Learning Go Lang Faster

In this project i will try to learn go programming language as fast as i can. I will list down every important aspects of Go programming language and then i will read some docs of that particular topic and write some codes and try to understand how things work out. 

I will try to learn and document the whole journey in this [*github repo*](https://github.com/StrandedDev/Learning-Go-faster), all the garbage code i'll write will be pushed in here with folders and files. Ypu can jump through clicking different urls to navigate to specific topics and also jump straight to the related folders and notes i've made. Remember this is not Go lang [docs](https://go.dev/doc/) this is just me. 

# Tools 
 

| Code Editor | Plugin | Code preservation |
|:----------------:|:-----------:|:----------------------:|
|[VS Code](https://code.visualstudio.com/download) | [Official Go plugin](https://marketplace.visualstudio.com/items?itemName=golang.Go) | [Github](https://www.github.com) |


# Important 

**The examples below does not have proper imports as well as print functions which may result into errors if you just copy and paste the code without the print functions.**

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



# Core concepts of  any programming language

- [Installation](#Installing-Go-in-your-machine)
    - [Jump to folder &rarr;](https://github.com/StrandedDev/Learning-Go-faster/tree/main/Topics/Installation)
    - [Jump to notes &rarr;](https://github.com/StrandedDev/Learning-Go-faster/blob/main/Topics/Installation/installation_notes.md)

- [Variables](#Go-Variables)
    - [Jump to folder &rarr;](https://github.com/StrandedDev/Learning-Go-faster/tree/main/Topics/Variables)
    - [Jump to notes &rarr;](https://github.com/StrandedDev/Learning-Go-faster/blob/main/Topics/Variables/variables_notes.md)

- Data types

- Operators

- Control flow statements

- Functions 

- Loops 

- Input/Output

- Arrays/Lists

- Types system

- Error handling 

- Packages/Modules/Libraries

- OOP 


## Installing Go in your machine

Head over to the [official downloads page](https://go.dev/doc/install) then download and run the installer file. This will install Go on your machine. 

Run below command in cmd or terminal to verify installation. 

```bash
go version
```

Here is the dedicated folder for [installation](https://github.com/StrandedDev/Learning-Go-faster/tree/main/Topics/Installation) and my first [hello world](https://github.com/StrandedDev/Learning-Go-faster/blob/main/Topics/Installation/hello_world.go) file.


## Go: Variables

- Variables can be declared by `var` or `const` or `:=` keywords
- Once defined, the program has to use the variable, unused variables will throw errors in compile time.
- Const variables cannot be overwritten after initialization. 
- In loose typing, Go will assign variable type based on the type of value assigned to the variable.

### Variable syntax

```go
var variableName type = value
```

**General Example:** 

```go

var age int = 24            // Strict type - type declared
var name = "John doe"       // Loose type - type not declared - Inferred to string
const light bool = false    // Constant will have strict type
education := true           // Does not require keyword or type declaration

```

- Go encourages to explicitly declare the data type of variable while declaration
- Go supports both loose typing and strict typing method 
- Go can handle multiple variable declaration with comma separation in single line
- Go also supports multi-line variable declaration with `var()` method. (Type must be declared explicitly)

**Extended examples:**

```go

var a,b,c int = 1,2,3   // Same type, Multiple variables
d,e,f := 4,"John",true  // Different type, loose typing


var(
    a int = 5           // Integer 
    b string = "John"   // String
    c bool              // Initialized boolean
)

c = false               // Assigned outside 



```

**To check more details about variables:** <br/>

| [Variables folder &rarr;](Topics/Variables) | [Variables notes &rarr;](Topics/Variables/variables_notes.md) |
|:-:|:-:|






    
## Authors

[Stranded Dev](https://github.com/StrandedDev)



