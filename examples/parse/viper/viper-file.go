// Go in action
// @jeffotoni
// 2019-02-01

package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	// viper.SetConfigType("toml")
	// viper.SetConfigName("servert") // name of config file (without extension)

	// viper.SetConfigType("yaml")
	// viper.SetConfigName("servery") // name of config file (without extension)

	viper.SetConfigType("json")
	viper.SetConfigName("serverj") // name of config file (without extension)

	viper.AddConfigPath(".") // optionally look for config in the working directory

	//viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.Set("Verbose", true)
	//viper.Set("LogFile", LogFile)

	fmt.Println("datastore.metric.host: ", viper.GetString("datastore.metric.host"))
	fmt.Println("datastore.warehouse.host: ", viper.GetString("datastore.warehouse.host"))

	fmt.Println("version.json: ", viper.GetString("version"))
	fmt.Println("info.json: ", viper.GetString("info.description"))
	fmt.Println("server1.host.json: ", viper.GetString("server1.host"))
	fmt.Println("server1.user.json: ", viper.GetString("server1.user"))

	fmt.Println("server2.host.json: ", viper.GetString("server2.host"))
	fmt.Println("server2.user.json: ", viper.GetString("server2.user"))
}
