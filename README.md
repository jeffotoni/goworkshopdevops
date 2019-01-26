# Golang Workshop DevOps

Workshop material for 8 hours using golang

This is a material in Golang to be presented face-to-face in a "hand in hand" Workshop that will be done in 8 hours.
The content and references used are from the [Golang Official Site] website (https://golang.org) and the material being developed 
which is a compilation of all Golang language and can be checked here [jeffotoni/Compilation]( https://github.com/jeffotoni/gocompilation#installation)

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
- [Installation](#installation)
  - [Introduction](#introduction-installation)
    - [Linux](#linux)
    - [Test your installation](#test-your-installation)
    - [$GOPATH](#gopath)
    - [Workspace](#workspace)
    - [Outside GOPATH](#outside-gopath)
    - [Func Main](#func-main)
- [Installation Docker](#installation-docker)
- [Go commands](#go-commands)
  - [Go commands introduction](#go-commands-introduction)
  - [go run](#go-run) 
  - [go build](#go-build)
  - [go get](#go-get)
  - [go test](#go-test)
  - [go install](#go-install)
  - [Go mod](#go-mod)
  - [go mod init](#gomodinit)
  - [go mod vendor](#go-mod-vendor)
  - [GO111MODULE](#go111module)

## Lab 02 Types with Golang

- [Types](#Types)
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
