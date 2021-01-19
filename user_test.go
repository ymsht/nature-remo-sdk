package nature_remo_sdk_test

import (
	"testing"
	sdk "ymsht/nature-remo-sdk"
)

func TestGetMe(t *testing.T) {
	sdk := sdk.NatureRemoSdk{Token: "Bearer "}
	_, _ = sdk.GetMe()
}
