package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

type studentConfig struct {
	Name string
	City string
}

func main() {
	var conf studentConfig
	_, err := toml.DecodeFile("a.toml", &conf)
	if err != nil {
		log.Fatal(err)
	}
	print(conf.Name)

}
