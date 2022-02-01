package params

import (
	"github.com/aliworkshop/terra-sdk/client"
	"github.com/aliworkshop/terra-sdk/codec"
	"github.com/aliworkshop/terra-sdk/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}
