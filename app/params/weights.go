package params

// Default simulation operation weights for messages and gov proposals.
const (
	DefaultWeightMsgCreateFixedAmountPlan int = 0
	DefaultWeightMsgCreateRatioPlan       int = 0
	DefaultWeightMsgStake                 int = 0
	DefaultWeightMsgUnstake               int = 0
	DefaultWeightMsgHarvest               int = 0
	DefaultWeightMsgRemovePlan            int = 0

	DefaultWeightMsgCreatePair       int = 10
	DefaultWeightMsgCreatePool       int = 10
	DefaultWeightMsgCreateRangedPool int = 15
	DefaultWeightMsgDeposit          int = 20
	DefaultWeightMsgWithdraw         int = 20
	DefaultWeightMsgLimitOrder       int = 80
	DefaultWeightMsgMarketOrder      int = 60
	DefaultWeightMsgMMOrder          int = 40
	DefaultWeightMsgCancelOrder      int = 20
	DefaultWeightMsgCancelAllOrders  int = 20
	DefaultWeightMsgCancelMMOrder    int = 20

	DefaultWeightAddPublicPlanProposal    int = 5
	DefaultWeightUpdatePublicPlanProposal int = 5
	DefaultWeightDeletePublicPlanProposal int = 5

	DefaultWeightMsgLiquidStake   int = 80
	DefaultWeightMsgLiquidUnstake int = 30

	DefaultWeightAddWhitelistValidatorsProposal    int = 50
	DefaultWeightUpdateWhitelistValidatorsProposal int = 5
	DefaultWeightDeleteWhitelistValidatorsProposal int = 5
	DefaultWeightCompleteRedelegationUnbonding     int = 30
	DefaultWeightTallyWithLiquidStaking            int = 30

	DefaultWeightMsgClaim int = 0

	DefaultWeightMsgLiquidFarm   int = 50
	DefaultWeightMsgLiquidUnfarm int = 10
	DefaultWeightMsgPlaceBid     int = 20
	DefaultWeightMsgRefundBid    int = 5

	DefaultWeightMsgApplyMarketMaker int = 20
	DefaultWeightMsgClaimIncentives  int = 10

	DefaultWeightMarketMakerProposal  int = 20
	DefaultWeightChangeIncentivePairs int = 5
	DefaultWeightChangeDepositAmount  int = 2
)
