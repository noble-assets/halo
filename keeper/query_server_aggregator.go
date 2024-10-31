// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/halo/v2/types/aggregator"
)

var _ aggregator.QueryServer = &aggregatorQueryServer{}

type aggregatorQueryServer struct {
	*Keeper
}

func NewAggregatorQueryServer(keeper *Keeper) aggregator.QueryServer {
	return &aggregatorQueryServer{Keeper: keeper}
}

func (k aggregatorQueryServer) Reporter(ctx context.Context, req *aggregator.QueryReporter) (*aggregator.QueryReporterResponse, error) {
	if req == nil {
		return nil, errors.ErrInvalidRequest
	}

	reporter, err := k.Keeper.Reporter.Get(ctx)
	return &aggregator.QueryReporterResponse{Reporter: reporter}, err
}
