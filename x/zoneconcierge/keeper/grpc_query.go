package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/babylonchain/babylon/x/zoneconcierge/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ChainList(c context.Context, req *types.QueryChainListRequest) (*types.QueryChainListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	chainIDs := []string{}
	store := k.chainInfoStore(ctx)
	pageRes, err := query.Paginate(store, req.Pagination, func(key, value []byte) error {
		chainID := string(key)
		chainIDs = append(chainIDs, chainID)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &types.QueryChainListResponse{
		ChainIds:   chainIDs,
		Pagination: pageRes,
	}
	return resp, nil
}

// ChainsInfo returns the latest info for a list of chains with given IDs
func (k Keeper) ChainsInfo(c context.Context, req *types.QueryChainsInfoRequest) (*types.QueryChainsInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if len(req.ChainIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "chain IDs cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	var chainsInfo []*types.ChainInfo
	for _, chainID := range req.ChainIds {
		chainInfo, err := k.GetChainInfo(ctx, chainID)
		if err != nil {
			return nil, err
		}

		chainsInfo = append(chainsInfo, chainInfo)
	}

	resp := &types.QueryChainsInfoResponse{ChainsInfo: chainsInfo}
	return resp, nil
}

// Header returns the header and fork headers at a given height
func (k Keeper) Header(c context.Context, req *types.QueryHeaderRequest) (*types.QueryHeaderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if len(req.ChainId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "chain ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)

	header, err := k.GetHeader(ctx, req.ChainId, req.Height)
	if err != nil {
		return nil, err
	}
	forks := k.GetForks(ctx, req.ChainId, req.Height)
	resp := &types.QueryHeaderResponse{
		Header:      header,
		ForkHeaders: forks,
	}

	return resp, nil
}

// EpochChainInfo returns the info of a chain with given ID in a given epoch
func (k Keeper) EpochChainInfo(c context.Context, req *types.QueryEpochChainInfoRequest) (*types.QueryEpochChainInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if len(req.ChainId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "chain ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)

	// find the chain info of the given epoch
	chainInfo, err := k.GetEpochChainInfo(ctx, req.ChainId, req.EpochNum)
	if err != nil {
		return nil, err
	}
	resp := &types.QueryEpochChainInfoResponse{ChainInfo: chainInfo}
	return resp, nil
}

// ListHeaders returns all headers of a chain with given ID, with pagination support
func (k Keeper) ListHeaders(c context.Context, req *types.QueryListHeadersRequest) (*types.QueryListHeadersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if len(req.ChainId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "chain ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)

	headers := []*types.IndexedHeader{}
	store := k.canonicalChainStore(ctx, req.ChainId)
	pageRes, err := query.Paginate(store, req.Pagination, func(key, value []byte) error {
		var header types.IndexedHeader
		k.cdc.MustUnmarshal(value, &header)
		headers = append(headers, &header)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &types.QueryListHeadersResponse{
		Headers:    headers,
		Pagination: pageRes,
	}
	return resp, nil
}

// ListEpochHeaders returns all headers of a chain with given ID
// TODO: support pagination in this RPC
func (k Keeper) ListEpochHeaders(c context.Context, req *types.QueryListEpochHeadersRequest) (*types.QueryListEpochHeadersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if len(req.ChainId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "chain ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)

	headers, err := k.GetEpochHeaders(ctx, req.ChainId, req.EpochNum)
	if err != nil {
		return nil, err
	}

	resp := &types.QueryListEpochHeadersResponse{
		Headers: headers,
	}
	return resp, nil
}

func (k Keeper) FinalizedChainInfo(c context.Context, req *types.QueryFinalizedChainInfoRequest) (*types.QueryFinalizedChainInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if len(req.ChainId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "chain ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	resp := &types.QueryFinalizedChainInfoResponse{}

	// find the last finalised chain info and the earliest epoch that snapshots this chain info
	finalizedEpoch, chainInfo, err := k.GetLastFinalizedChainInfo(ctx, req.ChainId)
	if err != nil {
		return nil, err
	}
	resp.FinalizedChainInfo = chainInfo

	// find the epoch metadata of the finalised epoch
	resp.EpochInfo, err = k.epochingKeeper.GetHistoricalEpoch(ctx, finalizedEpoch)
	if err != nil {
		return nil, err
	}

	rawCheckpoint, err := k.checkpointingKeeper.GetRawCheckpoint(ctx, finalizedEpoch)

	if err != nil {
		return nil, err
	}

	resp.RawCheckpoint = rawCheckpoint.Ckpt

	// find the raw checkpoint and the best submission key for the finalised epoch
	_, resp.BtcSubmissionKey, err = k.btccKeeper.GetBestSubmission(ctx, finalizedEpoch)
	if err != nil {
		return nil, err
	}

	// if the query does not want the proofs, return here
	if !req.Prove {
		return resp, nil
	}

	// generate all proofs
	resp.Proof, err = k.proveFinalizedChainInfo(ctx, chainInfo, resp.EpochInfo, resp.BtcSubmissionKey)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (k Keeper) FinalizedChainInfoUntilHeight(c context.Context, req *types.QueryFinalizedChainInfoUntilHeightRequest) (*types.QueryFinalizedChainInfoUntilHeightResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if len(req.ChainId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "chain ID cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	resp := &types.QueryFinalizedChainInfoUntilHeightResponse{}

	// find and assign the last finalised chain info and the earliest epoch that snapshots this chain info
	finalizedEpoch, chainInfo, err := k.GetLastFinalizedChainInfo(ctx, req.ChainId)
	if err != nil {
		return nil, err
	}
	resp.FinalizedChainInfo = chainInfo

	if chainInfo.LatestHeader.Height <= req.Height { // the requested height is after the last finalised chain info
		// find and assign the epoch metadata of the finalised epoch
		resp.EpochInfo, err = k.epochingKeeper.GetHistoricalEpoch(ctx, finalizedEpoch)
		if err != nil {
			return nil, err
		}

		rawCheckpoint, err := k.checkpointingKeeper.GetRawCheckpoint(ctx, finalizedEpoch)

		if err != nil {
			return nil, err
		}

		resp.RawCheckpoint = rawCheckpoint.Ckpt

		// find and assign the raw checkpoint and the best submission key for the finalised epoch
		_, resp.BtcSubmissionKey, err = k.btccKeeper.GetBestSubmission(ctx, finalizedEpoch)
		if err != nil {
			return nil, err
		}
	} else { // the requested height is before the last finalised chain info
		// starting from the requested height, iterate backward until a timestamped header
		closestHeader, err := k.FindClosestHeader(ctx, req.ChainId, req.Height)
		if err != nil {
			return nil, err
		}
		// assign the finalizedEpoch, and retrieve epoch info, raw ckpt and submission key
		finalizedEpoch = closestHeader.BabylonEpoch
		chainInfo, err = k.GetEpochChainInfo(ctx, req.ChainId, finalizedEpoch)
		if err != nil {
			return nil, err
		}
		resp.FinalizedChainInfo = chainInfo
		resp.EpochInfo, err = k.epochingKeeper.GetHistoricalEpoch(ctx, finalizedEpoch)
		if err != nil {
			return nil, err
		}

		rawCheckpoint, err := k.checkpointingKeeper.GetRawCheckpoint(ctx, finalizedEpoch)

		if err != nil {
			return nil, err
		}

		resp.RawCheckpoint = rawCheckpoint.Ckpt

		_, resp.BtcSubmissionKey, err = k.btccKeeper.GetBestSubmission(ctx, finalizedEpoch)
		if err != nil {
			return nil, err
		}
	}

	// if the query does not want the proofs, return here
	if !req.Prove {
		return resp, nil
	}

	// generate all proofs
	resp.Proof, err = k.proveFinalizedChainInfo(ctx, chainInfo, resp.EpochInfo, resp.BtcSubmissionKey)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
