package main

import (
	"fmt"
	"os"
)

// handing injection

//func main() {
//	message := pkg.NewMessage()
//	greeter := pkg.NewGreeter(message)
//	event := pkg.NewEvent(greeter)
//	event.Start()
//}

// wire
func main() {
	event, err := InitializeEvent("this is a single line")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	event.Start()
}
