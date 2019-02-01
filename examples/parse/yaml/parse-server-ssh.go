// Go in action
// @jeffotoni
// 2019-02-01

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type ServersSsh struct {
	Version string `yaml:"version"`

	Info struct {
		Description string `yaml:"description"`
	}

	Server1 struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		User    string `yaml:"user"`
		FilePem string `yaml:"filepem"`
		KeyAws  string `yaml:"keyaws"`
	}

	Server2 struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		User    string `yaml:"user"`
		FilePem string `yaml:"filepem"`
		KeyAws  string `yaml:"keyaws"`
	}
}

// Our config case has no structure created
// the system will dynamically create
var YamlMemory = `
`

// func main
func main() {

	var yamlByte []byte
	var Yaml ServersSsh
	var err error

	// local
	file := "server.yaml"

	// test exist
	if !FileExist(file) {
		fmt.Println(file + " not exist!")
		return
	}

	if yamlByte, err = ioutil.ReadFile(file); err != nil {
		log.Println("Error: ", err)
	}

	// Unmarshal receives the file in a byte format and assigns the values passed to the fields in the structure.
	// If there is an error it displays an error message on the screen informing the error
	if err := yaml.Unmarshal(yamlByte, &Yaml); err != nil {
		log.Println("Error", err)
	}

	fmt.Println("Version: ", Yaml.Version)
	fmt.Println("Description: ", Yaml.Info.Description)
	fmt.Println("Server1.Host: ", Yaml.Server1.Host)
	fmt.Println("Server2.Host: ", Yaml.Server2.Host)
}

func FileExist(name string) bool {
	//if _, err := os.Stat(name); os.IsNotExist(err) {
	if stat, err := os.Stat(name); err == nil && !stat.IsDir() {
		return true
	}
	return false
}
