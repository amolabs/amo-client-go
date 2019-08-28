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
	Amount    Currency      `json:"amount"`
	Validator PubKeyEd25519 `json:"validator"`
}

type Delegate struct {
	Holder    Address
	Delegatee Address  `json:"delegatee"`
	Amount    Currency `json:"amount"`
}

type Parcel struct {
	Owner   Address `json:"owner"`
	Custody string  `json:"custody"`
	Info    string  `json:"info,omitempty"`
}

type Request struct {
	Payment Currency  `json:"payment"`
	Exp     time.Time `json:"exp"`
}

type Usage struct {
	Custody string    `json:"custody"`
	Exp     time.Time `json:"exp"`
}
