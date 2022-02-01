package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/aliworkshop/terra-sdk/codec"
	cryptoAmino "github.com/aliworkshop/terra-sdk/crypto/codec"
	"github.com/aliworkshop/terra-sdk/testutil/testdata"
	sdk "github.com/aliworkshop/terra-sdk/types"
	"github.com/aliworkshop/terra-sdk/x/auth/migrations/legacytx"
	"github.com/aliworkshop/terra-sdk/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "cosmos-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}
