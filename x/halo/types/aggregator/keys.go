package aggregator

import "encoding/binary"

const SubmoduleName = "halo-aggregator"

var (
	OwnerKey       = []byte("aggregator/owner")
	LastRoundIDKey = []byte("aggregator/last_round_id")
	NextPriceKey   = []byte("aggregator/next_price")
	RoundPrefix    = []byte("aggregator/round/")
)

func RoundKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(RoundPrefix, bz...)
}
