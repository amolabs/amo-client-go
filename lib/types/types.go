package types

import (
	"time"
	//"math/big"
)

// These types are borrowed from github.com/amolabs/amoabci/amo/types. The
// amoabci codes are depending on tendermint codes in trun. This file is for
// reducing tendermint dependency. Since this amo-client does not perform
// complex jobs dealing with various subtypes, some elementary types are
// replaced by simple string type. As the client function develops in the
// future the situation may change. In that case, amolabs/amoabci types must be
// sorted out to exclude complex tendermint native types and imported into this
// package.

type Currency string
type PubKeyEd25519 []byte
type Address string

type Stake struct {
	Amount    Currency      `json:"amount"`
	Validator PubKeyEd25519 `json:"validator"`
}

type Delegate struct {
	Holder    Address
	Amount    Currency `json:"amount"`
	Delegator Address  `json:"delegator"`
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
