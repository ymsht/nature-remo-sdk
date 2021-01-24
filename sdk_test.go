package nature_remo_sdk_test

import (
	"fmt"
	"testing"
	sdk "ymsht/nature-remo-sdk"
)

// TODO テストかく
func TestGetMe(t *testing.T) {
	sdk := sdk.NatureRemoSdk{Token: "Bearer "}
	user, err := sdk.GetMe()
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("%s\n", user.Nickname)
}

func TestDevice(t *testing.T) {
	sdk := sdk.NatureRemoSdk{Token: "Bearer "}
	devices, err := sdk.GetDevice()
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("%f\n", devices[0].Newest_events.Te.Val)
}
