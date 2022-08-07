package keeper

import (
	"context"

  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
  "github.com/tendermint/tendermint/crypto"

  "scavenge/x/scavenge/types"
)

func (k msgServer) SubmitScavenge(goCtx context.Context, msg *types.MsgSubmitScavenge) (*types.MsgSubmitScavengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	var scavenge = types.Scavenge{
		Index: msg.SolutionHash,
		Description: msg.Description,
		SolutionHash: msg.SolutionHash,
		Reward: msg.Reward,
	}

	_, isFound := k.GetScavenge(ctx, scavenge.SolutionHash)

	if isFound{
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Scavenge with that solution hash already exists")
	}

	//get address of the Scavenge module account
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	scavenger, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
			panic(err)
	}
	// convert tokens from string into sdk.Coins
	reward, err := sdk.ParseCoinsNormalized(scavenge.Reward)
	if err != nil {
			panic(err)
	}

	// send tokens from the scavenge creator to the module account
	sdkError := k.bankKeeper.SendCoins(ctx, scavenger, moduleAcct, reward)
	if sdkError != nil {
			return nil, sdkError
	}
	
	// write the scavenge to the store
	k.SetScavenge(ctx, scavenge)
	return &types.MsgSubmitScavengeResponse{}, nil
}
