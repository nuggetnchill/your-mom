package yourmom_test

import (
	"testing"

	keepertest "github.com/nuggetnchill/your-mom/testutil/keeper"
	"github.com/nuggetnchill/your-mom/testutil/nullify"
	"github.com/nuggetnchill/your-mom/x/yourmom"
	"github.com/nuggetnchill/your-mom/x/yourmom/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.YourmomKeeper(t)
	yourmom.InitGenesis(ctx, *k, genesisState)
	got := yourmom.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
