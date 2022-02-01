package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/aliworkshop/terra-sdk/client"
	"github.com/aliworkshop/terra-sdk/codec"
	sdk "github.com/aliworkshop/terra-sdk/types"
	sdkerrors "github.com/aliworkshop/terra-sdk/types/errors"
	"github.com/aliworkshop/terra-sdk/x/gov/types"
	"github.com/aliworkshop/terra-sdk/x/gov/types/v1beta2"
)

// NewQuerier creates a new gov Querier instance
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case v1beta2.QueryParams:
			return queryParams(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case v1beta2.QueryProposals:
			return queryProposals(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case v1beta2.QueryProposal:
			return queryProposal(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case v1beta2.QueryDeposits:
			return queryDeposits(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case v1beta2.QueryDeposit:
			return queryDeposit(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case v1beta2.QueryVotes:
			return queryVotes(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case v1beta2.QueryVote:
			return queryVote(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case v1beta2.QueryTally:
			return queryTally(ctx, path[1:], req, keeper, legacyQuerierCdc)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func queryParams(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	switch path[0] {
	case v1beta2.ParamDeposit:
		bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, keeper.GetDepositParams(ctx))
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		return bz, nil

	case v1beta2.ParamVoting:
		bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, keeper.GetVotingParams(ctx))
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		return bz, nil

	case v1beta2.ParamTallying:
		bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, keeper.GetTallyParams(ctx))
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		return bz, nil

	default:
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "%s is not a valid query request path", req.Path)
	}
}

// nolint: unparam
func queryProposal(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params v1beta2.QueryProposalParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	proposal, ok := keeper.GetProposal(ctx, params.ProposalID)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrUnknownProposal, "%d", params.ProposalID)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, proposal)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

// nolint: unparam
func queryDeposit(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params v1beta2.QueryDepositParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	deposit, _ := keeper.GetDeposit(ctx, params.ProposalID, params.Depositor)
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, deposit)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

// nolint: unparam
func queryVote(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params v1beta2.QueryVoteParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	vote, _ := keeper.GetVote(ctx, params.ProposalID, params.Voter)
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, vote)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

// nolint: unparam
func queryDeposits(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params v1beta2.QueryProposalParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	deposits := keeper.GetDeposits(ctx, params.ProposalID)
	if deposits == nil {
		deposits = v1beta2.Deposits{}
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, deposits)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

// nolint: unparam
func queryTally(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params v1beta2.QueryProposalParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	proposalID := params.ProposalID

	proposal, ok := keeper.GetProposal(ctx, proposalID)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrUnknownProposal, "%d", proposalID)
	}

	var tallyResult v1beta2.TallyResult

	switch {
	case proposal.Status == v1beta2.StatusDepositPeriod:
		tallyResult = v1beta2.EmptyTallyResult()

	case proposal.Status == v1beta2.StatusPassed || proposal.Status == v1beta2.StatusRejected:
		tallyResult = *proposal.FinalTallyResult

	default:
		// proposal is in voting period
		_, _, tallyResult = keeper.Tally(ctx, proposal)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, tallyResult)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

// nolint: unparam
func queryVotes(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params v1beta2.QueryProposalVotesParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	votes := keeper.GetVotes(ctx, params.ProposalID)
	if votes == nil {
		votes = v1beta2.Votes{}
	} else {
		start, end := client.Paginate(len(votes), params.Page, params.Limit, 100)
		if start < 0 || end < 0 {
			votes = v1beta2.Votes{}
		} else {
			votes = votes[start:end]
		}
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, votes)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryProposals(ctx sdk.Context, _ []string, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params v1beta2.QueryProposalsParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	proposals := keeper.GetProposalsFiltered(ctx, params)
	if proposals == nil {
		proposals = v1beta2.Proposals{}
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, proposals)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
