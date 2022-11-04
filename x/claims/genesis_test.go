package claims_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "gitshock/testutil/keeper"
	"gitshock/testutil/nullify"
	"gitshock/x/claims"
	"gitshock/x/claims/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ClaimsKeeper(t)
	claims.InitGenesis(ctx, *k, genesisState)
	got := claims.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
