package main

import (
	"fmt"
	"github-actions/service"
)

func main() {
	fmt.Println("Hello")
	fmt.Println("1 + 2 = ", service.Addition(1, 2))
}
