package main

import (
	"fmt"
)

func main() {
	token := "Bearer "
	sdk := &NatureRemoSdk{Token: token}

	// Get me
	user := sdk.GetMe()
	fmt.Print(user)
}
