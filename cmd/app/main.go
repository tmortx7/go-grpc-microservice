package main

import (
	"fmt"

	"github.com/tmortx7/go-grpc-microservice/config"
)

type Server struct {
	
}


func main() {
	config.ReadConfig()

	fmt.Println(config.C.Server.Address)
}
