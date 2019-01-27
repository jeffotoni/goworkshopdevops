# Golang Workshop DevOps

Workshop material for **8 hours** using **Golang**

This is a material in Golang to be presented **"face-to-face"** in a **"hand in hand"** Workshop that will be done in 8 hours.
The content and references used are from the [Golang Official Site](https://golang.org) and the material being developed 
which is a compilation of all Golang language and can be checked here [jeffotoni/Compilation]( https://github.com/jeffotoni/goworkshopdevops#installation)

There are thousands of references today regarding Golang, let's start at the beginning and we could not stop talking about [Golang Tour](https://tour.golang.org).
Well that site here [Play Golang](https://play.golang.org) we can play Golang online.

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
   - [Map Types](#map-types)
   - [Channel Types](#channel-types)
   - [Blank identifier](#blank-identifier)
   - [Interface Types](#interface-types)
	 - [Here's an interface as a method](#heres-an-interface-as-a-method)
	 - [Interface as type](#interface-as-type)
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
    - [Asynchronous Functions](#asynchromous-functions)
- [Defer](#defer)

- [Exercise one](#exercise-one)

## Lab 03 Json with Golang

- [Json](#Json)
  - [introduction](#)
    - [Json marshal](#jsonmarshal)
    - [Json Unmarshal](#jsonunmarshal)
    - [json Encode](#jsonencode)
    - [Json Decode](#jsondecode)
- [Parse Json](#Json)
  - [introduction](#)
    - [Toml](#jsontoml)
    - [Yaml](#jsonyaml)
    - [Gcfg](#jsongcfg)

- [Exercise three](#Exercise-three)

## Lab 04 Goroutine the power

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

## Lab 05 Using Golang to create Command Line programs

- [Golang Cli](#golang-cli)
  - [introduction](#)
    - [Go flag types](#go-flag-types)
    - [Go flag Parse](go-flag-parse)
    - [Go flag PrintDefaults](go-flag-printDefaults)
    - [Os Args](#os-args)
    - [func init](#func-init)
    - [Go exec Command](#go-exec-command)
    - [Stdout](#stdout)
    - [Stdin](#stdin)
    - [New Scanner and Scan](#new-scanner-and-scan)
    - [Parse Url](#parse-url)
    - [Parse Yaml](#parse-yaml)

- [Exercise five](#Exercise-five)

## Lab 06 building apis with net/http

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
- [net/http Server Pages](#)
  - [introduction](#)
    - [http.FileServer](#)
    - [http.NotFound](#)
    - [Disable http.FileServer](#)
    - [http.Dir](#)
    - [http.StripPrefix](#)     
- [net/http Client](#)
  - [introduction](#)
    - [http.Transport](#)
    - [http.Client](#)
    - [http.Get](#)
    - [http,Post](#)
    - [http.NewRequest](#)
    - [Context.WithCancel](#)

- [Exercise six](#Exercise-six)


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

```bash
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName .
Tag           = string_lit .

// An empty struct.
struct {}

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

```bash
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

```bash
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

A goroutine is a lightweight thread of execution.


### Defer


### Exercise one
---

