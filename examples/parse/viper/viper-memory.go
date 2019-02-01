// Go in action
// @jeffotoni
// 2019-02-01

package main

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
version: 1.0
info: 
 description: this is an example of yaml file to ...
server1:
  host: 127.0.0.1
  port: 22
  user: ubuntu
  type: env
  file: key_1.pem
  env: KEY_AWS
  Clouds: [Aws, Google, Azure]
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	fmt.Println(viper.Get("version"))
	fmt.Println(viper.Get("info.description"))
	fmt.Println(viper.Get("server1.host"))
	fmt.Println(viper.Get("server1.user"))
	mapsc := viper.GetStringSlice("server1.clouds")

	fmt.Println(mapsc[0])
}
