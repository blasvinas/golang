- [Go Programming Language](#go-programming-language)
  - [Simple Go program](#simple-go-program)
# Go Programming Language
## Simple Go program
Let's start using the classic hello world program to explain the basic syntax of a go program.
```go
package main

import "fmt"

func main() {
 fmt.Println("Hello, Gophers!")
}
```

Even though this is a very short program, it contains most of the go programs' typical layout. Let's check the program line by line.

* go package main. Every go file starts with a package clause. A package is a collection of code that is related. For example, the package math contains a collection of functions related to math, the package strings includes code on string manipulation, etc. When you create your go files, the convention is to follow the package clause with the file's name (without the extension). In this example, the file is named main.go, so we start with package main. All the code in this file belongs to the main package.
* import "fmt." This line import the fmt package, which contains the Println function. Go does not allow to import packages that are not used in the program. Most editors or IDEs will automatically delete the import line if the package is not used.
* func main {. main is a special function in go, and it marks the entry point of the program. The program that you run from a terminal must have the main function.
* Curly braces {}. Go uses curly braces to mark a block of code. On other programming languages, you can write the main function as follow.

```go
func main()
{
 fmt.Println("Hello, Gophers!")
}
```
This will cause the following error in go at compile time: "syntax error: unexpected semicolon or newline before {. " The opening curly brace "{" must be on the same line as func main(). To understand why this error happens, let's talk about semicolons. As you can see, you do not need to end a line with a semicolon as you do in other languages. You can add the semicolons, and the program will compile, but that is not considered a best practice in go, and most likely, your editor will remove the semicolons if you try to add them. When you compile the program, the compiler will add the semicolons at the end of each line ( you won't see this in the source code), so the function above will be like this.
```go
func main();
{
 fmt.Println("Hello, Gophers!");
}
```
As you can see, that will create a function without a body, so you get the error.