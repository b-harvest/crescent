package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmosquad-labs/squad/x/liquidstaking/types"
	"github.com/stretchr/testify/require"
)

func TestDivideByWeight(t *testing.T) {
	testCases := []struct {
		whitelistedVals  []types.WhitelistedValidator
		addStakingAmt    sdk.Int
		currentDelShares []sdk.Int
		expectedOutputs  []sdk.Int
		expectedCrumb    sdk.Int
	}{
		{
			whitelistedVals: []types.WhitelistedValidator{
				{
					ValidatorAddress: "a",
					TargetWeight:     sdk.NewInt(1),
				},
				{
					ValidatorAddress: "b",
					TargetWeight:     sdk.NewInt(1),
				},
				{
					ValidatorAddress: "c",
					TargetWeight:     sdk.NewInt(1),
				},
			},
			addStakingAmt:    sdk.NewInt(10 * 1000000),
			currentDelShares: []sdk.Int{sdk.NewInt(2000000), sdk.NewInt(2000000), sdk.NewInt(1000000)},
			expectedOutputs:  []sdk.Int{sdk.NewInt(3333333), sdk.NewInt(3333333), sdk.NewInt(3333333)},
			expectedCrumb:    sdk.NewInt(1),
		},
		{
			whitelistedVals: []types.WhitelistedValidator{
				{
					ValidatorAddress: "a",
					TargetWeight:     sdk.NewInt(2),
				},
				{
					ValidatorAddress: "b",
					TargetWeight:     sdk.NewInt(2),
				},
				{
					ValidatorAddress: "c",
					TargetWeight:     sdk.NewInt(1),
				},
			},
			addStakingAmt:    sdk.NewInt(10 * 1000000),
			currentDelShares: []sdk.Int{sdk.NewInt(1000000), sdk.NewInt(1000000), sdk.NewInt(1000000)},
			expectedOutputs:  []sdk.Int{sdk.NewInt(4000000), sdk.NewInt(4000000), sdk.NewInt(2000000)},
			expectedCrumb:    sdk.NewInt(0),
		},
		{
			whitelistedVals: []types.WhitelistedValidator{
				{
					ValidatorAddress: "a",
					TargetWeight:     sdk.NewInt(1),
				},
				{
					ValidatorAddress: "b",
					TargetWeight:     sdk.NewInt(1),
				},
				{
					ValidatorAddress: "c",
					TargetWeight:     sdk.NewInt(1),
				},
			},
			addStakingAmt:    sdk.NewInt(10),
			currentDelShares: []sdk.Int{sdk.NewInt(3), sdk.NewInt(2), sdk.NewInt(1)},
			expectedOutputs:  []sdk.Int{sdk.NewInt(3), sdk.NewInt(3), sdk.NewInt(3)},
			expectedCrumb:    sdk.NewInt(1),
		},
	}

	for _, tc := range testCases {
		fmt.Println("------")
		require.IsType(t, []types.WhitelistedValidator{}, tc.whitelistedVals)
		require.IsType(t, sdk.Int{}, tc.addStakingAmt)
		require.IsType(t, sdk.Int{}, tc.expectedCrumb)
		require.IsType(t, []sdk.Int{}, tc.expectedOutputs)

		totalTargetAmt := sdk.ZeroInt()
		valMap := types.GetWhitelistedValMap(tc.whitelistedVals)
		var liquidVals types.LiquidValidators
		for _, v := range tc.whitelistedVals {
			liquidVals = append(liquidVals, types.LiquidValidator{
				OperatorAddress: v.ValidatorAddress,
				//Status:          0,
				//Weight:          sdk.Int{},
			})
		}
		outputs, crumb := types.DivideByWeight(liquidVals, tc.addStakingAmt, valMap)
		for k, v := range outputs {
			fmt.Println(k, v.String())
			totalTargetAmt = totalTargetAmt.Add(v)
		}
		require.EqualValues(t, tc.expectedOutputs, outputs)
		require.EqualValues(t, tc.addStakingAmt, totalTargetAmt.Add(crumb))
		require.Equal(t, tc.expectedCrumb.String(), crumb.String())
	}
}

// TODO: refactor to keeper level
//func TestRebalancing(t *testing.T) {
//	lvs := types.LiquidValidators{
//		{
//			OperatorAddress: "cosmosvaloper10e4vsut6suau8tk9m6dnrm0slgd6npe3jx5xpv",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(100 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper1ld6vlyy24906u3aqp5lj54f3nsg2592nm9nj5c",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(200 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper18hfzxheyknesfgcrttr5dg50ffnfphtwtar9fz",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(300 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper1nmfag3hmkx3qyhpmq7jx5996k8uhgh87xhcqfq",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(400 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//	}
//	moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte("rebalancing")))
//	keeper.Rebalancing(moduleAcc, lvs, sdk.NewDec(10000))
//}
//
//func TestRebalancingWithDelisting(t *testing.T) {
//	lvs := types.LiquidValidators{
//		{
//			OperatorAddress: "cosmosvaloper10e4vsut6suau8tk9m6dnrm0slgd6npe3jx5xpv",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(100 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper1ld6vlyy24906u3aqp5lj54f3nsg2592nm9nj5c",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(200 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper18hfzxheyknesfgcrttr5dg50ffnfphtwtar9fz",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(300 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper180d0fe0w0eqnn04mwhx8h66hnttgqw32fsr6jg",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(0 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper1nmfag3hmkx3qyhpmq7jx5996k8uhgh87xhcqfq",
//			Status:          2,
//			LiquidTokens:    sdk.NewIntFromUint64(400 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//	}
//	moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte("rebalancing")))
//	keeper.Rebalancing(moduleAcc, lvs, sdk.NewDec(10000))
//}
//
//func TestRebalancingUnderThreshold(t *testing.T) {
//	lvs := types.LiquidValidators{
//		{
//			OperatorAddress: "cosmosvaloper10e4vsut6suau8tk9m6dnrm0slgd6npe3jx5xpv",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(100 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper1ld6vlyy24906u3aqp5lj54f3nsg2592nm9nj5c",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(100 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper18hfzxheyknesfgcrttr5dg50ffnfphtwtar9fz",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(100 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper1nmfag3hmkx3qyhpmq7jx5996k8uhgh87xhcqfq",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(101 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//	}
//	moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte("rebalancing")))
//	keeper.Rebalancing(moduleAcc, lvs, sdk.NewDec(1*1000000))
//}
//
//func TestRebalancingDiffWeight(t *testing.T) {
//	lvs := types.LiquidValidators{
//		{
//			OperatorAddress: "cosmosvaloper10e4vsut6suau8tk9m6dnrm0slgd6npe3jx5xpv",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(100 * 1000000),
//			Weight:          sdk.NewInt(20),
//		},
//		{
//			OperatorAddress: "cosmosvaloper1ld6vlyy24906u3aqp5lj54f3nsg2592nm9nj5c",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(200 * 1000000),
//			Weight:          sdk.NewInt(20),
//		},
//		{
//			OperatorAddress: "cosmosvaloper18hfzxheyknesfgcrttr5dg50ffnfphtwtar9fz",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(300 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper1nmfag3hmkx3qyhpmq7jx5996k8uhgh87xhcqfq",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(400 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//	}
//	moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte("rebalancing")))
//	keeper.Rebalancing(moduleAcc, lvs, sdk.NewDec(10000))
//}
//
//func TestRebalancingWithDelistingDiffWeight(t *testing.T) {
//	lvs := types.LiquidValidators{
//		{
//			OperatorAddress: "cosmosvaloper10e4vsut6suau8tk9m6dnrm0slgd6npe3jx5xpv",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(100 * 1000000),
//			Weight:          sdk.NewInt(30),
//		},
//		{
//			OperatorAddress: "cosmosvaloper1ld6vlyy24906u3aqp5lj54f3nsg2592nm9nj5c",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(200 * 1000000),
//			Weight:          sdk.NewInt(20),
//		},
//		{
//			OperatorAddress: "cosmosvaloper18hfzxheyknesfgcrttr5dg50ffnfphtwtar9fz",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(300 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper180d0fe0w0eqnn04mwhx8h66hnttgqw32fsr6jg",
//			Status:          1,
//			LiquidTokens:    sdk.NewIntFromUint64(0 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//		{
//			OperatorAddress: "cosmosvaloper1nmfag3hmkx3qyhpmq7jx5996k8uhgh87xhcqfq",
//			Status:          2,
//			LiquidTokens:    sdk.NewIntFromUint64(400 * 1000000),
//			Weight:          sdk.NewInt(10),
//		},
//	}
//	moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte("rebalancing")))
//	keeper.Rebalancing(moduleAcc, lvs, sdk.NewDec(10000))
//}
//
//
//func TestDivideByCurrentWeight(t *testing.T) {
//	testCases := []struct {
//		whitelistedVals      types.LiquidValidators
//		addStakingAmt   sdk.Int
//		expectedOutputs []sdk.Int
//		expectedCrumb   sdk.Int
//	}{
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2 * 1000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2 * 1000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1 * 1000000),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt:   sdk.NewInt(10 * 1000000),
//			expectedOutputs: []sdk.Int{sdk.NewInt(4 * 1000000), sdk.NewInt(4 * 1000000), sdk.NewInt(2 * 1000000)},
//			expectedCrumb:   sdk.NewInt(0),
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1 * 1000000),
//					Weight:          sdk.NewInt(2),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1 * 1000000),
//					Weight:          sdk.NewInt(2),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1 * 1000000),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt:   sdk.NewInt(10 * 1000000),
//			expectedOutputs: []sdk.Int{sdk.NewInt(3333333), sdk.NewInt(3333333), sdk.NewInt(3333333)},
//			expectedCrumb:   sdk.NewInt(1),
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(3),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt:   sdk.NewInt(10),
//			expectedOutputs: []sdk.Int{sdk.NewInt(4), sdk.NewInt(3), sdk.NewInt(1)},
//			expectedCrumb:   sdk.NewInt(2),
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(10000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(3000001),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt:   sdk.NewInt(10000000),
//			expectedOutputs: []sdk.Int{sdk.NewInt(6666666), sdk.NewInt(1333333), sdk.NewInt(2000000)},
//			expectedCrumb:   sdk.NewInt(1),
//		},
//	}
//
//	for _, tc := range testCases {
//		fmt.Println("------")
//		require.IsType(t, types.LiquidValidators{}, tc.whitelistedVals)
//		require.IsType(t, sdk.Int{}, tc.addStakingAmt)
//		require.IsType(t, sdk.Int{}, tc.expectedCrumb)
//		require.IsType(t, []sdk.Int{}, tc.expectedOutputs)
//
//		totalTargetAmt := sdk.ZeroInt()
//		outputs, crumb := keeper.DivideByCurrentWeight(tc.whitelistedVals, tc.addStakingAmt)
//		for k, v := range outputs {
//			fmt.Println(k, v.String())
//			totalTargetAmt = totalTargetAmt.Add(v)
//		}
//		require.EqualValues(t, tc.expectedOutputs, outputs)
//		require.EqualValues(t, tc.addStakingAmt, totalTargetAmt.Add(crumb))
//		require.Equal(t, tc.expectedCrumb.String(), crumb.String())
//	}
//}
//
//func TestDivideByCurrentWeightDec(t *testing.T) {
//	testCases := []struct {
//		whitelistedVals      types.LiquidValidators
//		addStakingAmt   sdk.Dec
//		expectedOutputs []sdk.Dec
//		expectedCrumb   sdk.Dec
//	}{
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2 * 1000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2 * 1000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1 * 1000000),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt:   sdk.NewDec(10 * 1000000),
//			expectedOutputs: []sdk.Dec{sdk.NewDec(4 * 1000000), sdk.NewDec(4 * 1000000), sdk.NewDec(2 * 1000000)},
//			expectedCrumb:   sdk.NewDec(0),
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1 * 1000000),
//					Weight:          sdk.NewInt(2),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1 * 1000000),
//					Weight:          sdk.NewInt(2),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1 * 1000000),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt:   sdk.NewDec(10 * 1000000),
//			expectedOutputs: []sdk.Dec{sdk.MustNewDecFromStr("3333333.333333333333000000"), sdk.MustNewDecFromStr("3333333.333333333333000000"), sdk.MustNewDecFromStr("3333333.333333333333000000")},
//			expectedCrumb:   sdk.MustNewDecFromStr("0.000000000001000000"),
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(3),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt:   sdk.NewDec(10),
//			expectedOutputs: []sdk.Dec{sdk.MustNewDecFromStr("4.999999999999999998"), sdk.MustNewDecFromStr("3.333333333333333332"), sdk.MustNewDecFromStr("1.666666666666666666")},
//			expectedCrumb:   sdk.MustNewDecFromStr("0.000000000000000004"),
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(10000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(3000001),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt:   sdk.NewDec(10000000),
//			expectedOutputs: []sdk.Dec{sdk.MustNewDecFromStr("6666666.222222251850000000"), sdk.MustNewDecFromStr("1333333.244444450370000000"), sdk.MustNewDecFromStr("2000000.533333297777225185")},
//			expectedCrumb:   sdk.MustNewDecFromStr("0.000000000002774815"),
//		},
//	}
//
//	for _, tc := range testCases {
//		fmt.Println("------")
//		require.IsType(t, types.LiquidValidators{}, tc.whitelistedVals)
//		require.IsType(t, sdk.Dec{}, tc.addStakingAmt)
//		require.IsType(t, sdk.Dec{}, tc.expectedCrumb)
//		require.IsType(t, []sdk.Dec{}, tc.expectedOutputs)
//
//		totalTargetAmt := sdk.ZeroDec()
//		outputs, crumb := keeper.DivideByCurrentWeightDec(tc.whitelistedVals, tc.addStakingAmt)
//		for k, v := range outputs {
//			fmt.Println(k, v.String())
//			totalTargetAmt = totalTargetAmt.Add(v)
//		}
//		require.EqualValues(t, tc.expectedOutputs, outputs)
//		require.EqualValues(t, tc.addStakingAmt, totalTargetAmt.Add(crumb))
//		require.Equal(t, tc.expectedCrumb.String(), crumb.String())
//	}
//}
//
//func TestAddStakingTargetMap(t *testing.T) {
//	testCases := []struct {
//		whitelistedVals    types.LiquidValidators
//		addStakingAmt sdk.Int
//		expectedMap   map[string]sdk.Int
//	}{
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2 * 1000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2 * 1000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1 * 1000000),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt: sdk.NewInt(10 * 1000000),
//			expectedMap: map[string]sdk.Int{
//				"a": sdk.NewInt(3 * 1000000),
//				"b": sdk.NewInt(3 * 1000000),
//				"c": sdk.NewInt(4 * 1000000),
//			},
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(3),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt: sdk.NewInt(10),
//			expectedMap: map[string]sdk.Int{
//				"a": sdk.NewInt(3),
//				"b": sdk.NewInt(3),
//				"c": sdk.NewInt(4),
//			},
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(8),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(7),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(4),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt: sdk.NewInt(10),
//			expectedMap: map[string]sdk.Int{
//				"a": sdk.NewInt(3),
//				"b": sdk.NewInt(2),
//				"c": sdk.NewInt(5),
//			},
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(10),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(5),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt: sdk.NewInt(10),
//			expectedMap: map[string]sdk.Int{
//				"b": sdk.NewInt(3),
//				"c": sdk.NewInt(7),
//			},
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(10),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(3),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(1),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt: sdk.NewInt(10),
//			expectedMap: map[string]sdk.Int{
//				"b": sdk.NewInt(4),
//				"c": sdk.NewInt(6),
//			},
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(10),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(3),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt: sdk.NewInt(10),
//			expectedMap: map[string]sdk.Int{
//				"b": sdk.NewInt(5),
//				"c": sdk.NewInt(5),
//			},
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(10),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(3),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt: sdk.NewInt(10),
//			expectedMap: map[string]sdk.Int{
//				"b": sdk.NewInt(6),
//				"c": sdk.NewInt(4),
//			},
//		},
//		{
//			whitelistedVals: types.LiquidValidators{
//				{
//					OperatorAddress: "a",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(10000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "b",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(2000000),
//					Weight:          sdk.NewInt(1),
//				},
//				{
//					OperatorAddress: "c",
//					Status:          types.ValidatorStatusActive,
//					LiquidTokens:    sdk.NewIntFromUint64(3000001),
//					Weight:          sdk.NewInt(1),
//				},
//			},
//			addStakingAmt: sdk.NewInt(10000000),
//			expectedMap: map[string]sdk.Int{
//				"b": sdk.NewInt(5500001),
//				"c": sdk.NewInt(4499999),
//			},
//		},
//	}
//
//	for _, tc := range testCases {
//		fmt.Println("------")
//		require.IsType(t, types.LiquidValidators{}, tc.whitelistedVals)
//		require.IsType(t, sdk.Int{}, tc.addStakingAmt)
//		require.IsType(t, map[string]sdk.Int{}, tc.expectedMap)
//
//		totalTargetAmt := sdk.ZeroInt()
//		resMap := keeper.AddStakingTargetMap(tc.whitelistedVals, tc.addStakingAmt)
//		for k, v := range resMap {
//			fmt.Println(k, v.String())
//			totalTargetAmt = totalTargetAmt.Add(v)
//		}
//		require.Equal(t, resMap, tc.expectedMap)
//		require.Equal(t, tc.addStakingAmt, totalTargetAmt)
//	}
//}
