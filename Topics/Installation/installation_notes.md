
# Learning Go Lang Faster

## Hello world

I've written my first hello world program in Go. There is some keypoints i've noticed. 


- I have to declare the _main_ package on the top of the file that will ultimately call the **main** function. The code i will write will be inside the **main** function. This _main_ defines my current file/package. The main function/package name cannot be changed.

- The keyword to declare a function is _func_

```go
package main

func main(){

}
```


- To diplay output, I have to _import_ a Go package named **"fmt"** (format package) that stores the **println()** function

- There are also more functions to print data such as **print()** , **printf()** .

- It is not necessary to put semicolon after every line but every Go statement can be separated using a semicolon. In that case we can trim off all the whitespace and also create something like this **one liner**:

```go
package main; import ("fmt"); func main() { fmt.Println("Hello World!");}
```


## Running and Building files 

- After writing the Go source code, save the file.

- To **run** the code:

```bash
go run hello_world.go
```

or to **build** an executable file: 

```bash
go build hello_world.go
```


## Commenting 

Comments are just small lines inside our code that is intented not to get read by the compiler. 

- For single line comments:

```go
// This code will be ignored by the compiler
```

- For multi-line comments:

```bash
/*
These lines
will be
ignored
by the
compiler
*/
```

> **Tip for vscode**: Use `ctrl` + `/` to comment out the whole line from where the cursor is at. 

## Next Section: Variables

[Jump to variables notes &rarr;]()