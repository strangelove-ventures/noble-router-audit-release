package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k *Keeper) Prune(ctx sdk.Context) {
	mints := k.GetAllMints(ctx)

	height := ctx.BlockHeight()
	params := k.GetParams(ctx)

	for _, mint := range mints {
		if uint64(height)-mint.Height > params.MintPruneBlocks {
			if _, found := k.GetIBCForward(ctx, mint.SourceDomain, mint.Nonce); !found {
				k.DeleteMint(ctx, mint.SourceDomain, mint.Nonce)
			}
		}
	}
}
