package group

import (
	"github.com/aliworkshop/terra-sdk/codec/types"
	sdk "github.com/aliworkshop/terra-sdk/types"
	"github.com/aliworkshop/terra-sdk/types/tx"
)

func (p *Proposal) GetMsgs() []sdk.Msg {
	msgs, err := tx.GetMsgs(p.Msgs, "proposal")
	if err != nil {
		panic(err)
	}
	return msgs
}

func (p *Proposal) SetMsgs(msgs []sdk.Msg) error {
	anys, err := tx.SetMsgs(msgs)
	if err != nil {
		return err
	}
	p.Msgs = anys
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	return tx.UnpackInterfaces(unpacker, p.Msgs)
}
