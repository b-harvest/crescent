package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/crescent-network/crescent/v5/x/amm/types"
)

func (k Keeper) GetLastPoolId(ctx sdk.Context) (poolId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastPoolIdKey)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetLastPoolId(ctx sdk.Context, poolId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastPoolIdKey, sdk.Uint64ToBigEndian(poolId))
}

func (k Keeper) GetNextPoolIdWithUpdate(ctx sdk.Context) uint64 {
	poolId := k.GetLastPoolId(ctx)
	poolId++
	k.SetLastPoolId(ctx, poolId)
	return poolId
}

func (k Keeper) GetLastPositionId(ctx sdk.Context) (positionId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastPositionIdKey)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetLastPositionId(ctx sdk.Context, positionId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastPositionIdKey, sdk.Uint64ToBigEndian(positionId))
}

func (k Keeper) GetNextPositionIdWithUpdate(ctx sdk.Context) uint64 {
	positionId := k.GetLastPositionId(ctx)
	positionId++
	k.SetLastPositionId(ctx, positionId)
	return positionId
}

func (k Keeper) GetPool(ctx sdk.Context, poolId uint64) (pool types.Pool, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPoolKey(poolId))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshal(bz, &pool)
	return pool, true
}

func (k Keeper) SetPool(ctx sdk.Context, pool types.Pool) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&pool)
	store.Set(types.GetPoolKey(pool.Id), bz)
}

func (k Keeper) GetPosition(ctx sdk.Context, positionId uint64) (position types.Position, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPositionKey(positionId))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshal(bz, &position)
	return position, true
}

func (k Keeper) GetPositionByParams(ctx sdk.Context, poolId uint64, ownerAddr sdk.AccAddress, lowerTick, upperTick int32) (position types.Position, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPositionIndexKey(poolId, ownerAddr, lowerTick, upperTick))
	if bz == nil {
		return
	}
	return k.GetPosition(ctx, sdk.BigEndianToUint64(bz))
}

func (k Keeper) SetPosition(ctx sdk.Context, position types.Position) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&position)
	store.Set(types.GetPositionKey(position.Id), bz)
}

func (k Keeper) SetPositionIndex(ctx sdk.Context, position types.Position) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPositionIndexKey(
		position.PoolId, sdk.MustAccAddressFromBech32(position.Owner),
		position.LowerTick, position.UpperTick),
		sdk.Uint64ToBigEndian(position.Id))
}

func (k Keeper) GetTickInfo(ctx sdk.Context, poolId uint64, tick int32) (tickInfo types.TickInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTickInfoKey(poolId, tick))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshal(bz, &tickInfo)
	return tickInfo, true
}

func (k Keeper) SetTickInfo(ctx sdk.Context, poolId uint64, tick int32, tickInfo types.TickInfo) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&tickInfo)
	store.Set(types.GetTickInfoKey(poolId, tick), bz)
}
