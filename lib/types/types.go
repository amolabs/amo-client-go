package types

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"
)

// These types are borrowed from github.com/amolabs/amoabci/amo/types. The
// amoabci codes are depending on tendermint codes in trun. This file is for
// reducing tendermint dependency. Since this amo-client does not perform
// complex jobs dealing with various subtypes, some elementary types are
// replaced by simple string type. As the client function develops in the
// future the situation may change. In that case, amolabs/amoabci types must be
// sorted out to exclude complex tendermint native types and imported into this
// package.

type Currency struct {
	big.Int
}

func (c *Currency) UnmarshalJSON(b []byte) error {
	if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
		return errors.New(
			"Currency should be represented as double-quoted string(hex:" +
				hex.EncodeToString(b) +
				",str:" +
				string(b) +
				").")
	}
	err := json.Unmarshal(b[1:len(b)-1], &c.Int)
	return err
}

func (c *Currency) String() string {
	oneAMO := new(big.Float)
	oneAMO.SetInt64(1000000000000000000)
	amo := new(big.Float)
	amo.SetInt(&c.Int)
	amo = amo.Quo(amo, oneAMO)
	return fmt.Sprintf("%s mote (%s AMO)", c.Int.String(), amo.String())
}

type PubKeyEd25519 []byte
type Address string

type Stake struct {
	Validator PubKeyEd25519 `json:"validator"`
	Amount    Currency      `json:"amount"`
	Delegates []Delegate    `json:"delegates"`
}

type Delegate struct {
	Delegator Address  `json:"delegator"`
	Delegatee Address  `json:"delegatee"`
	Amount    Currency `json:"amount"`
}

type Parcel struct {
	Owner    Address      `json:"owner"`
	Custody  string       `json:"custody"`
	Info     string       `json:"info,omitempty"`
	Requests []*RequestEx `json:"requests,omitempty"`
	Usages   []*UsageEx   `json:"usages,omitempty"`
}

type Request struct {
	Payment Currency  `json:"payment"`
	Exp     time.Time `json:"exp"`
}

type RequestEx struct {
	*Request
	Buyer Address `json:"buyer"`
}

type Usage struct {
	Custody string    `json:"custody"`
	Exp     time.Time `json:"exp"`
}

type UsageEx struct {
	*Usage
	Buyer Address `json:"buyer"`
}

type IncentiveInfo struct {
	BlockHeight int64    `json:"block_height"`
	Address     Address  `json:"address"`
	Amount      Currency `json:"amount"`
}

type AMOAppConfig struct {
	MaxValidators           uint64   `json:"max_validators"`
	WeightValidator         uint64   `json:"weight_validator"`
	WeightDelegator         uint64   `json:"weight_delegator"`
	MinStakingUnit          Currency `json:"min_staking_unit"`
	BlkReward               Currency `json:"blk_reward"`
	TxReward                Currency `json:"tx_reward"`
	PenaltyRatioM           float64  `json:"penalty_ratio_m"` // malicious validator
	PenaltyRatioL           float64  `json:"penalty_ratio_l"` // lazy validators
	LazinessCounterWindow   int64    `json:"laziness_counter_window"`
	LazinessThreshold       float64  `json:"laziness_threshold"`
	BlockBoundTxGracePeriod uint64   `json:"block_bound_tx_grace_period"`
	LockupPeriod            uint64   `json:"lockup_period"`
	DraftOpenCount          uint64   `json:"draft_open_count"`
	DraftCloseCount         uint64   `json:"draft_close_count"`
	DraftApplyCount         uint64   `json:"draft_apply_count"`
	DraftDeposit            Currency `json:"draft_deposit"`
	DraftQuorumRate         float64  `json:"draft_quorum_rate"`
	DraftPassRate           float64  `json:"draft_pass_rate"`
	DraftRefundRate         float64  `json:"draft_refund_rate"`
}
