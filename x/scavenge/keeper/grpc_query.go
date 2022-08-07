package keeper

import (
	"scavenge/x/scavenge/types"
)

var _ types.QueryServer = Keeper{}
