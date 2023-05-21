package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/crescent-network/crescent/v5/types"
	exchangetypes "github.com/crescent-network/crescent/v5/x/exchange/types"
)

func SqrtPriceAtTick(tick int32) sdk.Dec {
	return utils.DecApproxSqrt(exchangetypes.PriceAtTick(tick))
}

func NewTickInfo() TickInfo {
	return TickInfo{
		GrossLiquidity:              utils.ZeroInt,
		NetLiquidity:                utils.ZeroInt,
		FeeGrowthOutside:            sdk.DecCoins{},
		FarmingRewardsGrowthOutside: sdk.DecCoins{},
	}
}

func (tickInfo TickInfo) Validate() error {
	if tickInfo.GrossLiquidity.IsNegative() {
		return fmt.Errorf("gross liquidity must not be negative: %s", tickInfo.GrossLiquidity)
	}
	if tickInfo.NetLiquidity.IsZero() {
		return fmt.Errorf("net liquidity must not be 0")
	}
	if err := tickInfo.FeeGrowthOutside.Validate(); err != nil {
		return fmt.Errorf("invalid fee growth outside: %w", err)
	}
	if err := tickInfo.FarmingRewardsGrowthOutside.Validate(); err != nil {
		return fmt.Errorf("invalid farming rewards growth outside: %w", err)
	}
	return nil
}
