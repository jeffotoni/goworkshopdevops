// Go in action
// @jeffotoni
// 2019-02-01

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type ConfToml struct {
	Version string
	Info    Info
	Server1 Server1
	Server2 Server2
	Clients Clients
}

// Info from config file
type Info struct {
	Description string
}

// Server1
type Server1 struct {
	Host    string
	Port    int64
	User    string
	FilePem string
	KeyAws  string
}

// Server2
type Server2 struct {
	Host    string
	Port    int64
	User    string
	FilePem string
	KeyAws  string
}

// info from clients file
type Clients struct {
	Ping  string
	Data  [][]interface{} // very beautiful
	Hosts []string
}

func main() {

	var Toml ConfToml

	file := "server.toml"

	if !FileExist(file) {
		fmt.Println(file + " not exist!")
		return
	}

	if _, err := toml.DecodeFile(file, &Toml); err != nil {
		log.Fatal(err)
	}

	// config data
	fmt.Println("version:", Toml.Version)
	fmt.Println("Info.Description: ", Toml.Info.Description)
	fmt.Println("Server1.Host: ", Toml.Server1.Host)
	fmt.Println("Server2.Host:", Toml.Server2.Host)

	fmt.Println(Toml.Clients.Data)
	fmt.Println(Toml.Clients.Data[0])
	fmt.Println(Toml.Clients.Data[0][0])
	fmt.Println(Toml.Clients.Data[0][1])

	fmt.Println(Toml.Clients.Hosts)
	fmt.Println(Toml.Clients.Hosts[0])
	fmt.Println(Toml.Clients.Hosts[1])
}

func FileExist(name string) bool {
	//if _, err := os.Stat(name); os.IsNotExist(err) {
	if stat, err := os.Stat(name); err == nil && !stat.IsDir() {
		return true
	}
	return false
}
