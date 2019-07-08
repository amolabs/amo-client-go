package rpc

import (
	"github.com/amolabs/amo-client-go/lib/keys"
)

// Tx broadcast in AMO context

func Transfer(to string, amount string, key keys.Key) (TmBroadcastResult, error) {
	ret, err := SignSendTx("transfer", struct {
		To     string `json:"to"`
		Amount string `json:"amount"`
	}{to, amount}, key)
	return ret, err
}

func Stake(validator string, amount string, key keys.Key) (TmBroadcastResult, error) {
	ret, err := SignSendTx("stake", struct {
		Validator string `json:"validator"`
		Amount    string `json:"amount"`
	}{validator, amount}, key)
	return ret, err
}

func Withdraw(amount string, key keys.Key) (TmBroadcastResult, error) {
	ret, err := SignSendTx("withdraw", struct {
		Amount string `json:"amount"`
	}{amount}, key)
	return ret, err
}

func Delegate(to string, amount string, key keys.Key) (TmBroadcastResult, error) {
	ret, err := SignSendTx("delegate", struct {
		To     string `json:"to"`
		Amount string `json:"amount"`
	}{to, amount}, key)
	return ret, err
}

func Retract(amount string, key keys.Key) (TmBroadcastResult, error) {
	ret, err := SignSendTx("retract", struct {
		Amount string `json:"amount"`
	}{amount}, key)
	return ret, err
}

func Register(target string, custody string, key keys.Key) (TmBroadcastResult, error) {
	ret, err := SignSendTx("register", struct {
		Target  string `json:"target"`
		Custody string `json:"custody"`
	}{target, custody}, key)
	return ret, err
}

func Discard(target string, key keys.Key) (TmBroadcastResult, error) {
	ret, err := SignSendTx("discard", struct {
		Target string `json:"target"`
	}{target}, key)
	return ret, err
}

func Request(target string, payment string, key keys.Key) (TmBroadcastResult, error) {
	ret, err := SignSendTx("register", struct {
		Target  string `json:"target"`
		Payment string `json:"payment"`
	}{target, payment}, key)
	return ret, err
}

func Cancel(target string, key keys.Key) (TmBroadcastResult, error) {
	ret, err := SignSendTx("cancel", struct {
		Target string `json:"target"`
	}{target}, key)
	return ret, err
}

func Grant(target string, grantee string, custody string, key keys.Key) (TmBroadcastResult, error) {
	ret, err := SignSendTx("grant", struct {
		Target  string `json:"target"`
		Grantee string `json:"grantee"`
		Custody string `json:"custody"`
	}{target, grantee, custody}, key)
	return ret, err
}

func Revoke(target string, grantee string, key keys.Key) (TmBroadcastResult, error) {
	ret, err := SignSendTx("register", struct {
		Target  string `json:"target"`
		Grantee string `json:"grantee"`
	}{target, grantee}, key)
	return ret, err
}
