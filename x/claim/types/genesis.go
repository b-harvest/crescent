package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Airdrops:     []Airdrop{},
		ClaimRecords: []ClaimRecord{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	for _, a := range gs.Airdrops {
		if err := a.Validate(); err != nil {
			return err
		}
	}

	for _, r := range gs.ClaimRecords {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Validate validates an airdrop object.
func (a Airdrop) Validate() error {
	if _, err := sdk.AccAddressFromBech32(a.SourceAddress); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(a.TerminationAddress); err != nil {
		return err
	}

	if !a.EndTime.After(a.StartTime) {
		return errors.New("end time must be greater than start time")
	}

	for _, c := range a.Conditions {
		switch c {
		case ConditionTypeDeposit, ConditionTypeSwap, ConditionTypeFarming:
		default:
			return fmt.Errorf("unknown action type %T", c)
		}
	}
	return nil
}

// Validate validates claim record object.
func (r ClaimRecord) Validate() error {
	if _, err := sdk.AccAddressFromBech32(r.Recipient); err != nil {
		return err
	}

	if !r.InitialClaimableCoins.IsAllPositive() {
		return fmt.Errorf("initial claimable amount must be positive: %s", r.InitialClaimableCoins.String())
	}

	if !r.ClaimableCoins.IsAllPositive() {
		return fmt.Errorf("claimable amount must be positive: %s", r.InitialClaimableCoins.String())
	}
	return nil
}
