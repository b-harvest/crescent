package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/crescent-network/crescent/v4/x/antehandlers"
)

const (
	ProposalTypeBootstrap string = "Bootstrap"
)

// Implements Proposal Interface
var _ gov.Content = &BootstrapProposal{}
var _ antehandlers.ProposalExtended = &BootstrapProposal{}

func init() {
	gov.RegisterProposalType(ProposalTypeBootstrap)
	gov.RegisterProposalTypeCodec(&BootstrapProposal{}, "crescent/BootstrapProposal")
}

// NewBootstrapProposal creates a new BootstrapProposal object.
func NewBootstrapProposal(
	title string,
	description string,
	proposerAddress string,
	offerCoins sdk.Coins,
	baseCoinDenom string,
	quoteCoinDenom string,
	minPrice sdk.Dec,
	maxPrice sdk.Dec,
	pairId uint64,
	initialOrders []InitialOrder,
	startTime time.Time,
	stageDuration time.Duration,
	numOfStages uint32,
) *BootstrapProposal {
	return &BootstrapProposal{
		Title:           title,
		Description:     description,
		ProposerAddress: proposerAddress,
		OfferCoins:      offerCoins,
		BaseCoinDenom:   baseCoinDenom,
		QuoteCoinDenom:  quoteCoinDenom,
		MinPrice:        &minPrice,
		MaxPrice:        &maxPrice,
		PairId:          pairId,
		InitialOrders:   initialOrders,
		StartTime:       startTime,
		StageDuration:   stageDuration,
		NumOfStages:     numOfStages,
	}
}

func (p *BootstrapProposal) GetTitle() string { return p.Title }

func (p *BootstrapProposal) GetDescription() string { return p.Description }

func (p *BootstrapProposal) GetProposerAddress() string { return p.ProposerAddress }

func (p *BootstrapProposal) GetProposer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(p.ProposerAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

func (p *BootstrapProposal) ProposalRoute() string { return RouterKey }

func (p *BootstrapProposal) ProposalType() string { return ProposalTypeBootstrap }

func (p *BootstrapProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(p.ProposerAddress)
	if err != nil {
		return err
	}

	if err = p.OfferCoins.Validate(); err != nil {
		return err
	}

	if err = sdk.ValidateDenom(p.QuoteCoinDenom); err != nil {
		return err
	}

	if p.MinPrice != nil && !p.MinPrice.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "min price should be positive")
	}

	if p.MaxPrice != nil && !p.MaxPrice.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "max price should be positive")
	}

	if p.MinPrice != nil && p.MaxPrice != nil && p.MaxPrice.LTE(*p.MinPrice) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "max price should be greater than min price")
	}

	if p.PairId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pair id")
	}

	// validate initial orders, ascending, price
	lastPrice := sdk.ZeroDec()
	sumOfOfferCoin := sdk.NewCoins()
	for _, io := range p.InitialOrders {
		if err = io.OfferCoin.Validate(); err != nil {
			return err
		}
		if p.MinPrice != nil && io.Price.LT(*p.MinPrice) || p.MaxPrice != nil && io.Price.GT(*p.MaxPrice) {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "initial orders price should be between min price and max price")
		}
		if lastPrice.GT(io.Price) {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "initial orders should be sorted ascending order")
		}
		lastPrice = io.Price
		sumOfOfferCoin = sumOfOfferCoin.Add(io.OfferCoin)
	}

	// validate sum of order amount must be equal to the offer coin amount
	if !p.OfferCoins.IsEqual(sumOfOfferCoin) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "sum of order amount must be equal to the offer coin amount")
	}

	// TODO: StartTime, NumOfStages, StageDuration
	return gov.ValidateAbstract(p)
}

// TODO: check with testcode
func (p BootstrapProposal) String() string {
	return fmt.Sprintf(`Bootstrap Proposal:
  Title:         %s
  Description:   %s
  ProposerAddress: %v
  OfferCoins:     %v
  QuoteCoinDenom:%s
  MinPrice:      %v
  MaxPrice:      %v
  PairId:        %v
  InitialOrders: %v
`, p.Title, p.Description, p.ProposerAddress, p.OfferCoins, p.QuoteCoinDenom, p.MinPrice, p.MaxPrice, p.PairId,
		p.InitialOrders)
}
