package main

import (
	"Feed/producer/service"
	"fmt"
)

func main() {

	s, err := service.NewService()
	if err != nil {
		fmt.Println("err:", err.Error())
	}

	s.Start()

}
