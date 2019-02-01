// Go in action
// @jeffotoni
// 2019-02-01

package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	interval     int    `mapstructure:"Interval"`
	statsdPrefix string `mapstructure:"statsd_prefix"`
	groups       []group
}
type group struct {
	group        string `mapstructure:"group"`
	targetPrefix string `mapstructure:"target_prefix"`
	targets      []target
}
type target struct {
	target string   `mapstructure:"target"`
	hosts  []string `mapstructure:"hosts"`
}

func main() {

	var C config

	err := viper.Unmarshal(&C)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	C.interval = 209
	C.statsdPrefix = "logi_"

	g1add := group{group: "Aws"}
	g2add := group{group: "Devopsbh"}

	gall := []group{}

	gall = append(gall, g1add)
	gall = append(gall, g2add)

	// recive same type
	//C.groups = gall
	C.groups = []group{{group: "@devopsbr"}, {group: "@go_br"}}
	fmt.Println(C)
	//fmt.Printf("%v", gall)
}
