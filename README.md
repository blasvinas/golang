- [Go Programming Language](#go-programming-language)
- [Setting up Go](#setting-up-go)
  - [Installing Go](#installing-go)
  - [Organizing Your Projects](#organizing-your-projects)
- [A Simple Go program](#a-simple-go-program)
  - [Go Standard format](#go-standard-format)
  - [Go Style](#go-style)
  - [Running the program program](#running-the-program-program)
- [Data Types](#data-types)
  - [Booleans](#booleans)
  - [Integers](#integers)
  - [Floating Point](#floating-point)
  - [Complex Numbers](#complex-numbers)
  - [Strings](#strings)
  - [Runes](#runes)
- [Operators](#operators)
  - [Arithmetic Operators](#arithmetic-operators)
  - [Comparison Operators](#comparison-operators)
  - [Logical Operators](#logical-operators)
- [Variables](#variables)
- [Constants](#constants)
- [Arrays](#arrays)
- [Slices](#slices)
  - [append](#append)
  - [len](#len)
  - [cap](#cap)
  - [make](#make)
  - [Slicing](#slicing)
- [Maps](#maps)
  - [delete](#delete)
- [Structs](#structs)
- [if](#if)
- [for](#for)
- [switch](#switch)
- [Functions](#functions)
  - [Variadic Parameters](#variadic-parameters)
  - [Returning Multiple Values](#returning-multiple-values)
  - [Functions as Values](#functions-as-values)
  - [Closures](#closures)
  - [defer](#defer)
- [Pointers](#pointers)
- [Methods](#methods)
- [iota](#iota)
- [Composition](#composition)
- [Interfaces](#interfaces)
  - [The Empty Interface](#the-empty-interface)
  - [Type Assertions](#type-assertions)
  - [Type Switches](#type-switches)
- [Errors](#errors)
  - [How to Handle Errors](#how-to-handle-errors)
  - [panic and recover](#panic-and-recover)
- [Repositories,  Modules, and Packages](#repositories--modules-and-packages)
  - [Create a Go Module](#create-a-go-module)
  - [Building a Package](#building-a-package)
    - [Using a Package](#using-a-package)
  - [Package Name Conflicts](#package-name-conflicts)
- [Concurrency](#concurrency)
  - [Goroutines](#goroutines)
  - [Channels](#channels)
  - [select](#select)
# Go Programming Language

# Setting up Go

## Installing Go

You can find the latest version of Go and instructions on installing it on your operating system in https://golang.org/dl/.  

The installation program should remove any previous version of  Go, install Go in the right location, and add to go command to the path.

To verify that Go has been installed correctly, run the following from the command line.

```shell
go version
go version go1.16.2 darwin/amd6
```

You should get an output similar to the one above,  displaying the version of Go installed.

## Organizing Your Projects

Over the years, the way Go projects are organized has changed.  For modern Go projects,  you are free to organize your code any way it works for you.

By default, the third-party tools installed running the command go install are located in \$HOME/go. It is recommended that you set the \$GOPATH environment variable to \$HOME/go.

In Linux systems using bash, add the following lines to your .profile.

```shell
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

If you are using a MAC with the zsh shell, add the line to .zshrc.

On Windows, run the following commands at the prompt:

```shell
setx GOPATH %USERPROFILE%\go
setx path %path%;%USERPROFILE%\bin
```


You can get the go environment running the go env command.

```shell
go env
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/blasvinas/Library/Caches/go-build"
GOENV="/Users/blasvinas/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/blasvinas/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/blasvinas/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GOVCS=""
GOVERSION="go1.16.2"
GCCGO="gccgo"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/dev/null"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/vs/t98y17fn2gsdjllr5zmqg7q00000gn/T/go-build2207945809=/tmp/go-build -gno-record-gcc-switches -fno-common"
```

# A Simple Go program
Let's start using the classic hello world program to explain the basic syntax of a go program.
```go
package main
import "fmt"

func main() {
 fmt.Println("Hello, Gophers!")
}
```

Even though this is a very short program, it contains most of the go programs' typical layout. Let's check the program line by line.

* go package main. Every go file starts with a package clause. A package is a collection of code that is related. For example, the package math contains a collection of functions related to math. The package strings include code on string manipulation, etc. When you create your go files, the convention is to follow the package clause with the file's name (without the extension). In this example, the file is named main.go, so we start with package main. All the code in this file belongs to the main package.
* import "fmt." This line import the fmt package, which contains the Println function. Go does not allow to import packages that are not used in the program. Most editors or IDEs will automatically delete the import line if the package is not used.
* func main {. main is a special function in go, and it marks the entry point of the program. The program that you run from a terminal must have the main function.

## Go Standard format

Go uses curly braces to mark a block of code. On other programming languages, you can write the main function as follow.

```go
func main()
{
 fmt.Println("Hello, Gophers!")
}
```
This will cause the following error in go at compile time: "syntax error: unexpected semicolon or newline before {. " The opening curly brace "{" must be on the same line as func main(). To understand why this error happens, let's talk about semicolons. As you can see, you do not need to end a line with a semicolon as you do in other languages. You can add the semicolons. The program will compile, but that is not considered a best practice in go. Most likely, your editor will remove the semicolons if you try to add them. When you compile the program, the compiler will add the semicolons at the end of each line ( you won't see this in the source code), so the function above will be like this.

```go
func main();
{
 fmt.Println("Hello, Gophers!");
}
```
As you can see, that will create a function without a body, so you get the error.

Go enforces a standard format.  The go fmt command will reformat your code to math the standars.  An enhanced version of go fmt, called goimports, cleans up your import statements in addition to reformating the code.  You can install goimports as shown below.

```shell
go install golang.org/x/tools/cmd/goimports@latest
```

Usually, your editor runs go fmt or goimports when you save the program,  so you do not need to run these commands manually, but you can run goimports manually from the directory where your source code is as follows

```shell
goimports -l -w .
```

Note: go fmt won't fix braces on the wrong line.

## Go Style

You should read Effective Go (https://golang.org/doc/effective_go) to understand how idiomatic Go code looks like.

There are tools to help to enforce the Go style.  One of these tools is golint which you can download as shown below.

```shell
go install golang.org/x/lint/golint@latest
```

Once install it,   you can run with:

```shell
golint ./...
```

This will runs golint over the entire project.

## Running the program program

You can run a Go program in n different ways.  Supposed that you called the program above main.go, you use the go run command as shown below:

```shell
go run main.go           
Hello, Gophers!
```

Using the go run command, the program is compiled to a temporary location, executed, and then the binaries are deleted.

If you want to compile the program ans save the binaries, run go build main.go.  This command will create a binary called main in MacOS and Linux and main.exe in Windows.  You can use the -o flag to create the binary with a different name.  For example:

```shell
go build -o hello main.go
```

This will create a binary named hello (hello.exe in Windows).

You can install a Go program from your git repository in your computer using the go install command as follow:

```shell
go install github.com/git_user/golang/hello_world@latest 
```

This will download the files from the git repository in github.com/git_user/golang/hello_world and installed in \$GOPATH/bin/hello_world


# Data Types
Go has the following basic types.
* Booleans
* Integers
* Floating point
* Strings
* Runes
* Complex numbers

## Booleans

A boolean can only have two values: true or false. The zero value (the default value when the variable is not initialized) is false.  You can declare a boolean variable using the bool type as shown below.

```go
package main

import (
    "fmt"
)

func main() {
    var isOk bool // false (zero value)
    fmt.Println(isOk) // false
    isOk = true
    fmt.Println(isOk)   // true
}
```

## Integers

Go has the following integer types.

| Type | Value Range |
|------|-------------|
| uint8 | 0 to 255 |
| uint16 | 0 to 65535 |
| uint32 | 0 to 4294967295 |
| uint64 | 0 to 18446744073709551615 |
| int8 | -128 to 127 |
| int16 | -32768 to 32767 |
| int32 | -2147483648 to 2147483647 |
| int64 | -9223372036854775808 to 9223372036854775807 |

Go define some aliases to some of of the integer types as described below.

* byte alias to uint8.
* int alias to int32 or int64 depending on the number of bits of the CPU where you are running the program.
* uint alias to unsigned int.

Unless you have a specific need about the size or the sign,  use int when declaring your variables.

```go
package main

import "fmt"

func main() {
    var myNumber int // 0 (zero value)
    myNumber = 5
    fmt.Println(myNumber) // 5
}
```

## Floating Point

Go has two floating point types.

| Type | Value Range |
|------|-------------|
| float32 | 1.401298464324817070923729583289916131280e-45 to 3.40282346638528859811704183484516925440e+38 |
| float64 | 4.940656458412465441765687928682213723651e-324 to 1.797693134862315708145274237317043567981e+308 |

The following is a sample program using float32 and float64

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    var myFloat32 float32 // 0 (zero value)
    var myFloat64 float64 // 0 (zero value)

    myFloat32 = math.MaxFloat32
    fmt.Println(myFloat32) // 3.4028235e+38
    myFloat64 = math.MaxFloat64
    fmt.Println(myFloat64) // 1.7976931348623157e+308
}
```

You float64 unless you have a specific need to use float32.

## Complex Numbers

Go defines two complex number types.

* complex64 uses float32 to represent the real and imaginary parts.
* complex128 uses float64 to represent the real and imaginary parts.

Complex numbers are numbers that consist of two parts, a real number, and an imaginary number.  They are represented in the form a + bi.  Below is an example of a program that declares a complex number.

```go
package main

import (
    "fmt"
)

func main() {
    var myNumber complex64 // (0+0i) (zero value)
    myNumber = complex(3.5, 2.3) 
    fmt.Println(myNumber) // (3.5+2.3i)
}
```

## Strings

In Go, a string is a series of bytes. You can put any Unicode value into a string. Strings are immutable; once created, you cannot change their value.  So in a program, every time you change the content of a string variable, you are creating a new string in reality.  Go uses double-quotes for strings, as shown in the program below.

```go
package main

import (
    "fmt"
)

func main() {
    var message string
    message = "Hello Go programmers!"
    fmt.Println(message) // Hello Go programmers!
}
```

You can use slice expressions with strings the same way you use them with slices (See Slices).  

```go
    s1 := "Hello Go!"
    s2 := s1[0:5]
    fmt.Println(s2) // Hello
```

You can also use the [] to access an individual element of the string, but the results might not be what you expect.

```go
    s := "Hello Go!"
    fmt.Println(s[1]) // 101
```

You might be expecting the output to be 'e', but it is  101, which is the Unicode for 'e'.  Some characters take more than one byte, so if you try to get their value using [], you will get something else.

Remember that strings are immutable, so a program with the following lines won't compile.

```go
    s := "Hello Go!"
    s[1] = 'a'  // Error: cannot assign to s[1] (strings are immutable)
```

## Runes

Runes are used to represent single characters.  The rune type is an alias for the int32 type.  In Go, runes are written using single quotes.

```go
package main

import (
    "fmt"
)

func main() {
    var r rune
    r = 'A'
    fmt.Println(r) // 65
}
```

When you print a run, the output will be the numeric code representing the rune and not the original character. So, the program above with print 65 instead of A.

# Operators

## Arithmetic Operators

Go has the following arithmetic operators.

| Operator | Description | Data Types |
|----------|-------------|------------|
|+ | sum | integers, floats, complex values, strings|
| - |   difference |            integers, floats, complex values |
| * |   product |               integers, floats, complex values |
| / |   quotient |              integers, floats, complex values |
| % |   remainder |             integers |
| & |   bitwise AND |           integers |
| \||    bitwise OR |            integers |
| ^ |   bitwise XOR |           integers |
| &^|   bit clear (AND NOT)|    integers |
| <<|   left shift |            integer << unsigned integer |
| >>|   right shift  |          integer >> unsigned integer |

You can combine the arithmetic operator with = to modify a variable, for example +=, *=, -+, /=, ++, --.  See the example below.

## Comparison Operators

The table below shows the comparison operators supported by Go.

| Operator | Description |
|----------|-------------|
| ==  |  equal |
| !=  |  not equal |
| <   |  less |
| <=  |  less or equal |
| >  |   greater |
| >= |   greater or equal| 

## Logical Operators

Below is a list of the logical operators

| Operator | Description |
|----------|-------------|
| && | AND |
| \|\| | OR |
| ! | NOT |

```go
package main

import "fmt"

func main() {
    //Arithmetic operators
    a, b := 4, 2
    r := (a + b) / (a - b) * 2
    fmt.Println(r) //6

    r = 9 % a
    fmt.Println(r) // 1

    //Assignment operators
    a, b = 2, 3
    a += b
    fmt.Println(a) // 5

    b -= 2
    fmt.Println(b) // 1

    b *= 10
    fmt.Println(b) // 10

    b /= 5
    fmt.Println(b) // 2

    a %= 3
    fmt.Println(a) // 2

    x := 1
    x++
    fmt.Println(x) // 2

    x--
    fmt.Println(x) // 1

    //Comparison Operators
    a, b = 5, 10
    fmt.Println(a == b) // false
    fmt.Println(a != b) // true
    fmt.Println(a > 5) // false
    fmt.Println(a >= 5) // true
    fmt.Println(b < a, 10 <= b) // false true

    //Logical Operators
    a, b = 5, 10
    fmt.Println(a > 1 && b == 10) // true
    fmt.Println(a == 5 || b == 100) // true
    fmt.Println(!(a > 0)) // false
}
```

# Variables

In Go, you have several ways to declare a variable. The most verbose way is using the var keyword, explicitly specify the type, and assign a value. 

```go
var a int = 5
```

The previous example declares a variable named a as an integer and assigns 5 to it.

When you assign a value to a variable during the declaration, you can omit the type, and Go will infer the type from the value. You can rewrite the line above as follow.

```go
var a = 5
```

If you just need to declare a variable;e and assign a value later, you need to specify the type. The variable will be initialized to the zero value.

```go
var a int // initialized to 0 (zero value)
```

You can declare several variables in one line.

```go
var x, y float64 // x = 0.0, y = 0.0
var a, b = 4, 3 // a = 4, b = 3
var c, d = "one", "two" // c = "one", d = "two"
```

If you need to declare multiple variables simultaneously, you can use the syntax shown below.

```go
var (
 x int
 y float64
 i int = 5
 d = "Hello"
)
```

Another way to declare a variable is using the := operator. 

```go
i := 5
x := 6.3
d := "hello"
a, b := 4, 3
```
You can only use the := operator inside a function. You cannot use it to declare variables at the package level. Also, using :=, you can assign values to an existing variable as long as one variable on the lefthand side of the := is new.

```go
x = 5
x, y := 3, 2
a = 4
y, a = 6, 8 // Error, a and y already exist
```

In Go, you must use a variable. If you declare a variable in a function and don't use it, you will get an error at compilation time.

```go
// This program won't compile
package main

import (
 "fmt"
)

func main() {
 var a int
 b := 5
 fmt.Println(b)
}
```

If you try to compile this program, you will get the following error: a declared but not used.

# Constants

You can declare a constant using the const keyword.

```go
const a = 5
a = 3 // Error: cannot assign to a (declared const)
```

In Go, const can only hold values that are known at compilation time.

When you declare a constant without specified a type, the constant is untyped, for example.

```go
const a = 5
var x float64 = a
var y int = a
```

If you declare a constant with a type, it can only be assigned to a variable of the same type.

```go
 const a int = 5
 var x float64 = a // Error: cannot use a (type int) as type float64 in assignment
 var y int = a
```

# Arrays

Most of the time, you will not use arrays directly in go since they have several limitations.  For the vast majority of the cases, use slices (which we will explain in the next section) instead of arrays.

Below are several ways you can use to declare an array.

```go
var a[5] int // array of 5 integers each one initialized to 0
    var a = [3]int{1, 2, 3} //arrray of 3 integers wich values of 1,2, and 3
    var b = [...]float64{1.1, 5.3, 7.8}  // Array of 3 float64 numbers.
    var c = [10]int{1,2,5:8,8:2} //{1,2,0,0,8,0,0,2,0,0}
```
In the last declaration, you can specify the index to be initialized using index:value.   The elements not defined will be initialized to the zero value.

When you declare an array, the array's size is part of the type, so [5]int is not the same type as [3]int.

You can access the elements of an array using the brackets and the index.

```go
a[5] = 8
```

The length of the array must be known at compilation time, so you cannot use a variable to define the number of elements.  The following won't compile.

```go
package main

import (
    "fmt"
)

func main() {
    var len = 3
    var a[len]int // error, you can not use a variable to define the number of elements.
    fmt.Println(a)
}
```

The following program shows how to declare and use an array.

```go
package main

import (
    "fmt"
)

func main() {
    var x [5]int

    x[0] = 1
    x[1] = 3
    x[2] = 5
    x[3] = 7
    x[4] = 11
    fmt.Println(x) // [1,3,5,7,11]
}
```

You can have an array containing another array.

```go
package main

import (
    "fmt"
)

func main() {
    var x = [2][3]int{{1, 2, 3}, {4, 5, 6}}
    fmt.Println(x[1][2]) // 6

}
```

# Slices

Slices look similar to arrays, but they do not have the limitations of arrays like the fixed size, for example.  Below is an example of declaring slices.

```go
    var x[]int
    var y = []int{1,2,3}
```

From the previous example, x is initialized to nil, which is the zero value of a slice. The nil value means that the slice does not contain anything.

Notice that the declaration of the slice y is very similar to the way we declare arrays, but without the ... inside the brackets.

## append

You can add values at the end of a slice using append. Append takes two parameters: the name of the array and the value or values to add.

```go
append(x,5)
append(y,7,8,9)
```

Also, you can append one slice to another slice using the ... operator.

```go
x := []int{1,2,3}
y = append(y, x...)
```

## len

You can get the number of elements in a slice using len.

```go
x := []int{1,2,3}
fmt.Println(len(x)) // 3
```

## cap

Elements in a slide are store in consecutive memory locations. The cap function returns the capacity, which is the number of consecutive memory allocations that have been reserved.  Capacity is different than length, which is the current number of elements in the slice.  When the length reaches the capacity, there is no more room to add elements, so Go will create a new slice with a bigger capacity and copy the original slice's values.

When Go needs to increase the capacity, it will double the size if the capacity is less than 1024, and it will increase 25% after that.

The following program shows the difference between length and capacity.

```go
package main

import (
    "fmt"
)

func main() {
    x := []int{1,2,3}
    fmt.Println(len(x), cap(x)) // 3 3
     x = append(x,4)
    fmt.Println(len(x), cap(x)) // 4 6
    x = append(x,5,6)
    fmt.Println(len(x), cap(x)) // 6 6
    x = append(x,7)
    fmt.Println(len(x), cap(x)) // 7 12

}
```

After the slice initialization, both the length and the capacity are the same, 3.  Since length reached the capacity, when we add a value to the slice, the capacity needs to be increased.  Following the rules explained below, the capacity is doubled, so now the length is 4, and the capacity is 6.  After adding two more elements to the slice, the length and the capacity are 6.  At this point, again, the length reached the capacity. When we add another element, the capacity needs to be increased (doubled), so the last Println will display 7 for the length and 12 for the capacity.

##  make

You can define the initial length and optionally the capacity using the make function.

```go
x := make([]int, 10)
y := make([]int,10, 15)
```

The first line creates a slice with a length and capacity of 10.  The second line creates a slice with a length of 10 and a capacity of 15.

## Slicing

You can create a new slice from an existing one with slice_name[start:end].

```go
    x := []int{1,2,3,4,5,6}
    y := x[2:4]
    fmt.Println(y) //  [3, 4]
```

In the example above, x[2:4] will return a slice from position 2 until position3 of x.  The ending offset in [start:end] is not included.

Notice that slices created with slicing share the memory with the original slice, so if you make any changes in one of the slices, the changes will be reflected on the other.

```go
    x := []int{1,2,3,4,5,6}
    y := x[2:4]
    fmt.Println(y) //  [3, 4]
    x[3] = 40
    fmt.Println(x) // [1 2 3 40 5 6]
    fmt.Println(y) // [3 40]
    y[0] = 30
    fmt.Println(x) // [1 2 30 40 5 6]
    fmt.Println(y) // [30 40]
```

If you need to create a slice that does not share the memory with the original slice,  use the copy function.  If you modify the previous example, replacing the line y := x[2:4] with lines y := make([]int,2) and copy(y,x[2:4]), now y is independent of x, and the changes on one slice do not affect the other.

```go
package main

import (
    "fmt"
)

func main() {
    x := []int{1,2,3,4,5,6}
    y := make([]int,2)
    copy(y,x[2:4])
    fmt.Println(y) //  [3 4]
    x[3] = 40
    fmt.Println(x) // [1 2 3 40 5 6]
    fmt.Println(y) // [3 4]
    y[0] = 30
    fmt.Println(x) // [1 2 3 40 5 6]
    fmt.Println(y) // [30 4]
}
```

# Maps

Use maps when you need to associate one value with another.  Maps are key-value pairs.  You can declare a map using map[key_type]value_type.

```go
    numbers := map[string]int{}
    numbers["one"] = 1
    numbers["two"] = 2
    numbers["three"] = 3
```

Also, you can assign values when declaring the map.

```go
numbers := map[string]int{"one": 1, "two":  2, "three": 3}
```

If you know how many pairs your map has, you can declare the map using the make function.

```go
numbers := make(map[string]int,5)
```

The previous example creates a map with a default size of 5.

If you access a non-existing key, it will return the zero value, for example.

```go
    numbers := make(map[string]int,3)
    numbers["one"] = 1
    numbers["two"] = 2
    numbers["three"] = 3
    
    fmt.Println(numbers["two"]) // 2
    fmt.Println(numbers["four"]) // 0
```

Since there is no element with the "four" key, zero is displayed if you try to access the map using this key.  There are times when you need to know if the key didn't exist or if the value was zero.  In these cases, you can use the comma ok idiom.

```go
    numbers := map[string]int{}
    numbers["one"] = 1
    numbers["two"] = 2
    numbers["three"] = 3
    
    value, ok := numbers["four"]
    fmt.Println(value, ok) // 0 false
```
Println displays 0 false.  The ok variable is true if the key exists or false if it doesn't.

## delete

You can delete an element from the map using the delete function.  The delete function takes two parameters:  the map and the key to delete.

```go
    numbers := map[string]int{}
    numbers["one"] = 1
    numbers["two"] = 2
    numbers["three"] = 3
    delete(numbers, "two")
    fmt.Println(numbers) // map[one:1 three:3]
```

In the example above, we deleted the element with the key "two."

# Structs

Use a struct when you have related data that needs to be group together. The example below shows the syntax to declare a struct in Go.

```go
type car struct {
    year int
    make string
    model string
}
```

Here, we declared a struct named car that has three variables.

```go
    type car struct {
        year  int
        make  string
        model string
    }

    var car1 car
    car2 := car{2020, "Infinity", "Q50"}
    car3 := car{
        model: "Altima",
        make: "Nissan",
        year: 2017,
    }
    

    car1.year = 2018
    car1.make = "Ford"
    car1.model = "Explorer"

    fmt.Println(car1) // {2018 Ford Explorer}
    fmt.Println(car2) // {2020 Infinity Q50}
    fmt.Println(car3) // {2017 Nissan Altima}
```

The previous example shows how to define a struct type and several ways to declare and initialize the struct's variables.

There is another way to declare a struct named anonymous struct, shown in the example below.

```go
    car1 := struct {
        year  int
        make  string
        model string
    }{
        model: "Altima",
        make:  "Nissan",
        year:  2017,
    }

    fmt.Println(car1) // {2017 Nissan Altima}
```

# if

The if statement in Go is similar to other languages.  One difference is that Go does not use parenthesis around the condition.

```go
	grade := 85
	if grade >= 90 {
		fmt.Println("A")
	} else if grade >= 80 {
		fmt.Println("B")
	} else if grade >= 70 {
		fmt.Println("C")
	} else if grade >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
```

If you  want a variable to be local just to the if stament you can declare it in the if stament as shown below.

```go
	if grade := 85; grade >= 90 {
		fmt.Println("A")
	} else if grade >= 80 {
		fmt.Println("B")
	} else if grade >= 70 {
		fmt.Println("C")
	} else if grade >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
    // fmt.Println(grade). Error: grade is not defined.
```
If you uncomment the last line, the program won't compile becuase grade is not defined. 

# for

In Go, for is the only loop statement. Go does not support while or do-while like other languages.  However, you can use the for statement in different ways that accomplish the same results than the missing loop statements.

The first loop statement that we are going to check is the c-style for.  See example below.

```go
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
```

The previous example will print the numbers from 0 to 4.  Notice that there are not parenthesis.

If you simulate  the behaviour of a while loop with a confition-only for loop.  The next lines of code produce the same results that the previious example.

```go
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
```

If you use for without any condition, you create a infinite loop.
```go
	for {
		fmt.Println("this will run forever!")
	}
```

You can use the break statement to end a infinite for for, or any  for loop.

```go
	i:=1
	for {
		fmt.Println("this will run forever!")
		i++
		if i > 10 {
			break
		}
	}
```

Also, you can use the continue stament if you want to skip the rest of the body in the for loop and continue with the next iteration.  The next example print the odd numbers between 0 and 10.

```go
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
```

One more way of the for statement is the for-range statement.  The for-range statement is use to iterate over collections like arrays, slices, maps, etc.  The next example iterate over a slice, in this  case  range returns the index and the value.

```go
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range x {
		fmt.Println(i, v)
	}
```

The lines above print 2 numbers per line.  The first number is the index and the second is the value.

```shell
0 1
1 2
2 3
3 4
4 5
5 6
6 7
7 8
8 9
```

If you only need the value, use an undersore instead of i as follow.

```go
	for _, v := range x {
		fmt.Println(v)
	}
```

The following example iterate over a map.  In this case range returns the key and the value.

```go
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}

	for k, v := range numbers {
		fmt.Println(k, v)
	}
```
The output is similar to the one  below.

```shell
one 1
two 2
three 3
```

If you are interested only on the values, use an underscore instead of k.

```go
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}

	for _, v := range numbers {
		fmt.Println(v)
	}
```

If you need only the key, you can leave the second variable off.

```go
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}

	for k := range numbers {
		fmt.Println(v)
	}
```

# switch

Sometimes, when you have several if-else if statements, it is better to use a switch statement.

```go
 dayOfWeek := 6
 switch dayOfWeek {
 case 1:
 fmt.Println("Monday")
 case 2:
 fmt.Println("Tuesday")
 case 3:
 fmt.Println("Wednesday")
 case 4:
 fmt.Println("Thursday")
 case 5, 6, 7:
 fmt.Println("Weekend!!!")
 default:
 fmt.Println("Invalid day")
 }
```

As you can see in the previous example, in Go, there is no need to use the break statement since switch statements do not fall through in Go. You can still use break if you want to exit early from a case.

You can combine multiple matches, separating them with commas, like the case for the Weekend in the previous example.

Also, cases can use any logical tests, as shown below.

```go
 dayOfWeek := 6
 switch {
 case dayOfWeek > 5:
 fmt.Println("I am freeeee!!!!!")
 case dayOfWeek >= 4:
 fmt.Println("I can see the light at the end of the tunnel")
 case dayOfWeek >= 1:
 fmt.Println("is it Friday yet?")
 }
```

In the previous example, you can declare the dayOfWeek variable inside the switch as shown next.

```
switch dayOfWeek := 6; {
 ...
}
```

# Functions

Below is the syntax to declare a function in Go.

```go
    func name(arguments) return_type {
        ...
        return ...
    }
```

The example below shows how to declare and call a function.

```go
package main

import (
    "fmt"
)

func main() {
    x := sum(8,6)
    fmt.Println(x) // 14
}

func sum(a int, b int) int {
    return a + b
}
```

In the previous example, we declare a function named sum, which takes two integer parameters and returns an integer with the two numbers' addition. 

Since the parameters, a and b, have the same type, you could have declared the function as follow.

```go
func sum(a , b int) int {
    return a + b
}
```

Unlike other languages, Go does not have neither named nor optional parameters.

## Variadic Parameters

Continuing with the previous example, what if we do not know how many numbers we need to add?  In those cases, we can use variadic parameters, as shown below.

```go
package main

import (
    "fmt"
)

func main() {
    x := sum(8, 6)
    y := sum(6,5,9)
    z := sum(7,5,3,4)
    fmt.Println(x) // 14
    fmt.Println(y) // 20
    fmt.Println(z) // 19
    
}

func sum(a, b int, numbers ...int) int {
    sum := a + b
    for _, v := range numbers {
        sum += v
    }
    return sum
}
```

As you can see in the example, you can define a variable number of parameters using the syntax ...type.  Now you can call sum with a different number of parameters as shown in the example.

## Returning Multiple Values

Functions in Go can return multiple values.  Below is the syntax.

```go
    func name(arguments) (return_type1, return_type2, ...) {
        ...
        return value1, value2, ...
    }
```

The function in the following example returns the addition and the subtraction of two integers.

```go
package main

import (
    "fmt"
)

func main() {
    x, y := addAndSubtract(8, 6)

    fmt.Println(x) // 14
    fmt.Println(y) // 2
}

func addAndSubtract(a, b int) (int, int) {
    sum := a + b
    diff := a - b
    return sum, diff
}
```

You can ignore a return value using an underscore.  In the previous example, if you want to ignore the value for the sum, you can call the function as follow.

```go
_,y := addAndSubtract(8, 6)
```

You can name the return the return values as shown below.

```go
func addAndSubtract(a, b int) (sum int, diff int) {
    sum = a + b
    diff = a - b
    return sum, diff
}
```

When you do that, you do not need to declare them inside the function, and they will be initialized to the zero value.

## Functions as Values

In Go,  functions are similar to other values and you can use them in collections or pass them to other functions as long as the signature of the function match.  The signature of a function is the combination of the type of the parameters and the return value.


```go
package main

import "fmt"

func main() {
    ops := []func(int, int) int{add, sub, mul, div}

    for _, op := range ops {
        fmt.Println(op(6, 3))
    }
}

func add(a int, b int) int {
    return a + b
}

func sub(a int, b int) int {
    return a - b
}

func mul(a int, b int) int {
    return a * b
}

func div(a int, b int) int {
    if b != 0 {
        return a / b
    }
    return 0
}
```

This program creates a list of functions that take two integers as parameters and return an integer value.  Then, the program iterates over the list and call each function.

You can create a type with the function signature, so in the previous program, you can declare the list of functions as follow.

```go
    type opFunc func(int, int) int
    ops := []opFunc{add, sub, mul, div}
```


## Closures

A closure is a function declared inside another function that can access the variables of the outer function.  Closures are useful when they are passed to other functions or when they are returned by other functions. Below is an example that uses closures.

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    type Employee struct {
        name   string
        salary float64
    }

    employees := []Employee{
        {"John", 50000},
        {"Nancy", 75000},
        {"Mary", 45000},
        {"Fred", 70000},
    }

    fmt.Println(employees) // [{John 50000} {Nancy 75000} {Mary 45000} {Fred 70000}]

    //Sort by salaries
    sort.Slice(employees, func(i int, j int) bool {
        return employees[i].salary < employees[j].salary
    })
    fmt.Println(employees) // [{Mary 45000} {John 50000} {Fred 70000} {Nancy 75000}]
    
    // Sort by names
    sort.Slice(employees, func(i int, j int) bool {
        return employees[i].name < employees[j].name
    })
    fmt.Println(employees) // [{Fred 70000} {John 50000} {Mary 45000} {Nancy 75000}]
}
```

The sort.Slice function takes a slice and a closure with the logic to compare the values and sort the slice.  As you can see, the closure has access to the employees slice, which is in the outer function.

The following example shows a function that returns a closure.

```go
package main

import (
    "fmt"
)

func seq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextNumber1 := seq()
    fmt.Println(nextNumber1()) // 1
    fmt.Println(nextNumber1()) // 2
    fmt.Println(nextNumber1()) // 3
    
    nextNumber2 := seq()
    fmt.Println(nextNumber2()) // 1


}
```

In the previous example, the seq() function returns a closure that increments an integer and returns it.  Notice that the closure has access to i which is declared in the outer function.  When we assign seq() to nextNumber1, nextNumber1 captures the value of the variable i.  nextNumber1 will remember the previous value of i and increment it.  When we assign seq() to nextNumber2, it will capture its own i value.

## defer

Usually, you use defer when you need to release resources like files, network connections, database connections, etc.  When you call a function using defer, it won't be executed until the calling function ends.  In the example below, the closeDB function will be called when the addRecord() function exits.

```go
package main

import (
    "fmt"
)

func closeDB() {
    fmt.Println("Closing the database")
}

func addRecord() {
    defer closeDB()
    fmt.Println("Opening the database")
    fmt.Println("Add a new record")
}

func main() {
    addRecord()
}
```

Instead of having a closeDB() function, you can use a closure, as shown below.

```go
package main

import (
    "fmt"
)

func addRecord() {
    defer func() {
        fmt.Println("Closing the database")
    }()
    fmt.Println("Opening the database")
    fmt.Println("Add a new record")
}

func main() {
    addRecord()
}
```

# Pointers

A pointer contains the address where a variable is stored.  You declare a variable of pointer type using a * before the type.  To get the address of a variable, you can use the address operator, &.  To get the content of the address referenced by a pointer, use the indirection operator, *.


```go
    var p *int  // Pointer to an integer
    x := 5

    p = &x // assign to p the location in memory of x
    fmt.Println(p). // Something similar to this:  0xc0000be010
    fmt.Println(*p) // 5
```

In Go, parameters a passed by value, which means that the parameter is a copy of the original value,  so if you modify the parameter inside the function, that won't affect the original value.  This is shown with update1 in the example below. After calling update1, the value of x is still the same, 5.  If you need to change the original value, you need to pass a reference as a parameter, as shown with update2.   After calling update2, the value of x will change to 10.

```go
package main

import (
    "fmt"
)

func update1(x int) {
    x = 10
}

func update2(x *int) {
    *x = 10
}

func main() {
    x := 5
    update1(x)
    fmt.Println(x) // 5

    update2(&x)
    fmt.Println(x) // 10

}
```

The zero value of pointers is nil, which means no value has been assigned to a variable yet.

# Methods

Methods are similar to functions, but they work on user-defined types.  The following example shows an Employee struct and a method named Info() to display the employee information.

```go
package main

import (
    "fmt"
)

type Employee struct {
    firstName string
    lastName  string
    salary    float64
}

func (e Employee) Info() string {
    return fmt.Sprintf("%s %s %f", e.firstName, e.lastName, e.salary)
}

func main() {
    e := Employee{
        firstName: "Joe",
        lastName:  "Doe",
        salary:    50000,
    }

    fmt.Println(e.Info())
}
```

As you can see, the method Info() looks very similar to a function, with the only difference being the receiver.  The receiver appears after func and before the name of the method.  In the example above, the receiver is (e Employee).  The receiver is what identifies that the Info() method work on the Employee type.

Let's add the following method to increase the employee's salary.

```go
func (e Employee) IncreaseSalary(amount float64) {
    e.salary += amount
}
```

Add the line to call IncreaseSalary in main(), as shown below.

```go
func main() {
    e := Employee{
        firstName: "Joe",
        lastName:  "Doe",
        salary:    50000,
    }
    fmt.Println(e.Info())

    e.IncreaseSalary(20000)
    fmt.Println(e.Info())
}
```

If we run this program, we get the following output.

```shell
Joe Doe 50000.000000
Joe Doe 50000.000000
```

So,  the salary is still 50000 instead of 70000.  As explained when we discussed pointers, the reason is that the receiver in IncreaseSalary is passed by value, which means that the method makes a copy of the struct and changes the copy without affecting the original struct.  If we need to change the original struct, we need to declare the receiver as a pointer, as shown below.

```go
func (e *Employee) IncreaseSalary(amount float64) {
    e.salary += amount
}
```

Now, if you execute the program again, you get the following output.

```shell
Joe Doe 50000.000000
Joe Doe 70000.000000
```

Below is the complete program.

```go
package main

import (
    "fmt"
)

type Employee struct {
    firstName string
    lastName  string
    salary    float64
}

func (e *Employee) Info() string {
    return fmt.Sprintf("%s %s %f", e.firstName, e.lastName, e.salary)
}

func (e *Employee) IncreaseSalary(amount float64) {
    e.salary += amount
}

func main() {
    e := Employee{
        firstName: "Joe",
        lastName:  "Doe",
        salary:    50000,
    }
    fmt.Println(e.Info())

    e.IncreaseSalary(20000)
    fmt.Println(e.Info())
}
```

Notice that in main(), when we call SalaryIncrease, we use the variable e, which is a value type instead of a pointer. Go automatically converts it to a pointer type.  In this case, e.IncreaseSalary is converted to (&e).IncreaseSalary.

Also, notice that we change the Info() method to take a pointer for the receiver even though it does not need to change the value.  Go best practices suggest that if you change any methods to take a pointer receiver, change all methods to use pointer receiver for consistency.


# iota

You can define a block of constants using iota, as shown in the program below. Iota gives the first constant the value of zero and increments the value by one for each subsequent constant.  In this example, we do not care about the value of zero, so we use _ to ignore it.

```go
    type DaysOfWeek int

    const (
        _ DaysOfWeek = iota
        Monday      // 1
        Tuesday     // 2
        Wednesday   // 3
        Thursday    // 4
        Friday      // 5
        Saturday    // 6
        Sunday      // 7
    )
    
    fmt.Println(Sunday). // 7
```

# Composition

Go favors composition over inheritance.  The following program shows how to use composition.

```go
package main

import (
    "fmt"
)

type Person struct {
    firstName string
    lastName  string
}

func (p Person) PersonInfo() string {
    return fmt.Sprintf("Name: %s %s.", p.firstName, p.lastName)
}

type Employee struct {
    Person
    salary float64
}

func (e Employee) SalaryInfo() string {
    return fmt.Sprintf("Salary %f", e.salary)
}

func main() {

    e := Employee{
        Person: Person{
            firstName: "Joe",
            lastName:  "Doe",
        },
        salary: 50000,
    }

    fmt.Println(e.PersonInfo())
    fmt.Println(e.SalaryInfo())
}
```

Note that Employee contains a field of type Employee.  All the fields and methods of Person can be accessed by Employee.

# Interfaces

An interface is an abstract type that defines a list of methods that must be implemented by a concrete type that meets the interface. In Go, the name of the interface usually ends with "er." The methods defined by an interface are called the method set of the interface.

```go
type Shaper Interface {
    Area() float64
}
```

In Go, interfaces are implemented implicitly. A concrete type does not declare that it implements the interface.  If the concrete type contains all the interface methods, the concrete type implements the interface.  If a concrete type implements an interface, it can be assigned to a variable of the interface type.

```go
package main

import (
    "fmt"
)

type Shaper interface {
    Area() float64
}

type Square struct {
    size float64
}

func (s Square) Area() float64 {
    return s.size * s.size
}

type Rectangle struct {
    length float64
    width  float64
}

func (r Rectangle) Perimeter() float64 {
    return (2 * r.length) + (2 * r.width)
}

func main() {
    s := Square{
        size: 5,
    }

    r := Rectangle{
        length: 10,
        width:  5,
    }

    var i Shaper = s

    fmt.Println(i.Area())
    
    i = r // cannot use r (type Rectangle) as type Shaper in assignment: Rectangle does not implement Shaper (missing Area method)
    fmt.Println(i.Perimeter())

}
```

In the previous example, the Square type implements the interface because it defines the Area method. Square can be assigned to a variable of the interface type Shaper. In the main() function above,  the Square variable s  is assigned to the Shaper interface variable i.  When we try to do the same with the Rectangle variable, we get an error because Rectangle does not implement the interface since it does not define the Area() method.

## The Empty Interface

An empty interface type can store a type that implements zero or more methods.   In other words, an empty interface can store a value of any type.

```go
    var i interface{}
    
    i = 5
    fmt.Println(i)
    
    i = "Hello World"
    fmt.Println(i)
    
    i = 7.8
    fmt.Println(i)
    
    i = []int {1,2,3}
    fmt.Println(i)
```

The variable i is an empty interface, and as you can see in the example, it can store any type.

## Type Assertions

A type assertions check if an interface has a specific concrete type.  The syntax is interface.(type).

```go
package main

import (
    "fmt"
)

type Shaper interface {
    Area() float64
}

type Square struct {
    size float64
}

func (s Square) Area() float64 {
    return s.size * s.size
}

type Rectangle struct {
    length float64
    width  float64
}

func (r Rectangle) Perimeter() float64 {
    return (2 * r.length) + (2 * r.width)
}

func main() {
    s := Square{
        size: 5,
    }

    r := Rectangle{
        length: 10,
        width:  5,
    }

    var i interface{}

    i = s

    i2, ok := i.(Square)

    if !ok {
        fmt.Println("i2 does not have a Square type")
    } else {
        fmt.Println(i2.Area())
    }

    i = r
    i3, ok := i.(Square)

    if !ok {
        fmt.Println("i3 does not have a Square type")
    } else {
        fmt.Println(i3.Area())
    }
}
```

In this example, the first type assertion i.(Square) is correct, and i2 is of type Square, but the second type assertion is wrong because i is storing r, which is not the type Square, i3 is not of type Square, and a message is displayed.


## Type Switches

You can use a type switch to verify the type of the interface, as shown below.

```go
package main

import "fmt"

func checkType(i interface{}) {
    switch j := i.(type) {
    case nil:
        fmt.Println("i is nil")
    case int:
        fmt.Println("i is int")
        fmt.Println(j)
    case float64:
        fmt.Println("i is float64")
        fmt.Println(j)

    case string:
        fmt.Println("i is string")
        fmt.Println(j)
    }
}

func main() {
    x := 6.3
    checkType(x)
}
```

This program will display:

```shell
i is float64
6.3
```

# Errors

## How to Handle Errors

If there is an error inside a function in Go, the last return value should be a value of type error.  If there are no errors, nil is returned.

```go
package main

import (
    "errors"
    "fmt"
    "os"
)

func division(numerator float64, denominator float64) (float64, error) {
    if denominator == 0 {
        return 0, errors.New("denominator cannot be zero")
    }
    return numerator / denominator, nil
}

func main() {
    result, err := division(5.2, 2)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)

    }
    fmt.Println(result) // 2.6

    result, err = division(8.7, 0)

    if err != nil {
        fmt.Println(err) // denominator cannot be zero
        os.Exit(1)

    }
    fmt.Println(result) // 2.6
}
```

In the example above, the function division() checks if the denominator is zero. If that is the case, it will return an error for its last value; otherwise, it will return nil.  The first time division is called inside main() it is successful, and the result of the division is displayed.  The second time the denominator was zero, so an error is displayed, and the program exists.

The built-in error interface defines the Error() method as shown below.

```go
type error interface {
    Error() string
}
```

When you create an error by calling the New function in the errors package, as shown in the example above.  The errors.New function takes a string and returns an error.  If you pass an error to fmt.Println, it automatically calls the Error method.

Another way to create an error is using fmt.Errorf("message") as shown below. If you use fmt.Errorf, you do not need to import the errors package.

```go
func division(numerator float64, denominator float64) (float64, error) {
    if denominator == 0 {
        return 0, fmt.Errorf("denominator cannot be zero")
    }
    return numerator / denominator, nil
}
```

You can define your own errors type as long as they implement the error interface, as shown below.

```go
package main

import (
    "fmt"
    "os"
)

type AppErr struct {
    code    int
    message string
}

func (e AppErr) Error() string {
    return e.message
}

func division(numerator float64, denominator float64) (float64, error) {
    if denominator == 0 {
        return 0, AppErr{code: 1, message: "division by zero"}
    }
    return numerator / denominator, nil
}

func main() {
    result, err := division(5.2, 0)

    if err != nil {
        appErr := err.(AppErr)
        fmt.Println(appErr.code, appErr.message)
        os.Exit(1)

    }
    fmt.Println(result) // 2.6

}
```

In this example, we create a type named AppErr that includes an error code and the message. We are using type assertions, but another way is using the errors.As function as shown below.

```go
    if err != nil {
        var appErr AppErr
        if errors.As(err, &appErr) {
            fmt.Println(appErr.code, appErr.message)
        }
        os.Exit(1)

    }
```

The errors.As function returns true if a returned error matches a specific type.  Notice that the second argument passed to errors.As is a pointer to a variable of the type we are checking.

You can use the function errors.Is to check if an error matches a custom error type.  For example, in the previous example, you can check that error matches AppErr, as shown below.

```go
if errors.Is(err, appErr) {
    fmt.Println("Yes")
} else {
    fmt.Println("No")
}
```

## panic and recover

A panic is generated when there is a situation where Go is unable to figure out what to do next.  For example, the server ran out of memory, read a non-existing file, etc.  When panic occurs, the function immediately exits, and any defers attached to the function starts running.  You can create your own panic using the panic function, as shown below.

```go
package main

import "fmt"

func division(numerator float64, denominator float64) float64 {
    if denominator == 0 {
        panic("division by zero")
    }
    return numerator / denominator
}

func main() {
    result := division(5, 0)
    fmt.Println(result)
}
```

If you run this program, a panic message will be print out, followed by a stack trace.

You can gracefully handle panic situations using the built-in function recover.  The recover function must be called from a defer since once the panic occurs, only defer is executed, for example.

```go
package main

import "fmt"

func division(numerator float64, denominator float64) float64 {
    defer func() {
        if v := recover(); v != nil {
            fmt.Println(v)
        }
    }()
    if denominator == 0 {
        panic("division by zero")
    }
    return numerator / denominator
}

func main() {
    result := division(5, 0)
    fmt.Println(result)
    fmt.Println("end of the program")
}
```

In this example, when a panic occurs, the defer function is invoked, and the error message is printed. The program continues with its execution and prints "end of the program." Usually, when a panic occurs, you do not want to continue executing the program but gracefully shutdown.

# Repositories,  Modules, and Packages

A repository is a place in a version control system where source code for a project is stored.  A module is the root of a Go library or application.  Modules consist of one or more packages.

Every module has a globally unique name.  In Go, we usually use the path to the module repository.  For example https://github.com/blasvinas/golang

## Create a Go Module

In the following example, we will create a module with some basic math operations functions. Then we will the function from another module.  

First, create a directory for your module.

```shell
% mkdir mathop  
% cd mathop 
```
To create a module, you need to create a go.mod file in the module's root directory.  You can create this file with the command go mod init modle_path.

```shell
% go mod init github.com/blasvinas/golang/mathops
go: creating new go.mod: module github.com/blasvinas/golang/mathops

% ls -la
total 8
drwxr-xr-x@ 3 blasvinas  staff   96 Apr 16 08:12 .
drwxr-xr-x@ 9 blasvinas  staff  288 Apr 16 08:09 ..
-rw-r--r--  1 blasvinas  staff   51 Apr 16 08:12 go.mod
```

The go mod init command created the file go.mod which has the following content.

```shell
module github.com/blasvinas/golang/mathops

go 1.16
```

The go.mod file starts with the word module and the path.  The following line specifies the version that your code supports.

## Building a Package

Now let create the package with the math functions.  Let's name it mathops.go

```go
package mathops

func Add(x int, y int) int {
    return x + y
}

func Subtract(x int, y int) int {
    return x - y
}
```

The first line is the package clause.  It consists of the keyword package and the name of the package.  Then, you have the functions that are part of the package.

Notice that the name of the functions is capitalized.  This is the way Go makes the functions public and can be used by other modules.   Functions and variables that are not capitalized are private and can only be used inside the module.

As a general rule, you should make the package's name match the directory's name, except for the main program, since you cannot import main.

### Using a Package

For this example, we are creating a directory named prog at the same level that the directory mathops, where the package is. Hence, the root directory has the directory mathops and prog.

```shell
% ls
mathops     prog
```

We need to create the go.mod file in the prog directory similar to the mathops module.

```shell
% go mod init github.com/blasvinas/golang/prog 
go: creating new go.mod: module github.com/blasvinas/golang/prog
go: to add module requirements and sums:
    go mod tidy
```

Now, let's create a file called main.go with the following content in the prog directory.

```go
package main

import (
    "fmt"

    "github.com/blasvinas/golang/mathops"
)

func main() {
    result := mathops.Add(5, 2)
    fmt.Println(result)
    result = mathops.Subtract(5, 2)
    fmt.Println(result)
}
```

To use a package, you need to import it using import as shown above.  In this example, we are importing the fmt and mathops packages.

For productions, your mathops module should be published to GitLab, and Go can download the module from there.  In this example, we are working with a local copy, so the following steps are needed to resolve the dependencies.

```shell
% go mod edit -replace=github.com/blasvinas/golang/mathops=../mathops
```

The go.mod should look similar to the one below.

```shell
module github.com/blasvinas/golang/prog

go 1.16

replace github.com/blasvinas/golang/mathops => ../mathops
```


We need to run the go mod tiddy command.

```shell
% go mod tidy
go: found github.com/blasvinas/golang/mathops in github.com/blasvinas/golang/mathops v0.0.0-00010101000000-000000000000
```

Now the go.mod should look similar to this.

```shell
module github.com/blasvinas/golang/prog

go 1.16

replace github.com/blasvinas/golang/mathops => ../mathops

require github.com/blasvinas/golang/mathops v0.0.0-00010101000000-000000000000
```

The program is ready to run now.   


```shell
% go run .
7
3
```

## Package Name Conflicts

There are situations when you might need to import packages with the same name.  For example, the standard library includes the packages:  crypto/rand and math rand.  Suppose you need to import these packages in your program. In that case, it can resolve the conflict by providing an alternative name for one of the packages, as shown below.

```go
import (
    "fmt"
    crand "crypto/rand"
    "math/rand"
)
```

We can call the functions that belong to the crypto/rand packages using the crand name in the program.

# Concurrency

Before using concurrency, be sure that your program benefits from it. Using concurrency does not always make your program faster. Actually, in some cases, concurrency can make your program run slower. Use concurrency when you want to combine data from multiple operations that can operate independently. Also, you won't get too much benefit from concurrency if the process doesn't take a lot of time. Some in-memory algorithms are so fast the using concurrency is an overhead. Concurrency is often used for I/O to disk or network since these operations are usually slower than in-memory processes.

## Goroutines

Before working with goroutines, there are some essential concepts that we need to understand. The first is a process. A process is an instance of a program that is being run by the operating system. The operating system assigns resources like memory and CPU to the process. It makes sure that other processes can't access them. A process is composed of one or more threads. A thread is a unit of execution that is given some time to run by the operating system. Threads within a process share access to the resources. A CPU can execute instructions from one or more threads depending on the number of cores.

Goroutines are lightweight processes managed by the Go runtime. When a Go program starts, it creates several threads and lunches a single goroutine to run your program. All the goroutines created by your program are assigned to these threads. Go managing the threads instead of the operating system has several benefits like faster goroutine creation, thread stack sizes can grow as needed, faster switching between goroutines, etc.

You can lunch a goroutine by placing the go keyword before a function.

```go
go func process() {
 ...
}
```

## Channels

Goroutines communicate using channels. Channels are reference types. 

You can declare a channel as follow.

```go
ch := make(chan int)
```

By default, channels are unbuffered, so writing to a channel causes a writing goroutine to pause until another goroutine reads from the channel. The same applies to reading. A reading goroutine pause until another goroutine writes to the channel.

In Go, you can create a buffered channel specifying the capacity when declaring the channel.

```go
ch := make(chan int, 20)
```

The function len returns how many values are in the channel, and the function cap returns the maximum capacity.

Use the <- to read from and write to a channel.

```go
x := <-ch // read from channel
ch <- x // write to a channel
```

Also, you can read from a channel using a for-range loop.

```go
for v := range ch {
 fmt.print(v)
}
```

Use the close function to close a channel.

```go
close(ch)
```

You can use the ok operator when reading from a channel to make use the channel is still open.

```go
v, ok := <- ch
```

If ok is true, the channel is open. If ok is false, the channel is closed.

If you try to write to or close a closed channel, the program will crash.

## select

When a goroutine performs operations on different channels, use the select statement.

```go
select {
case v := <-ch:
 fmt.Println(v)
case v := <-ch2:
 fmt.Println(v)
case ch3 <- x:
 fmt.Println("wrote", x)
case <-ch4:
 fmt.Println("got value on ch4, but ignored it")
}
```

If there are several channels ready, the select statement picks one randomly.

