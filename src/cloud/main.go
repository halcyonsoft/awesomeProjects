package main

import (
	"fmt"
	"github.com/halcyonsoft/awesomeProjects/cloud/api"
)

func main() {
	hub := api.CloudHub{Host: "localhost", Port: "9098"}
	fmt.Println(hub)
}
