<h2 align="center">
  <br />
  <img src="https://github.com/jeffotoni/gocompilation/blob/master/golang-compilation.png" alt="logo" width="670" />
  <br />
  <br />
  <br />
</h2>

# DevOps BootCamp

Material for **8 hours** Practical Immersion **with Golang**
This is a material in Golang to be presented **"face-to-face"** in a **"hand in hand"** Workshop that will be done in 8 hours.

# Golang Workshop DevOps
---

All content aims at the basic level of the student many practical examples were made with details richness to make life easier than it is initiating.
If you know little and almost nothing programming will not be problem every manual was made to level starting to advanced.
All the difficulties I had when I started trying to contemplate this material.
We will try to improve the material all the time so we can have a reference when it comes to Go.

I hope you all enjoy it and can serve as a base for learning and help several possible Gophers.

The content and references used are from the [Golang Official Site](https://golang.org) and the material being developed 
which is a compilation of all Golang language and can be checked here [jeffotoni/Compilation]( https://github.com/jeffotoni/goworkshopdevops#installation)

There are thousands of references today regarding Golang, let's start at the beginning and we could not stop talking about [Golang Tour](https://tour.golang.org).
Well that site here [Play Golang](https://play.golang.org/p/ayud1GTDJ0G) or [Play Go Space](https://goplay.space/#ayud1GTDJ0G) we can play Golang online.

We have a very interesting link that we have been able to search for packages written in Golang  check out this link: [Go Doc](https://godoc.org/)

We have this link that presents us as a manual all libs developed in Golang [Dev Docs](https://devdocs.io/go/)

Soon below some channels that I participate and can find me online.

### Telegram:
   - [gobr](https://t.me/go_br)
### Slack: 
   - [gophers.slack.com](https://gophers.slack.com)
      - brazil
      - brasil
      - general
      - go-kit
      - gotimefm
      
## Lab 01 Install and commands Golang

- [Overview](#overview)
- [Introduction Installation](#introduction-installation)
  - [Installation](#installation)
    - [Linux](#linux)
    - [$GOPATH](#gopath)
    - [Test your installation](#test-your-installation)
    - [Workspace](#workspace)
    - [Outside GOPATH](#outside-gopath)
- [Installation Docker](#installation-docker)
	- [Install Docker to Golang](#install-docker-to-golang)
	- [Compile your app inside the Docker container](#compile-your-app-inside-the-docker-container)
	- [Cross-compile your app inside the Docker container](#cross-compile-your-app-inside-the-docker-container)
- [Introduction golang](#introduction-golang)
  - [Golang language](#golang-language)
    - [Keywords](#keywords)
    - [Operators and punctuation](#operators-and-punctuation)
    - [Println Print](#println-print)
    - [Bufio NewWriter](#bufio-newWriter)
    - [Func Main](#func-main)
- [Go commands](#go-commands)
   - [go commands introduction](#go-commands-introduction)
   - [go run](#go-run) 
   - [go build](#go-build)
   - [go install](#go-install)
   - [go get](#go-get)
   - [go mod](#go-mod)
   - [go mod init](#go-mod-init)
   - [go mod vendor](#go-mod-vendor)
   - [GO111MODULE](#go111module)
   - [go test](#go-test)

## Lab 02 The golang types

- [Types](#types)
   - [Numeric Types](#numeric-types)
   - [String types](#string-types)
   - [Pointer Types](#pointer-types)
   - [Array Types](#array-types)
   - [Slice Types](#slice-types)
   - [Struct Types](#struct-types)
   - [Struct in C](#struct-in-c)
   - [Struct Type Tags Json](#struct-type-tags-json)
   - [Fatih Structs to Map](#fatih-structs-to-map)
   - [Map Types](#map-types)
   - [Map literals continued](#map-literals-continued)
   - [Channel Types](#channel-types)
   - [Blank identifier](#blank-identifier)
   - [Interface Types](#interface-types)
	 - [Here's an interface as a method](#heres-an-interface-as-a-method)
	 - [Interface as type](#interface-as-type)
   - [Exercise one](#exercise-one)
- [Control structures](#control-structures)
  - [Control](#control)
    - [Control Return](#control-return)
    - [Control Goto](#control-goto)
    - [Control if else](#control-if-else)
    - [Control for break continue](#control-for-break-continue)
    - [Control Switch case break](#control-switch-case-break)
    - [Control Label](#control-label)
    - [Control Range](#control-range)
- [Errors](#Errors)
  - [Introduction Errors](#introduction-Errors)
    - [How Error Control Works](#how-error-control-works)
    - [Errors.New](#errors-new)
    - [Custom Errors](#custom-errors)    
    - [fmt Errorf](#fmt-errorf)
- [Functions](#functions)
  - [Introduction Function](#introduction-function)
    - [Return multiple values](#return-multiple) 
    - [Variadic Functions](#variadic-functions) 
    - [functions as a parameter](#functions-as-a-parameter) 
    - [Closures](#closures)
    - [Recursion](#recursion)
    - [Asynchronous Functions](#asynchronous-functions)
- [Defer](#defer)
- [Exercise two](#exercise-two)

## Lab 03 Json with Golang

- [Json](#Json)
  - [introduction](#)
    - [Json marshal Encode](#json-marshal-encode)
    - [json MarshalIndent](#json-marshalIndent)
    - [Option Omitempty](#option-omitempty)
    - [Initialized collections of data](#initialized-collections-of-data)
    - [Json NewEncoder](#json-newencoder)
    - [Json Unmarshal Decode](#json-unmarshal-decode)
    - [Generic JSON with interface{} and assertion](#generic-json-with-interface-and-assertion)
    - [Dynamic type](#dynamic-type)
    - [What is reflection?](#what-is-reflection?)
- [Parse Json](#Json)
	- [Reading and Parsing a JSON File](#reading-and-parsing-a-json-file)
	- [Reading the JSON File](#reading-the-json-file)
	- [Parsing with Structs](#parsing-with-structs)
	- [Parsing with Map and Interface](#parsing-with-map-and-interface)
    - [Json Toml](#json-toml)
    - [Json Yaml](#json-yaml)
    - [Json-Gcfg](#json-gcfg)
- [Links Json to Golang](#links-json-to-golang)
- [Exercise three](#Exercise-three)

## Lab 04 building apis with net/http

- [net/http Server](#)
  - [introduction](#)
    - [http.NewServeMux](#)
    - [http.HandlerFunc](#)
    - [http.Handle](#)
    - [http.Handler](#)
    - [http.Server](#)
    - [next.ServeHTTP](#)
    - [ListenAndServe](#)
    - [ListenAndServeTLS](#)
    - [Server.Shutdown](#)
    - [Middleware](#)
- [Exercise five](#Exercise-five)

- [net/http Client](#)
  - [introduction](#)
    - [http.Transport](#)
    - [http.Client](#)
    - [http.Get](#)
    - [http,Post](#)
    - [http.NewRequest](#)
    - [Context.WithCancel](#)
- [Exercise six](#Exercise-six)

- [net/http Server Pages](#)
  - [introduction](#)
    - [http.FileServer](#)
    - [http.NotFound](#)
    - [Disable http.FileServer](#)
    - [http.Dir](#)
    - [http.StripPrefix](#)    
- [Exercise seven](#Exercise-seven)

## Lab 05 Using Golang to create Command Line programs

- [Golang Cli](#golang-cli)
  - [introduction](#)
  	- [Environment Variables](#environment-variables)
  	- [Open and Read File](#open-write-file)
  	- [Write File](#write-file)
  	- [New Scanner and Scan](#new-scanner-and-scan)
  	- [Stdout](#stdout)
    - [Stdin](#stdin)
    - [Go flag types](#go-flag-types)
    - [Go flag Parse](go-flag-parse)
    - [Go flag PrintDefaults](go-flag-printDefaults)
    - [Os Args](#os-args)
    - [func init](#func-init)
    - [Go exec Command](#go-exec-command)
    - [Parse Url](#parse-url)
    - [Parse Yaml](#parse-yaml)

- [Exercise five](#Exercise-five)

## Lab 06 Goroutine the power

- [Goroutine](#)
  - [introduction](#)
    - [Channel](#)
    - [Parallelism](#)
    - [GOMAXPROCS](#)
    - [WaitGroup](#)
    - [Race Conditions](#)
    - [Worker](#)
    - [Context](#)
    - [Ticker](#)
    - [Singleton Connect Thread Safe](#)

- [Exercise four](#Exercise-four)

### Overview
---

Go is a powerful language when it comes to competition and high performance, with a clean and efficient architecture. It grows year after 
year and every day the communities grow even more.

Some paradigms have been broken to make it a high performance language, where competition is one of its strengths. Go makes it easy to 
create programs that take full advantage of multicore and networked machines, while the new type system allows you to build flexible and modular programs.

It is a fast and statically compiled language that looks like a dynamically interpreted language. This feature Golang becomes 
a unique language as the subject is web.

Go is a compiled, competing, strong and statically typed programming language.
It is a "General Use" language that can be used to solve various problems and in different areas.
Problems involving competition, web applications, high performance applications, development of APIs, communications sockets etc ... 
Is where language is increasingly becoming prominent in the market and in communities.

### Introduction installation

In golang the installation is all very simple and practical, for Linux, Mac and Windows.

Just copy the files to the correct directory for each operating system and export the paths to the environment and prompt, golang is installed.

Let's take a look at how we do this.

### Installation 
---

We will download the file, unpack it and install it in /usr/local/go, if we have golang already installed in the machine we will have to remove the existing one to leave our installation as unique.
Let's create our directory in our workspace and test to see if everything went well

### Linux

```bash
$ sudo rm -rf /usr/local/go
$ wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

### $GOPATH

$GOPATH is the golang in your $HOME, this is necessary for your projects to use pkg and build properly. This was mandatory for all versions before version 1.11. The cool thing is that from now on we will not have to create projects in $GOPATH, we can create in any other directory that is not in $GOPATH.

Here is the link to the versioning proposal [Proposal: Versioned Go Modules](https://go.googlesource.com/proposal/+/master/design/24301-versioned-go.md/) or [Go 1.11 Modules](https://github.com/golang/go/wiki/Modules/)

We'll detail how to work with **go mod**, it was one of the best experiences I had for versioning projects using Golang.

Let's set up our environment to run Go. Add **/usr/local/go/bin** to the PATH environment variable. You can do this by adding this line to your **/etc/profile** (for a system-wide installation) or **$HOME/.profile**. 

```bash
$ export PATH=$PATH:/usr/local/go/bin
```

**Note**: changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile. 

```bash
$ echo "export GOPATH=$HOME/go" >> $HOME/.profile
$ echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.profile
$ echo "export PATH=$PATH:$GOPATH/bin" >> $HOME/.profile
```

### Test your installation

Let's run go version to see if everything is correct.

```bash
$ go version
go version go1.11.5 linux/amd64
```

Check that Go is installed correctly by setting up a workspace and building a simple program, as follows. 

Create your **workspace** directory, $HOME/go. (If you'd like to use a different directory, you will need to set the $GOPATH environment variable.)

Next, make the directory src/hello inside your workspace, and in that directory create a file named hello.go that looks like:

### Workspace

Workspace is our place of work, where we will organize our directories with our projects. As shown above, until **Go version 1.11** we were forced to do everything under the Workspace. $GOPATH Down Projects.

**Example hello**
```bash
$ export GOPATH=$HOME/go
$ mkdir $HOME/go
$ mkdir $HOME/go/src
$ mkdir $HOME/go/src/hello
$ vim $HOME/go/src/hello/hello.go
```

```bash
$GOPATH/
  |-src
    |-hello
      |-hello.go
```

**Example Project**
```bash
$ export GOPATH=$HOME/go
$ mkdir $HOME/go/src/project1
$ mkdir $HOME/go/src/project1/my-pkg
$ mkdir $HOME/go/src/project1/my-cmd
$ mkdir $HOME/go/src/project1/my-vendor
$ mkdir $HOME/go/src/project1/my-logs
$ mkdir $HOME/go/src/project1/my-models
$ mkdir $HOME/go/src/project1/my-repo
$ mkdir $HOME/go/src/project1/my-handler
```

```bash
$GOPATH/
  |-src
    |-github.com/user/project1/
        |-cmd (of project1)
          |-main.go
        |-vendor
        |-logs
        |-models
        |-repo
        |-handler
    |-github.com/user/project2/
      ....
      ....
```

The $GOPATH environment variable tells the Go tool where your workspace is located. 

```go
$ go get github.com/user/project1
```

The **go get** command fetches source repositories from the internet and places them in your workspace.
Package paths matter to the Go tool. Using "github.com/..." means the tool knows how to fetch your repository. 

In the scenario above everything would have to stay in our **$GOPATH** so that our projects worked correctly.

### Outside $GOPATH

Now we can do our projects without being in $GOPATH, we can, for example, do it in any directory.

**Project outside GOPATH**

```bash
$ export GOPATH=$HOME/go
$ mkdir $HOME/2019/project1
$ mkdir $HOME/2019/project1/my-pkg
$ mkdir $HOME/2019/project1/my-cmd
$ mkdir $HOME/2019/project1/my-logs
$ mkdir $HOME/2019/project1/my-models
$ mkdir $HOME/2019/project1/my-repo
$ mkdir $HOME/2019/project1/my-handler
```
```bash
$HOME/
  |-2019
    |-github.com/user/project1/
      |-cmd
        |-main.go
      |-vendor
      |-logs
      |-models
      |-repo
      |-handler
```

We can put our project in any directory now.

```bash
$HOME/
  |-any-directory
    |-github.com/user/project1/
      |-cmd
        |-main.go
      |-vendor
      |-logs
      |-models
      |-repo
      |-handler
```

For the above scenario, we will have to use **go mod** in our project so that all external packages can work correctly, in this way we will be able to manage them correctly and version.
More information can be found here: [Wiki Go Modules](https://github.com/golang/go/wiki/Modules)

Practical example of how you will proceed:
```go
$ go mod init github.com/user/project1
```

**Note**: 
When we use go mod in $GOPATH we will have to enable using GO111MODULE=on, so that it can work within the $GOPATH structure.
So our program can compile successfully.

```go
$ GO111MODULE=on go run cmd/main.go
$ GO111MODULE=on go build -o project1 cmd/main.go
```

### Installation Docker

If we do not want to install directly on our operating system golang we can install it in a docker container.

We can upload a docker container with the installed language and compile and run our programs from this container.

Let's check how we can do it below.

More information and details you can visit this link: [hub.docker](https://hub.docker.com/_/golang)

### Install Docker to Golang

```bash
$ docker pull golang
```

### Compile your app inside the Docker container

There may be occasions where it is not appropriate to run your app inside a container. To compile, but not run your app inside the 
Docker instance, you can write something like:

```bash
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.11.5 go build -v
```

This will add your current directory as a volume to the container, set the working directory to the volume, and run the command go build which will tell go to compile the project in the working directory and output the executable to myapp. Alternatively, if you have a Makefile, you can run the make command inside your container.

```bash
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.11.5 make
```

### Cross-compile your app inside the Docker container
If you need to compile your application for a platform other than linux/amd64 (such as windows/386):

```bash
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=windows \
-e GOARCH=386 golang:1.11.5 go build -v
```

### Example main.go

Let's do our test program, let's call it main.go

```go
package main

import "fmt"

func main(){
	fmt.Println("My first program being compiled by a docker container!")
}
```

Now let's run a program to see if it works correctly

```bash
$ docker run --rm -v "$PWD":/usr/src/main -w /usr/src/main golang:1.11.5 go run main.go
```

Output:
```bash
My first program being compiled by a docker container!
```

Check the version:
```bash
$ docker run --rm -v "$PWD":/usr/src/main -w /usr/src/main golang:1.11.5 go versio
```

Output:
```bash
go version go1.11.5 linux/amd64
```

### Introduction golang
---

Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. 
Programs are constructed from packages, whose properties allow efficient management of dependencies.

The grammar is compact and regular, allowing for easy analysis by automatic tools such as integrated development environments.

### Golang language
---

### Keywords

The following keywords are reserved and may not be used as identifiers. 

```bash
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Operators and punctuation

The following character sequences represent operators (including assignment operators) and punctuation: 

```bash
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```

### Println Print

Let's learn how to send data to screen which is actually **stdout** standard output we will see more ahead with details on **stdout** and **stdin**.

Let's know **print, println and fmt.Println**

Current implementations provide several built-in functions useful during bootstrapping. These functions are documented for completeness but are not guaranteed to stay in the language. They do not return a result. 

Implementation restriction: **print** and **println** need not accept arbitrary argument types, but printing of boolean, numeric, and string types must be supported. 

**println is an built-in function** (into the runtime) which may eventually be removed, while the **fmt package** is in the standard library, which will persist.


```bash
Function   Behavior

print      prints all arguments; formatting of arguments is implementation-specific
println    like print but prints spaces between arguments and a newline at the end
```

using print:
```go
// test print
package main

func main() {
   print("debugging my system with print")
}
```

Output:
```bash
debugging my system with print
```

using println:
```go
// test println
package main

func main() {
   println("debugging my system with println")
}
```

Output:
```bash
debugging my system with println
```

using fmt.Println:
```go
package main

import "fmt"

func main() {
   fmt.Println("debugging my system with fmt.Println")
}
```

Output:
```bash
debugging my system with fmt.Println
```

The goal of starting and running the print, println or fmt.Println command is to help us with the tests we will be performing from now on at every step of our Go learning.


### Bufio NewWriter

```bash
bufio.Writer
```

Doing many small writes can hurt performance. Each write is ultimately a syscall and if doing frequently can put burden on the CPU. Devices like disks work better dealing with block-aligned data. To avoid the overhead of many small write operations Golang is shipped with bufio.Writer. Data, instead of going straight to destination (implementing io.Writer interface) are first accumulated inside the buffer and send out when buffer is full:

Let’s visualise how buffering works with nine writes (one character each) when buffer has space for 4 characters:

```bash
producer         buffer           destination (io.Writer)
 
   a    ----->   a
   b    ----->   ab
   c    ----->   abc
   d    ----->   abcd
   e    ----->   e      ------>   abcd
   f    ----->   ef               abcd
   g    ----->   efg              abcd
   h    ----->   efgh             abcd
   i    ----->   i      ------>   abcdefgh
```

Check out the example below
```go
package main

import (
	"bufio"
	"os"
)

// creating the write object pointer
// so that we can receive value in every
// scope of our program
var writer *bufio.Writer

func main() {
	// All screen output will be redirected
	// to bufio.NewWriter
	writer = bufio.NewWriter(os.Stdout)
	s := "How many stars does Orion have?\n"
	var b byte = 'H'

	writer.WriteString(s)
	writer.WriteByte(b)
	writer.WriteString("\n")

	// when all the functions finishes it closes
	// the buffer and sends to the.Stdout
	defer writer.Flush()
}
```

Output:
```bash
How many stars does Orion have?
H
```

### Func Main

```go
package main

import "fmt"

func main() {
  fmt.Printf("hello, Gophers\n")
}
```

Then **build** it with the **go tool**: 

```go
$ cd $HOME/go/src/hello
$ go build
```

Or we can compile like this:
```go
$ cd $HOME/go/src/hello
$ go build -o hello hello.go
```

The command above will build an executable named hello in the directory alongside your source code. Execute it to see the greeting: 

```go
$ ./hello
hello, Gophers
```

Check also the command **run** it with the go: 

```go
$ go run hello.go
hello, Gophers
```

If you see the **"hello, Gophers"** message then your Go installation **is working**.

You can run **go install** to install the binary into your workspace's **bin** directory or **go clean -i** to remove it.

Example: go install
```go
$ pwd
$ $HOME/go/src/hello
$ cd $HOME/go/src/hello
$ go install
$ ls -lhs $HOME/go/bin
-rwxrwxr-x 1 user user 2,9M nov  8 03:11 hello
```

Example: go clean -i

```go
$ go clean -i 
$ ls -lhs $HOME/go/bin
```

### Go commands
---

### Go commands introduction

In golang we have an arsenal to help us when it comes to compiling, testing, documenting, managing Profiling etc.

```bash
bug         start a bug report
build       compile packages and dependencies
clean       remove object files and cached files
doc         show documentation for package or symbol
env         print Go environment information
fix         update packages to use new APIs
fmt         gofmt (reformat) package sources
generate    generate Go files by processing source
get         download and install packages and dependencies
install     compile and install packages and dependencies
list        list packages or modules
mod         module maintenance
run         compile and run Go program
test        test packages
tool        run specified go tool
version     print Go version
vet         report likely mistakes in packages
```

Use "go help " for more information about a command.


### Go run
---

Usage:
```bash
go run [build flags] [-exec xprog] package [arguments...]
```

Run compiles and runs the named main Go package. Typically the package is specified as a list of .go source files, but it may also be an import path, file system path, or pattern matching a single known package, as in 'go run .' or 'go run my/cmd'.

By default, 'go run' runs the compiled binary directly: 'a.out arguments...'. If the -exec flag is given, 'go run' invokes the binary using xprog: 

If the -exec flag is not given, GOOS or GOARCH is different from the system default, and a program named go_$GOOS_$GOARCH_exec can be found on the current search path, 'go run' invokes the binary using that program, for example 'go_nacl_386_exec a.out arguments...'. This allows execution of cross-compiled programs when a simulator or other execution method is available.

The exit status of Run is not the exit status of the compiled binary.

For more about build flags, see 'go help build'. For more about specifying packages, see 'go help packages'.

See below an example:

```go
// test println
package main

func main() {
   println("Debugging my system with println")
}
```
Go run:
```bash
go run println.go
```

Output:
```bash
Debugging my system with println
```

### Go build
---

Build compiles the packages named by the import paths, along with their dependencies, but it does not install the results. 

When compiling packages, build ignores files that end in '_test.go'.

The -o flag, only allowed when compiling a single package, forces build to write the resulting executable or object to the named output file, instead of the default behavior described in the last two paragraphs.

The -i flag installs the packages that are dependencies of the target.

```go
$ go build [-o output] [-i] [build flags] [packages]
```

See an example:
```go
package main

import "fmt"

func main() {
  fmt.Println("Workshop Golang 2019.")
}
```

Output:
```bash
Workshop Golang 2019.
```

Normal compilation
```go
go build -o hello hello.go
```

Output:
```bash
$ ls -lh 
-rwxrwxr-x 1 root root **1,9M** jan 18 12:42 hello
-rw-rw-r-- 1 root root   75 jan 17 12:04 hello.go
```

Leaving the file smaller after compiling
```go
go build -ldflags="-s -w" hello.go
```

Output:
```bash
$ ls -lh 
-rwxrwxr-x 1 root root **1,3M** jan 18 12:42 hello
-rw-rw-r-- 1 root root   75 jan 17 12:04 hello.go
```

### Go Install
---

Install packages and dependencies

Usage:
```bash
$ go install [-i] [build flags] [packages]
```

Install compiles and installs the packages named by the import paths.

The **-i flag** installs the dependencies of the named packages as well.

For more about the build flags, see 'go help build'. For more about specifying packages, see 'go help packages'.

### Go Get
---

The **'go get'** command changes behavior depending on whether the go command is running in module-aware mode or legacy GOPATH mode. This help text, accessible as 'go help module-get' even in legacy GOPATH mode, describes 'go get' as it operates in module-aware mode.

Usage:
```bash
$ go get [-d] [-m] [-u] [-v] [-insecure] [build flags] [packages]
```

Get downloads the packages named by the import paths, along with their dependencies. It then installs the named packages, like 'go install'.

Look at the flags accepted below:
```bash
The -d flag instructs get to stop after downloading the packages; that is, it instructs get not to install the packages.

The -f flag, valid only when -u is set, forces get -u not to verify that each package has been checked out from the source control repository implied by its import path. This can be useful if the source is a local fork of the original.

The -fix flag instructs get to run the fix tool on the downloaded packages before resolving dependencies or building the code.

The -insecure flag permits fetching from repositories and resolving custom domains using insecure schemes such as HTTP. Use with caution.

The -t flag instructs get to also download the packages required to build the tests for the specified packages.

The -u flag instructs get to use the network to update the named packages and their dependencies. By default, get uses the network to check out missing packages but does not use it to look for updates to existing packages.

The -v flag enables verbose progress and debug output.
```

Examples:
```bash
$ go get -v github.com/guptarohit/asciigraph
$ go get -u github.com/mxk/go-sqlite
$ go get -v github.com/google/uuid
$ go get -v github.com/sirupsen/logru
```


### Go mod
---

A module is a collection of related Go packages. Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules. Modules replace the old GOPATH-based approach to specifying which source files are used in a given build. 

Usage:

```bash
$ go mod <command> [arguments]
```
A module is defined by a tree of Go source files with a **go.mod** file in the tree's root directory. The directory containing the go.mod file is called the module root. Typically the module root will also correspond to a source code repository root (but in general it need not). The module is the set of all Go packages in the module root and its subdirectories, but excluding subtrees with their own go.mod files.

The "module path" is the import path prefix corresponding to the module root. The go.mod file defines the module path and lists the specific versions of other modules that should be used when resolving imports during a build, by giving their module paths and versions.

For example, this go.mod declares that the directory containing it is the root of the module with path example.com/m, and it also declares that the module depends on specific versions of golang.org/x/text and gopkg.in/yaml.v2: 

```bash
$ go mod init github.com/user/gomyproject

require (
  golang.org/x/text v0.3.0
  gopkg.in/yaml.v2 v2.1.0
)
```
The go.mod file can also specify replacements and excluded versions that only apply when building the module directly; they are ignored when the module is incorporated into a larger build. For more about the go.mod file, see 'go help go.mod'.

To start a new module, simply create a go.mod file in the root of the module's directory tree, containing only a module statement. The 'go mod init' command can be used to do this: 

```bash
$ go mod init github.com/user/gomyproject
```
In a project already using an existing dependency management tool like **godep, glide, or dep, 'go mod init'** will also add require statements matching the existing configuration.

Once the go.mod file exists, no additional steps are required: go commands like **'go build'**, **'go test'**, or even **'go list'** will automatically add new dependencies as needed to satisfy imports.

The commands are: 

```bash
download    download modules to local cache
edit        edit go.mod from tools or scripts
graph       print module requirement graph
init        initialize new module in current directory
tidy        add missing and remove unused modules
vendor      make vendored copy of dependencies
verify      verify dependencies have expected content
why         explain why packages or modules are needed
```
Use "go help mod <command>" for more information about a command.


### Go Mod Init

Initialize new module in current directory

Usage:

```bash
$ go mod init [module]
```

Init initializes and writes a new **go.mod** to the current directory, in effect creating a new module rooted at the current directory. The file go.mod must not already exist. If possible, init will guess the module path from import comments (see 'go help importpath') or from version control configuration. To override this guess, supply the module path as an argument. 


```bash
$ go mod init github.com/user/gomyproject2

require (
  github.com/dgrijalva/jwt-go v3.2.0+incompatible
  github.com/didip/tollbooth v4.0.0+incompatible
  github.com/go-sql-driver/mysql v1.4.1
  github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
  golang.org/x/crypto v0.0.0-20190103213133-ff983b9c42bc
  golang.org/x/time v0.0.0-20181108054448-85acf8d2951c // indirect
)
```

### Go mod vendor

The go mod vendor command will download all dependencies to the "vendor" directory.
When using go mod init the packages are not in your directory.

```bash
$ cd gomyproject2
$ go mod vendor
```

Output:
```bash
$ ls -lh vendor
total 8,0K
drwxrwxr-x 3 root root 4,0K jan 27 01:47 github.com
-rw-rw-r-- 1 root root  137 jan 27 01:47 modules.txt
```

### GO111MODULE

Go 1.11 includes preliminary support for Go modules, including a new module-aware 'go get' command. We intend to keep revising this support, while preserving compatibility, until it can be declared official (no longer preliminary), and then at a later point we may remove support for work in GOPATH and the old 'go get' command.

The quickest way to take advantage of the new Go 1.11 module support is to check out your repository into a directory outside GOPATH/src, create a go.mod file (described in the next section) there, and run go commands from within that file tree.

For more fine-grained control, the module support in Go 1.11 respects a temporary environment variable, GO111MODULE, which can be set to one of three string values: off, on, or auto (the default). If GO111MODULE=off, then the go command never uses the new module support. Instead it looks in vendor directories and GOPATH to find dependencies; we now refer to this as "GOPATH mode." If GO111MODULE=on, then the go command requires the use of modules, never consulting GOPATH. We refer to this as the command being module-aware or running in "module-aware mode". If GO111MODULE=auto or is unset, then the go command enables or disables module support based on the current directory. Module support is enabled only when the current directory is outside GOPATH/src and itself contains a go.mod file or is below a directory containing a go.mod file.

In module-aware mode, GOPATH no longer defines the meaning of imports during a build, but it still stores downloaded dependencies (in GOPATH/pkg/mod) and installed commands (in GOPATH/bin, unless GOBIN is set).


Check below how we use the command:
```bash
$ GO111MODULE=on go run myprogram.go
$ GO111MODULE=on go build myprogram.go
```
When our project is not in our **$GOPATH** it is not necessary to use **GO111MODULE**, but when our project is in **$GOPATH** and we want to use **"go mod"** we need to inform this to the compiler using **GO111MODULE**...


### Go Test
---

Test packages

Usage:

```go
$ go test [build/test flags] [packages] [build/test flags & test binary flags]
```

Go **test** automates testing the packages named by the import paths. It prints a summary of the test results in the format: 

```bash
=== RUN   TestWhatever
--- PASS: TestWhatever (0.00s)
PASS
ok    command-line-arguments  0.001s
```

The test package runs side-by-side with the go test command. The package test should have the suffix "\_test.go".
We can split the tests into several files following this convention. For example: "myprog1_test.go" and "myprog2_test.go".

We should put our test functions in these test files.

Each test function is an exported public function whose name begins with **"Test"**, accepts a pointer to a **testing.T** object, and returns nothing. Like this:

Example one / myprog1_test:
```go
package main

import "testing"

func TestWhatever(t *testing.T) {
    // Your test code goes here
}
```

```bash
$ go test -v
```

Output:
```bash
=== RUN   TestWhatever
--- PASS: TestWhatever (0.00s)
PASS
ok    command-line-arguments  0.001s
```

The T object provides several methods that we can use to indicate failures or log errors.

Example two / myprog2_test:
```go
package main

import "testing"

func TestSum(t *testing.T) {
  x := 1 + 1
  if x != 11 { // forcing the error
    t.Error("Error! 1 + 1 is not equal to 2, I got", x)
  }
}
```

```bash
$ go test -v
```

Output:
```bash
=== RUN   TestSum
-- FAIL: TestSum (0.00s)
    myprog1_test.go:12: Error! 1 + 1 is not equal to 2, I got 2
FAIL
FAIL  command-line-arguments  0.001s
```

In this example we will make an examination as it would be in our projects.

In this program I will pass parameter at compile time or in our execution to facilitate and also serve as an example the use of **"ldflags"** that we can use in both **go run -ldflags ** and **go build -ldflags**

From a check in our code below / main.go:
```go
import "strconv"

import (
  "github.com/jeffotoni/goworkshopdevops/examples/tests/pkg/math"
)

var (
  x, y   string
  xi, yi int
)

func init() {
  xi, _ = strconv.Atoi(x)
  yi, _ = strconv.Atoi(y)
}

func main() {
  println(math.Sum(xi, yi))
}
```

Now we have a Sum function in a pkg that we create in **pkg/math/sum.go**

```go
package math

func Sum(x, y int) int {
  return x + y
}
```

We created our test file in **pkg/math/sum_test.go**

```go
package math

import "testing"

func TestSum(t *testing.T) {
  type args struct {
    x int
    y int
  }
  tests := []struct {
    name string
    args args
    want int
  }{
    // TODO: Add test cases.
    {"test_1: ", args{2, 2}, 4},
    {"test_2: ", args{-2, 6}, 4},
    {"test_3: ", args{-4, 8}, 4},
    {"test_4: ", args{5, 7}, 12},
    {"test_5: ", args{8, 8}, 15}, // forcing the error
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := Sum(tt.args.x, tt.args.y); got != tt.want {
        t.Errorf("Sum() = %v, want %v", got, tt.want)
      }
    })
  }
}
```

```bash
$ cd pkg/math/
$ go test -v
```

Output:
```bash
=== RUN   TestSum
=== RUN   TestSum/test_1:_
=== RUN   TestSum/test_2:_
=== RUN   TestSum/test_3:_
=== RUN   TestSum/test_4:_
=== RUN   TestSum/test_5:_
--- FAIL: TestSum (0.00s)
    --- PASS: TestSum/test_1:_ (0.00s)
    --- PASS: TestSum/test_2:_ (0.00s)
    --- PASS: TestSum/test_3:_ (0.00s)
    --- PASS: TestSum/test_4:_ (0.00s)
    --- FAIL: TestSum/test_5:_ (0.00s)
        sum_test.go:29: Sum() = 16, want 15
FAIL
exit status 1
FAIL  github.com/jeffotoni/goworkshopdevops/examples/tests/pkg/pkg/math  0.001s
```

It converts to json the output of the tests

```bash
$ go test -v -json
```

check your output on your terminal screen to view json output.

---

Now that we've saved our pkg / math / sum.go let's do a main.go by making the call in this packet.
But first let's run go mod to manage our packages and versions correctly.

Check the command below:
```bash
$ go mod init github.com/jeffotoni/goworkshopdevops/examples/tests/pkg
```

Output:
```bash
go: finding github.com/jeffotoni/goworkshopdevops/examples/tests/pkg/math latest
go: finding github.com/jeffotoni/goworkshopdevops/examples/tests latest
go: finding github.com/jeffotoni/goworkshopdevops/examples latest
go: finding github.com/jeffotoni/goworkshopdevops latest
go: downloading github.com/jeffotoni/goworkshopdevops v0.0.0-20190127023912-a2fa53299867
0
```

Now we can do **build** or **run** in our **main.go**.
Let's run go run using the **"-ldflags"** flag to pass parameter to our code at compile time.

```bash
$ go run -ldflags "-X main.x=2 -X main.y=3" main.go
```

Output:
```bash
5
```

```bash
$ go run -ldflags "-X main.x=7 -X main.y=3" main.go
```

Output:
```bash
10
```

## Lab 02 Types with Golang
---

### Types
---

A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a type name, if it has one, or specified using a type literal, which composes a type from existing types. 

The language predeclares certain type names. Others are introduced with type declarations. Composite types—array, struct, pointer, function, interface, slice, map, and channel types—may be constructed using type literals.

Each type T has an underlying type: If T is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is T itself. Otherwise, T's underlying type is the underlying type of the type to which T refers in its type declaration. 

Example:
```bash
type (
    A1 = string
    A2 = A1
)

type (
    B1 string
    B2 B1
    B3 []B1
    B4 B3
)
```

The underlying type of string, A1, A2, B1, and B2 is string. The underlying type of []B1, B3, and B4 is []B1. 


### Numeric Types

A numeric type represents sets of integer or floating-point values. The predeclared architecture-independent numeric types are: 

```bash
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```
The value of an n-bit integer is n bits wide and represented using two's complement arithmetic.

There is also a set of predeclared numeric types with implementation-specific sizes:

```bash
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

To avoid portability issues all numeric types are defined types and thus distinct except byte, which is an alias for uint8, and rune, which is an alias for int32. Conversions are required when different numeric types are mixed in an expression or assignment. For instance, int32 and int are not the same type even though they may have the same size on a particular architecture. 

### String types

A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is string; it is a defined type.

The length of a string s (its size in bytes) can be discovered using the built-in function len. The length is a compile-time constant if the string is a constant. A string's bytes can be accessed by integer indices 0 through len(s)-1. It is illegal to take the address of such an element; if s[i] is the i'th byte of a string, &s[i] is invalid. 

```go
package main

import "fmt"

type S string

var (
  String = "@jeffotoni"
)

func main() {
  var text string
  var str S

  mypicture := "@Photograph-jeffotoni"
  str = "@workshop-devOpsBh"
  text = "@jeffotoni-golang"

  fmt.Println(str)
  fmt.Println(String)
  fmt.Println(text)
  fmt.Println(mypicture)

  // example string
  s := "日本語"
  fmt.Printf("Glyph:             %q\n", s)
  fmt.Printf("UTF-8:             [% x]\n", []byte(s))
  fmt.Printf("Unicode codepoint: %U\n", []rune(s))
}
```
Output:

```go
@workshop-devOpsBh
@jeffotoni
@jeffotoni-golang
@Photograph-jeffotoni
Glyph:             "日本語"
UTF-8:             [e6 97 a5 e6 9c ac e8 aa 9e]
Unicode codepoint: [U+65E5 U+672C U+8A9E]
```
### Pointer types

Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

A pointer type denotes the set of all pointers to variables of a given type, called the base type of the pointer. The value of an uninitialized pointer is nil.

```bash
PointerType = "*" BaseType .
BaseType    = Type .
```

```bash
*Point
*[4]int
```

Example:
```go
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
  fmt.Println(p.Y)
}
```
Output:
```bash
{1000000000 2}
2
```
For every type that is declared, either by you or the language itself, you get for free a complement pointer type you can use for sharing. There already exists a built-in type named int so there is a complement pointer type called *int.

All pointer types have the same two characteristics. First, they start with the character *. Second, they all have the same memory size and representation, which is a 4 or 8 bytes that represent an address. On 32bit architectures (like the playground), pointers require 4 bytes of memory and on 64bit architectures (like your machine), they require 8 bytes of memory.

Example:
```go
package main

func main() {

  var a int
  inc := &a
  *inc = 2
  *inc++
  println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
```

Output:
```bash
inc:  Value Of[ 0xc000036778 ]  Addr Of[ 0xc000036780 ]  Value Points To[ 3 ]
```

For an operand x of type T, the address operation &x generates a pointer of type *T to x. The operand must be addressable, that is, either a variable, pointer indirection, or slice indexing operation; or a field selector of an addressable struct operand; or an array indexing operation of an addressable array. As an exception to the addressability requirement, x may also be a (possibly parenthesized) composite literal. If the evaluation of x would cause a run-time panic, then the evaluation of &x does too.

For an operand x of pointer type *T, the pointer indirection *x denotes the variable of type T pointed to by x. If x is nil, an attempt to evaluate *x will cause a run-time panic.

```bash
&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // causes a run-time panic
&*x  // causes a run-time panic
```

See the example below:
```go
package main

import "fmt"

func main() {
  var a int = 100  /* actual variable declaration */
  var pa *int      /* pointer variable declaration */
  var pointer *int /* pointer variable declaration */

  pa = &a /* store address of a in pointer variable*/

  fmt.Printf("Address of a variable: %x\n", &a)

  /* address stored in pointer variable */
  fmt.Printf("Address stored in ip variable: %x\n", pa)

  /* access the value using the pointer */
  fmt.Printf("Value of *ip variable: %d\n", *pa)

  /* succeeds if p is not nil */
  if pa != nil {
    fmt.Println("success is not nil")
  }

  /* succeeds if p is null */
  if (pointer) == nil {
    fmt.Println("success pointer is nil")
  }
}
```

Output:
```bash
Address of a variable: c0000160c8
Address stored in ip variable: c0000160c8
Value of *ip variable: 100
success is not nil
success pointer is nil
```

### Array types

An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length and is never negative.

```bash
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
```

The length is part of the array's type; it must evaluate to a non-negative constant representable by a value of type int. The length of array a can be discovered using the built-in function len. The elements can be addressed by integer indices 0 through len(a)-1. Array types are always one-dimensional but may be composed to form multi-dimensional types.

```
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays. 

Example:
```go
package main

import "fmt"

func main() {

  // var a []string // wrong

  // An array of 10 integers
  var a1 [5]int
  a1[0] = 5
  a1[1] = 4
  a1[2] = 3
  a1[3] = 2
  a1[4] = 1
  fmt.Println(a1)
}
```

Output:
```bash
[5 4 3 2 1]
```

### Slice Types

A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The value of an uninitialized slice is nil. 

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type **[]T** is a slice with elements of **type T**.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```bash
a[low : high]
```

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a: 

```bash
a[1:4]
```

Example:

```go
package main

import "fmt"

func main() {
  primes := [7]int{2, 3, 5, 7, 11, 13, 14}
  
  var p []int = primes[2:5]
  fmt.Println(p)
}
```
Output:
```bash
[5 7 11]
```

A new, initialized slice value for a given element type T is made using the built-in function make, which takes a slice type and parameters specifying the length and optionally the capacity. A slice created with make always allocates a new, hidden array to which the returned slice value refers. That is, executing.

```bash
make([]T, length, capacity)
```

Produces the same slice as allocating an array and slicing it, so these two expressions are equivalent:

```bash
make([]int, 50, 100)
new([100]int)[0:50]
```

Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. With arrays of arrays, the inner arrays are, by construction, always the same length; however with slices of slices (or arrays of slices), the inner lengths may vary dynamically. Moreover, the inner slices must be initialized individually.

 Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array: 

```go
a := make([]int, 4)  // len(a)=4
```
Example:

```go
package main

import "fmt"

func main() {
  a := make([]int,4)
  a[0]=12
  fmt.Println("a", a)

  b := make([]int, 0, 5)
  fmt.Println("b", b)
  
  c := b[:2]
  fmt.Println("c", c)
}
```
Output:
```bash
a [12 0 0 0]
b []
c [0 0]
```

A slice of type T is declared using []T. For example, Here is how you can declare a slice of type int -

```go
// Slice of type `int`
var slice []int

// Slice of type `string`
var slice []string

// Slice of type `string` with parameter variadic
var lang = [...]string{"Erlang", "Elixir", "Haskell", "Clojure", "Scala"}
```

You can create a slice using a slice literal like this -

```go
// Creating a slice using a slice literal
var s = []int{3, 5, 7, 9, 11, 13, 17}
```

The expression on the right-hand side of the above statement is a slice literal. The slice literal is declared just like an array literal, except that you do not specify any size in the square brackets [].

When you create a slice using a slice literal, it first creates an array and then returns a slice reference to it.

Let’s see:
```go
package main

import "fmt"

func main() {
	// Creating a slice using a slice literal
	var s = []string{"@jeffotoni", "@awsbrasil", "@devopsbh", "@go_br"}

	// Short hand declaration
	t := []int{2, 4, 8, 16, 32, 64}

	fmt.Println("s = ", s, len(s))
	fmt.Println("t = ", t, len(t))
}
```
Output:
```bash
s =  [@jeffotoni @awsbrasil @devopsbh @go_br] 4
t =  [2 4 8 16 32 64] 6
```

Since a slice is a segment of an array, we can create a slice from an array.

To create a slice from an array a, we specify two indices low (lower bound) and high (upper bound) separated by a colon

```bash
// Obtaining a slice from an array `a`
a[low:high]
```

The above expression selects a slice from the array a. The resulting slice includes all the elements starting from index low to high, but excluding the element at index high.

```go
package main

import "fmt"

func main() {
	// Creating a slice using a slice literal
	var groups = [5]string{"@awsbrasil", "@devopsbh", "@go_br", "@devopsbr", "@docker"}

	// Creating a slice from the array
	var s []string = groups[2:5]

	s2 := s[1:3]
	s3 := s[:3]
	s4 := s[2:]
	s5 := s[:]

	fmt.Println("Array groups = ", groups, "len:", len(groups), "cap:", cap(groups))
	fmt.Println("Slice s = ", s, "len:", len(s), "cap:", cap(s))
	fmt.Println("Slice s = ", s2, "len:", len(s2), "cap:", cap(s2))
	fmt.Println("Slice s = ", s3, "len:", len(s3), "cap:", cap(s3))
	fmt.Println("Slice s = ", s4, "len:", len(s4), "cap:", cap(s4))
	fmt.Println("Slice s = ", s5, "len:", len(s5), "cap:", cap(s5))
}
```

Output:
```bash
Array groups =  [@awsbrasil @devopsbh @go_br @devopsbr @docker] len: 5 cap: 5
Slice s =  [@go_br @devopsbr @docker] len: 3 cap: 3
Slice s =  [@devopsbr @docker] len: 2 cap: 2
Slice s =  [@go_br @devopsbr @docker] len: 3 cap: 3
Slice s =  [@docker] len: 1 cap: 1
Slice s =  [@go_br @devopsbr @docker] len: 3 cap: 3
```

The copy() function copies elements from one slice to another. Its signature looks like this

```go
func copy(dst, src []T) int
```

It takes two slices - a destination slice, and a source slice. It then copies elements from the source to the destination and returns the number of elements that are copied.

Note that the elements are copied only if the destination slice has sufficient capacity.

```go
package main

import "fmt"

func main() {

	src := []string{"Erlang", "Elixir", "Haskell", "Clojure", "Scala"}
	dest := make([]string, 2)

	numElementsCopied := copy(dest, src)

	fmt.Println("src = ", src)
	fmt.Println("dest = ", dest)
	fmt.Println("Number of elements copied from src to dest = ", numElementsCopied)
}
```

Output:
```bash
src =  [Erlang Elixir Haskell Clojure Scala]
dest =  [Erlang Elixir]
Number of elements copied from src to dest =  2
```
The append() function appends new elements at the end of a given slice. Following is the signature of append function.

```go
func append(s []T, x ...T) []T
```

It takes a slice and a variable number of arguments x …T. It then returns a new slice containing all the elements from the given slice as well as the new elements.

If the given slice doesn’t have sufficient capacity to accommodate new elements then a new underlying array is allocated with bigger capacity. All the elements from the underlying array of the existing slice are copied to this new array, and then the new elements are appended.

However, if the slice has enough capacity to accommodate new elements, then the append() function re-uses its underlying array and appends new elements to the same array.

Let’s see an example:
```go
package main

import "fmt"

func main() {
	slice1 := []string{"Clojure", "Scala", "Elixir"}
	slice2 := append(slice1, "Assembly", "Rust", "Go")
	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))
	slice1[0] = "C++"
	fmt.Println("\nslice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	
	// slice nil
	var s []string

	// Appending to a nil slice
	s = append(s, "Java", "C", "Lisp", "Haskell")
	fmt.Printf("\ns = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
```

Output:
```bash
slice1 = [Clojure Scala Elixir], len = 3, cap = 3
slice2 = [Clojure Scala Elixir Assembly Rust Go], len = 6, cap = 6

slice1 =  [C++ Scala Elixir]
slice2 =  [Clojure Scala Elixir Assembly Rust Go]

s = [Java C Lisp Haskell], len = 4, cap = 4
```

### Struct types

A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique.

```go
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName .
Tag           = string_lit .

// An empty struct.
struct{}

// A struct with 6 fields.
struct {
  x, y int
  u float32
  _ float32  // padding
  A *[]int
  F func()
}
```

A field declared with a type but no explicit field name is called an embedded field. An embedded field must be specified as a type name T or as a pointer to a non-interface type name *T, and T itself may not be a pointer type. The unqualified type name acts as the field name.

```go
// A struct with four embedded fields of types T1, *T2, P.T3 and *P.T4
struct {
  T1        // field name is T1
  *T2       // field name is T2
  P.T3      // field name is T3
  *P.T4     // field name is T4
  x, y int  // field names are x and y
}
```

The following declaration is illegal because field names must be unique in a struct type:

```go
struct {
  T     // conflicts with embedded field *T and *P.T
  *T    // conflicts with embedded field T and *P.T
  *P.T  // conflicts with embedded field T and *T
}
```
A struct is a collection of fields. 

```go
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{10, 201}
  v.X = 4
  fmt.Println(v)
}
```

Output:
```bash
{4 201}
```

### Struct in C 

Before we will see the origin of everything in C, let's see the code written in C and the comments of how it would be in Golang.
And we have it written in Golang, so that you can see how dynamic you were in Go, the C inheritance was simply fantastic and light.

```c
#include <stdio.h>
#include <stdlib.h>

struct Login {
    char *User;
    char *Email;
};

/**
  // In Go would look like this
  type Login struct {
    User string
    Email string
  }
 */

int main() {
/** 
 // In Go 
func main() {
*/
    
    struct Login login, *pointerToLogin;

    /**
    // In Go
    login := Login{User:"jeffotoni", Email:"jef@m.com"}
    */

    login.User = "jeffotoni";
    login.Email = "jef@m.com";

    pointerToLogin = malloc(sizeof(struct Login));
    pointerToLogin->User = "pike";
    pointerToLogin->Email = "pike@g.com";

    /**
    var pointerToLogin  = new(Login)
    pointerToLogin.User  = "pike"
    pointerToLogin.Email = "pike@g.com"
    */
    printf("login vals: %s %s\n", login.User, login.Email);
    printf("pointerToLogin: %p %s %s\n", pointerToLogin, pointerToLogin->User, pointerToLogin->Email);

    /**
     fmt.Printf("login vals: %s %s\n", login.User, login.Email)
     fmt.Printf("pointerToLogin: %v %s %s\n", pointerToLogin, pointerToLogin.User, pointerToLogin.Email);
    */

    free(pointerToLogin);
    return 0;
}
```

Compile the program in C:
```
$ gcc -o struct-1-c struct-1-c.c
```

Output:
```bash
login vals: jeffotoni jef@m.com
pointerToLogin: 0x5627788a0260 pike pike@g.com
```

Confira agora o Código em Go abaixo:

```go
package main

import "fmt"
// #include <stdio.h>
// #include <stdlib.h>

// struct Login {
//     char *User;
//     char *Email;
// };

//In Go would look like this
type Login struct {
    User  string
    Email string
}

// int main() {

// In Go
func main() {

    //struct Login login, *pointerToLogin;

    // In Go
    login := Login{User: "jeffotoni", Email: "jef@m.com"}

    login.User = "jeffotoni"
    login.Email = "jef@m.com"

    // pointerToLogin = malloc(sizeof(struct Login));
    // pointerToLogin->User = "pike";
    // pointerToLogin->Email = "pike@g.com";

    var pointerToLogin = new(Login)
    pointerToLogin.User = "pike"
    pointerToLogin.Email = "pike@g.com"

    //printf("login vals: %s %s\n", login.User, login.Email);
    //printf("pointerToLogin: %p %s %s\n", pointerToLogin, pointerToLogin->User, pointerToLogin->Email);

    fmt.Printf("login vals: %s %s\n", login.User, login.Email)
    fmt.Printf("pointerToLogin: %v %s %s\n", pointerToLogin, pointerToLogin.User, pointerToLogin.Email)

    //free(pointerToLogin);
    return
}
```

```bash
login vals: jeffotoni jef@m.com
pointerToLogin: &{pike pike@g.com} pike pike@g.com
```

### Struct Type Tags Json

We use json tags with omitempty or not to represent our json string, when generated by **json.Marshal** which we will see soon below.

Example:
```go
import (
	//"encoding/json"
	"fmt"
)

// We use the omitempty tag in 'json: field, omitempty'
// when we want the field not to appear after
// making the marshal if it is empty.
type D struct {
	Height int
	Width  int `json:"width,omitempty"`
}

// Fields that do not have the "omitempty" tag will display 
// the same field being empty when generating 
// json through marshal.
type Cat struct {
	Breed    string `json:"breed,omitempty"`
	WeightKg int    `json:WeightKg`
	Size     D      `json:"size,omitempty"`
}

func main() {
	d := Cat{
		//Breed:    "Persa",
		WeightKg: 23,
		//Size:     D{20, 60},
	}
	//b, _ := json.Marshal(d)
	//fmt.Println(string(b))
	fmt.Println(d)
}
```

Output:
```bash
{ 23 {20 60}}
```

In this example we have a struct "ApiLogin" and inside it we have And1 \*struct", ie, referencing a pointer from another struct, beautiful is not.
To initialize and access the "And1" field, we have to use the "&" operator and create the struct at its initialization because it has not been defined yet so it would look like this: **"And1: & struct {City string} {City: "BH"}"**
```go
type ApiLogin struct {
	And1  *struct {
		City string
	}
}
```

In this example we only time the pointer of a struct already defined above.
To access it we do not need to create it again, just refer to it with "&" and this way "And2: &Anddress{}", very nice, right?
```go
type ApiLogin struct {
	And2 *Anddress
}
```

The example below is a way to display various forms of initialization using struct.
```go
package main

import "fmt"

type Anddress struct {
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Zipcode      string `json:"zipcode"`
}

type ApiLogin struct {
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Login string `json:"login"`
	Email string `json:"email"`
	And1  *struct {
		City string
	}

	And2 *Anddress
}

func main() {

	// different ways to inizialize a struct
	//
	//

	apilogin1 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34", And1: &struct{ City string }{City: "BH"}}
	fmt.Println(apilogin1)
	fmt.Println(apilogin1.Name)
	fmt.Println(apilogin1.And1)
	fmt.Println(apilogin1.And1.City)

	apilogin2 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34", And2: &Anddress{City: "BH"}}
	fmt.Println(apilogin2)
	fmt.Println(apilogin2.Name)
	fmt.Println(apilogin2.And2)
	fmt.Println(apilogin2.And2.City)

	var apilogin3 ApiLogin
	fmt.Println(apilogin3)

	apilogin4 := ApiLogin{}
	fmt.Println(apilogin4)

	apilogin5 := &ApiLogin{}
	fmt.Println(apilogin5)

	apilogin6 := new(ApiLogin)
	fmt.Println(apilogin6)
}
```

Output:
```bash
&{@jeffotoni 093.393.334-34   0xc00000e1e0 <nil>}
@jeffotoni
&{BH}
BH
&{@jeffotoni 093.393.334-34   <nil> 0xc000060150}
@jeffotoni
&{BH  }
BH
{    <nil> <nil>}
{    <nil> <nil>}
&{    <nil> <nil>}
&{    <nil> <nil>}
```

The example below is to demonstrate the ease we have when we manipulate structs in Golang.
We are initializing the struct with values and displaying on the screen.

Take a look:
```go
package main

import "fmt"

type jsoninput []struct {
	Data string `json:"data"`
}

func main() {

	data := &jsoninput{{Data: "some data"}, {Data: "some more data"},
		{Data: "some more data"}}
	fmt.Println(data)

	// output:
	//  &[{some data} {some more data} {some more data}]
}
```

```bash
&[{some data} {some more data} {some more data}]
```

Example Struct AWS Sqs Json
```go
type SqsJson struct {
	Type      string `json:"type"`
	MessageId string `json:"messageid"`
	TopicArn  string `json:"topicarn"`
	Message   string `json:"message"`
	//Message          JsonMessage
	Timestamp        string `json:"timestamp"`
	SignatureVersion string `json:"signatureversion"`
	Signature        string `json:"signature"`
	SigningCertURL   string `json:"signingcerturl"`
	UnsubscribeURL   string `json:"unsubscribeurl"`
}

type JsonMessage struct {
	NotificationType string `json:"notificationType"`
	Bounce           struct {
		BounceType        string `json:"bounceType"`
		BounceSubType     string `json:"bounceSubType"`
		BouncedRecipients []struct {
			EmailAddress   string `json:"emailAddress"`
			Action         string `json:"action"`
			Status         string `json:"status"`
			DiagnosticCode string `json:"diagnosticCode"`
		} `json:"bouncedRecipients"`
		Timestamp    time.Time `json:"timestamp"`
		FeedbackID   string    `json:"feedbackId"`
		RemoteMtaIP  string    `json:"remoteMtaIp"`
		ReportingMTA string    `json:"reportingMTA"`
	} `json:"bounce"`

	Mail struct {
		Timestamp        time.Time `json:"timestamp"`
		Source           string    `json:"source"`
		SourceArn        string    `json:"sourceArn"`
		SourceIP         string    `json:"sourceIp"`
		SendingAccountID string    `json:"sendingAccountId"`
		MessageID        string    `json:"messageId"`
		Destination      []string  `json:"destination"`
		HeadersTruncated bool      `json:"headersTruncated"`
		Headers          []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"headers"`

		CommonHeaders struct {
			From    []string `json:"from"`
			ReplyTo []string `json:"replyTo"`
			To      []string `json:"to"`
			Subject string   `json:"subject"`
		} `json:"commonHeaders"`
	} `json:"mail"`
}
```

When the subject is struct we have several possibilities to deal with and work with this feature in Golang, it is practically embedded in everything we will build in Golang, Structs is something powerful in Go to manipulate data and send data all the time between channels using goroutine, be it in queues, database recordings, json and database reads, GraphQL, REST API, SOAP etc ...


### Fatih Structs to Map

Below is a lib that works directly with struct, converting to Maps.
I do not use it in production but for our course it is interesting to analyze.
We know that the more native we are in Golang it will be a good option, but at times we will need some libs to help us.
This project is not maintained anymore and is archived. Feel free to fork and make your own changes if needed.

Structs contains various utilities to work with Go (Golang) structs. It was initially used by me to convert a struct into a map[string]interface{}. With time I've added other utilities for structs. It's basically a high level package based on primitives from the reflect package. Feel free to add new functions or improve the existing code.

Install
```go
go get github.com/fatih/structs
```

To test low rotate the code below:
```go

type Server struct {
	Name        string `json:"name,omitempty"`
	ID          int
	Enabled     bool
	Users       []string // not exported
	http.Server          // embedded
}

func main() {

	// Create a new struct type:

	server := &Server{
		Name:    "gopher",
		ID:      12345678,
		Users:   []string{"jeffotoni", "pike", "dennis", "ken"},
		Enabled: true,
	}

	// struct
	fmt.Println(server)

	// create struct fatih
	s := structs.New(server)
	m := s.Map() // Get a map[string]interface{}
	fmt.Println(m)

	v := s.Values() // Get a []interface{}
	fmt.Println(v)

	f := s.Fields() // Get a []*Field
	fmt.Println(f)

	n := s.Names() // Get a []string
	fmt.Println(n)

	name := s.Field("Name") // Get a *Field based on the given field name

	// Get the underlying value,  value => "gopher"
	value := name.Value().(string)
	fmt.Println(value)

	tagValue := name.Tag("json")
	fmt.Println(tagValue)

	f1, ok := s.FieldOk("Name") // Get a *Field based on the given field name
	fmt.Println(f1, ok)

	n2 := s.Name() // Get the struct name
	fmt.Println(n2)

	h := s.HasZero() // Check if any field is uninitialized
	fmt.Println(h)
}
```

Output:
```bash
&{gopher 12345678 true [jeffotoni pike dennis ken] 
{ <nil> <nil> 0s 0s 0s 0s 0 map[] <nil> <nil> 0 0 {{0 0} 0} <nil> {0 0} map[] map[] <nil> []}}
map[Server:map[TLSConfig:<nil> ReadHeaderTimeout:0s IdleTimeout:0s Addr: Handler:<nil> 
]MaxHeaderBytes:0 TLSNextProto:map[] ConnState:<nil> ErrorLog:<nil> ReadTimeout:0s WriteTimeout:0s] 
Name:gopher ID:12345678 Enabled:true Users:[jeffotoni pike dennis ken]]
[gopher 12345678 true [jeffotoni pike dennis ken]  <nil> <nil> 0s 0s 0s 0s 0 map[] <nil> <nil>]
[0xc000106000 0xc000106090 0xc000106120 0xc0001061b0 0xc000106240]
[Name ID Enabled Users Server]
gopher
name,omitempty
&{{0x626180 0xc0000f6000 408} {Name  0x626180 json:"name,omitempty" 0 [0] false} structs} true
Server
true
```

### Map types

A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type. The value of an uninitialized map is nil.

```bash
MapType     = "map" "[" KeyType "]" ElementType .
KeyType     = Type .
```

The comparison operators == and != must be fully defined for operands of the key type; thus the key type must not be a function, map, or slice. If the key type is an interface type, these comparison operators must be defined for the dynamic key values; failure will cause a run-time panic.

```bash
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
```

The number of map elements is called its length. For a map m, it can be discovered using the built-in function len and may change during execution. Elements may be added during execution using assignments and retrieved with index expressions; they may be removed with the delete built-in function.

A new, empty map value is made using the built-in function make, which takes the map type and an optional capacity hint as arguments:

```bash
make(map[string]int)
make(map[string]int, 100)
```

The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them, with the exception of nil maps. A nil map is equivalent to an empty map except that no elements may be added.

Some examples of map initialization:
```go
package main

import "fmt"

type linkResult struct {
	body string
	urls []string
}

type linkFetcher map[string]*linkResult

func main() {
	// Required to initialize
	// the map with values
	var m1 map[string]int
	var m2 = make(map[string]int)
	var m3 = map[string]int{"population": 500000}
	var m4 = m3
	var m5 map[string]string
	/* create a map*/
	m5 = make(map[string]string)
	fmt.Println(m1, m2, m3, m4, m5)

	var l = linkFetcher{
		"https://golang.org/": &linkResult{
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &linkResult{
			"Packages",
			[]string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		}}

	fmt.Println(l)
}
```
Output:
```bash
map[] map[] map[population:500000] map[population:500000] map[]
```

A map is declared using the following syntax:
```go
var m map[KeyType]ValueType

// For example, Here is how you can declare a map of string keys to int values
var m map[string]int
```

The zero value of a map is nil. A nil map has no keys. Moreover, any attempt to add keys to a nil map will result in a runtime error.

Let’s see an example:
```go
package main

import "fmt"

func main() {

	// Required to initialize
	// the map with values
	var m map[string]int
	fmt.Println(m)
	if m == nil {
		fmt.Println("is nil")
	}
	// Attempting to add keys
	// to a nil map will result in a runtime error
	//m["population"] = 500000
	//fmt.Println(m)
}
```

Output:
```bash
map[]
is nil
```

You can initialize a map using the built-in make() function. You just need to pass the type of the map to the make() function as in the example below. The function will return an initialized and ready to use map.

```go
// Initializing a map using the built-in make() function
var m = make(map[string]int)
```

Example:
```go
package main

import "fmt"

func main() {

	// Required to initialize
	// the map with values
	var m = make(map[string]int)
	fmt.Println(m)
	if m == nil {
		fmt.Println("is nil")
	}
	m["population"] = 500000
	fmt.Println(m)
}
```
Output:
```bash
map[]
map[population:500000]
```

The example below introduces the creation of a map getting a struct done.
When do we use an empty struct?
There are some scenarios when we have a large amount of access in our API, but when I say large is> 10k = 10k of requests per second, in this scenario when we do our handler, we can implement a channel receiving an empty struct {} so that we can get put on a channel and process everything in with more security.
We will show more ahead this very legal approach.

Example:
```go
package main

import "fmt"

func main() {
	s := "key"
	seen := make(map[string]struct{}) // set of strings
	// ...
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
		// ...first time seeing s...
	}
	fmt.Println(seen)
}
```

Output:
```go
map[key:{}]
```

we can also make our map key receive an empty struct.
Well we know that it gets any type ie the struct can be complete without it's done too.

Example:
```go
type T struct{}

func main() {
	s := T{}
	seen := make(map[struct{}]struct{}) // set of strings
	// ...
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
		// ...first time seeing s...
	}
	fmt.Println(seen)
}
```

Output:
```go
map[{}:{}]
```

### Map literals continued

A map literal is a very convenient way to initialize a map with some data. You just need to pass the key-value pairs separated by colon inside curly braces.

```go
package main

import "fmt"

func main() {
	var m = map[string]string{
		"Brasil": "Brasilia",
		"EUA":    "Washington, D.c",
		"Italy":  "Roma",
		"France": "Paris",
		"Japan":  "Toquio",
	}

	fmt.Println(m)
}
```

Output:
```bash
map[Italy:Roma France:Paris Japan:Toquio Brasil:Brasilia EUA:Washington, D.c]
```

So you can check for the existence of a key in a map by using the following two-value assignment.
The boolean variable ok will be true if the key exists, and false otherwise.

```go
value, ok := m[key]
```

Consider the following map for example:
```go
var C = map[string]string{
		"Brasil": "Brasilia",
		"EUA":    "Washington, D.c",
		"Italy":  "Roma",
		"France": "Paris",
		"Japan":  "Toquio",
	}
```

```go
capital, ok := C["EUA"]  // "Washington, D.c", true
```

However, If you try to access a key that doesn’t exist, then the map will return an empty string "" (zero value of strings), and false

```go
capital, ok := C["África do Sul"]  // "", false
```

You can delete a key from a map using the built-in delete() function. The syntax looks like this.

```go
// Delete the `key` from the `map`
delete(map, key)
```
Example:
```go
package main

import "fmt"

func main() {
	var country = map[string]string{
		"Brasil": "Brasilia",
		"EUA":    "Washington, D.c",
		"Italy":  "Roma",
		"France": "Paris",
		"Japan":  "Toquio",
	}
	delete(country, "Japan")
	delete(country, "Italy")
	fmt.Println(country)
}
```

Output:
```bash
map[Brasil:Brasilia EUA:Washington, D.c France:Paris]
```
 
If the top-level type is just a type name, you can omit it from the elements of the literal. 
```go
package main

import "fmt"

type Login struct {
	User, Login, Email string
}

// passing a struct as parameter
// for our struct map
var m = map[string]Login{
	"jeffotoni": {"jeffotoni", "jeff", "jeff@gm.com"},
	"Google":    {"root", "super", "google@gm.com"},
}

func main() {
	fmt.Println(m)
}
```go

```bash
map[jeffotoni:{jeffotoni jeff jeff@gm.com} Google:{root super google@gm.com}]
```

### Channel Types

A channel provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type. The value of an uninitialized channel is nil.

```bash
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .
```

The optional <- operator specifies the channel direction, send or receive. If no direction is given, the channel is bidirectional. A channel may be constrained only to send or only to receive by conversion or assignment.

```bash
chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints
```

The <- operator associates with the leftmost chan possible:

```bash
chan<- chan int    // same as chan<- (chan int)
chan<- <-chan int  // same as chan<- (<-chan int)
<-chan <-chan int  // same as <-chan (<-chan int)
chan (<-chan int)
```

A new, initialized channel value can be made using the built-in function make, which takes the channel type and an optional capacity as arguments:

```bash
make(chan int, 100)
```

The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready. Otherwise, the channel is buffered and communication succeeds without blocking if the buffer is not full (sends) or not empty (receives). A nil channel is never ready for communication.

A channel may be closed with the built-in function close. The multi-valued assignment form of the receive operator reports whether a received value was sent before the channel was closed.

A single channel may be used in send statements, receive operations, and calls to the built-in functions cap and len by any number of goroutines without further synchronization. Channels act as first-in-first-out queues. For example, if one goroutine sends values on a channel and a second goroutine receives them, the values are received in the order sent.

Let me show you an example:
```go

package main

import (
  "fmt"
  "os"
  "time"
)

type Promise struct {
  Result chan string
  Error  chan error
}

var (
  ch1  = make(chan *Promise)  // received a pointer from the structure
  ch2  = make(chan string, 1) // allows only 1 channels
  ch3  = make(chan int, 2)    // allows only 2 channels
  ch4  = make(chan float64)   // has not been set can freely receive
  ch5  = make(chan []byte)    // by default the capacity is 0
  ch6  = make(chan bool, 1)   // non-zero capacity
  ch7  = make(chan time.Time, 2)
  ch8  = make(chan struct{}, 2)
  ch9  = make(chan struct{})
  ch10 = make(map[string](chan int)) // map channel
  ch11 = make(chan error)
  ch12 = make(chan error, 2)
  // receives a zero struct
  ch14 <-chan struct{}
  ch15 = make(<-chan bool)          // can only read from
  ch16 = make(chan<- []os.FileInfo) // // can only write to
  // holds another channel as its value
  ch17 = make(chan<- chan bool) // // can read and write to
)

// Parameters of Func
// (jobs <-chan int, results chan<- int)

// Receives Value, only read
// jobs <-chan int //receives the value

// Receives Channel, only write
// results chan<- int // receive channel
// or
// results chan int // receive channel

// Receives Channel variadic
// results ...<-chan int

func main() {

  ch2 <- "okay"
  defer close(ch2)
  fmt.Println(ch2, &ch2, <-ch2)

  ch7 <- time.Now()
  ch7 <- time.Now()
  fmt.Println(ch7, &ch7, <-ch7)
  fmt.Println(ch7, &ch7, <-ch7)
  defer close(ch7)

  ch3 <- 1 // okay
  ch3 <- 2 // okay
  // deadlock
  // ch3 <- 3 // does not accept any more values, if you do it will error : deadlock
  defer close(ch3)
  fmt.Println(ch3, &ch3, <-ch3)
  fmt.Println(ch3, &ch3, <-ch3)

  ch10["lambda"] = make(chan int, 2)
  ch10["lambda"] <- 100
  defer close(ch10["lambda"])
  fmt.Println(<-ch10["lambda"])
}
```

Output:
```bash
0xc000056180 0x55bb00 okay
0xc0000561e0 0x55bb28 2019-01-25 15:11:41.982906669 -0200 -02 m=+0.000147197
0xc0000561e0 0x55bb28 2019-01-25 15:11:41.982906922 -0200 -02 m=+0.000147409
0xc00001e0e0 0x55bb08 1
0xc00001e0e0 0x55bb08 2
100
```

### Blank identifier

The blank identifier is represented by the underscore character **_**. It serves as an anonymous placeholder instead of a regular (non-blank) identifier and has special meaning in declarations, as an operand, and in assignments.

Example:
```bash
// function statement
func f() (int, string, error)

// function return
_, _, _ := f()
```

### Interface types

**An interface is two things:**
 - it is a set of methods
 - but it is also a type

The __interface{} type__, the empty interface is the interface that has __no methods__

Since there is no implements keyword, all types implement at least zero methods, and satisfying an interface is done automatically, all types satisfy the empty interface.
That means that if you write a function that takes an interface{} value as a parameter, you can supply that function with any value.

Example:
```go
func DoSomething(v interface{}) {
   // ...
}

var Msg interface{}

type Stringer interface {
    String() string
}
```

### Here's an interface as a method

An interface type specifies a method set called its interface. A variable of interface type can store a value of any type with a method set that is any superset of the interface. Such a type is said to implement the interface. The value of an uninitialized variable of interface type is nil.


```bash
InterfaceType      = "interface" "{" { MethodSpec ";" } "}" .
MethodSpec         = MethodName Signature | InterfaceTypeName .
MethodName         = identifier .
InterfaceTypeName  = TypeName .
```

As with all method sets, in an interface type, each method must have a unique non-blank name.

```go
// A simple File interface
interface {
  Read(b Buffer) bool
  Write(b Buffer) bool
  Close()
}
```

More than one type may implement an interface. For instance, if two types S1 and S2 have the method set

```bash
func (p T) Read(b Buffer) bool { return … }
func (p T) Write(b Buffer) bool { return … }
func (p T) Close() { … }
```

(where T stands for either S1 or S2) then the File interface is implemented by both S1 and S2, regardless of what other methods S1 and S2 may have or share.

A type implements any interface comprising any subset of its methods and may therefore implement several distinct interfaces. For instance, all types implement the empty interface:

```bash
interface{}
```

Similarly, consider this interface specification, which appears within a type declaration to define an interface called Locker:

```go
type Locker interface {
  Lock()
  Unlock()
}
```

If S1 and S2 also implement

```bash
func (p T) Lock() { … }
func (p T) Unlock() { … }
```

they implement the Locker interface as well as the File interface.

An interface T may use a (possibly qualified) interface type name E in place of a method specification. This is called embedding interface E in T; it adds all (exported and non-exported) methods of E to the interface T.

```go
type ReadWriter interface {
  Read(b Buffer) bool
  Write(b Buffer) bool
}

type File interface {
  ReadWriter  // same as adding the methods of ReadWriter
  Locker      // same as adding the methods of Locker
  Close()
}

type LockedFile interface {
  Locker
  File        // illegal: Lock, Unlock not unique
  Lock()      // illegal: Lock not unique
}
```

An interface type T may not embed itself or any interface type that embeds T, recursively.
```go
// illegal: Bad cannot embed itself
type Bad interface {
  Bad
}

// illegal: Bad1 cannot embed itself using Bad2
type Bad1 interface {
  Bad2
}
type Bad2 interface {
  Bad1
}
```

Example: 
```go
package main

import (
  "fmt"
)

type R struct {
  R string
}

type Iread interface {
  Read() string
}

func (r *R) Read() string {
  return fmt.Sprintf("Only: call Read")
}

func Call(ir Iread) string {
  return fmt.Sprintf("Read: %s", ir.Read())
}

func main() {
  var iread Iread
  r := R{"hello interface"}
  // A way to use Interface
  iread = &r
  fmt.Println(iread, r)
  fmt.Println(iread.Read())

  // Second way to access interface
  r2 := R{"hello interface call"}
  fmt.Println(Call(&r2))
}
```

Output:
```bash
&{hello interface} {hello interface}
Only: call Read
Read: Only: call Read
```

###  Interface as type

Interfaces as type __interface{}__ means you can put value of any type, including your own custom type. All types in Go satisfy an empty interface (interface{} is an empty interface).
In your example, Msg field can have value of any type. 


```go
var val interface{} // element type of m is assignable to val
``` 

```go
type Empty interface {
    /* it has no methods */
}

// Because, Empty interface has no methods, 
// following types satisfy the Empty interface
var a Empty

a = 60
a = 10.5
a = "Lambda Man"
```

Interfaces as types looks at another example below:
```go
package main

import (
  "fmt"
)

type MyStruct struct {
  Msg interface{}
}

func main() {
  b := MyStruct{}
  // string
  b.Msg = "5"
  fmt.Printf("%#v %T \n", b.Msg, b.Msg) // Output: "5" string

  // int
  b.Msg = 5
  fmt.Printf("%#v %T", b.Msg, b.Msg) //Output:  5 int

  // map
  b.Msg = map[string]string{"population": "500000", "language": "sueco"}
  fmt.Printf("%#v %T", b.Msg, b.Msg) //Output:  5 int
}
```

### Exercise one
Exercise:
Fill in the struct JsonMessage AWS above, initialize the struct and fill in the fields, and make a fmt.Println to display the filled fields.
To be more readable you can separate into each struct type struct.

### Control structures
---

### Control

The control structures are:

__For, If, else, else if__

And some statments between them: __break, continue, switch, case and goto__.

### Control Return

Statements control execution.

```bash
Statement =
  Declaration | LabeledStmt | SimpleStmt |
  GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
  FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
  DeferStmt .

SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
```

A terminating statement prevents execution of all statements that lexically appear after it in the same block. The following statements are terminating:

1. A "return" or "goto" statement.

Return:
```go
package main

func main() {
  println(Lambda())
  return
}

func Lambda() string {

  return "Lambda"
}
```

Output:
```bash
Lambda
```

### Control Goto

Goto:
```go
package main

import "fmt"

func main() {
  n := 0

LOOP1:
  n++
  if n == 10 {
    println("fim")
    return
  }

  if n%2 == 0 {
    goto LOOP2
  } else {

    fmt.Println("n", n, "LOOP1 here...")
    goto LOOP1
  }

LOOP2:
  fmt.Println("n", n, "LOOP2 here...")
  goto LOOP1

}
```

Output:
```bash
n 1 LOOP1 here...
n 2 LOOP2 here...
n 3 LOOP1 here...
n 4 LOOP2 here...
n 5 LOOP1 here...
n 6 LOOP2 here...
n 7 LOOP1 here...
n 8 LOOP2 here...
n 9 LOOP1 here...
fim
```

### Control if else

2. An "if" statement in which:
      - the "else" branch is present, and
      - both branches are terminating statements.
```go
package main

func main() {
  n := 100
  if n > 0 && n <= 55 {
    println("n > 0 or n <= 55")
  } else if n > 56 && n < 70 {
    println("n > 56 and n < 70")
  } else {

    if n >= 100 {
      println(" else here.. n > 100")
    } else {
      println(" else here.. n > 70")
    }
  }
}
```

Output:
```bash
else here.. n > 100
```

### Control For break continue

3. A "for" statement in which:
      - there are no "break" statements referring to the "for" statement, and
      - the loop condition is absent.
      - there are "continue"
      - A "break" statement terminates execution of the innermost "for", "switch", or "select" statement within the same 

```go
package main

func main() {
  // will be looping infinitely
  // for {
  // }

  // will run only once and exit
  for {
    break
  }

  n := 5
  for n > 0 {
    n--
    println(n)
  }
  // Output:
  // 4
  // 3
  // 2
  // 1
  // 0

  // declaring i no and increasing i
  for i := 0; i < 5; i++ {
    println(i)
  }
  // Output:
  // 0
  // 1
  // 2
  // 3
  // 4

  n = 5
  for i := 0; i < n; i++ {
    if i <= 2 {
      continue
    } else {
      println("i > 2 = ", i)
    }
  }

  // Output:
  // i > 2 =  3
  // i > 2 =  4

  n = 5
  for i := 0; i < n; i++ {
    if i == 2 {
      break
    } else {
      println("i: ", i)
    }
  }
  // Output:
  // i:  0
  // i:  1
  
  // infinitely
  for ; ; i++ {
    println("i: ", i)
  }
  // Output:
  // i:  1
  // i:  2
  // ..
  // ..
}
```

### Control Switch case break

4. A "switch" statement in which:
      - there are no "break" statements referring to the "switch" statement,
      - there is a default "case", and
      - the statement lists in each case, including the default, end in a terminating statement, or a possibly labeled "fallthrough" statement.

```go
package main

func main() {
  j := 10
  i := 0
  switch j {
  case 11:
    println("here: 11")
    break
  default:
    println("here default")
    break
  }

  // infinitely
  for ; ; i++ {

    switch i {
    case 5:
      goto LABELS
    case i:
      println("i: ", i)
      break
    default:
      println("default: ", i)
    }
  }

LABELS:
  f()

}

func f() {
  println("goto fim")
}
```

Output:
```bash
here default
i:  0
i:  1
i:  2
i:  3
i:  4
goto fim
```

### Control Label

5. A labeled statement labeling a terminating statement.

```go
package main

func main() {
  i := 0
  // infinitely
  for ; ; i++ {
    for {
      if i == 10 {
        goto LABEL
      }
      i++
    }
  }

LABEL:
  f(i)

}

func f(i int) {
  println("label fim i:", i)
}
```

Output:
```bash
label fim i: 10
```

All other statements are not terminating.

A statement list ends in a terminating statement if the list is not empty and its final non-empty statement is terminating. 

A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map, or values received on a channel. For each entry it assigns iteration values to corresponding iteration variables if present and then executes the block. 

```bash
RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .
```

### Control Range

The expression on the right in the "range" clause is called the range expression, which may be an array, pointer to an array, slice, string, map, or channel permitting receive operations. As with an assignment, if present the operands on the left must be addressable or map index expressions; they denote the iteration variables. If the range expression is a channel, at most one iteration variable is permitted, otherwise there may be up to two. If the last iteration variable is the blank identifier, the range clause is equivalent to the same clause without that identifier. 

```bash
Range expression                          1st value          2nd value

array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
string          s  string type            index    i  int    see below  rune
map             m  map[K]V                key      k  K      m[k]       V
channel         c  chan E, <-chan E       element  e  E
```

See an example below, with various uses using Range:
```go

package main

import "fmt"

func main() {

  // string arrays / slice
  var lang = [...]string{"Erlang", "Elixir", "Haskell", "Clojure", "Scala"}

  // screen list
  fmt.Println(lang)

  // list the positions srtring arrays
  for k, v := range lang {
    fmt.Println(k, v)
  }

  /* create a map*/
  countryCapitalMap := map[string]string{"Brasil": "Brasilia", "EUA AMERICA": "Washington, D.C.", "France": "Paris", "Italy": "Roma", "Japan": "Tokyo"}

  /* print map using key-value*/
  for country, capital := range countryCapitalMap {
    fmt.Println("Capital of", country, "is", capital)
  }

  // channel
  jobs := make(chan int, 3)

  // for channel
  for j := 1; j <= 3; j++ {
    jobs <- j
  }
  // println(<-jobs)
  // println(<-jobs)
  // println(<-jobs)

  // close
  /* date is required for range to work*/
  close(jobs)

  /* This syntax is valid too. */
  for range jobs {
  }

  /* it is mandatory to close the channels to be able to scroll */
  for ch := range jobs {
    println(ch)
  }

  // it is not an array struct, it will range from error.
  sa := struct{ nick string }{"@jeffotoni"}
  fmt.Println(sa.nick)

  // here the range will be able to list all struct
  a := []struct{ nick string }{{"@devopsbr"}, {"@go_br"}, {"@awsbrasil"}, {"@go_br"}, {"@devopsbh"}}
  for i, v := range a {
    fmt.Println(i, v.nick)
  }

  // struct pointer
  var testdata *struct {
    a *[3]int
  }
  for i := range testdata.a {
    // testdata.a is never evaluated; len(testdata.a) is constant
    // i ranges from 0 to 2
    fmt.Println(i)
  }

  // new example interface and range
  var key string
  var val interface{} // element type of m is assignable to val
  m := map[string]int{"mon": 0, "tue": 1, "wed": 2, "thu": 3, "fri": 4, "sat": 5, "sun": 6}
  for key, val = range m {
    fmt.Println(key, val)
  }
}
```

Output:
```bash
[Erlang Elixir Haskell Clojure Scala]
0 Erlang
1 Elixir
2 Haskell
3 Clojure
4 Scala
Capital of Brasil is Brasilia
Capital of EUA AMERICA is Washington, D.C.
Capital of France is Paris
Capital of Italy is Roma
Capital of Japan is Tokyo
@jeffotoni
0 @devopsbr
1 @go_br
2 @awsbrasil
3 @go_br
4 @devopsbh
0
1
2
sat 5
sun 6
mon 0
tue 1
wed 2
thu 3
fri 4
```
### Errors
---
 
 The predeclared type error is defined as 

 ```bash
 type error interface {
	Error() string
}
 ```
It is the conventional interface for representing an error condition, with the nil value representing no error. For instance, a function to read data from a file might be defined: 

```bash
func Read(f *File, b []byte) (n int, err error)
```

### Introduction Errors

Error handling in Golang is very simple to deal with. There is no try, catch or exceptions.
The error is handled with every call of some function.
As we show "error" is an interface, and much used.

When we talk about handling errors in Golang everything is very simple, a good practice is to return an error in the functions that we created and treat them.

Let's see in practice how it works.

```go
package main

import "fmt"

func main() {
	var error error
	fmt.Println(error)
}
```

Output:
```bash
<nil>
```
### How Error Control Works

Example:
```go
package main

import (
	"encoding/json"
	"fmt"
)

var Error error
var b []byte

func main() {
	fmt.Println(Error)
	cpu := make(chan int)
	// Create JSON from the instance data.
	b, Error = json.Marshal(cpu)
	if Error != nil {
		fmt.Println(Error)
	} else {
		fmt.Println(string(b))
	}
}
```

Output:
```bash
<nil>
json: unsupported type: chan int
```

### Errors.New

```go
package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(fvalue float64) (float64, error) {
	if fvalue < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt [" + fmt.Sprintf("%.2f", fvalue) + "]")
	}
	return math.Sqrt(fvalue), nil
}

func main() {
	result, err := Sqrt(-33)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt (-33) = [", result, "]")
	}

	result, err = Sqrt(81)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt(81) = [", result, "]")
	}
}
```
Output:
```bash
Math: negative number passed to Sqrt [-33.00]
Sqrt(81) = [ 9 ]
```

### Custom Errors

```go
package main

import "fmt"

// It's possible to use custom types as `error`s by
// implementing the `Error()` method on them.
type MyError struct {
    line int
    msg  string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("%d - %s", e.line, e.msg)
}

func MyFunc(line int) (int, error) {
    if line < 100 {

        // In this case we use `&MyError` syntax to build
        // a new struct, supplying values for the two
        // fields `arg` and `prob`.
        return -1, &MyError{line, "can't work with it"}
    }
    return line, nil
}

func main() {

    // The one loops below test out each of our
    // error-returning functions.
    for _, i := range []int{200, 99} {
        if r, e := MyFunc(i); e != nil {
            fmt.Println("MyFunc failed:", e)
        } else {
            fmt.Println("MyFunc worked in line: ", r)
        }
    }
}
```

Output:
```bash
MyFunc worked in line:  200
MyFunc failed: 99 - can't work with it
```

### fmt Errorf

```go
package main

import (
    "fmt"
    "math"
)

func circleArea(radius float64) (float64, error) {
    if radius < 0 {
        return 0, fmt.Errorf("Failed, radius %0.2f is less than zero", radius)
    }
    return math.Pi * radius * radius, nil
}

func main() {
    radius := -80.0
    area, err := circleArea(radius)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Area of circle: %0.2f", area)
}
```
Output:
```bash
Area calculation failed, radius -80.00 is less than zero
```

### Functions 
---

Declaring and Calling Functions in Golang. 
In Golang, we declare a function using the func keyword. 
A function has a name, a list of comma-separated input parameters along with their types, the result type(s), and a body. 
The input parameters and return type(s) are optional for a function.

Example of declaring and Calling Functions in Golang:
```go
func Sum(x float64, y float64) float64 {
	return (x + y) / 2
}
```

### Introduction Function

Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
A function type denotes the set of all functions with the same parameter and result types. The value of an uninitialized variable of function type is nil. 

Some possibilities:
```bash
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```

```go
package main

func F1(name string) string {
    return "Hello, " + name
}

func main() {

    println(F1("@go_br"))
}
```

Output:
```bash
Hello, @go_br
```

### Return multiple values

Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

```go
package main

import "fmt"

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
func F1() (string, int, error) {
    return "@go_br", 100, nil
}

func main() {

    // Here we use the 2 different return values from the
    // call with _multiple assignment_.
    a, b, err := F1()
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(err)

    // If you only want a subset of the returned values,
    // use the blank identifier `_`.
    a, _, err = F1()
    fmt.Println(a)
    fmt.Println(err)
}
```

Output:
```bash
@go_br
100
<nil>
@go_br
<nil>
```

### Variadic Functions

Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
Here’s a function that will take an arbitrary number of ints as arguments.
Variadic functions can be called in the usual way with individual arguments.

```go
package main

import (
	"fmt"
	"strings"
)

// This variadic function takes an arbitrary number of ints as arguments.
func Show(names ...string) {
	fmt.Print("The Len of ", len(names)) // Also a variadic function.
	appends := ""
	for _, name := range names {
		appends += name + ","
	}
	appends = strings.Trim(appends, ",")
	fmt.Println(" is: [", appends, "]") // Also a variadic function.
}

func main() {

	// Variadic functions can be called in the usual way with individual
	// arguments.
	Show("C", "C++")
	Show("Clojure", "Elixir", "Scala")

	// If you already have multiple args in a slice, apply them to a variadic
	// function using func(slice...) like this.
	nums := []string{"Algol", "C", "C++", "Golang"}
	Show(nums...)
}
```

Output:
```bash
The Len of 2 is: [ C,C++ ]
The Len of 3 is: [ Clojure,Elixir,Scala ]
The Len of 4 is: [ Algol,C,C++,Golang ]
```

### functions as a parameter

You can pass function as parameter to a Go function. Here is an example of passing function as parameter to another Go function.

```go
package main

import "fmt"

// convert types take an int
// and return a string value.
type fn func(int) string

func f1(param int) string {
	return fmt.Sprintf("param is %v", param)
}

func f2(param int) string {
	return fmt.Sprintf("param is %v", param)
}

func test(f fn, val int) {
	fmt.Println(f(val))
}

func main() {
	test(f1, 432)
	test(f2, 874)
}
```

Output:
```bash
param is 432
param is 874
```

```go
package main

import "fmt"

// --------------------------------
func Square(num int) int {
	return num * num
}

func Mapp(f func(int) int, List []int) []int {
	var a = make([]int, len(List), len(List))
	for index, val := range List {

		a[index] = f(val)
	}
	return a
}

func main() {
	list := []int{454, 455, 86, 988}
	result := Mapp(Square, list)
	fmt.Println(result)
}
```

Output:
```bash
[206116 207025 7396 976144]
```

### Closures

Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.
This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.

Example:
```go
package main

import "fmt"

func PlusX() func(v int) int {
    return func(v int) int {
        return v + 5
    }
}

func plusXandY(x int) func(v int) int {
    return func(v int) int {
        return v + x
    }
}

func main() {
    p := PlusX()
    fmt.Printf("5+15: %d\n", p(15))

    px := plusXandY(6)
    fmt.Printf("6+10: %d\n", px(10))
}
```

Output:
```bash
5+15: 20
6+10: 16
```

### Recursion

Go supports recursive functions. Here’s a classic factorial example.

A simple example:
```go
package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
```

Output:
```bash
5040
```

Listing all subdirectories directories:
```go
package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// Listing our example directory
// Listing our example directory recursively
func main() {

	// Capturing our path that is in the environment
	gopath := os.Getenv("PWD")

	// directory we want to list
	gopath += "/examples"

	// Making call in function
	list := ListDir(gopath)

	// listing the function return
	for i, p := range list {
		fmt.Printf("[%d:%s===%s]\n", i, path.Dir(p), path.Base(p))
	}
}

// This function uses pkg filepath.Walk, it is
// a recursive function, where it will go through
// our directory and its subfolders.
func ListDir(rootpath string) []string {

	list := make([]string, 0)

	// recursive call
	// This function receives a function as parameter and after going through all levels it ends.
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".git" && filepath.Ext(path) != ".svn" {
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return list
}
```

Output:
```bash
[0:$(pwd)/examples/bufio.writer===main.go]
[1:$(pwd)/examples/error===error1.go]
[2:$(pwd)/examples/error===error2.go]
[3:$(pwd)/examples/error===error3.go]
[4:$(pwd)/examples/error===error4.go]
...
...
```

### Asynchronous Functions

In golang to perform asynchronous functions we use the keyword **"go"** it is responsible for putting the functions to be executed concurrently.
A **"go"** statement starts the execution of a function call as an independent concurrent thread of control, or goroutine, within the same address space. 

```go
GoStmt = "go" Expression .
```

```go
go Server()
go func(ch chan<- bool) { for { sleep(10); ch <- true }} (c)
```

Example:
```go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.Intn(500)

func pinger() {
	time.Sleep(time.Duration(r) * time.Microsecond)
	fmt.Println("pinger")
}

func ponger() {
	time.Sleep(time.Duration(r) * time.Microsecond)
	fmt.Println("ponger")
}

func printer() {
	time.Sleep(time.Duration(r) * time.Microsecond)
	fmt.Println("printer")
}

func main() {
	// making functions
	// calls asynchronously
	go pinger()
	go ponger()
	go printer()

	// Waiting to press any key to end
	var input string
	fmt.Scanln(&input)
}
```
Output one:
```bash
ponger
pinger
printer
```

Output two:
```bash
pinger
ponger
printer
```

### Defer

A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking. 

```go
DeferStmt = "defer" Expression .
```

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for expression statements.

Each time a "defer" statement executes, the function value and parameters to the call are evaluated as usual and saved anew but the actual function is not invoked. Instead, deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred. If a deferred function value evaluates to nil, execution panics when the function is invoked, not when the "defer" statement is executed. 

Examples:
```go
defer unlock(l)
defer myFunc()
defer close(channel)
defer fmt.Print(x)
defer db.Close()
defer f.Close()
defer res.Body.Close()
```

## Lab 03 Json with Golang

### Json
---

Package json implements encoding and decoding of JSON as defined in **RFC 7159**. The mapping between JSON and Go values is described in the documentation for the **Marshal** and **Unmarshal** functions.

### introduction

**JSON** (JavaScript Object Notation) is a simple data interchange format. Syntactically it resembles the objects and lists of JavaScript. It is most commonly used for communication between web back-ends and JavaScript programs running in the browser, but it is used in many other places, too. Its home page, json.org, provides a wonderfully clear and concise definition of the standard. 

With the [json package](https://golang.org/pkg/encoding/json/) it's a snap to read and write **JSON** data from your Go programs.

### Json Marshal Encode

Marshal returns the JSON **encoding** of **v**.

Marshal traverses the value v recursively. If an encountered value implements the Marshaler interface and is not a nil pointer, Marshal calls its MarshalJSON method to produce JSON. If no MarshalJSON method is present but the value implements encoding.TextMarshaler instead, Marshal calls its MarshalText method and encodes the result as a JSON string.

```go
func Marshal(v interface{}) ([]byte, error)
```

Given the Go data structure, Message, 

```go
type ApiLogin struct {
   Name  string `json:"name"`
   Cpf   string `json:"cpf"`
}
```

and an instance of Message 

```go
m := ApiLogin{"Jefferson", "033.343.434-89"}
```

we can marshal a JSON-encoded version of m using json.Marshal: 

```go
b, err := json.Marshal(m)
```

If all is well, err will be nil and b will be a []byte containing this JSON data: 

```go
b == []byte(`{"Name":"Lambda Man","Cpf":"033.343.434-89"}`)
```

Check out the complete code:
```go
import (
	"encoding/json"
	"fmt"
	"log"
)

type ApiLogin struct {
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

func main() {

	a := ApiLogin{"Jefferson", "033.343.434-89"}
	fmt.Println(a)

	m, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
	}
	// show bytes
	fmt.Println(m)

	// show string json
	fmt.Println(string(m))
}
```

Output:
```bash
{Jefferson 033.343.434-89}
[123 34 110 97 109 101 34 58 34 74 101 102 102 
101 114 115 111 110 34 44 34 99 112 102 34 58 
34 48 51 51 46 51 52 51 46 52 51 52 45 56 57 34 125]
{"name":"Jefferson","cpf":"033.343.434-89"}
```

### json MarshalIndent

MarshalIndent is like Marshal but applies Indent to format the output. Each JSON element in the output will begin on a new line beginning with prefix followed by one or more copies of indent according to the indentation nesting. 

```go
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
```

Example:
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ApiLogin struct {
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

func main() {

	a := ApiLogin{"Jefferson", "033.343.434-89"}
	// improving output for json format viewing
	json, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}
```

Output:
```bash
{
	"name": "Jefferson",
	"cpf": "033.343.434-89"
}
```

### Option Omitempty

The "omitempty" option specifies that the field should be omitted from the encoding if the field has an empty value, defined as false, 0, a nil pointer, a nil interface value, and any empty array, slice, map, or string. 

```go
// Field appears in JSON as key "login".
Login string `json:"login"`

// Field appears in JSON as key "email" and
// the field is omitted from the object if its value is empty,
// as defined above.
Email string `json:"email,omitempty"`

// Field appears in JSON as key "nick" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Nick string `json:",omitempty"`

// Field is ignored by this package.
Level int `json:"-"`

// Field appears in JSON as key "-".
LastEmail string `json:"-,"`
```
Look at the example below:
```go

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Login struct {
	// Field appears in JSON as key "login".
	Login string `json:"login"`

	// Field appears in JSON as key "email" and
	// the field is omitted from the object if its value is empty,
	// as defined above.
	Email string `json:"email,omitempty"`

	// Field appears in JSON as key "nick" (the default), but
	// the field is skipped if empty.
	// Note the leading comma.
	Nick string `json:",omitempty"`

	// Field is ignored by this package.
	Level int `json:"-"`

	// Field appears in JSON as key "-".
	LastEmail string `json:"-,"`
}

func main() {

	l := Login{Login: "Austin", Email: "austin@go.com", Nick: 
	 "", Level: 1000, LastEmail: "austin@gmail.com"}
	fmt.Println(l)

	m, err := json.Marshal(l)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(m))
}
```
Output:
```bash
{Austin austin@go.com Aust 1000 austin@gmail.com}
{"login":"Austin","email":"austin@go.com","-":"austin@gmail.com"}
```

Only data structures that can be represented as valid JSON will be encoded:

```bash
- JSON objects only support strings as keys; to encode a Go map type it must be of 
  the form map[string]T (where T is any Go type supported by the json package).
- Channel, complex, and function types cannot be encoded.
- Cyclic data structures are not supported; they will cause Marshal to go into an infinite loop.
- Pointers will be encoded as the values they point to (or 'null' if the pointer is nil).
```

The json package only accesses the exported fields of struct types (those that begin with an uppercase letter). Therefore only the the exported fields of a struct will be present in the JSON output. 

In this example we work with pointers to reference the struct within another struct, and another point is that we declare the struct within the struct itself.
With this we have different ways to initialize and fill the fields of our structs.
Let's see how it works? Check out the example below.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Anddress struct {
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Zipcode      string `json:"zipcode"`
}

type ApiLogin struct {
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Login string `json:"login"`
	Email string `json:"email"`
	And1  *struct {
		City string
	}

	And2 *Anddress
}

func main() {

	apilogin1 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34",
		And1: &struct{ City string }{City: "BH"}, And2: &Anddress{City: "BH"}}
	//m, err := json.Marshal(apilogin1)
	m, err := json.MarshalIndent(apilogin1, "", "\t")

	if err != nil {
		log.Println(err)
	}
	//fmt.Println("apilogin1 initialized")
	//fmt.Println(apilogin1)

	//fmt.Println("\njson.Marshal returning bytes")
	//fmt.Println(m)

	fmt.Println("\njson.Marshal as string")
	fmt.Println(string(m))
}
```

```bash
json.Marshal as string
{
	"name": "@jeffotoni",
	"cpf": "093.393.334-34",
	"login": "",
	"email": "",
	"And1": {
		"City": "BH"
	},
	"And2": {
		"city": "BH",
		"neighborhood": "",
		"zipcode": ""
	}
}
```

In this other example we take a short excerpt from the json of an AWS SES service, it notifies via Json the Bounces of the sent emails, we are going to check the completion of our fields in the struct and transform them into json.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	type BouncedRecipients []struct {
		EmailAddress string `json:"emailAddress"`
		Action       string `json:"action"`
		Status       string `json:"status"`
		//DiagnosticCode string `json:"diagnosticCode"`
	}

	type Bounce struct {
		BounceType string `json:"bounceType"`
		//BounceSubType string    `json:"bounceSubType"`
		//Timestamp     time.Time `json:"timestamp"`
		BR BouncedRecipients
	}

	type JsonMessage struct {
		NotificationType string `json:"notificationType"`
		B                Bounce
		From             []string `json:"from"`
	}

	l := &JsonMessage{NotificationType: "38733773737xxxx",
		B: Bounce{BounceType: "bounce type",
			BR: BouncedRecipients{
				{"devops@g.com", "permanet", "error"}, {"lambdaman@g.com", "complaint", "success"}}},
		From: []string{"from1@m.com", "from2@gm.com"}}

	m, err := json.Marshal(l)

	if err != nil {
		log.Println(err)
	}

	//fmt.Println("\033[1;40mlinkFetcher initialized\033[0m")
	//fmt.Println(l)

	//fmt.Println("\n\033[1;42mjson.Marshal returning bytes\033[0m")
	//fmt.Println(m)

	fmt.Println("\n\033[1;41mjson.Marshal as string\033[0m")
	fmt.Println(string(m))
}
```

Output:
```bash
json.Marshal as string
{
	"notificationType": "38733773737xxxx",
	"B": {
		"bounceType": "bounce type",
		"BR": [
			{
				"emailAddress": "devops@g.com",
				"action": "permanet",
				"status": "error"
			},
			{
				"emailAddress": "lambdaman@g.com",
				"action": "complaint",
				"status": "success"
			}
		]
	},
	"from": [
		"from1@m.com",
		"from2@gm.com"
	]
}
```

In the example below there is an entire field in lowercase, this field the **json.Marshal** function **will not be** able to do marshal, because the field initializes with the lowercase letter **"myname"**, so it works the first letter has that is **"Myname"**
Another legal point our struct is set to receive "N" values as an array of struct **v[]struct"**
And inside the struct we have a string-type field **"[]string"**.
too cool is not? Let's see how we do the initialization and fill in the values, look at the code below.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type linkResult []struct {
	Body   string   `json:"body"`
	Urls   []string `json:"urls"`
	myname string   `json:"myname"`
}

func main() {
	
	// creating our struct
	// with some fields
	// interesting as [] string
	// and showing how to initialize them
	var ll = linkResult{
		{
			Body:   "The Go Programming Language",
			Urls:   []string{"https://golang.org/pkg/", "https://golang.org/cmd/"},
			myname: "lambda man",
		},
		{
			Body:   "Package",
			Urls:   []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/fmt/"},
			myname: "go_br in action",
		}}

	//m0, err := json.Marshal(ll)
	m0, err := json.MarshalIndent(ll, "", "\t")

	if err != nil {
		log.Println(err)
	}
	//fmt.Println("linkFetcher initialized")
	//fmt.Println(ll)

	//fmt.Println("\njson.Marshal returning bytes")
	//fmt.Println(m0)

	fmt.Println("\njson.Marshal as string")
	fmt.Println(string(m0))
}
```

Output:
```bash
linkFetcher initialized
[{The Go Programming Language [https://golang.org/pkg/ https://golang.org/cmd/] lambda man} {Package [https://golang.org/ https://golang.org/cmd/ https://golang.org/pkg/fmt/] go_br in action}]

json.Marshal returning bytes
[91 10 9 123 10 9 9 34 98 111 100 121 34 58 32 34 84 104 101 32 71 111 32 80 114 111 103 114 97 109 109 105 110 103 32 76 97 110 103 117 97 103 101 34 44 10 9 9 34 117 114 108 115 34 58 32 91 10 9 9 9 34 104 116 116 112 115 58 47 47 103 111 108 97 110 103 46 111 114 103 47 112 107 103 47 34 44 10 9 9 9 34 104 116 116 112 115 58 47 47 103 111 108 97 110 103 46 111 114 103 47 99 109 100 47 34 10 9 9 93 10 9 125 44 10 9 123 10 9 9 34 98 111 100 121 34 58 32 34 80 97 99 107 97 103 101 34 44 10 9 9 34 117 114 108 115 34 58 32 91 10 9 9 9 34 104 116 116 112 115 58 47 47 103 111 108 97 110 103 46 111 114 103 47 34 44 10 9 9 9 34 104 116 116 112 115 58 47 47 103 111 108 97 110 103 46 111 114 103 47 99 109 100 47 34 44 10 9 9 9 34 104 116 116 112 115 58 47 47 103 111 108 97 110 103 46 111 114 103 47 112 107 103 47 102 109 116 47 34 10 9 9 93 10 9 125 10 93]

json.Marshal as string
[
	{
		"body": "The Go Programming Language",
		"urls": [
			"https://golang.org/pkg/",
			"https://golang.org/cmd/"
		]
	},
	{
		"body": "Package",
		"urls": [
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/"
		]
	}
]
```

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type crud struct {
    Create   bool
    Retrieve bool
    Update   bool
    Delete   bool
}

type crudWithDetail struct {
    crud
    Detail bool
}

type acceleration struct {
    crudWithDetail
    Participant crudWithDetail
    Challenge   crudWithDetail
}

type acl struct {
    User         crud
    Acceleration acceleration
}

func main() {

    // on this boot we show exactly how it would be using non-embedded structs but all separated and making type of each one of them
    // all initialization is more friendly, although having nesting between structs is much better if we have built structs ie
    // structs being declared directly inside structs what they call embedded.

    // Example embedded strcut would
    // type my struct {
    //    M struct {
    //       Name string
    //   }
    // }

    j := acl{
        User: crud{Create: true, Retrieve: false, Update: false, Delete: true},
        Acceleration: acceleration{
            Participant:    crudWithDetail{Detail: true, crud: crud{Create: false, Retrieve: false, Update: true, Delete: false}},
            Challenge:      crudWithDetail{Detail: false, crud: crud{Create: true, Retrieve: true, Update: true, Delete: false}},
            crudWithDetail: crudWithDetail{Detail: false, crud: crud{Create: true, Retrieve: false, Update: true, Delete: false}}},
    }

    json, err := json.MarshalIndent(j, "", "\t")
    if err != nil {
        log.Println(err)
    }
    fmt.Println(string(json))
}
```

Output:
```bash
{
	"User": {
		"Create": true,
		"Retrieve": false,
		"Update": false,
		"Delete": true
	},
	"Acceleration": {
		"Create": true,
		"Retrieve": false,
		"Update": true,
		"Delete": false,
		"Detail": false,
		"Participant": {
			"Create": false,
			"Retrieve": false,
			"Update": true,
			"Delete": false,
			"Detail": true
		},
		"Challenge": {
			"Create": true,
			"Retrieve": true,
			"Update": true,
			"Delete": false,
			"Detail": false
		}
	}
}
```

Let's now present 3 ways to declare a map [string] interfaces {} inside a struct and generate a json of it all after initializing our struct.

Take a look at the example below:
```go

package main

import (
	"encoding/json"
	"fmt"
)

// create type map
type body map[string]interface{}

// struct Message
type Message struct {
	Subject  string                 `json:"subject"`
	TimeSent int64                  `json:"sent"`
	Body     body                   `json:"body"`  // embedded => Body of the type map[string]interface{}
	Body2    map[string]interface{} `json:"body2"` // Body of the type map[string]interface{}
	Status   string                 `json:"status"`
}

func main() {

	// our interface map body
	// This map can be dynamic,
	// create types as needed, because the interface{}
	// is for this purpose, it creates dynamic types.
	b := body{
		"Data": body{
			"login":  "jeffotoni",
			"level":  1000,
			"create": true,
		}, "Anddress": body{
			"City": "Jersey City",
			"Zip":  "34.566-333",
			"Fone": body{
				"fone1": "87-77047345",
				"fone2": "83-93483838",
			},
		},
	}

	// create type
	type out map[string]interface{}

	// a way to initialize the struct
	// and the map contained within the struct
	s := &Message{
		Subject:  "Test Struct to Map",
		TimeSent: 345,
		Body:     b,
		Body2:    out{"Instance": "c5.xlarge", "vCpu": "4", "Ecu": "16", "Mem": "8"},
		Status:   "success",
	}
	fmt.Println(s)

	// Another way to implement is to generate
	// the map interface {} inside the initialization
	// of the struture.
	s2 := &Message{
		Subject:  "Test Struct to Map",
		TimeSent: 345,
		Body:     b,
		Body2:    map[string]interface{}{"Instance": "c5.xlarge", "vCpu": "4", "Ecu": "16", "Mem": "8"},
		Status:   "success",
	}
	fmt.Println(s2)

	// converting everything to Json
	m, err := json.Marshal(s2)
	if err != nil {
		fmt.Println(err)
	}

	// sending json to the screen
	fmt.Println("##### json")
	fmt.Println(string(m))
}
``` 

Output:
```bash
&{Test Struct to Map 345 map[Data:map[create:true login:jeffotoni level:1000] Anddress:map[City:Jersey City Zip:34.566-333 Fone:map[fone1:87-77047345 fone2:83-93483838]]] map[Instance:c5.xlarge vCpu:4 Ecu:16 Mem:8] success}
&{Test Struct to Map 345 map[Data:map[login:jeffotoni level:1000 create:true] Anddress:map[City:Jersey City Zip:34.566-333 Fone:map[fone1:87-77047345 fone2:83-93483838]]] map[Ecu:16 Mem:8 Instance:c5.xlarge vCpu:4] success}
##### json
{
	"subject": "Test Struct to Map",
	"sent": 345,
	"body": {
		"Anddress": {
			"City": "Jersey City",
			"Fone": {
				"fone1": "87-77047345",
				"fone2": "83-93483838"
			},
			"Zip": "34.566-333"
		},
		"Data": {
			"create": true,
			"level": 1000,
			"login": "jeffotoni"
		}
	},
	"body2": {
		"Ecu": "16",
		"Instance": "c5.xlarge",
		"Mem": "8",
		"vCpu": "4"
	},
	"status": "success"
}
```

### Initialized collections of data

Some ways to anonymously declare and initialize types and collections of types in marshal to transform into Json.

Take a look at the example below:
```go
func main() {

    // let's try something simple
    // to understand what's
    // really going on
    m, _ := json.Marshal(false)
    fmt.Println(m)         // show as bytes
    fmt.Println(string(m)) // show as string

    m2, _ := json.Marshal(-1)
    fmt.Println(string(m2))

    m3, _ := json.Marshal(0)
    fmt.Println(string(m3))

    m4, _ := json.Marshal(102.3)
    fmt.Println(string(m4))

    m5, _ := json.Marshal("DevOpsBH")
    fmt.Println(string(m5))

    // For this to occur we need to pass a collection
    // with fields, something like a struct, a slice,
    // maps, interfaces {} Let's see below an example
    m6, _ := json.Marshal([...]string{"Elixir", "Scala"})
    fmt.Println(string(m6))

    m7, _ := json.Marshal(map[string]string{"twitter": "@jeffotoni"})
    fmt.Println(string(m7))

    m8, _ := json.Marshal(map[string]interface{}{"instagram": "jeffotoni", "langs": struct{ G string }{G: "gophers"}})
    fmt.Println(string(m8))

    m9, _ := json.Marshal(map[string]struct{ L string }{"jeff": {L: "@Ricardo Maraschini"}})
    fmt.Println(string(m9))
}
```

Output:
```bash
```
### Json NewEncoder

NewEncoder returns a new encoder that writes to w. 

```go
func NewEncoder(w io.Writer) *Encoder
```

With this function we can implement our own custom Marshal function

Check out the example below:
```go

package main

import (
    "bytes"
    "encoding/json"
    "fmt"
)

const (
    empty = ""
    tab   = "\t"
)

func main() {

    //json, err := MyJson(map[string]string{"twitter": "@jeffotoni"})
    json, err := MyJson(map[string]interface{}{"instagram": "jeffotoni", "langs": struct{ G string }{G: "gophers"}})

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(json)
}

func MyJson(data interface{}) (string, error) {
    buffer := new(bytes.Buffer)
    encoder := json.NewEncoder(buffer)
    encoder.SetIndent(empty, tab)
    // encode...
    err := encoder.Encode(data)
    if err != nil {
        return empty, err
    }
    return buffer.String(), nil
}
```

Output:
```
{
	"instagram": "jeffotoni",
	"langs": {
		"G": "gophers"
	}
}
```
The above example shows the use of the json.NewEncoder () function. Encode () to transform the data collection into JSON, very cool, right?


### Json Unmarshal Decode

Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an "InvalidUnmarshalError".

Unmarshal uses the inverse of the encodings that Marshal uses, allocating maps, slices, and pointers as necessary, with the following additional rules:

To unmarshal JSON into a pointer, Unmarshal first handles the case of the JSON being the JSON literal null. In that case, Unmarshal sets the pointer to nil. Otherwise, Unmarshal unmarshals the JSON into the value pointed at by the pointer. If the pointer is nil, Unmarshal allocates a new value for it to point to.

To unmarshal JSON into a value implementing the Unmarshaler interface, Unmarshal calls that value's UnmarshalJSON method, including when the input is a JSON null. Otherwise, if the value implements encoding.TextUnmarshaler and the input is a JSON quoted string, Unmarshal calls that value's UnmarshalText method with the unquoted form of the string.

To unmarshal JSON into a struct, Unmarshal matches incoming object keys to the keys used by Marshal (either the struct field name or its tag), preferring an exact match but also accepting a case-insensitive match. By default, object keys which don't have a corresponding struct field are ignored (see Decoder.DisallowUnknownFields for an alternative).

To unmarshal a JSON array into a slice, Unmarshal resets the slice length to zero and then appends each element to the slice. As a special case, to unmarshal an empty JSON array into a slice, Unmarshal replaces the slice with a new empty slice. 

Consider this JSON data, stored in the variable b: 

 ```go
b := []byte(`{"Name":"DevOps","Age":10,"Parents":["Google","Aws"]}`)
 ```

Without knowing this data's structure, we can decode it into an interface{} value with Unmarshal:

```go
var f interface{}
err := json.Unmarshal(b, &f)
```

To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:

```go
bool, for JSON booleans
float64, for JSON numbers
string, for JSON strings
[]interface{}, for JSON arrays
map[string]interface{}, for JSON objects
nil for JSON null
```

All of the content below has been described and created examples using unmarshal with interfaces {} which is one of the most complex ways of handling a value that we do not know the type they can be born dynamically, and for this we need to understand a little more how interfaces { } and how to do type assertions in Go.

Soon after we will return in unmarshal using structs.

### Generic JSON with interface{} and assertion

The interface{} (empty interface) type describes an interface with zero methods. Every Go type implements at least zero methods and therefore satisfies the empty interface.

The empty interface serves as a general container type:

```go
var i interface{}
i = "DevOps BH"
i = 2019
i = 9.5
```

A type assertion accesses the underlying concrete type: 

```go
r := i.(float64)
fmt.Println("the circle's area", math.Pi*r*r)
```

Or, if the underlying type is unknown, a type switch determines the type: 

```go
switch v := i.(type) {
case int:
    fmt.Println("twice i is", v*2)
case float64:
    fmt.Println("the reciprocal of i is", 1/v)
case string:
    h := len(v) / 2
    fmt.Println("i swapped by halves is", v[h:]+v[:h])
default:
    // i isn't one of the types above
}
```

The code below was perfect for us to understand, and to know how we make insertions using Golang.

Example:
```go
package main

import "fmt"

func main() {
	var i interface{} = "DevOpsBh"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}
```

Output:
```bash
DevOpsBh
DevOpsBh true
0 false
```

**Type switches**

A type switch is a construct that permits several type assertions in series.
A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

```bash
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```
The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type. 

Check out the code below:
```go
package main

import "fmt"

func doInterface(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	doInterface(33)
	doInterface("DevOpsBH")
	doInterface("Lambda")
	doInterface(true)
}
```

Output:
```bash
Twice 33 is 66
"DevOpsBH" is 8 bytes long
"Lambda" is 6 bytes long
I don't know about type bool!
```

The code below will initialize the interface causing it to receive several different types.

Check out all the code below:
```go
package main

import "fmt"

func main() {
	var i interface{}
	fmt.Println(i)

	i = "DevOps BH"
	fmt.Println(i)

	i = 2019
	fmt.Println(i)

	i = 9.5
	fmt.Println(i)

	i = [...]string{"Pike", "Jeffotoni", "Thompson", "Griesemer"}
	fmt.Println(i)

	i = map[string]interface{}{"Lang": []string{"Go", "Rust", "Scala", "Elixir"}}
	fmt.Println(i)

	//struct{ City string }{City: "BH"}
	i = map[struct{ L string }]interface{}{{L: "Lang"}: []string{"Go", "Rust", "Scala", "Elixir"}}
	fmt.Println(i)
}
```

Output:
```bash
<nil>
DevOps BH
2019
9.5
[Pike Jeffotoni Thompson Griesemer]
map[Lang:[Go Rust Scala Elixir]]
map[{Lang}:[Go Rust Scala Elixir]]
```
It is noticed that the interface created accepted all the types that was passed to it, was created dynamically.

**Syntax of type assertion is defined as:** 

```go
PrimaryExpression.(Type)
```

At this point the Go value in f would be a map whose keys are strings and whose values are themselves stored as empty interface values: 

```go
f = map[string]interface{}{
    "Name": "DevOps",
    "Age":  10,
    "Parents": []interface{}{
        "Google",
        "Aws",
    },
}
```

To access this data we can use a type assertion to access f's underlying map[string]interface{}: 

```go
m := f.(map[string]interface{})
```

We can then iterate through the map with a range statement and use a type switch to access its values as their concrete types: 

```go
for k, v := range m {
    switch vv := v.(type) {
    case string:
        fmt.Println(k, "is string", vv)
    case float64:
        fmt.Println(k, "is float64", vv)
    case []interface{}:
        fmt.Println(k, "is an array:")
        for i, u := range vv {
            fmt.Println(i, u)
        }
    default:
        fmt.Println(k, "is of a type I don't know how to handle")
    }
}
```

In this way you can work with unknown JSON data while still enjoying the benefits of type safety.

See full code below:
```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	b := []byte(`{"Name":"DevOps","Age":10,"Company":["Google","Aws"]}`)

	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		fmt.Println(err)
	}

	fmt.Println(f)

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	//fmt.Println()
}
```

```bash
map[Name:DevOps Age:10 Company:[Google Aws]]
Name is string DevOps
Age is float64 10
Company is an array:
0 Google
1 Aws
```

When we work with empty interfaces, we have to do asserts, so we can use them and treat them correctly, because in Golang everything is static and typed, if we do not assert we will not be able to capture the values of types that are in the interface that can be any type.

Check the code below:
```go
	case string:
	// ..
	case float64:
	// ..
	case [] interface {}:
	// ..
	case int:
	// ..
	case bool:
	// ..
	case map[string]interface{}:
```

### Dynamic type

Besides static type that all variables have (it’s a type from variable’s declaration), variables of interface type also have a dynamic type. It’s a type of value currently set in interface type variable. Over the course of program execution variable of interface type has the same static type but its dynamic type can change as different values implementing desired interface will be assigned:

Check out the example:
```go
package main

import "fmt"

type I interface {
	comeon()
}

type A struct{}

func (a A) comeon() {}

type B struct{}

func (b B) comeon() {}

func main() {
	var i I
	i = A{} // dynamic type of i is A
	fmt.Printf("%T\n", i.(A))
	i = B{} // dynamic type of i is B
	fmt.Printf("%T\n", i.(B))
}
```

```bash
main.A
main.B
```

### What is reflection?

Reflection is the ability of a program to inspect its variables and values at run time and find their type. You might not understand what this means but that's alright. You will get a clear understanding of reflection by the end of this section, so stay with me.

Reflection is a very powerful and advanced concept in Go and it should be used with care. It is very difficult to write clear and maintainable code using reflection. It should be avoided wherever possible and should be used only when absolutely necessary.

This is very powerful, we managed to sweep our entire struct using asserts and reflect, we got the name of the struct, name of the fields, their values, their tags, this feature is used to do Parse in Files, like Yaml, Toml, Json etc ......

Let's take another example when using Unmarshal in Interfaces{}
Example:
```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// map more interface
	input := map[string]interface{}{
		"Key1": []string{"some", "key"},
		"key3": nil,
		"val":  2,
		"val2": "str",
		"val3": 4,
	}

	fmt.Println(input)
	for key, value := range input {
		slice, ok := value.([]string)
		if ok {
			input["Count"+key] = len(slice)
		}
	}
	fmt.Println(input)

	// Encode to Json
	m, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
	}

	// convert byte to string
	fmt.Println(string(m))

	// interface empty
	var f interface{}

	// Decode json
	if err := json.Unmarshal(m, &f); err != nil {
		fmt.Println(err)
	}

	// show screen
	fmt.Println(f)

	fmt.Println("### loop interface")
	m2 := f.(map[string]interface{})
	loopI(m2)
}

// loop in interface
func loopI(m2 map[string]interface{}) {
	for k, v := range m2 {

		// assert interface
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
```

Output:
```bash
map[key3:<nil> Key1:[some key] val:2 val2:str val3:4]
map[key3:<nil> Key1:[some key] val:2 val2:str val3:4 CountKey1:2]
{"CountKey1":2,"Key1":["some","key"],"key3":null,"val":2,"val2":"str","val3":4}
map[CountKey1:2 Key1:[some key] key3:<nil> val:2 val2:str val3:4]
### loop interface
Key1 is an array:
0 some
1 key
key3 is of a type I don't know how to handle
val is float64 2
val2 is string str
val3 is float64 4
CountKey1 is float64 2
```

### Json Unmarshal Structs

We saw the most complex part of using Unmarshal in dynamic composites using interfaces {} and had to make assertions to capture the values.
Now let's use Unmarshal in structures, in this scenario the structure is something static, it has to be always predefined.

When we develop our API, we will realize that we will have to do Unmarshal and Marshal all the time of receiving and sending information using json.

So if we get angry at Unmarshal, Marshal, Interfaces {} and Structs, our APIs will be a little tricky to develop.

Let's take a look at the code below that represents exactly the way we use Unmarshal with Structs.

Example:
```go
import (
	"encoding/json"
	"fmt"
)

func main() {

	// byte with json
	var jsonBlob = []byte(`[
	{"Name": "Golang", "Level": "10"},
	{"Name": "Rust",    "Level": "7"}
]`)

	// struct that is our model
	type Lang struct {
		Name  string
		Level string
	}

	var langs []Lang

	// convert Json to Struct
	err := json.Unmarshal(jsonBlob, &langs)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("\n%+v", langs)
	fmt.Printf("\n%+v", langs[0].Name)
	fmt.Printf("\n%+v", langs[0].Level)
	fmt.Printf("\n%+v", langs[1].Name)
	fmt.Printf("\n%+v", langs[1].Level)
}
```

Output:
```bash
[{Name:Golang Level:10} {Name:Rust Level:7}]
Golang
10
Rust
7
```

In other words, we get a string of bytes that is a Json and fill in the values of the struct with those elements.
Using the json.Unmarshal function, Fantastic it.
In other words, Json has to be compatible with the struct that is our model.

### MarshalJSON and UnmarshalJSON

There is the possibility of customizing the native method of Marshal and Unmarshal, that is to override with other functionalities without changing their behavior as a whole, but to increase in the already existing method.
Very cool this feature, I do not need to rewrite the whole method in yes add a behavior.

First we create our struct with the fields we want to add the new rule as the example below:

```go
type Email struct {
Email string
}
```

Soon after the initialization of my struct where it feeds the information will receive this struct as parameter as the example below:

```go
a: = struct {
Login string
Email
Time Time
}
```

Ready now we simply implement our method always with the name **MarshalJSON** referencing this struct that we created **"Email"**.

```go
func (and Email) MarshalJSON () ([] byte, error) {
return [] byte (fmt.Sprintf (`"% s "`, checkEmail (e.Email))), nil
}
```

The checkEmail function is a regular expression that we created to validate the email.
Ready when running the Marshal function it will icorporate this method that you implemented.
See the example below with the complete code.

```go
package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

type Email struct {
	Email string
}

type Time struct {
	time.Time
}

// layout
const layout = "2006-01-02"

var regxmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Time.Format(layout))), nil
}

func (e Email) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, checkEmail(e.Email))), nil
}

func checkEmail(email string) string {
	if regxmail.MatchString(email) {
		return email
	} else {
		return "-invalid-"
	}
}

func main() {
	a := struct {
		Login string
		Email Email
		Time  Time
	}{
		Login: "devops",
		Email: Email{"jeffotoni-go.com"},
		Time:  Time{time.Now()},
	}

	fmt.Println(a)
	json, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json))
}
```

Output:
```bash
{devops {jeffotoni-go.com} 2019-01-30 22:41:00.506602023 -0200 -02 m=+0.000508436}
{
	"Login": "devops",
	"Email": "-invalid-",
	"Time": "2019-01-30"
}
```

The Unmarshal version is a little more complex, since we will need to scan the received structures. The email was added in a string map, to scan the field and add the functionality to validate the email.

Check out the complete code below.
```go
package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Email struct {
	Email string
}

type Time struct {
	time.Time
}

// layout
const layout = "2006-01-02"

// regex to email
var regxmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if string(b) == `null` {
		*t = Time{}
		return
	}
	t.Time, err = time.Parse(layout, string(b))
	return
}

// Unmarshal Json
func (t *Email) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if string(b) == `null` {
		return
	}
	var stuff map[string]string
	err = json.Unmarshal(b, &stuff)
	if err != nil {
		return err
	}
	for key, value := range stuff {
		if strings.ToLower(key) == "email" {
			t.Email = checkEmail(value)
		}
	}
	return nil
}

func checkEmail(email string) string {
	if regxmail.MatchString(email) {
		return email
	} else {
		return "-invalid-"
	}
}

func main() {
	a := struct {
		Login string
		Email Email
		Time  Time
	}{}

	b := []byte(`{"Login":"devops","Email":{"Email":"devops-go.com"},"Time":"2019-01-30"}`)
	err := json.Unmarshal(b, &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
```

Output:
```bash
{devops {-invalid-} 2019-01-30 00:00:00 +0000 UTC}
```

### Parse Json
---

###  Reading and Parsing a JSON File

Let's do parse in Json file, there are several libs to do this, let's create our own, simple to understand the whole process and of course get mad at Parse.

The file below is saved in the folder **parsejson** and we will create our parse within this directory.

```json
{
  "users": [
    {
      "name" : "Devopsbh",
      "type" : "Reader",
      "age" : 29,
      "social" : {
        "facebook" : "https://facebook.com/devopsbh",
        "twitter" : "https://twitter.com/devopsbh",
        "instagram" : "https://instagram.com/devopsbh"
      }
    },
    {
      "name" : "Jefferson",
      "type" : "Author",
      "age" : 160,
      "social" : {
        "facebook" : "https://facebook.com/jeffotoni",
        "twitter" : "https://twitter.com/jeffotoni"
        "instagram" : "https://instagram.com/jeffotoni"
      }
    }
  ]
}
```

We'll be using the packages in order to open up our users.json file from our filesystem. Once we have opened the file, we'll defer the closing of the file to the end of the function so that we can work with the data inside of it.

Check the code below:
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Getwd())
	// Open our jsonFile
	jsonFile, err := os.Open("./users.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
}
```

### Parsing with Structs	

We have a few options when it comes to parsing the JSON that is contained within our users.json file. We could either unmarshal the JSON using a set of predefined structs, or we could unmarshal the JSON using a map[string]interface{} to parse our JSON into strings mapped against arbitrary data types.

If you know the structure that you are expecting then I would recommend going down the verbose route and defining your structs like so:

```go
package main

func main() {
	// Users struct which contains
	// an array of users
	type Users struct {
		Users []User `json:"users"`
	}

	// Social struct which contains a
	// list of links
	type Social struct {
		Facebook string `json:"facebook"`
		Twitter  string `json:"twitter"`
		Instagram  string `json:"instagram"`
	}

	// User struct which contains a name
	// a type and a list of social links
	type User struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Age    int    `json:"Age"`
		Social Social `json:"social"`
	}
}
```

Once we have these in place, we can use them to unmarshal our JSON.
Unmarshalling our JSON - Once we’ve used the os.Open function to read our file into memory, we then have to convert it toa byte array using ioutil.ReadAll. Once it’s in a byte array we can pass it to our json.Unmarshal() method.

Remember if the json file is a wrong comma, it will not convert the file to json format that has it right.
Now you can check the code below.

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook  string `json:"facebook"`
	Twitter   string `json:"twitter"`
	Instagram string `json:"instagram"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

func main() {

	//fmt.Println(os.Getwd())
	// Open our jsonFile
	jsonFile, err := os.Open("./users.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// show json
	fmt.Println(string(byteValue))

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	fmt.Println(users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		fmt.Println("Twitter Url: " + users.Users[i].Social.Twitter)
		fmt.Println("Instagram Url: " + users.Users[i].Social.Instagram)
	}
}
```

Output:
```bash
{[{Devopsbh Reader 29 {https://facebook.com/devopsbh https://twitter.com/devopsbh https://instagram.com/devopsbh}} {Jefferson Author 160 {https://facebook.com/jeffotoni https://twitter.com/jeffotoni https://instagram.com/jeffotoni}}]}
User Type: Reader
User Age: 29
User Name: Devopsbh
Facebook Url: https://facebook.com/devopsbh
Twitter Url: https://twitter.com/devopsbh
Instagram Url: https://instagram.com/devopsbh
User Type: Author
User Age: 160
User Name: Jefferson
Facebook Url: https://facebook.com/jeffotoni
Twitter Url: https://twitter.com/jeffotoni
Instagram Url: https://instagram.com/jeffotoni
```
### Parsing with Map and Interface

Sometimes, going through the process of creating structs for everything can be somewhat time consuming and overly verbose for the problems you are trying to solve. In this instance, we can use standard interfaces{} in order to read in any JSON data:

Format Json Example:

```json
{
  "name" : "Devopsbh",
  "city" : "Belo Horizonte",
  "age" : 29
}
```

**How can we do Unmarshal in this file and access the fields?**

When we read json and do the parse, there are several boring points we have to deal with, and an excellent strategy is to divide to conquer. The above example is an arrayless json with no complexity, a json in its simplest possible format.

Our Unmarshal received **var result map[string]interface{}** to be able to play all the data in an interface map so we have something dynamic without having to define the fields the interface makes the magic of accepting any type.
In this example we have no problem in directly accessing the map fields, in json format we have the keys and values and ready we can direct access.

Check the code below:
```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	// Open our jsonFile
	byteJson, err := ioutil.ReadFile("./user-only-0.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}

	// show screen
	fmt.Println("Successfully Opened users.json")

	// defining the interface map that will receive
	// the file and the format is dynamic
	var result = make(map[string]interface{})

	// let's now run our Unmarshal and convert to objects
	json.Unmarshal(byteJson, &result)

	// show all
	fmt.Println(result)

	// how is a map I can do exactly
	// the syntax below
	val, ok := result["name"].(string)
	fmt.Println(val, ok)

	// type string
	fmt.Println(result["name"].(string))

	// type string
	fmt.Println(result["city"].(string))

	// type float64
	fmt.Println(result["age"].(float64))
}
```

Output:
```bash
Successfully Opened users.json
map[age:3 name:Devopsbh city:Belo Horizonte]
Devopsbh true
Devopsbh
Belo Horizonte
3
```

Now let's complicate our Json file a bit, let's put an Array of users and let's see how we should proceed to access the keys and values of this Json.

```json
{
  "users":
  [
    {
    "name" : "Andre Almar",
    "type" : "knowledge",
    "age" : 25,
    "nick": "@andrealmar"
    },
    {
    "name" : "Jefferson",
    "type" : "knowledge",
    "age" : 160,
    "nick": "@jeffotoni"
    },
    {
    "name" : "Ieso Dias",
    "type" : "knowledge",
    "age" : 23,
    "nick": "@iesodias"
    }
  ]
}
```

Now our file is an Array of values, let's see how we do it to access the keys and values of our new Json.

Let's take a look at the code:
```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {

	// We can use the ioutil.ReadFile which would already
	// return bytes and we would not need to use two functions
	// ioutil.ReadFile("./user-only.json")

	// Open our jsonFile
	jsonFile, err := os.Open("./user-only.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened users.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// reading archive content
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// defining the interface map that will receive
	// the file and the format is dynamic
	var result map[string]interface{}

	// the syntax below is also allowed
	//var result = make(map[string]interface{})

	// let's now run our Unmarshal and convert to objects
	json.Unmarshal([]byte(byteValue), &result)

	// Let's access the interface on our map
	obj := result["users"]

	// Here completely changes from the
	// previous example we now have an [] interface,
	// ie an array of interfaces a slice.
	objInterface := obj.([]interface{})

	// Now as we have multiple users,
	// we'll loop through to fetch them
	// and access each user's key
	// and value in the slice
	for line, keyOne := range objInterface {

		fmt.Println("#### line ", line)

		// ValueOf returns a new Value initialized
		// to the concrete value stored in the interface i.
		val := reflect.ValueOf(keyOne)

		// A Kind represents the
		// specific kind of type
		// that a Type represents.
		if val.Kind() == reflect.Map {

			// Loop the map
			for _, key := range val.MapKeys() {

				// Index of Map
				v := val.MapIndex(key)

				// get the value of type Interface
				switch value := v.Interface().(type) {

				// v.Interface().(type)
				// here test the types
				case int:
					fmt.Println(key, value)
				case float64:
					fmt.Println(key, value)
				case string:
					fmt.Println(key, value)
				case bool:
					fmt.Println(key, value)
				default:
					fmt.Println("not found")
				}
			}
		}
	}
}
```

Output:
```bash
Successfully Opened users.json
k: 0 v: map[name:Andre Almar city:Sao Paulo age:25 nick:@andrealmar]
name Andre Almar
city Sao Paulo
age 25
nick @andrealmar
k: 1 v: map[name:Jefferson city:Belo Horizonte age:160 nick:@jeffotoni]
name Jefferson
city Belo Horizonte
age 160
nick @jeffotoni
k: 2 v: map[age:23 nick:@iesodias name:Ieso Dias city:Mato Grosso]
name Ieso Dias
city Mato Grosso
age 23
nick @iesodias
```

**What if I have a recursion in Json?**

```json
{
  "users":
  [
    {
      "name" : "Joelson",
      "city" : "Porto Alegre",
      "age" : 39,
      "social" : {
        "facebook" : "https://facebook.com/joelson",
        "twitter" : "https://twitter.com/joelson",
        "instagram" : "https://instagram.com/joelson"
      },
      "fone" : {
        "cell" : "5531987387246",
        "resid1" : "55314565678",
        "job" : "55314785679"
      }
    },
    {
      "name" : "Jefferson",
      "city" : "Belo Horizonte",
      "age" : 160,
      "social" : {
        "facebook" : "https://facebook.com/jeffotoni",
        "twitter" : "https://twitter.com/jeffotoni",
        "instagram" : "https://instagram.com/jeffotoni"
      },
      "fone" : {
        "cell" : "5531987387246",
        "resid1" : "55314565678",
        "job" : "55314785679"
      }
    }
  ]
}
```

```go

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {

	// Open our jsonFile
	byteJson, err := ioutil.ReadFile("./users.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}

	// show screen
	fmt.Println("Successfully Opened users.json")

	// defining the interface map that will receive
	// the file and the format is dynamic
	var result = make(map[string]interface{})

	// let's now run our Unmarshal and convert to objects
	json.Unmarshal(byteJson, &result)

	obj := result["users"]

	// get []interface
	objData := obj.([]interface{})

	// recursive
	DecodeMapVetInterface(objData)
}

func DecodeMapVetInterface(objData []interface{}) {
	//fmt.Println(objData)
	for line, v := range objData {

		fmt.Println("#### line: ", line)

		// ValueOf returns a new Value initialized
		// to the concrete value stored in the interface i.
		val := reflect.ValueOf(v)

		// A Kind represents the
		// specific kind of type
		// that a Type represents.
		if val.Kind() == reflect.Map {

			// Loop the map
			for _, key := range val.MapKeys() {

				fmt.Println("..............................................")

				// Index
				v := val.MapIndex(key)
				switch iv := v.Interface().(type) {

				// v.Interface().(type)
				// here test the types
				case int:
					fmt.Println(key, iv)
				case float64:
					fmt.Println(key, iv)
				case string:
					fmt.Println(key, iv)
				case bool:
					fmt.Println(key, iv)
				default:
					fmt.Println(key)
					DecodeMapInterface(iv)
				}
			}
		}
	}
}

func DecodeMapInterface(v interface{}) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Map {
		for _, e := range val.MapKeys() {
			v := val.MapIndex(e)
			switch t := v.Interface().(type) {
			case int:
				fmt.Println(e, t)
			case float64:
				fmt.Println(e, t)
			case string:
				fmt.Println(e, t)
			case bool:
				fmt.Println(e, t)
			default:
				fmt.Println("not found!")
			}
		}
	}
}
```

Output:
```bash
Successfully Opened users.json
#### line:  0
..............................................
name Joelson
..............................................
city Porto Alegre
..............................................
age 39
..............................................
social
facebook https://facebook.com/joelson
twitter https://twitter.com/joelson
instagram https://instagram.com/joelson
..............................................
fone
cell 5531987387246
resid1 55314565678
job 55314785679
#### line:  1
..............................................
name Jefferson
..............................................
city Belo Horizonte
..............................................
age 160
..............................................
social
instagram https://instagram.com/jeffotoni
facebook https://facebook.com/jeffotoni
twitter https://twitter.com/jeffotoni
..............................................
fone
cell 5531987387246
resid1 55314565678
job 55314785679
```

We were able to recursively list our Json, with several layers.
But every Json file will have to be handled, structures change as needed, so this recursion will only work for this file template.

Let's take another example, with a new json with several layers.

Let's take a look at the code below:
```json
{
"payload": 
  [
    {
        "products": {
            "house": {
                "rectangular bed": {
                    "price": 99.33,
                    "weight": 44
                },
                "Corner table": {
                    "price": 100.00,
                    "weight": 400
                }
            },
            "car": {
                "bmw": {
                    "price": 400.000,
                    "weight": 2000
                },
                "jaguar": {
                    "price": 600.000,
                    "weight": 3000
                }
            },
            "robotics": {
                "arduino circuit board": {
                    "price": 3000,
                    "weight": 140
                },
                "proximity sensor...": {
                    "price": 100,
                    "weight": 34
                },
                "temperature sensor": {
                    "price": 344,
                    "weight": 55
                }
            }
        }
    }
  ]
}
```

Complete code:
```go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	// open file json
	data, err := ioutil.ReadFile("./payload.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// make map interface
	payload := make(map[string]interface{})

	// Unmarshal data
	err = json.Unmarshal(data, &payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	// call recursive
	recursivMap(payload)
}

// recursive Map
func recursivMap(payload map[string]interface{}) {
	for k, v := range payload {
		//fmt.Printf("%v %T: %v\n", k, v, v)
		switch v.(type) {
		case []interface{}:
			recursivSlice(v.([]interface{}))
		case map[string]interface{}:
			//fmt.Printf("%v %T: %v\n", k, v, v)
			fmt.Println("----------------------")
			fmt.Printf("%v\n", k)
			recursivMap(v.(map[string]interface{}))

		default:
			fmt.Printf("%v=%v\n", k, v)
		}
	}
}

// recursive Slice
func recursivSlice(pauload []interface{}) {
	for _, v := range pauload {
		switch v.(type) {
		case []interface{}:
			recursivSlice(v.([]interface{}))
		case map[string]interface{}:
			recursivMap(v.(map[string]interface{}))
		}
	}
}
```

Output:
```bash
----------------------
products
----------------------
house
----------------------
rectangular bed
price=99.33
weight=44
----------------------
Corner table
price=100
weight=400
----------------------
car
----------------------
bmw
price=400
weight=2000
----------------------
jaguar
price=600
weight=3000
----------------------
robotics
----------------------
arduino circuit board
price=3000
weight=140
----------------------
proximity sensor...
price=100
weight=34
----------------------
temperature sensor
price=344
weight=55
```

Now, we did something totally recursive, presenting all the keys and values from our Json file.


    - [Json Toml](#json-toml)
    - [Json Yaml](#json-yaml)
    - [Json-Gcfg](#json-gcfg)
    
- [Links Json to Golang](#links-json-to-golang)
- [Exercise three](#Exercise-three)

### Links Json to Golang

Below I am making available some links to convert from Json to Struct in Golang, it gets a json or you write your Json and it mounts the struct for you.
Of course it helps when you know what you're doing, and it's very useful sometimes to find some more complex json.

 - [Mholt Json to Go](https://mholt.github.io/json-to-go/)
 - [Transform json to Go](https://transform.now.sh/json-to-go/)
 - [Json2struct](http://json2struct.mervine.net/)