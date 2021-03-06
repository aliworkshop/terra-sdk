package client

import (
	"github.com/aliworkshop/terra-sdk/x/distribution/client/cli"
	govclient "github.com/aliworkshop/terra-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal)
)
