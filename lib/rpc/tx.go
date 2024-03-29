package rpc

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/amolabs/amo-client-go/lib/keys"
	"github.com/amolabs/amo-client-go/lib/types"
)

// Tx broadcast in AMO context

func Transfer(udc uint32, to, asset string, isParcel bool, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	to = toUpper(to)
	amount := ""
	parcel := ""
	if isParcel {
		parcel = asset
	} else {
		amount = asset
	}
	return SignSendTx("transfer", struct {
		To     string `json:"to"`
		UDC    uint32 `json:"udc,omitempty"`
		Amount string `json:"amount,omitempty"`
		Parcel string `json:"parcel,omitempty"`
	}{to, udc, amount, parcel}, key, fee, lastHeight)
}

func Issue(udcID, amount string, desc string, operators []string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	udcIDUint32, err := types.ConvIDFromStr(udcID)
	if err != nil {
		return TmTxResult{}, err
	}
	for i, op := range operators {
		operators[i] = toUpper(op)
	}
	return SignSendTx("issue", struct {
		UDC       uint32   `json:"udc"`
		Desc      string   `json:"desc,omitempty"`
		Operators []string `json:"operators,omitempty"`
		Amount    string   `json:"amount"`
	}{udcIDUint32, desc, operators, amount}, key, fee, lastHeight)
}

func Burn(udcID, amount string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	udcIDUint32, err := types.ConvIDFromStr(udcID)
	if err != nil {
		return TmTxResult{}, err
	}
	return SignSendTx("burn", struct {
		UDC    uint32 `json:"udc"`
		Amount string `json:"amount"`
	}{udcIDUint32, amount}, key, fee, lastHeight)
}

func Lock(udcID, holder, amount string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	udcIDUint32, err := types.ConvIDFromStr(udcID)
	if err != nil {
		return TmTxResult{}, err
	}
	holder = toUpper(holder)
	return SignSendTx("lock", struct {
		UDC    uint32 `json:"udc"`
		Holder string `json:"holder"`
		Amount string `json:"amount"`
	}{udcIDUint32, holder, amount}, key, fee, lastHeight)
}

func Stake(validator, amount string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	validator = toUpper(validator)
	return SignSendTx("stake", struct {
		Validator string `json:"validator"`
		Amount    string `json:"amount"`
	}{validator, amount}, key, fee, lastHeight)
}

func Withdraw(amount string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	return SignSendTx("withdraw", struct {
		Amount string `json:"amount"`
	}{amount}, key, fee, lastHeight)
}

func Delegate(to, amount string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	to = toUpper(to)
	return SignSendTx("delegate", struct {
		To     string `json:"to"`
		Amount string `json:"amount"`
	}{to, amount}, key, fee, lastHeight)
}

func Retract(amount string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	return SignSendTx("retract", struct {
		Amount string `json:"amount"`
	}{amount}, key, fee, lastHeight)
}

func Propose(draftID, config, desc string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	draftIDUint32, err := types.ConvIDFromStr(draftID)
	if err != nil {
		return TmTxResult{}, err
	}
	return SignSendTx("propose", struct {
		DraftID uint32          `json:"draft_id"`
		Config  json.RawMessage `json:"config"`
		Desc    string          `json:"desc"`
	}{draftIDUint32, []byte(config), desc}, key, fee, lastHeight)
}

func Vote(draftID string, approve bool, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	draftIDUint32, err := types.ConvIDFromStr(draftID)
	if err != nil {
		return TmTxResult{}, err
	}
	return SignSendTx("vote", struct {
		DraftID uint32 `json:"draft_id"`
		Approve bool   `json:"approve"`
	}{draftIDUint32, approve}, key, fee, lastHeight)
}

func Setup(storageID, url, regFee, hostFee string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	storageIDUint32, err := types.ConvIDFromStr(storageID)
	if err != nil {
		return TmTxResult{}, err
	}
	return SignSendTx("setup", struct {
		Storage         uint32 `json:"storage"`
		Url             string `json:"url"`
		RegistrationFee string `json:"registration_fee"`
		HostingFee      string `json:"hosting_fee"`
	}{storageIDUint32, url, regFee, hostFee}, key, fee, lastHeight)
}

func Close(storageID string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	storageIDUint32, err := types.ConvIDFromStr(storageID)
	if err != nil {
		return TmTxResult{}, err
	}
	return SignSendTx("close", struct {
		Storage uint32 `json:"storage"`
	}{storageIDUint32}, key, fee, lastHeight)
}

func Register(target, custody, proxy, extra string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	target = toUpper(target)
	custody = toUpper(custody)
	proxy = toUpper(proxy)
	return SignSendTx("register", struct {
		Target       string          `json:"target"`
		Custody      string          `json:"custody"`
		ProxyAccount string          `json:"proxy_account,omitempty"`
		Extra        json.RawMessage `json:"extra,omitempty"`
	}{target, custody, proxy, []byte(extra)}, key, fee, lastHeight)
}

func Discard(target string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	target = toUpper(target)
	return SignSendTx("discard", struct {
		Target string `json:"target"`
	}{target}, key, fee, lastHeight)
}

func Request(target, payment, recipient, dealer, dealerFee, extra string,
	key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	if dealer != "" && dealerFee == "" {
		return TmTxResult{}, errors.New("'dealer_fee' is missing")
	}
	if dealer == "" && dealerFee != "" {
		return TmTxResult{}, errors.New("'dealer' is missing")
	}

	target = toUpper(target)
	recipient = toUpper(recipient)
	dealer = toUpper(dealer)

	return SignSendTx("request", struct {
		Target    string          `json:"target"`
		Payment   string          `json:"payment"`
		Recipient string          `json:"recipient,omitempty"`
		Dealer    string          `json:"dealer,omitempty"`
		DealerFee string          `json:"dealer_fee,omitempty"`
		Extra     json.RawMessage `json:"extra,omitempty"`
	}{target, payment, recipient, dealer, dealerFee, []byte(extra)}, key, fee, lastHeight)
}

func Cancel(target, recipient string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	target = toUpper(target)
	recipient = toUpper(recipient)
	return SignSendTx("cancel", struct {
		Target    string `json:"target"`
		Recipient string `json:"recipient,omitempty"`
	}{target, recipient}, key, fee, lastHeight)
}

func Grant(target, recipient, custody, extra string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	target = toUpper(target)
	recipient = toUpper(recipient)
	custody = toUpper(custody)
	return SignSendTx("grant", struct {
		Target    string          `json:"target"`
		Recipient string          `json:"recipient"`
		Custody   string          `json:"custody"`
		Extra     json.RawMessage `json:"extra,omitempty"`
	}{target, recipient, custody, []byte(extra)}, key, fee, lastHeight)
}

func Revoke(target, recipient string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	target = toUpper(target)
	recipient = toUpper(recipient)
	return SignSendTx("revoke", struct {
		Target    string `json:"target"`
		Recipient string `json:"recipient"`
	}{target, recipient}, key, fee, lastHeight)
}

func DIDClaim(target, docFile string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	b, err := ioutil.ReadFile(docFile)
	if err != nil {
		return TmTxResult{}, err
	}

	return SignSendTx("did.claim", struct {
		Target   string          `json:"target"`
		Document json.RawMessage `json:"document"`
	}{target, b}, key, fee, lastHeight)
}

func DIDDismiss(target string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	return SignSendTx("did.dismiss", struct {
		Target string `json:"target"`
	}{target}, key, fee, lastHeight)
}

func DIDIssue(target, docFile string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	b, err := ioutil.ReadFile(docFile)
	if err != nil {
		return TmTxResult{}, err
	}

	return SignSendTx("did.issue", struct {
		Target     string          `json:"target"`
		Credential json.RawMessage `json:"credential"`
	}{target, b}, key, fee, lastHeight)
}

func DIDRevoke(target string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	return SignSendTx("did.revoke", struct {
		Target string `json:"target"`
	}{target}, key, fee, lastHeight)
}
