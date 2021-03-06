package types

import (
	"github.com/aliworkshop/terra-sdk/codec"
	cryptocodec "github.com/aliworkshop/terra-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/gov module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/gov and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	cryptocodec.RegisterCrypto(amino)

	// v1beta1 and v1beta2 will each add their own Amino registrations inside
	// their init() functions.
}
