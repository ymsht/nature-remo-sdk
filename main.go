package main

import (
	"fmt"
	"ymsht/nature-remo-sdk/user"
)

func main() {
	// Get me
	user := user.GetMe()
	fmt.Print(user)
}
