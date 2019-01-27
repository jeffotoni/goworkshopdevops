# Golang Workshop DevOps

Workshop material for **8 hours** using **Golang**

This is a material in Golang to be presented **"face-to-face"** in a **"hand in hand"** Workshop that will be done in 8 hours.
The content and references used are from the [Golang Official Site](https://golang.org) and the material being developed 
which is a compilation of all Golang language and can be checked here [jeffotoni/Compilation]( https://github.com/jeffotoni/goworkshopdevops#installation)

There are thousands of references today regarding Golang, let's start at the beginning and we could not stop talking about [Golang Tour](https://tour.golang.org).
Well that site here [Play Golang](https://play.golang.org) we can play Golang online.

Soon below some channels that I participate and can find me online.

#### Telegram:
   - [gobr](https://t.me/go_br)
#### Slack: 
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
- [Go commands](#go-commands)
  - [Go commands introduction](#go-commands-introduction)
  - [go run](#go-run) 
  - [go build](#go-build)
  - [go install](#go-install)
  - [go get](#go-get)
  - [Go mod](#go-mod)
  - [go mod init](#gomodinit)
  - [go mod vendor](#go-mod-vendor)
  - [GO111MODULE](#go111module)
  - [go test](#go-test)
- [Func Main](#func-main)

## Lab 02 Println and Types with Golang and For
- [Println](#println)
- [Types](#types)
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
- [For Break Range If else](#for-break-range-if-else)

- [Exercise one](#Exercise-one)

- [Erros](#erros)
  - [introduction](#)
    - [How Error Control Works](#)
    - [Custom Errors](#customerrors)
    - [Interface Errors](#)
    - [Errors.New](#)
    - [fmt.Errorf](#)

- [Functions](#functions)
  - [introduction](#)
    - [return multiple values](#returnmulti) 
    - [Variadic Functions](#variadicfunc) 
    - [functions as a parameter](#funcparameter) 
    - [Closures](#closures)
    - [Recursion](#recursion)
    - [Asynchronous Functions](#asynchromous)
- [Defer](#defer)

- [Exercise two](#Exercise-two)

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
