package gitshock_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "gitshock/testutil/keeper"
	"gitshock/testutil/nullify"
	"gitshock/x/gitshock"
	"gitshock/x/gitshock/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GitshockKeeper(t)
	gitshock.InitGenesis(ctx, *k, genesisState)
	got := gitshock.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
