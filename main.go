package main

import (
	"fmt"

	cdi "github.com/container-orchestrated-devices/container-device-interface"
)

func main() {
	fmt.Println("Welcome to CDI test")
	found, err := cdi.HasDevice("myDevice")
	if err != nil {
		fmt.Printf("Not found")
	}
	if found {
		fmt.Printf("It was found")
	}
}
