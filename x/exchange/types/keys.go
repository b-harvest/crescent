package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/crescent-network/crescent/v5/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "exchange"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

var (
	LastOrderIdKey              = []byte{0x01}
	SpotMarketKeyPrefix         = []byte{0x02}
	SpotLimitOrderKeyPrefix     = []byte{0x03}
	SpotOrderBookOrderKeyPrefix = []byte{0x04}
)

func GetSpotMarketKey(marketId string) []byte {
	return utils.Key(SpotMarketKeyPrefix, []byte(marketId))
}

func GetSpotLimitOrderKey(marketId string, orderId uint64) (key []byte) {
	return utils.Key(SpotLimitOrderKeyPrefix, utils.LengthPrefixString(marketId), sdk.Uint64ToBigEndian(orderId))
}

func GetSpotOrderBookOrderKey(marketId string, isBuy bool, price sdk.Dec, orderId uint64) (key []byte) {
	var orderIdBytes []byte
	if isBuy {
		orderIdBytes = sdk.Uint64ToBigEndian(-orderId)
	} else {
		orderIdBytes = sdk.Uint64ToBigEndian(orderId)
	}
	return utils.Key(
		SpotOrderBookOrderKeyPrefix,
		utils.LengthPrefixString(marketId),
		isBuyToBytes(isBuy),
		sdk.SortableDecBytes(price),
		orderIdBytes)
}

func GetSpotOrderBookIteratorPrefix(marketId string, isBuy bool) (key []byte) {
	return utils.Key(
		SpotOrderBookOrderKeyPrefix,
		utils.LengthPrefixString(marketId),
		isBuyToBytes(isBuy))
}

func GetSpotOrderBookIteratorEndBytes(marketId string, isBuy bool, price sdk.Dec) []byte {
	prefix := utils.Key(
		SpotOrderBookOrderKeyPrefix,
		utils.LengthPrefixString(marketId),
		isBuyToBytes(isBuy),
		sdk.SortableDecBytes(price))
	if isBuy {
		return prefix
	}
	return sdk.PrefixEndBytes(prefix)
}

var (
	buyBytes  = []byte{1}
	sellBytes = []byte{0}
)

func isBuyToBytes(isBuy bool) []byte {
	if isBuy {
		return buyBytes
	}
	return sellBytes
}
