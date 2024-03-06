package main

import "github.com/andremelinski/pos-goexpert/8-API/configs"

func main(){
	config, _ := configs.LoadConfig(".")
	println(config)
}