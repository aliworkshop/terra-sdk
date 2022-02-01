package client

import (
	govclient "github.com/aliworkshop/terra-sdk/x/gov/client"
	"github.com/aliworkshop/terra-sdk/x/upgrade/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal)
