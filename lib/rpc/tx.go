package rpc

import (
	"encoding/json"

	"github.com/amolabs/amo-client-go/lib/keys"
	"github.com/amolabs/amo-client-go/lib/types"
)

// Tx broadcast in AMO context

func Transfer(to string, amount string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	ret, err := SignSendTx("transfer", struct {
		To     string `json:"to"`
		Amount string `json:"amount"`
	}{to, amount}, key, fee, lastHeight)
	return ret, err
}

func Stake(validator string, amount string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	ret, err := SignSendTx("stake", struct {
		Validator string `json:"validator"`
		Amount    string `json:"amount"`
	}{validator, amount}, key, fee, lastHeight)
	return ret, err
}

func Withdraw(amount string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	ret, err := SignSendTx("withdraw", struct {
		Amount string `json:"amount"`
	}{amount}, key, fee, lastHeight)
	return ret, err
}

func Delegate(to string, amount string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	ret, err := SignSendTx("delegate", struct {
		To     string `json:"to"`
		Amount string `json:"amount"`
	}{to, amount}, key, fee, lastHeight)
	return ret, err
}

func Retract(amount string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	ret, err := SignSendTx("retract", struct {
		Amount string `json:"amount"`
	}{amount}, key, fee, lastHeight)
	return ret, err
}

func Propose(draftID, config, desc string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	draftIDUint32, err := types.ConvIDFromStr(draftID)
	if err != nil {
		return TmTxResult{}, err
	}
	ret, err := SignSendTx("propose", struct {
		DraftID uint32          `json:"draft_id"`
		Config  json.RawMessage `json:"config"`
		Desc    string          `json:"desc"`
	}{draftIDUint32, []byte(config), desc}, key, fee, lastHeight)
	return ret, err
}

func Vote(draftID string, approve bool, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	draftIDUint32, err := types.ConvIDFromStr(draftID)
	if err != nil {
		return TmTxResult{}, err
	}
	ret, err := SignSendTx("vote", struct {
		DraftID uint32 `json:"draft_id"`
		Approve bool   `json:"approve"`
	}{draftIDUint32, approve}, key, fee, lastHeight)
	return ret, err
}

func Register(target string, custody string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	ret, err := SignSendTx("register", struct {
		Target  string `json:"target"`
		Custody string `json:"custody"`
	}{target, custody}, key, fee, lastHeight)
	return ret, err
}

func Discard(target string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	ret, err := SignSendTx("discard", struct {
		Target string `json:"target"`
	}{target}, key, fee, lastHeight)
	return ret, err
}

func Request(target string, payment string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	ret, err := SignSendTx("request", struct {
		Target  string `json:"target"`
		Payment string `json:"payment"`
	}{target, payment}, key, fee, lastHeight)
	return ret, err
}

func Cancel(target string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	ret, err := SignSendTx("cancel", struct {
		Target string `json:"target"`
	}{target}, key, fee, lastHeight)
	return ret, err
}

func Grant(target string, grantee string, custody string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	ret, err := SignSendTx("grant", struct {
		Target  string `json:"target"`
		Grantee string `json:"grantee"`
		Custody string `json:"custody"`
	}{target, grantee, custody}, key, fee, lastHeight)
	return ret, err
}

func Revoke(target string, grantee string, key keys.KeyEntry, fee, lastHeight string) (TmTxResult, error) {
	ret, err := SignSendTx("revoke", struct {
		Target  string `json:"target"`
		Grantee string `json:"grantee"`
	}{target, grantee}, key, fee, lastHeight)
	return ret, err
}
