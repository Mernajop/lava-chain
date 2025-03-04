package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/lavanet/lava/v2/x/projects/keeper"
	"github.com/lavanet/lava/v2/x/projects/types"
)

func SimulateMsgSetPolicy(
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSetPolicy{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SetPolicy simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SetPolicy simulation not implemented"), nil, nil
	}
}
