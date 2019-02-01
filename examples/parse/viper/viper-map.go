// Go in action
// @jeffotoni
// 2019-02-01

package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	v1, err := readConfig("enviroment", map[string]interface{}{
		"version": 1.0,
		"info": map[string]string{
			"description": "this is an example of yaml file to server an ssh server or",
		},
		"server1": map[string]string{
			"host": "localhost",
			"user": "titpetric",
		},
	})
	if err != nil {
		panic(fmt.Errorf("Error when reading config: %v\n", err))
	}

	version := v1.GetFloat64("version")
	info := v1.GetString("info.description")
	server1 := v1.GetStringMapString("server1")

	fmt.Printf("version: %0.1f\n", version)
	fmt.Printf("info.description: %s\n", info)
	fmt.Printf("server1:%#v\n", server1)
	fmt.Printf("server1.host: %s\n", server1["host"])
}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}
