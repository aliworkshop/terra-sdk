package client

import (
	govclient "github.com/aliworkshop/terra-sdk/x/gov/client"
	"github.com/aliworkshop/terra-sdk/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
