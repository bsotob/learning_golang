package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Usuario    string
	Contraseña string
}

func main() {
	byteValue, err := ioutil.ReadFile("./playbooks/config.json")
	if err != nil {
		fmt.Println("ha habido algun problema", err)
	} else {
		//fmt.Println(string(file))
		var config Config
		json.Unmarshal(byteValue, &config)
		fmt.Println(string(config.Contraseña))
		fmt.Println(config)
	}
	playbook, err := ioutil.ReadFile("./playbooks/config.json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(playbook))
	}
}
