package keeper_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/aliworkshop/terra-sdk/simapp"
	sdk "github.com/aliworkshop/terra-sdk/types"
	authtypes "github.com/aliworkshop/terra-sdk/x/auth/types"
)

// returns context and app with params set on account keeper
func createTestApp(t *testing.T, isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(t, isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())

	return app, ctx
}
