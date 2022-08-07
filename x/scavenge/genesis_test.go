package scavenge_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "scavenge/testutil/keeper"
	"scavenge/testutil/nullify"
	"scavenge/x/scavenge"
	"scavenge/x/scavenge/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ScavengeKeeper(t)
	scavenge.InitGenesis(ctx, *k, genesisState)
	got := scavenge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
