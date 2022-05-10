package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
)

func (k Keeper) GetAllowedCU(ctx sdk.Context, entry *epochstoragetypes.StakeEntry) (uint64, error) {
	var allowedCU uint64 = 0
	found := false
	stakeToMaxCUMap := k.StakeToMaxCUList(ctx).List

	for _, stakeToCU := range stakeToMaxCUMap {
		if entry.Stake.IsGTE(stakeToCU.StakeThreshold) {
			allowedCU = stakeToCU.MaxComputeUnits
			found = true
		} else {
			break
		}
	}
	if found {
		return allowedCU, nil
	}
	return 0, nil
}

func (k Keeper) EnforceClientCUsUsageInEpoch(ctx sdk.Context, chainID string, clientEntry *epochstoragetypes.StakeEntry, totalCUInEpochForUserProvider uint64, providerAddr sdk.AccAddress) error {
	allowedCU, err := k.GetAllowedCU(ctx, clientEntry)
	if err != nil {
		panic(fmt.Sprintf("user %s, allowedCU was not found for stake of: %d", clientEntry, clientEntry.Stake.Amount.Int64()))
	}

	if allowedCU == 0 {
		// k.Logger(ctx).Error("!!!! bbbbbbxxxxxx !!!!!!!!!")
		return fmt.Errorf("user %s, MaxCU was not found for stake of: %d", clientEntry, clientEntry.Stake.Amount.Int64())
		// panic(fmt.Sprintf("user %s, MaxCU was not found for stake of: %d", clientEntry, clientEntry.Stake.Amount.Int64()))
	}
	allowedCU = allowedCU / k.ServicersToPairCount(ctx)
	if totalCUInEpochForUserProvider > allowedCU {
		return k.LimitClientPairingsAndMarkForPenalty(ctx, chainID, clientEntry, totalCUInEpochForUserProvider, allowedCU, providerAddr)
	}

	return nil
}

func getMaxCULimitsPercentage() (float64, float64) {
	// TODO: Get from param
	slashLimitP, unpayLimitP := 0.1, 0.2 // 10% , 20%
	return slashLimitP, unpayLimitP
}

func (k Keeper) GetOverusedCUSumPercentage(ctx sdk.Context, chainID string, clientEntry *epochstoragetypes.StakeEntry, providerAddr sdk.AccAddress) (overusedSumTotalP float64, overusedSumProviderP float64, err error) {
	//TODO: Caching will save a lot of time...
	clientAddr := sdk.AccAddress(clientEntry.Address)
	epochLast := k.epochStorageKeeper.GetEpochStart(ctx)

	// for every epoch in memory
	for epoch := k.epochStorageKeeper.GetEarliestEpochStart(ctx); epoch <= epochLast; epoch = k.epochStorageKeeper.GetNextEpoch(ctx, epoch) {
		// get epochPayments for this client
		clientPaymentStorage, found := k.GetClientPaymentStorage(ctx, k.GetClientPaymentStorageKey(ctx, chainID, epoch, sdk.AccAddress(clientAddr)))
		if !found { // no payments this epoch, continue + advance epoch
			continue
		}
		// for every unique payment of client for this epoch
		uniquePaymentStoragesClientProvider := clientPaymentStorage.UniquePaymentStorageClientProvider
		for _, uniquePaymentStorageClientProvider := range uniquePaymentStoragesClientProvider {
			paymentProviderAddr := k.GetProviderFromUniquePayment(ctx, *uniquePaymentStorageClientProvider)

			// get current stake of client for this epoch
			currentStakeEntry, stakeErr := k.epochStorageKeeper.GetStakeEntryForClientEpoch(ctx, chainID, clientAddr, epoch)
			if stakeErr != nil {
				panic(fmt.Sprintf("Got payment but not stake - epoch %s for client %s, chainID %s ", epoch, clientAddr, chainID))
			}
			// get allowed of client for this epoch
			allowedCU, allowedCUErr := k.GetAllowedCU(ctx, currentStakeEntry)
			if allowedCUErr != nil {
				panic(fmt.Sprintf("Could not find allowedCU for client %s , epoch %s, stakeEntry %s", clientAddr, epoch, currentStakeEntry))
			}
			var overusedPercent float64 = 0.0
			var providersCount float64 = 0.0 // Get from param
			// ? do we use UsedCU for this payment or do we calculate this over the sum for this epoch
			// ? do we divide the allowedCU by providerCount ? how do we calc this percentage
			overusedPercent = float64(uniquePaymentStorageClientProvider.UsedCU) / (float64(allowedCU) / float64(providersCount))
			// TODO: do the math! - calc precentage of overused for allowedCU

			// overused cu found !
			if overusedPercent > 0 {
				// add overused percentage to Total
				overusedSumTotalP += overusedPercent
				if paymentProviderAddr == providerAddr.String() { //TODO: compare AccAddress not strings
					// add overused percentage for the selected Provider
					overusedSumProviderP += overusedPercent
				}
			}
		}
		epoch = k.epochStorageKeeper.GetNextEpoch(ctx, epoch)
	}
	return overusedSumProviderP, overusedSumTotalP, nil
}

func (k Keeper) LimitClientPairingsAndMarkForPenalty(ctx sdk.Context, chainID string, clientEntry *epochstoragetypes.StakeEntry, totalCUInEpochForUserProvider uint64, allowedCU uint64, providerAddr sdk.AccAddress) error {
	implementationFinished := false
	if !implementationFinished {

		k.Logger(ctx).Error("lava_LimitClientPairingsAndMarkForPenalty not fully implemented")
		return fmt.Errorf("lava_LimitClientPairingsAndMarkForPenalty not paying provider")
	} else {

		overused := totalCUInEpochForUserProvider - allowedCU
		overusedP := float64(overused / allowedCU)
		slashLimitP, unpayLimitP := getMaxCULimitsPercentage()
		overusedSumProviderP, overusedSumTotalP, err := k.GetOverusedCUSumPercentage(ctx, chainID, clientEntry, providerAddr)
		if err != nil {
			panic(fmt.Sprintf("user %s, could not calculate overusedCU from memory: %s", clientEntry, clientEntry.Stake.Amount))
		}

		if overusedSumTotalP+overusedP > slashLimitP || overusedSumProviderP+overusedP > slashLimitP/float64(k.ServicersToPairCount(ctx)) {
			k.SlashUser(ctx, clientEntry.Address)
		}

		if overusedSumTotalP+overusedP < unpayLimitP && overusedSumProviderP+overusedP < unpayLimitP/float64(k.ServicersToPairCount(ctx)) {
			// overuse is under the limit - will allow provider to get payment
			return nil
		}
		// overused is not over the slashLimit, but over the unpayLimit - not paying provider
		return fmt.Errorf("user %s bypassed allowed CU %d by using: %d", clientEntry, allowedCU, totalCUInEpochForUserProvider)
	}
}

func (k Keeper) SlashUser(ctx sdk.Context, clientAddr string) {
	//TODO: jail user, and count problems
}
