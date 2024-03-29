package types

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
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

func (c Currency) MarshalJSON() ([]byte, error) {
	return []byte("\"" + c.Text(10) + "\""), nil
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

const PubKeyP256Size = 65

type (
	PubKeyEd25519 []byte
	PubKeyP256    [PubKeyP256Size]byte
	Address       string
)

type UDC struct {
	Owner     Address   `json:"owner"`     // required
	Desc      string    `json:"desc"`      // optional
	Operators []Address `json:"operators"` // optional
	Total     Currency  `json:"total"`     // required
}

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

type Draft struct {
	Proposer Address      `json:"proposer"`
	Config   AMOAppConfig `json:"config"`
	Desc     string       `json:"desc"`

	OpenCount  uint64   `json:"open_count"`
	CloseCount uint64   `json:"close_count"`
	ApplyCount uint64   `json:"apply_count"`
	Deposit    Currency `json:"deposit"`

	TallyQuorum  Currency `json:"tally_quorum"`
	TallyApprove Currency `json:"tally_approve"`
	TallyReject  Currency `json:"tally_reject"`
}

type DraftEx struct {
	*Draft
	Votes []*VoteInfo `json:"votes"`
}

type Vote struct {
	Approve bool `json:"approve"`
}

type VoteInfo struct {
	Voter Address `json:"voter"`
	*Vote
}

type Storage struct {
	Owner           Address  `json:"owner"`
	Url             string   `json:"url"`
	RegistrationFee Currency `json:"registration_fee"`
	HostingFee      Currency `json:"hosting_fee"`
	Active          bool     `json:"active"`
}

type Extra struct {
	Register json.RawMessage `json:"register,omitempty"`
	Request  json.RawMessage `json:"request,omitempty"`
	Grant    json.RawMessage `json:"grant,omitempty"`
}

type Parcel struct {
	Owner        Address `json:"owner"`
	Custody      string  `json:"custody"`
	ProxyAccount Address `json:"proxy_account,omitempty"`
	Extra        Extra   `json:"extra,omitempty"`
}

type ParcelEx struct {
	*Parcel
	Requests []*RequestEx `json:"requests,omitempty"`
	Usages   []*UsageEx   `json:"usages,omitempty"`
}

type Request struct {
	Payment   Currency `json:"payment"`
	Agency    Address  `json:"agency,omitempty"`
	Dealer    Address  `json:"dealer,omitempty"`
	DealerFee Currency `json:"dealer_fee,omitempty"`
	Extra     Extra    `json:"extra,omitempty"`
}

type RequestEx struct {
	*Request
	Recipient Address `json:"recipient"`
}

type Usage struct {
	Custody string `json:"custody"`
	Extra   Extra  `json:"extra,omitempty"`
}

type UsageEx struct {
	*Usage
	Recipient Address `json:"recipient"`
}

type DIDEntry struct {
	Document json.RawMessage `json:"document"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

type AMOAppConfig struct {
	MaxValidators          uint64   `json:"max_validators"`
	WeightValidator        float64  `json:"weight_validator"`
	WeightDelegator        float64  `json:"weight_delegator"`
	MinStakingUnit         Currency `json:"min_staking_unit"`
	BlkReward              Currency `json:"blk_reward"`
	TxReward               Currency `json:"tx_reward"`
	PenaltyRatioM          float64  `json:"penalty_ratio_m"` // malicious validator
	PenaltyRatioL          float64  `json:"penalty_ratio_l"` // lazy validators
	LazinessCounterWindow  int64    `json:"laziness_counter_window"`
	LazinessThreshold      float64  `json:"laziness_threshold"`
	BlockBindingWindow     uint64   `json:"block_binding_window"`
	LockupPeriod           uint64   `json:"lockup_period"`
	DraftOpenCount         uint64   `json:"draft_open_count"`
	DraftCloseCount        uint64   `json:"draft_close_count"`
	DraftApplyCount        uint64   `json:"draft_apply_count"`
	DraftDeposit           Currency `json:"draft_deposit"`
	DraftQuorumRate        float64  `json:"draft_quorum_rate"`
	DraftPassRate          float64  `json:"draft_pass_rate"`
	DraftRefundRate        float64  `json:"draft_refund_rate"`
	UpgradeProtocolHeight  int64    `json:"upgrade_protocol_height"`
	UpgradeProtocolVersion uint64   `json:"upgrade_protocol_version"`
}

func ConvIDFromStr(IDStr string) (uint32, error) {
	tmp, err := strconv.ParseInt(IDStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(tmp), nil
}

func (pubKey *PubKeyP256) UnmarshalJSON(data []byte) error {
	if len(data) != PubKeyP256Size*2+2 {
		return errors.New("Invalid public key format")
	}
	_, err := hex.Decode(pubKey[:], data[1:len(data)-1])
	if err != nil {
		return err
	}
	return nil
}
