// Copyright 2020 Findora, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package services

import (
	"context"

	"github/findoranetwork/findora-rosetta/configuration"

	"github.com/findoranetwork/rosetta-sdk-go/server"
	"github.com/findoranetwork/rosetta-sdk-go/types"
)

// MempoolAPIService implements the server.MempoolAPIServicer interface.
type MempoolAPIService struct {
	config *configuration.Configuration
	client Client
}

// NewMempoolAPIService creates a new instance of a MempoolAPIService.
func NewMempoolAPIService(
	config *configuration.Configuration,
	client Client,
) server.MempoolAPIServicer {
	return &MempoolAPIService{
		config: config,
		client: client,
	}
}

// Mempool implements the /mempool endpoint.
func (s *MempoolAPIService) Mempool(
	ctx context.Context,
	_ *types.NetworkRequest,
) (*types.MempoolResponse, *types.Error) {
	if s.config.Mode != configuration.Online {
		return nil, ErrUnavailableOffline
	}

	response, err := s.client.GetMempool(ctx)
	if err != nil {
		return nil, wrapErr(ErrFindora, err)
	}

	return response, nil
}

// MempoolTransaction implements the /mempool/transaction endpoint.
func (s *MempoolAPIService) MempoolTransaction(
	ctx context.Context,
	request *types.MempoolTransactionRequest,
) (*types.MempoolTransactionResponse, *types.Error) {
	return nil, wrapErr(ErrUnimplemented, nil)
}
