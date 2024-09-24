package data

import (
	"cosmossdk.io/math"
	"github.com/noble-assets/halo/v2/types/aggregator"
	"github.com/noble-assets/halo/v2/utils"
)

var EthereumRounds = []struct {
	Msg    aggregator.MsgReportBalance
	Answer math.Int
}{
	{
		// https://etherscan.io/tx/0xbcba4db502a72e51a05b378d2e5867be4c60936585cedaf5aad90002f0599428
		Msg: aggregator.MsgReportBalance{
			Principal:   utils.MustParseInt("4790909990"),
			Interest:    utils.MustParseInt("701123"),
			TotalSupply: utils.MustParseInt("46169872257060"),
			NextPrice:   utils.MustParseInt("103794354"),
		},
		Answer: utils.MustParseInt("103780685"),
	},
	{
		// https://etherscan.io/tx/0xd986b4f0367ae38ecca94452a6f3cb471379712f24f42fe85b53c0b420041b3b
		Msg: aggregator.MsgReportBalance{
			Principal:   utils.MustParseInt("4785004746"),
			Interest:    utils.MustParseInt("698930"),
			TotalSupply: utils.MustParseInt("46106890838455"),
			NextPrice:   utils.MustParseInt("103807973"),
		},
		Answer: utils.MustParseInt("103794328"),
	},
	{
		// https://etherscan.io/tx/0x1ce8bb109574bd8e39f979f87c6d516eb09a4ced111ef4b6a2c2d6b5a9dcf746
		Msg: aggregator.MsgReportBalance{
			Principal:   utils.MustParseInt("4785703676"),
			Interest:    utils.MustParseInt("676614"),
			TotalSupply: utils.MustParseInt("46107564218218"),
			NextPrice:   utils.MustParseInt("103821180"),
		},
		Answer: utils.MustParseInt("103807535"),
	},
	{
		// https://etherscan.io/tx/0xc2fe6053f8758f0968ec1595ab5018d60a3d4ae3dc0e47c5c1699b6f4156191d
		Msg: aggregator.MsgReportBalance{
			Principal:   utils.MustParseInt("4796380420"),
			Interest:    utils.MustParseInt("687997"),
			TotalSupply: utils.MustParseInt("46204549387934"),
			NextPrice:   utils.MustParseInt("103861146"),
		},
		Answer: utils.MustParseInt("103820937"),
	},
	{
		// https://etherscan.io/tx/0x25e2f62b50b35dd985ad4827013b3e365f68a565e0e6aaf78dde58ca934ce70e
		Msg: aggregator.MsgReportBalance{
			Principal:   utils.MustParseInt("4597451007"),
			Interest:    utils.MustParseInt("1981867"),
			TotalSupply: utils.MustParseInt("44282503451837"),
			NextPrice:   utils.MustParseInt("103874647"),
		},
		Answer: utils.MustParseInt("103861216"),
	},
	{
		// https://etherscan.io/tx/0x8d6146ed696752a8400f40750baf6515ebc63da7e102976f779da495356ec5cc
		Msg: aggregator.MsgReportBalance{
			Principal:   utils.MustParseInt("4599432874"),
			Interest:    utils.MustParseInt("659658"),
			TotalSupply: utils.MustParseInt("44284411639586"),
			NextPrice:   utils.MustParseInt("103888030"),
		},
		Answer: utils.MustParseInt("103874623"),
	},
	{
		// https://etherscan.io/tx/0xe8154863cf89175c9ec361999ec7ddeebf7c29297ee62325f777067409071303
		Msg: aggregator.MsgReportBalance{
			Principal:   utils.MustParseInt("4600092532"),
			Interest:    utils.MustParseInt("658503"),
			TotalSupply: utils.MustParseInt("44285046691710"),
			NextPrice:   utils.MustParseInt("103901390"),
		},
		Answer: utils.MustParseInt("103888005"),
	},
	{
		// https://etherscan.io/tx/0x1a628856cb74de37357a35c29ec22509b72b1fc826ac7bb1020c73a99f9f80fc
		Msg: aggregator.MsgReportBalance{
			Principal:   utils.MustParseInt("4600751035"),
			Interest:    utils.MustParseInt("657348"),
			TotalSupply: utils.MustParseInt("44285680550257"),
			NextPrice:   utils.MustParseInt("103914726"),
		},
		Answer: utils.MustParseInt("103901364"),
	},
	{
		// https://etherscan.io/tx/0x5cc762e8201084220a63bdf97dd8d07b87d1a95cbd528d388c41e054aa86e52a
		Msg: aggregator.MsgReportBalance{
			Principal:   utils.MustParseInt("4601408383"),
			Interest:    utils.MustParseInt("657442"),
			TotalSupply: utils.MustParseInt("44286313215676"),
			NextPrice:   utils.MustParseInt("103954812"),
		},
		Answer: utils.MustParseInt("103914725"),
	},
	{
		// https://etherscan.io/tx/0x5122f8dd032fccd723fc3f1bda4cf0b3074731ef44208ed13dee11b0b138ee92
		Msg: aggregator.MsgReportBalance{
			Principal:   utils.MustParseInt("4602370972"),
			Interest:    utils.MustParseInt("1976357"),
			TotalSupply: utils.MustParseInt("44289882403775"),
			NextPrice:   utils.MustParseInt("103968279"),
		},
		Answer: utils.MustParseInt("103954886"),
	},
}
