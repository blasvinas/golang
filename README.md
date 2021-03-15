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

A boolean can only have two values: true or false. The zero value (the default value when the variable is not initialized) is false.  You can declare a boolean valiable using the bool type as shown below.

```go
package main

import (
	"fmt"
)

func main() {
	var isOk bool // false (zero value)
	fmt.Println(isOk) // false
	isOk = true
	fmt.Println(isOk)	// true
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

Unless you have a specific need about the size or the sign,  use int when declaring your valiables.

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

You float64 unless you haver a specific need to use float32.

## Complex Numbers

Go defines two  complex number types.

* complex64 use float32 to represent the real and imaginary parts.
* complex128 use float64 to represent the real and imaginary parts.

Complex numbers are numbers that consisy of two parts, a real number and an imaginary number.  They are represented in the form a + bi.  Below is an example of a program that declares a complex number.

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
