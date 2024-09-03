package utils

import "cosmossdk.io/math"

func MustParseInt(s string) math.Int {
	res, _ := math.NewIntFromString(s)
	return res
}
