package cosmoblog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/example/CosmoBlog/x/cosmoblog/keeper"
	"github.com/example/CosmoBlog/x/cosmoblog/types"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePost) (*sdk.Result, error) {
	k.CreatePost(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
