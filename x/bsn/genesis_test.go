package bsn_test

import (
	"testing"

	keepertest "github.com/gokulsan/bsn/testutil/keeper"
	"github.com/gokulsan/bsn/testutil/nullify"
	"github.com/gokulsan/bsn/x/bsn"
	"github.com/gokulsan/bsn/x/bsn/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BsnKeeper(t)
	bsn.InitGenesis(ctx, *k, genesisState)
	got := bsn.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
