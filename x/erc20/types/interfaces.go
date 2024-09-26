// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/evmos/ethermint/x/evm/statedb"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

// AccountKeeper defines the expected interface needed to retrieve account info.
type AccountKeeper interface {
	GetSequence(context.Context, sdk.AccAddress) (uint64, error)
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
}

// StakingKeeper defines the expected interface needed to retrieve the staking denom.
type StakingKeeper interface {
	BondDenom(ctx context.Context) (string, error)
}

// EVMKeeper defines the expected EVM keeper interface used on erc20
type EVMKeeper interface {
	GetParams(ctx sdk.Context) evmtypes.Params
	GetAccountWithoutBalance(ctx sdk.Context, addr common.Address) *statedb.Account
	ApplyMessage(ctx sdk.Context, msg core.Message, tracer vm.EVMLogger, commit bool) (*evmtypes.MsgEthereumTxResponse, error)
	DeleteAccount(ctx sdk.Context, addr common.Address) error
	CallEVM(ctx sdk.Context, abi abi.ABI, from, contract common.Address, commit bool, method string, args ...interface{}) (*evmtypes.MsgEthereumTxResponse, error)
	CallEVMWithData(ctx sdk.Context, from common.Address, contract *common.Address, data []byte, commit bool) (*evmtypes.MsgEthereumTxResponse, error)
}

type (
	LegacyParams = paramtypes.ParamSet
	// Subspace defines an interface that implements the legacy Cosmos SDK x/params Subspace type.
	// NOTE: This is used solely for migration of the Cosmos SDK x/params managed parameters.
	Subspace interface {
		GetParamSet(ctx sdk.Context, ps LegacyParams)
		WithKeyTable(table paramtypes.KeyTable) paramtypes.Subspace
	}
)