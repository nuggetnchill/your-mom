package keeper_test

import (
	"testing"

	testkeeper "github.com/nuggetnchill/your-mom/testutil/keeper"
	"github.com/nuggetnchill/your-mom/x/yourmom/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.YourmomKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
