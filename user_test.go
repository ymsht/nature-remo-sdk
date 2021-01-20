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
	fmt.Printf(user.Nickname)
}
