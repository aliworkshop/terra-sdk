package nft

import (
	types "github.com/aliworkshop/terra-sdk/codec/types"
	sdk "github.com/aliworkshop/terra-sdk/types"
	"github.com/aliworkshop/terra-sdk/types/msgservice"
)

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
