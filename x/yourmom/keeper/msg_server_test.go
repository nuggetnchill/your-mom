package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/nuggetnchill/your-mom/testutil/keeper"
	"github.com/nuggetnchill/your-mom/x/yourmom/keeper"
	"github.com/nuggetnchill/your-mom/x/yourmom/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.YourmomKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
