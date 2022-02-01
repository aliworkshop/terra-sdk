package v040_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/aliworkshop/terra-sdk/client"
	"github.com/aliworkshop/terra-sdk/simapp"
	sdk "github.com/aliworkshop/terra-sdk/types"
	v038auth "github.com/aliworkshop/terra-sdk/x/auth/migrations/v038"
	v039auth "github.com/aliworkshop/terra-sdk/x/auth/migrations/v039"
	v036supply "github.com/aliworkshop/terra-sdk/x/bank/migrations/v036"
	v038bank "github.com/aliworkshop/terra-sdk/x/bank/migrations/v038"
	v040bank "github.com/aliworkshop/terra-sdk/x/bank/migrations/v040"
)

func TestMigrate(t *testing.T) {
	encodingConfig := simapp.MakeTestEncodingConfig()
	clientCtx := client.Context{}.
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithCodec(encodingConfig.Codec)

	coins := sdk.NewCoins(sdk.NewInt64Coin("stake", 50))
	addr1, _ := sdk.AccAddressFromBech32("cosmos1xxkueklal9vejv9unqu80w9vptyepfa95pd53u")
	acc1 := v038auth.NewBaseAccount(addr1, coins, nil, 1, 0)

	addr2, _ := sdk.AccAddressFromBech32("cosmos15v50ymp6n5dn73erkqtmq0u8adpl8d3ujv2e74")
	vaac := v038auth.NewContinuousVestingAccountRaw(
		v038auth.NewBaseVestingAccount(
			v038auth.NewBaseAccount(addr2, coins, nil, 1, 0), coins, nil, nil, 3160620846,
		),
		1580309972,
	)

	supply := sdk.NewCoins(sdk.NewInt64Coin("stake", 1000))

	bankGenState := v038bank.GenesisState{
		SendEnabled: true,
	}
	authGenState := v039auth.GenesisState{
		Accounts: v038auth.GenesisAccounts{acc1, vaac},
	}
	supplyGenState := v036supply.GenesisState{
		Supply: supply,
	}

	migrated := v040bank.Migrate(bankGenState, authGenState, supplyGenState)
	expected := `{"params":{"send_enabled":[],"default_send_enabled":true},"balances":[{"address":"cosmos1xxkueklal9vejv9unqu80w9vptyepfa95pd53u","coins":[{"denom":"stake","amount":"50"}]},{"address":"cosmos15v50ymp6n5dn73erkqtmq0u8adpl8d3ujv2e74","coins":[{"denom":"stake","amount":"50"}]}],"supply":[{"denom":"stake","amount":"1000"}],"denom_metadata":[]}`

	bz, err := clientCtx.Codec.MarshalJSON(migrated)
	require.NoError(t, err)
	require.Equal(t, expected, string(bz))
}
