package main

import (
	"fmt"
	"ymsht/nature-remo-sdk/user"
)

func main() {
	user := user.GetMe()
	fmt.Print(user)
}
