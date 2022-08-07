package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "scavenge/testutil/keeper"
	"scavenge/x/scavenge/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ScavengeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
