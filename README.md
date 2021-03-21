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

