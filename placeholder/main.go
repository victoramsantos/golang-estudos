package main

import (
	"fmt"
	"placeholder/controller"
)

func main() {

	var response = controller.GetDynamic()
	fmt.Println(response)
}
