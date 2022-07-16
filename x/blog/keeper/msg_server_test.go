package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/uwezukwechibuzor/planet/testutil/keeper"
	"github.com/uwezukwechibuzor/planet/x/blog/keeper"
	"github.com/uwezukwechibuzor/planet/x/blog/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlogKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
