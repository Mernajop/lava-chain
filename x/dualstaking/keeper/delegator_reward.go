package keeper

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/x/dualstaking/types"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
)

// SetDelegatorReward set a specific DelegatorReward in the store from its index
func (k Keeper) SetDelegatorReward(ctx sdk.Context, delegatorReward types.DelegatorReward) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegatorRewardKeyPrefix))
	index := types.DelegationKey(delegatorReward.Provider, delegatorReward.Delegator, delegatorReward.ChainId)
	b := k.cdc.MustMarshal(&delegatorReward)
	store.Set(types.DelegatorRewardKey(
		index,
	), b)
}

// GetDelegatorReward returns a DelegatorReward from its index
func (k Keeper) GetDelegatorReward(
	ctx sdk.Context,
	index string,
) (val types.DelegatorReward, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegatorRewardKeyPrefix))

	b := store.Get(types.DelegatorRewardKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDelegatorReward removes a DelegatorReward from the store
func (k Keeper) RemoveDelegatorReward(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegatorRewardKeyPrefix))
	store.Delete(types.DelegatorRewardKey(
		index,
	))
}

// GetAllDelegatorReward returns all DelegatorReward
func (k Keeper) GetAllDelegatorReward(ctx sdk.Context) (list []types.DelegatorReward) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegatorRewardKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DelegatorReward
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// The reward for servicing a consumer is divided between the provider and its delegators.
// The following terms are used when calculating the reward distribution between them:
//   1. TotalReward: The total reward before it's divided to the provider and delegators
//   2. ProviderReward: The provider's part in the reward
//   3. DelegatorsReward: The total reward for all delegators
//   4. DelegatorReward: The reward for a specific delegator
//   5. TotalDelegations: The total sum of delegations of a specific provider
//   6. effectiveStake: The provider's stake + TotalDelegations ("effective stake")
//   7. DelegationLimit: The provider's delegation limit (e.g. the max value of delegations the provider uses)
//   8. effectiveDelegations: min(TotalDelegations, DelegationLimit)

// CalcRewards calculates the provider reward and the total reward for delegators
// providerReward = totalReward * ((effectiveDelegations*commission + providerStake) / effectiveStake)
// delegatorsReward = totalReward - providerReward
func (k Keeper) CalcRewards(stakeEntry epochstoragetypes.StakeEntry, totalReward math.Int) (providerReward math.Int, delegatorsReward math.Int) {
	effectiveDelegations, effectiveStake := k.CalcEffectiveDelegationsAndStake(stakeEntry)

	providerReward = totalReward.
		Mul(stakeEntry.Stake.Amount.
			Add(effectiveDelegations.MulRaw(int64(stakeEntry.DelegateCommission)).QuoRaw(100))).
		Quo(effectiveStake)

	return providerReward, totalReward.Sub(providerReward)
}

// CalcEffectiveDelegationsAndStake calculates the effective stake and effective delegations (for delegator rewards calculations)
// effectiveDelegations = min(totalDelegations, delegateLimit)
// effectiveStake = effectiveDelegations + providerStake
func (k Keeper) CalcEffectiveDelegationsAndStake(stakeEntry epochstoragetypes.StakeEntry) (effectiveDelegations math.Int, effectiveStake math.Int) {
	effectiveDelegationsInt64 := math.Min(stakeEntry.DelegateTotal.Amount.Int64(), stakeEntry.DelegateLimit.Amount.Int64())
	effectiveDelegations = math.NewInt(effectiveDelegationsInt64)
	return effectiveDelegations, effectiveDelegations.Add(stakeEntry.Stake.Amount)
}

// CalcDelegatorReward calculates a single delegator reward according to its delegation
// delegatorReward = delegatorsReward * (delegatorStake / totalDelegations) = (delegatorsReward * delegatorStake) / totalDelegations
func (k Keeper) CalcDelegatorReward(delegatorsReward math.Int, totalDelegations math.Int, delegation types.Delegation) math.Int {
	return delegatorsReward.Mul(delegation.Amount.Amount).Quo(totalDelegations)
}
