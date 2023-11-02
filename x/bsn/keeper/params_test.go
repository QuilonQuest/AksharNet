package keeper_test

import (
	"testing"

	testkeeper "github.com/gokulsan/bsn/testutil/keeper"
	"github.com/gokulsan/bsn/x/bsn/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BsnKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
