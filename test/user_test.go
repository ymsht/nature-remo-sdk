package test

import (
	"fmt"
	nature_remo_sdk "ymsht/nature-remo-sdk"
)

func TestGetMe() {
	sdk := nature_remo_sdk.NatureRemoSdk{}
	sdk.Token = "Bearer "
	user := sdk.GetMe()
	fmt.Print(user)
}
