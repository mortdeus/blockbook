package rvn

import (
	"blockbook/bchain/coins/btc"

	"github.com/btcsuite/btcd/wire"
	"github.com/jakm/btcutil/chaincfg"
)

/*

   nDefaultPort = 8767;
   nPruneAfterHeight = 100000;

*/
const (
	MainnetMagic wire.BitcoinNet = 0x5241564e
	TestnetMagic wire.BitcoinNet = 0x52564e54
	RegtestMagic wire.BitcoinNet = 0x43524f57
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{60}
	MainNetParams.ScriptHashAddrID = []byte{122}
	MainNetParams.Bech32HRPSegwit = "rvn"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{111}
	TestNetParams.ScriptHashAddrID = []byte{196}
	TestNetParams.Bech32HRPSegwit = "trvn"
}

// RavencoinParser handle
type RavencoinParser struct {
	*btc.BitcoinParser
}

// NewRavencoinParser returns new RavencoinParser instance
func NewRavencoinParser(params *chaincfg.Params, c *btc.Configuration) *RavencoinParser {
	return &Ravencoin
	Parser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Ravencoin network,
// and the test Ravencoin network
func GetChainParams(chain string) *chaincfg.Params {
	// register bitcoin parameters in addition to Ravencoin parameters
	// Ravencoin has dual standard of addresses and we want to be able to
	// parse both standards
	if !chaincfg.IsRegistered(&chaincfg.MainNetParams) {
		chaincfg.RegisterBitcoinParams()
	}
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
