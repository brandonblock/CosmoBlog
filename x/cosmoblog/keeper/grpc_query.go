package keeper

import (
	"github.com/example/CosmoBlog/x/cosmoblog/types"
)

var _ types.QueryServer = Keeper{}
