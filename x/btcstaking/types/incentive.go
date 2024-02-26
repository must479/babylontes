package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewVotingPowerDistCache() *VotingPowerDistCache {
	return &VotingPowerDistCache{
		TotalVotingPower:  0,
		FinalityProviders: []*FinalityProviderDistInfo{},
	}
}

func (dc *VotingPowerDistCache) Empty() bool {
	return len(dc.FinalityProviders) == 0
}

func (dc *VotingPowerDistCache) AddFinalityProviderDistInfo(v *FinalityProviderDistInfo) {
	dc.FinalityProviders = append(dc.FinalityProviders, v)
}

// ApplyActiveFinalityProviders sorts all finality providers, counts the total voting
// power of top N finality providers, and records them in cache
func (dc *VotingPowerDistCache) ApplyActiveFinalityProviders(maxActiveFPs uint32) {
	// reset total voting power
	dc.TotalVotingPower = 0
	// sort finality providers
	SortFinalityProviders(dc.FinalityProviders)
	// calculate voting power of top N finality providers
	numActiveFPs := dc.GetNumActiveFPs(maxActiveFPs)
	for i := uint32(0); i < numActiveFPs; i++ {
		dc.TotalVotingPower += dc.FinalityProviders[i].TotalVotingPower
	}
}

func (dc *VotingPowerDistCache) GetNumActiveFPs(maxActiveFPs uint32) uint32 {
	return min(maxActiveFPs, uint32(len(dc.FinalityProviders)))
}

// GetActiveFinalityProviders returns the list of active finality providers
// i.e., top N of them in terms of voting power
func (dc *VotingPowerDistCache) GetActiveFinalityProviders(maxActiveFPs uint32) []*FinalityProviderDistInfo {
	numActiveFPs := dc.GetNumActiveFPs(maxActiveFPs)
	return dc.FinalityProviders[:numActiveFPs]
}

// FilterVotedDistCache filters out a voting power distribution cache
// with finality providers that have voted according to a map of given
// voters, and their total voted power.
func (dc *VotingPowerDistCache) FilterVotedDistCache(maxActiveFPs uint32, voterBTCPKs map[string]struct{}) *VotingPowerDistCache {
	activeFPs := dc.GetActiveFinalityProviders(maxActiveFPs)
	filteredFps := []*FinalityProviderDistInfo{}
	totalVotingPower := uint64(0)
	for _, v := range activeFPs {
		if _, ok := voterBTCPKs[v.BtcPk.MarshalHex()]; ok {
			filteredFps = append(filteredFps, v)
			totalVotingPower += v.TotalVotingPower
		}
	}

	return &VotingPowerDistCache{
		FinalityProviders: filteredFps,
		TotalVotingPower:  totalVotingPower,
	}
}

// GetFinalityProviderPortion returns the portion of a finality provider's voting power out of the total voting power
func (dc *VotingPowerDistCache) GetFinalityProviderPortion(v *FinalityProviderDistInfo) sdkmath.LegacyDec {
	return sdkmath.LegacyNewDec(int64(v.TotalVotingPower)).QuoTruncate(sdkmath.LegacyNewDec(int64(dc.TotalVotingPower)))
}

func NewFinalityProviderDistInfo(fp *FinalityProvider) *FinalityProviderDistInfo {
	return &FinalityProviderDistInfo{
		BtcPk:            fp.BtcPk,
		BabylonPk:        fp.BabylonPk,
		Commission:       fp.Commission,
		TotalVotingPower: 0,
		BtcDels:          []*BTCDelDistInfo{},
	}
}

func (v *FinalityProviderDistInfo) GetAddress() sdk.AccAddress {
	return sdk.AccAddress(v.BabylonPk.Address())
}

func (v *FinalityProviderDistInfo) AddBTCDel(btcDel *BTCDelegation) {
	btcDelDistInfo := &BTCDelDistInfo{
		BtcPk:         btcDel.BtcPk,
		BabylonPk:     btcDel.BabylonPk,
		StakingTxHash: btcDel.MustGetStakingTxHash().String(),
		VotingPower:   btcDel.TotalSat,
	}
	v.BtcDels = append(v.BtcDels, btcDelDistInfo)
	v.TotalVotingPower += btcDelDistInfo.VotingPower
}

func (v *FinalityProviderDistInfo) AddBTCDelDistInfo(d *BTCDelDistInfo) {
	v.BtcDels = append(v.BtcDels, d)
	v.TotalVotingPower += d.VotingPower
}

// GetBTCDelPortion returns the portion of a BTC delegation's voting power out of
// the finality provider's total voting power
func (v *FinalityProviderDistInfo) GetBTCDelPortion(d *BTCDelDistInfo) sdkmath.LegacyDec {
	return sdkmath.LegacyNewDec(int64(d.VotingPower)).QuoTruncate(sdkmath.LegacyNewDec(int64(v.TotalVotingPower)))
}

func (d *BTCDelDistInfo) GetAddress() sdk.AccAddress {
	return sdk.AccAddress(d.BabylonPk.Address())
}
