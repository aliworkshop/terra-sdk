package types

import (
	"github.com/aliworkshop/terra-sdk/codec"
	cryptocodec "github.com/aliworkshop/terra-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
