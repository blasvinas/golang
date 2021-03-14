- [Go Programming Language](#go-programming-language)
- [Setting up Go](#setting-up-go)
  - [Installing Go](#installing-go)
  - [Organizing Your Projects](#organizing-your-projects)
- [A Simple Go program](#a-simple-go-program)
- [Data Types](#data-types)
  - [Integers](#integers)
# Go Programming Language

# Setting up Go

## Installing Go

You can find the lastest version of Go and instructions of how to install it on your operating system in https://golang.org/dl/.  

The installation program should remove any previous  version of  Go, install Go in the right location and add to go command to the path.

To verify that go has been installed properly run the following from the command line.

```shell
go version
go version go1.16.2 darwin/amd6
```

You should get a output similar to the one above,  displaying the version of Go installed.

## Organizing Your Projects

Over the years the way Go projects are organized has change.  For modern Go projects,  you are free to organize your code any way it works for you.

By default, the third-party tools installed running go install are located in $HOME/go. It is recommend that you set the $GOPATH enviromental variable to $HOME/go.

In Linux systems using bash add the following lines to your .profile.

```shell
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

If you are using a MAC with the zsh shell add the line to .zshrc.

On Windows, run the following commands at the prompt:

setx GOPATH %USERPROFILE%\go
setx path %path%;%USERPROFILE%\bin


You can get the go enviroment running the go env command

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

# Data Types
Go has the following basic types.
* int int8 int16 int32 int64 uint uint8 uint16  uint64 uintptr
* float32 float64
* string
* bool
* byte
* rune
* complex64 complex128

## Integers
