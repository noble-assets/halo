package utils

import sdk "github.com/cosmos/cosmos-sdk/types"

func MustParseInt(s string) sdk.Int {
	res, _ := sdk.NewIntFromString(s)
	return res
}
