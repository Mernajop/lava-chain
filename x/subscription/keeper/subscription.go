package keeper

import (
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	legacyerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lavanet/lava/utils"
	planstypes "github.com/lavanet/lava/x/plans/types"
	projectstypes "github.com/lavanet/lava/x/projects/types"
	"github.com/lavanet/lava/x/subscription/types"
)

// GetSubscription returns the subscription of a given consumer
func (k Keeper) GetSubscription(ctx sdk.Context, consumer string) (val types.Subscription, found bool) {
	block := uint64(ctx.BlockHeight())

	var sub types.Subscription
	found = k.subsFS.FindEntry(ctx, consumer, block, &sub)

	return sub, found
}

// CreateSubscription creates a subscription for a consumer
func (k Keeper) CreateSubscription(
	ctx sdk.Context,
	creator string,
	consumer string,
	planIndex string,
	duration uint64,
	autoRenewalFlag bool,
) error {
	var err error

	block := uint64(ctx.BlockHeight())
	creatorAcct, plan, err := k.verifySubscriptionBuyInputAndGetPlan(ctx, block, creator, consumer, planIndex)
	if err != nil {
		return err
	}

	var sub types.Subscription
	found := k.subsFS.FindEntry(ctx, consumer, block, &sub)

	// Subscription creation:
	//   When: if not already exists for consumer address)
	//   What: find plan, create default project, set duration, calculate price,
	//         charge fees, save subscription.
	//
	// Subscription renewal:
	//   When: if already exists and existing plan is the same as current plans
	//         ("same" means same index and same block of creation)
	//   What: find plan, update duration (total and remaining), calculate price,
	//         charge fees, save subscription.
	//
	// Subscription upgrade:
	//   When: if already exists and existing plan, and new plan's price is higher than current plan's
	//   What: find subscription, check for sufficient funds, upgrade.
	//
	// Subscription downgrade: (TBD)

	if !found {
		sub, err = k.createNewSubscription(ctx, &plan, creator, consumer, block, autoRenewalFlag)
		if err != nil {
			return utils.LavaFormatWarning("failed to create subscription", err)
		}
	} else {
		// allow renewal with the same plan ("same" means both plan index);
		// otherwise, upgrade the plan.
		if plan.Index != sub.PlanIndex {
			currentPlan, err := k.GetPlanFromSubscription(ctx, consumer, block)
			if err != nil {
				return utils.LavaFormatError("failed to find plan for current subscription", err,
					utils.LogAttr("consumer", consumer),
					utils.LogAttr("block", block),
				)
			}
			if plan.Price.Amount.LT(currentPlan.Price.Amount) {
				return utils.LavaFormatError("New plan's price must be higher than the old plan", nil,
					utils.LogAttr("consumer", consumer),
					utils.LogAttr("currentPlan", currentPlan),
					utils.LogAttr("newPlan", plan),
					utils.LogAttr("block", block),
				)
			}
			originalPlanIndex := sub.PlanIndex
			originalPlanBlock := sub.PlanBlock
			err = k.upgradeSubscriptionPlan(ctx, duration, &sub, plan.Index, plan.Block)
			if err != nil {
				return err
			}

			// Upgrade worked, now we can decrease the old plan's refcount
			k.plansKeeper.PutPlan(ctx, originalPlanIndex, originalPlanBlock)

			details := map[string]string{
				"consumer": consumer,
				"oldPlan":  sub.PlanIndex,
				"newPlan":  plan.Index,
			}
			utils.LogLavaEvent(ctx, k.Logger(ctx), types.UpgradeSubscriptionEventName, details, "subscription upgraded")
		} else if plan.Block != sub.PlanBlock {
			// if the plan was changed since the subscription originally bought it, allow extension
			// only if the updated plan's price is up to 5% more than the original price
			subPlan, found := k.plansKeeper.FindPlan(ctx, plan.Index, sub.PlanBlock)
			if !found {
				return utils.LavaFormatError("cannot extend subscription", fmt.Errorf("critical: cannot find subscription's plan"),
					utils.Attribute{Key: "plan_index", Value: plan.Index},
					utils.Attribute{Key: "block", Value: strconv.FormatUint(sub.PlanBlock, 10)})
			}
			if plan.Price.Amount.GT(subPlan.Price.Amount.MulRaw(105).QuoRaw(100)) {
				// current_plan_price > 1.05 * original_plan_price
				return utils.LavaFormatError("cannot auto-extend subscription", fmt.Errorf("subscription's original plan price is lower by more than 5 percent than current plan"),
					utils.Attribute{Key: "sub", Value: sub.Consumer},
					utils.Attribute{Key: "plan", Value: plan.Index},
					utils.Attribute{Key: "original_plan_block", Value: strconv.FormatUint(subPlan.Block, 10)},
					utils.Attribute{Key: "original_plan_price", Value: subPlan.Price.String()},
					utils.Attribute{Key: "current_plan_block", Value: strconv.FormatUint(plan.Block, 10)},
					utils.Attribute{Key: "current_plan_block", Value: plan.Price.String()},
				)
			}

			// update sub plan
			sub.PlanBlock = plan.Block
		}

		// The total duration may not exceed MAX_SUBSCRIPTION_DURATION, but allow an
		// extra month to account for renewals before the end of current subscription
		if sub.DurationLeft+duration > types.MAX_SUBSCRIPTION_DURATION+1 {
			str := strconv.FormatInt(types.MAX_SUBSCRIPTION_DURATION, 10)
			return utils.LavaFormatWarning("duration would exceed limit ("+str+" months)",
				fmt.Errorf("subscription renewal failed"),
				utils.Attribute{Key: "duration", Value: sub.DurationLeft},
			)
		}
	}

	// update total (last requested) duration and remaining duration
	sub.DurationBought = duration
	sub.DurationLeft += duration

	if err := sub.ValidateSubscription(); err != nil {
		return utils.LavaFormatWarning("create subscription failed", err)
	}

	// subscription looks good; let's charge the creator
	price := plan.GetPrice()
	price.Amount = price.Amount.MulRaw(int64(duration))
	k.applyPlanDiscountIfEligible(duration, &plan, &price)

	err = k.chargeFromCreatorAccountToModule(ctx, creatorAcct, price)
	if err != nil {
		return err
	}

	if !found {
		expiry := uint64(utils.NextMonth(ctx.BlockTime()).UTC().Unix())
		sub.MonthExpiryTime = expiry
		k.subsTS.AddTimerByBlockTime(ctx, expiry, []byte(consumer), []byte{})
		err = k.subsFS.AppendEntry(ctx, consumer, block, &sub)
	} else {
		k.subsFS.ModifyEntry(ctx, consumer, sub.Block, &sub)
	}

	return err
}

func (k Keeper) verifySubscriptionBuyInputAndGetPlan(ctx sdk.Context, block uint64, creator, consumer, planIndex string) (creatorAcct sdk.AccAddress, plan planstypes.Plan, err error) {
	EMPTY_PLAN := planstypes.Plan{}

	if _, err := sdk.AccAddressFromBech32(consumer); err != nil {
		return nil, EMPTY_PLAN, utils.LavaFormatWarning("invalid subscription consumer address", err,
			utils.Attribute{Key: "consumer", Value: consumer},
		)
	}

	creatorAcct, err = sdk.AccAddressFromBech32(creator)
	if err != nil {
		return nil, EMPTY_PLAN, utils.LavaFormatWarning("invalid subscription creator address", err,
			utils.Attribute{Key: "creator", Value: creator},
		)
	}

	plan, found := k.plansKeeper.GetPlan(ctx, planIndex)
	if !found {
		return nil, EMPTY_PLAN, utils.LavaFormatWarning("cannot create subscription with invalid plan", nil,
			utils.Attribute{Key: "plan", Value: planIndex},
			utils.Attribute{Key: "block", Value: block},
		)
	}
	return
}

func (k Keeper) createNewSubscription(ctx sdk.Context, plan *planstypes.Plan, creator, consumer string,
	block uint64, autoRenewalFlag bool,
) (types.Subscription, error) {
	sub := types.Subscription{
		Creator:       creator,
		Consumer:      consumer,
		Block:         block,
		PlanIndex:     plan.Index,
		PlanBlock:     plan.Block,
		DurationTotal: 0,
		AutoRenewal:   autoRenewalFlag,
	}

	sub.MonthCuTotal = plan.PlanPolicy.GetTotalCuLimit()
	sub.MonthCuLeft = plan.PlanPolicy.GetTotalCuLimit()
	sub.Cluster = types.GetClusterKey(sub)

	// new subscription needs a default project
	err := k.projectsKeeper.CreateAdminProject(ctx, consumer, *plan)
	if err != nil {
		return types.Subscription{}, utils.LavaFormatWarning("failed to create default project", err)
	}

	return sub, nil
}

func (k Keeper) upgradeSubscriptionPlan(ctx sdk.Context, duration uint64, sub *types.Subscription, toPlanIndex string, toPlanBlock uint64) error {
	date := ctx.BlockTime()
	block := uint64(ctx.BlockHeight())

	k.addCuTrackerTimerForSubscription(ctx, block, sub)

	if sub.DurationLeft == 0 {
		// Subscription was already expired. Can't upgrade.
		k.handleZeroDurationLeftForSubscription(ctx, block, date, sub)
		return utils.LavaFormatError("Trying to upgrade subscription, but subscription was already expired", nil,
			utils.LogAttr("consumer", sub.Consumer))
	}

	// The "old" subscription's duration is now expired
	// If called from CreateSubscription, the duration will reset to the duration bought
	sub.DurationLeft = 0
	sub.Cluster = types.GetClusterKey(*sub)

	nextEpoch, err := k.epochstorageKeeper.GetNextEpoch(ctx, block)
	if err != nil {
		return utils.LavaFormatError("Trying to upgrade subscription, but got an error while trying to get next epoch", err,
			utils.LogAttr("consumer", sub.Consumer))
	}

	// Remove one refcount for previous plan
	k.plansKeeper.PutPlan(ctx, sub.PlanIndex, sub.PlanBlock)

	sub.PlanIndex = toPlanIndex
	sub.PlanBlock = toPlanBlock

	return k.resetSubscriptionDetailsAndAppendEntry(ctx, sub, nextEpoch, date)
}

func (k Keeper) advanceMonth(ctx sdk.Context, subkey []byte) {
	date := ctx.BlockTime()
	block := uint64(ctx.BlockHeight())
	consumer := string(subkey)

	var sub types.Subscription
	if !k.verifySubExists(ctx, consumer, block, &sub) {
		return
	}

	k.addCuTrackerTimerForSubscription(ctx, block, &sub)

	if sub.DurationLeft == 0 {
		k.handleZeroDurationLeftForSubscription(ctx, block, date, &sub)
		return
	}

	sub.DurationLeft -= 1

	if sub.DurationLeft > 0 {
		sub.DurationTotal += 1
		k.resetSubscriptionDetailsAndAppendEntry(ctx, &sub, block, date)
	} else {
		if sub.FutureSubscription != nil {
			// Consumer made advance purchase. Now we activate it.
			newSubInfo := sub.FutureSubscription

			plan, found := k.plansKeeper.FindPlan(ctx, newSubInfo.PlanIndex, newSubInfo.PlanBlock)
			if !found {
				utils.LavaFormatWarning("subscription advance purchase failed: could not find plan. removing subscription", nil,
					utils.Attribute{Key: "consumer", Value: sub.Consumer},
					utils.Attribute{Key: "planIndex", Value: newSubInfo.PlanIndex},
				)
				k.RemoveExpiredSubscription(ctx, consumer, block, sub.PlanIndex, sub.PlanBlock)
				return
			}

			sub.Creator = newSubInfo.Creator
			sub.PlanIndex = newSubInfo.PlanIndex
			sub.PlanBlock = newSubInfo.PlanBlock
			sub.DurationBought = newSubInfo.DurationBought
			sub.DurationLeft = newSubInfo.DurationBought
			sub.DurationTotal = 0
			sub.FutureSubscription = nil
			sub.MonthCuTotal = plan.PlanPolicy.TotalCuLimit

			k.resetSubscriptionDetailsAndAppendEntry(ctx, &sub, block, date)
		} else if sub.AutoRenewal {
			// apply the DurationLeft decrease to 0 and buy an extra month
			k.subsFS.ModifyEntry(ctx, sub.Consumer, sub.Block, &sub)
			err := k.CreateSubscription(ctx, sub.Creator, sub.Consumer, sub.PlanIndex, 1, sub.AutoRenewal)
			if err != nil {
				utils.LavaFormatWarning("subscription auto renewal failed. removing subscription", err,
					utils.Attribute{Key: "consumer", Value: sub.Consumer},
				)
				k.RemoveExpiredSubscription(ctx, consumer, block, sub.PlanIndex, sub.PlanBlock)
			}
		} else {
			k.RemoveExpiredSubscription(ctx, consumer, block, sub.PlanIndex, sub.PlanBlock)
		}
	}
}

func (k Keeper) verifySubExists(ctx sdk.Context, consumer string, block uint64, sub *types.Subscription) bool {
	if found := k.subsFS.FindEntry(ctx, consumer, block, sub); !found {
		// subscription (monthly) timer has expired for an unknown subscription:
		// either the timer was set wrongly, or the subscription was incorrectly
		// removed; and we cannot even return an error about it.
		utils.LavaFormatError("critical: month expiry for unknown subscription, skipping",
			fmt.Errorf("subscription not found"),
			utils.Attribute{Key: "consumer", Value: consumer},
			utils.Attribute{Key: "block", Value: block},
		)
		return false
	}
	return true
}

func (k Keeper) addCuTrackerTimerForSubscription(ctx sdk.Context, block uint64, sub *types.Subscription) {
	blocksToSave, err := k.epochstorageKeeper.BlocksToSave(ctx, block)
	if err != nil {
		utils.LavaFormatError("critical: failed assigning CU tracker callback, skipping", err,
			utils.Attribute{Key: "block", Value: block},
		)
	} else {
		k.cuTrackerTS.AddTimerByBlockHeight(ctx, block+blocksToSave-1, []byte(sub.Consumer), []byte(strconv.FormatUint(sub.Block, 10)))
	}
}

func (k Keeper) handleZeroDurationLeftForSubscription(ctx sdk.Context, block uint64, blockTime time.Time, sub *types.Subscription) {
	// subscription duration has already reached zero before and should have
	// been removed before. Extend duration by another month (without adding
	// CUs) to allow smooth operation.
	utils.LavaFormatError("critical: negative duration for subscription, extend by 1 month",
		fmt.Errorf("negative duration in AdvanceMonth()"),
		utils.Attribute{Key: "consumer", Value: sub.Consumer},
		utils.Attribute{Key: "plan", Value: sub.PlanIndex},
		utils.Attribute{Key: "block", Value: block},
	)
	// normally would panic! but can "recover" by auto-extending by 1 month
	// (don't bother to modify sub.MonthExpiryTime to minimize state changes)
	expiry := uint64(utils.NextMonth(blockTime).UTC().Unix())
	k.subsTS.AddTimerByBlockTime(ctx, expiry, []byte(sub.Consumer), []byte{})
}

func (k Keeper) resetSubscriptionDetailsAndAppendEntry(ctx sdk.Context, sub *types.Subscription, block uint64, blockTime time.Time) error {
	// reset projects CU allowance for this coming month
	k.projectsKeeper.SnapshotSubscriptionProjects(ctx, sub.Consumer, block)

	// reset subscription CU allowance for this coming month
	sub.MonthCuLeft = sub.MonthCuTotal
	sub.Block = block

	// restart timer and append new (fixated) version of this subscription
	expiry := uint64(utils.NextMonth(blockTime).UTC().Unix())
	sub.MonthExpiryTime = expiry
	k.subsTS.AddTimerByBlockTime(ctx, expiry, []byte(sub.Consumer), []byte{})

	// since the total duration increases, the cluster might change
	sub.Cluster = types.GetClusterKey(*sub)

	err := k.subsFS.AppendEntry(ctx, sub.Consumer, block, sub)
	if err != nil {
		// Remove new timer if failed
		k.subsTS.DelTimerByBlockTime(ctx, expiry, []byte(sub.Consumer))

		// normally would panic! but ok to ignore - the subscription remains
		// as is with same remaining duration (but not renewed CU)
		return utils.LavaFormatError("critical: failed to recharge subscription", err,
			utils.Attribute{Key: "consumer", Value: sub.Consumer},
			utils.Attribute{Key: "plan", Value: sub.PlanIndex},
			utils.Attribute{Key: "block", Value: block},
		)
	}

	return nil
}

func (k Keeper) applyPlanDiscountIfEligible(duration uint64, plan *planstypes.Plan, price *sdk.Coin) {
	if duration >= utils.MONTHS_IN_YEAR {
		// adjust cost if discount given
		discount := plan.GetAnnualDiscountPercentage()
		if discount > 0 {
			factor := int64(100 - discount)
			price.Amount = price.Amount.MulRaw(factor).QuoRaw(100)
		}
	}
}

func (k Keeper) chargeFromCreatorAccountToModule(ctx sdk.Context, creator sdk.AccAddress, price sdk.Coin) error {
	if k.bankKeeper.GetBalance(ctx, creator, k.stakingKeeper.BondDenom(ctx)).IsLT(price) {
		return utils.LavaFormatWarning("create subscription failed", legacyerrors.ErrInsufficientFunds,
			utils.LogAttr("creator", creator),
			utils.LogAttr("price", price),
		)
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, []sdk.Coin{price})
	if err != nil {
		return utils.LavaFormatError("create subscription failed. funds transfer failed", err,
			utils.LogAttr("creator", creator),
			utils.LogAttr("price", price),
		)
	}

	return nil
}

func (k Keeper) CreateFutureSubscription(ctx sdk.Context,
	creator string,
	consumer string,
	planIndex string,
	duration uint64,
) error {
	var err error

	block := uint64(ctx.BlockHeight())
	creatorAcct, plan, err := k.verifySubscriptionBuyInputAndGetPlan(ctx, block, creator, consumer, planIndex)
	if err != nil {
		return err
	}

	var sub types.Subscription
	found := k.subsFS.FindEntry(ctx, consumer, block, &sub)

	if !found {
		return utils.LavaFormatWarning("could not find active subscription. advanced purchase is only available for an active subscription", nil,
			utils.LogAttr("consumer", consumer),
			utils.LogAttr("block", block),
		)
	}

	createFutureSubscription := func(sub *types.Subscription, creator, planIndex string, planBlock, duration uint64) {
		sub.FutureSubscription = &types.FutureSubscription{
			Creator:        creator,
			PlanIndex:      planIndex,
			PlanBlock:      planBlock,
			DurationBought: duration,
		}
	}

	if sub.FutureSubscription != nil {
		// Consumer already has a future subscription
		// If the new plan's price > current future subscription's plan - change and charge the diff
		currentPlan, err := k.GetPlanFromSubscription(ctx, consumer, block)
		if err != nil {
			return utils.LavaFormatError("panic: could not get active subscription's plan. aborting", err,
				utils.Attribute{Key: "creator", Value: creator},
			)
		}

		newPlanCostForDuration := plan.Price.Amount.MulRaw(int64(duration))
		consumerPaid := currentPlan.Price.Amount.MulRaw(int64(sub.FutureSubscription.DurationBought))
		if newPlanCostForDuration.GT(consumerPaid) {
			priceDiff := sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), newPlanCostForDuration.Sub(consumerPaid))
			err = k.chargeFromCreatorAccountToModule(ctx, creatorAcct, priceDiff)
			if err != nil {
				return err
			}

			createFutureSubscription(&sub, creator, plan.Index, plan.Block, duration)

			details := map[string]string{
				"creator":      creator,
				"consumer":     consumer,
				"duration":     strconv.FormatUint(duration, 10),
				"oldPlanIndex": currentPlan.Index,
				"oldPlanBlock": strconv.FormatUint(currentPlan.Block, 10),
				"newPlanIndex": plan.Index,
				"newPlanBlock": strconv.FormatUint(plan.Block, 10),
			}
			utils.LogLavaEvent(ctx, k.Logger(ctx), types.AdvancedBuyUpgradeSubscriptionEventName, details, "advanced subscription upgraded")
			return nil
		} else {
			return utils.LavaFormatWarning("can't purchase another plan in advanced with a lower price", nil)
		}
	}

	price := plan.GetPrice()
	price.Amount = price.Amount.MulRaw(int64(duration))
	k.applyPlanDiscountIfEligible(duration, &plan, &price)

	err = k.chargeFromCreatorAccountToModule(ctx, creatorAcct, price)
	if err != nil {
		return err
	}

	createFutureSubscription(&sub, creator, plan.Index, plan.Block, duration)

	k.subsFS.ModifyEntry(ctx, consumer, sub.Block, &sub)
	return nil
}

func (k Keeper) RemoveExpiredSubscription(ctx sdk.Context, consumer string, block uint64, planIndex string, planBlock uint64) {
	// delete all projects before deleting
	k.delAllProjectsFromSubscription(ctx, consumer)

	// delete subscription effective now (don't wait for end of epoch)
	k.subsFS.DelEntry(ctx, consumer, block)

	// decrease plan ref count
	k.plansKeeper.PutPlan(ctx, planIndex, planBlock)

	details := map[string]string{"consumer": consumer}
	utils.LogLavaEvent(ctx, k.Logger(ctx), types.ExpireSubscriptionEventName, details, "subscription expired")
}

func (k Keeper) GetPlanFromSubscription(ctx sdk.Context, consumer string, block uint64) (planstypes.Plan, error) {
	var sub types.Subscription
	if found := k.subsFS.FindEntry(ctx, consumer, block, &sub); !found {
		return planstypes.Plan{}, fmt.Errorf("can't find subscription with consumer address: %s", consumer)
	}

	plan, found := k.plansKeeper.FindPlan(ctx, sub.PlanIndex, sub.PlanBlock)
	if !found {
		// normally would panic! but can "recover" by ignoring and returning error
		err := utils.LavaFormatError("critical: failed to find existing subscription plan", legacyerrors.ErrKeyNotFound,
			utils.Attribute{Key: "consumer", Value: consumer},
			utils.Attribute{Key: "planIndex", Value: sub.PlanIndex},
			utils.Attribute{Key: "planBlock", Value: sub.PlanBlock},
		)
		return plan, err
	}

	return plan, nil
}

func (k Keeper) AddProjectToSubscription(ctx sdk.Context, consumer string, projectData projectstypes.ProjectData) error {
	var sub types.Subscription
	if found := k.subsFS.FindEntry(ctx, consumer, uint64(ctx.BlockHeight()), &sub); !found {
		return legacyerrors.ErrKeyNotFound.Wrapf("AddProjectToSubscription_can't_get_subscription_of_%s", consumer)
	}

	plan, found := k.plansKeeper.FindPlan(ctx, sub.PlanIndex, sub.PlanBlock)
	if !found {
		err := utils.LavaFormatError("critical: failed to find existing subscription plan", legacyerrors.ErrKeyNotFound,
			utils.Attribute{Key: "consumer", Value: sub.Creator},
			utils.Attribute{Key: "planIndex", Value: sub.PlanIndex},
			utils.Attribute{Key: "planBlock", Value: sub.PlanBlock},
		)
		return err
	}

	return k.projectsKeeper.CreateProject(ctx, consumer, projectData, plan)
}

func (k Keeper) DelProjectFromSubscription(ctx sdk.Context, consumer, name string) error {
	var sub types.Subscription
	if found := k.subsFS.FindEntry(ctx, consumer, uint64(ctx.BlockHeight()), &sub); !found {
		return legacyerrors.ErrKeyNotFound.Wrapf("AddProjectToSubscription_can't_get_subscription_of_%s", consumer)
	}

	projectID := projectstypes.ProjectIndex(consumer, name)
	return k.projectsKeeper.DeleteProject(ctx, consumer, projectID)
}

func (k Keeper) delAllProjectsFromSubscription(ctx sdk.Context, consumer string) {
	allProjectsIDs := k.projectsKeeper.GetAllProjectsForSubscription(ctx, consumer)
	for _, projectID := range allProjectsIDs {
		err := k.projectsKeeper.DeleteProject(ctx, consumer, projectID)
		if err != nil {
			// TODO: should panic (should never fail because these are exactly the
			// subscription's current projects)
			utils.LavaFormatError("critical: delete project failed at subscription removal", err,
				utils.Attribute{Key: "subscription", Value: consumer},
				utils.Attribute{Key: "project", Value: projectID},
				utils.Attribute{Key: "block", Value: uint64(ctx.BlockHeight())},
			)
		}
	}
}

func (k Keeper) ChargeComputeUnitsToSubscription(ctx sdk.Context, consumer string, block, cuAmount uint64) (types.Subscription, error) {
	var sub types.Subscription
	if found := k.subsFS.FindEntry(ctx, consumer, block, &sub); !found {
		return sub, utils.LavaFormatError("can't charge cu to subscription",
			fmt.Errorf("subscription not found"),
			utils.Attribute{Key: "subscription", Value: consumer},
			utils.Attribute{Key: "block", Value: block},
		)
	}

	if sub.MonthCuLeft < cuAmount {
		sub.MonthCuLeft = 0
	} else {
		sub.MonthCuLeft -= cuAmount
	}

	k.subsFS.ModifyEntry(ctx, consumer, sub.Block, &sub)
	return sub, nil
}
