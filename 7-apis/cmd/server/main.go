package main

import "github.com/tallyto/curso-go/7-apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
