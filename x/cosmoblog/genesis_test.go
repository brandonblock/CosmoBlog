package cosmoblog_test

import (
	"testing"

	keepertest "github.com/example/CosmoBlog/testutil/keeper"
	"github.com/example/CosmoBlog/x/cosmoblog"
	"github.com/example/CosmoBlog/x/cosmoblog/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmoblogKeeper(t)
	cosmoblog.InitGenesis(ctx, *k, genesisState)
	got := cosmoblog.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
