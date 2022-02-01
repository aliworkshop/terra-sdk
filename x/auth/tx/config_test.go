package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/aliworkshop/terra-sdk/codec"
	codectypes "github.com/aliworkshop/terra-sdk/codec/types"
	"github.com/aliworkshop/terra-sdk/std"
	"github.com/aliworkshop/terra-sdk/testutil/testdata"
	sdk "github.com/aliworkshop/terra-sdk/types"
	"github.com/aliworkshop/terra-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
