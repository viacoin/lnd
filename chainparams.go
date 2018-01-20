package main

import (
	viacoinCfg "github.com/viacoin/viad/chaincfg"
	"github.com/roasbeef/btcd/chaincfg"
	bitcoinCfg "github.com/roasbeef/btcd/chaincfg"
	"github.com/roasbeef/btcd/chaincfg/chainhash"
	"github.com/roasbeef/btcd/wire"
)

// activeNetParams is a pointer to the parameters specific to the currently
// active bitcoin network.
var activeNetParams = bitcoinTestNetParams

// bitcoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type bitcoinNetParams struct {
	*bitcoinCfg.Params
	rpcPort string
}

// viacoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type viacoinNetParams struct {
	*viacoinCfg.Params
	rpcPort string
}

// bitcoinTestNetParams contains parameters specific to the 3rd version of the
// test network.
var bitcoinTestNetParams = bitcoinNetParams{
	Params:  &bitcoinCfg.TestNet3Params,
	rpcPort: "18334",
}

// bitcoinSimNetParams contains parameters specific to the simulation test
// network.
var bitcoinSimNetParams = bitcoinNetParams{
	Params:  &bitcoinCfg.SimNetParams,
	rpcPort: "18556",
}

// viaTestNetParams contains parameters specific to the 4th version of the
// test network.
var viaTestNetParams = viacoinNetParams{
	Params:  &viacoinCfg.TestNet3Params,
	rpcPort: "19224",
}

// regTestNetParams contains parameters specific to a local regtest network.
var regTestNetParams = bitcoinNetParams{
	Params:  &bitcoinCfg.RegressionNetParams,
	rpcPort: "18334",
}

// applyViacoinParams applies the relevant chain configuration parameters that
// differ for viacoin to the chain parameters typed for btcsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyViacoinParams(params *bitcoinNetParams) {
	params.Name = viaTestNetParams.Name
	params.Net = wire.BitcoinNet(viaTestNetParams.Net)
	params.DefaultPort = viaTestNetParams.DefaultPort
	params.CoinbaseMaturity = viaTestNetParams.CoinbaseMaturity

	copy(params.GenesisHash[:], viaTestNetParams.GenesisHash[:])

	// Address encoding magics
	params.PubKeyHashAddrID = viaTestNetParams.PubKeyHashAddrID
	params.ScriptHashAddrID = viaTestNetParams.ScriptHashAddrID
	params.PrivateKeyID = viaTestNetParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = viaTestNetParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = viaTestNetParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = viaTestNetParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], viaTestNetParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], viaTestNetParams.HDPublicKeyID[:])

	params.HDCoinType = viaTestNetParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(viaTestNetParams.Checkpoints))
	for i := 0; i < len(viaTestNetParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], viaTestNetParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: viaTestNetParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = viaTestNetParams.rpcPort
}
