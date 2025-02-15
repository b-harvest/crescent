package ante_test

import (
	"fmt"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	claimtypes "github.com/crescent-network/crescent/v5/x/claim/types"
)

type runTxMode uint8

const (
	runTxModeCheck    runTxMode = iota // Check a transaction
	runTxModeReCheck                   // Recheck a (pending) transaction after a commit
	runTxModeSimulate                  // Simulate a transaction
	runTxModeDeliver                   // Deliver a transaction
)

func (suite *AnteTestSuite) TestAnteHandlerSubmitProposalMsg() {
	suite.SetupTest(false) // reset

	// Same data for every test cases
	accounts := suite.CreateTestAccounts(2)
	feeAmount := testdata.NewTestFeeAmount()
	gasLimit := testdata.NewTestGasLimit()
	acc := accounts[0].acc.GetAddress()
	//accStr := acc.String()

	// Variable data per test case
	var (
		accNums []uint64
		msgs    []sdk.Msg
		privs   []cryptotypes.PrivKey
		accSeqs []uint64
	)

	content := govtypes.ContentFromProposalType("title", "description", govtypes.ProposalTypeText)
	suite.Require().NotNil(content)
	depositParams := govtypes.DefaultDepositParams()
	depositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin("ucre", sdk.NewInt(7500000000)))
	suite.app.GovKeeper.SetDepositParams(suite.ctx, depositParams)

	testCases := []TestCase{
		{
			"gov msg with min deposit",
			func() {
				msg, err := govtypes.NewMsgSubmitProposal(
					content,
					depositParams.MinDeposit,
					acc,
				)
				suite.Require().NoError(err)

				msgs = []sdk.Msg{msg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeCheck,
			true,
			nil,
		},
		{
			"gov msg with min deposit * minInitialDepositFraction 50%",
			func() {
				msg, err := govtypes.NewMsgSubmitProposal(
					content,
					sdk.NewCoins(sdk.NewCoin("ucre", sdk.NewInt(7500000000/2))),
					acc,
				)
				suite.Require().NoError(err)

				msgs = []sdk.Msg{msg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeCheck,
			true,
			nil,
		},
		{
			"gov msg with less then min deposit * minInitialDepositFraction 50%",
			func() {
				msg, err := govtypes.NewMsgSubmitProposal(
					content,
					sdk.NewCoins(sdk.NewCoin("ucre", sdk.NewInt(7500000000/2-1))),
					acc,
				)
				suite.Require().NoError(err)

				msgs = []sdk.Msg{msg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeCheck,
			false,
			sdkerrors.ErrInsufficientFunds,
		},
		{
			"gov msg with less then min deposit * minInitialDepositFraction 50% - authz nested",
			func() {
				msg, err := govtypes.NewMsgSubmitProposal(
					content,
					sdk.NewCoins(sdk.NewCoin("ucre", sdk.NewInt(7500000000/2-1))),
					acc,
				)
				suite.Require().NoError(err)

				authzMsg := authz.NewMsgExec(acc, []sdk.Msg{msg})
				msgs = []sdk.Msg{&authzMsg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeCheck,
			false,
			sdkerrors.ErrInsufficientFunds,
		},
		{
			"gov msg with - multi",
			func() {
				msg, err := govtypes.NewMsgSubmitProposal(
					content,
					sdk.NewCoins(sdk.NewCoin("ucre", sdk.NewInt(7500000000/2))),
					acc,
				)
				suite.Require().NoError(err)
				msg2, err := govtypes.NewMsgSubmitProposal(
					content,
					sdk.NewCoins(sdk.NewCoin("ucre", sdk.NewInt(1))),
					acc,
				)
				suite.Require().NoError(err)

				suite.Require().NoError(err)

				msgs = []sdk.Msg{msg, msg2}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeCheck,
			false,
			sdkerrors.ErrInsufficientFunds,
		},
		{
			"gov msg with less then min deposit * minInitialDepositFraction 50% - pass on deliver tx",
			func() {
				msg, err := govtypes.NewMsgSubmitProposal(
					content,
					sdk.NewCoins(sdk.NewCoin("ucre", sdk.NewInt(7500000000/2-1))),
					acc,
				)
				suite.Require().NoError(err)

				msgs = []sdk.Msg{msg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeDeliver,
			true,
			nil,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()
			tc.malleate()

			suite.RunTestCase(privs, msgs, feeAmount, gasLimit, accNums, accSeqs, suite.ctx.ChainID(), tc)
		})
	}
}

func (suite *AnteTestSuite) TestAnteHandlerDeprecatedMsg() {
	suite.SetupTest(false) // reset

	// Same data for every test cases
	accounts := suite.CreateTestAccounts(2)
	feeAmount := testdata.NewTestFeeAmount()
	gasLimit := testdata.NewTestGasLimit()
	acc := accounts[0].acc.GetAddress()
	accStr := acc.String()

	// Variable data per test case
	var (
		accNums []uint64
		msgs    []sdk.Msg
		privs   []cryptotypes.PrivKey
		accSeqs []uint64
	)

	testCases := []TestCase{
		{
			"good tx from one signer",
			func() {
				msg := testdata.NewTestMsg(acc)
				msgs = []sdk.Msg{msg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeDeliver,
			true,
			nil,
		},
		{
			"deprecated msg",
			func() {
				msg := &claimtypes.MsgClaim{
					AirdropId:     0,
					Recipient:     accStr,
					ConditionType: 0,
				}
				msgs = []sdk.Msg{msg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeDeliver,
			false,
			fmt.Errorf("/crescent.claim.v1beta1.MsgClaim is deprecated msg type"),
		},
		{
			"deprecated msg - authz nested",
			func() {
				msg := &claimtypes.MsgClaim{
					AirdropId:     0,
					Recipient:     accStr,
					ConditionType: 0,
				}

				authzMsg := authz.NewMsgExec(acc, []sdk.Msg{msg})
				msgs = []sdk.Msg{&authzMsg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeDeliver,
			false,
			fmt.Errorf("/crescent.claim.v1beta1.MsgClaim is deprecated msg type"),
		},
		{
			"deprecated msg - checkTx",
			func() {
				msg := &claimtypes.MsgClaim{
					AirdropId:     0,
					Recipient:     accStr,
					ConditionType: 0,
				}
				msgs = []sdk.Msg{msg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeCheck,
			false,
			fmt.Errorf("/crescent.claim.v1beta1.MsgClaim is deprecated msg type"),
		},
		{
			"deprecated msg - sim",
			func() {
				msg := &claimtypes.MsgClaim{
					AirdropId:     0,
					Recipient:     accStr,
					ConditionType: 0,
				}
				msgs = []sdk.Msg{msg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeSimulate,
			false,
			fmt.Errorf("/crescent.claim.v1beta1.MsgClaim is deprecated msg type"),
		},
		{
			"deprecated msg - recheck",
			func() {
				msg := &claimtypes.MsgClaim{
					AirdropId:     0,
					Recipient:     accStr,
					ConditionType: 0,
				}
				msgs = []sdk.Msg{msg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeReCheck,
			false,
			fmt.Errorf("/crescent.claim.v1beta1.MsgClaim is deprecated msg type"),
		},
		{
			"deprecated msg - multi",
			func() {
				msg := &claimtypes.MsgClaim{
					AirdropId:     0,
					Recipient:     accStr,
					ConditionType: 0,
				}
				msgs = []sdk.Msg{msg, msg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeDeliver,
			false,
			fmt.Errorf("/crescent.claim.v1beta1.MsgClaim is deprecated msg type"),
		},
		{
			"deprecated msg - multi with normal msg",
			func() {
				msgNormal := testdata.NewTestMsg(acc)
				msg := &claimtypes.MsgClaim{
					AirdropId:     0,
					Recipient:     accStr,
					ConditionType: 0,
				}
				msgs = []sdk.Msg{msgNormal, msg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeDeliver,
			false,
			fmt.Errorf("/crescent.claim.v1beta1.MsgClaim is deprecated msg type"),
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()
			tc.malleate()

			suite.RunTestCase(privs, msgs, feeAmount, gasLimit, accNums, accSeqs, suite.ctx.ChainID(), tc)
		})
	}
}

func (suite *AnteTestSuite) TestDoubleNestedAuthzMsg() {
	suite.SetupTest(false) // reset

	// Same data for every test cases
	accounts := suite.CreateTestAccounts(2)
	feeAmount := testdata.NewTestFeeAmount()
	gasLimit := testdata.NewTestGasLimit()
	acc := accounts[0].acc.GetAddress()
	//accStr := acc.String()

	// Variable data per test case
	var (
		accNums []uint64
		msgs    []sdk.Msg
		privs   []cryptotypes.PrivKey
		accSeqs []uint64
	)

	msg := testdata.NewTestMsg(acc)
	authzMsg := authz.NewMsgExec(acc, []sdk.Msg{msg})
	authzMsgNested := authz.NewMsgExec(acc, []sdk.Msg{&authzMsg})

	testCases := []TestCase{
		{
			"normal authz msg",
			func() {

				msgs = []sdk.Msg{&authzMsg}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeDeliver,
			true,
			nil,
		},
		{
			"double nested authz msg",
			func() {

				msgs = []sdk.Msg{&authzMsgNested}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeDeliver,
			false,
			fmt.Errorf("double nested /cosmos.authz.v1beta1.MsgExec is not allowed"),
		},
		{
			"double nested /cosmos.authz.v1beta1.MsgExec msg - check tx",
			func() {

				msgs = []sdk.Msg{&authzMsgNested}

				privs, accNums, accSeqs = []cryptotypes.PrivKey{accounts[0].priv}, []uint64{0}, []uint64{0}
			},
			runTxModeCheck,
			false,
			fmt.Errorf("double nested /cosmos.authz.v1beta1.MsgExec is not allowed"),
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()
			tc.malleate()

			suite.RunTestCase(privs, msgs, feeAmount, gasLimit, accNums, accSeqs, suite.ctx.ChainID(), tc)
		})
	}
}
