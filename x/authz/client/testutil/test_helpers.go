package testutil

import (
	"github.com/aliworkshop/terra-sdk/testutil"
	clitestutil "github.com/aliworkshop/terra-sdk/testutil/cli"
	"github.com/aliworkshop/terra-sdk/testutil/network"
	"github.com/aliworkshop/terra-sdk/x/authz/client/cli"
)

func ExecGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
