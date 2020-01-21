package rpc

import (
	"fmt"

	"github.com/amolabs/amo-client-go/lib/types"
)

// ABCI queries in AMO context

func QueryAppConfig() ([]byte, error) {
	ret, err := ABCIQuery("/config", nil)
	return ret, err
}

func QueryBalance(udc uint32, address string) ([]byte, error) {
	queryPath := "/balance"
	if udc != 0 {
		queryPath = fmt.Sprintf("%s"+"/%d", queryPath, udc)
	}
	ret, err := ABCIQuery(queryPath, address)
	if ret == nil {
		ret = []byte("0")
	}
	return ret, err
}

func QueryStake(address string) ([]byte, error) {
	return ABCIQuery("/stake", address)
}

func QueryDelegate(address string) ([]byte, error) {
	return ABCIQuery("/delegate", address)
}

func QueryDraft(draftID string) ([]byte, error) {
	draftIDUint32, err := types.ConvIDFromStr(draftID)
	if err != nil {
		return nil, err
	}
	return ABCIQuery("/draft", draftIDUint32)
}

func QueryVote(draftID, address string) ([]byte, error) {
	draftIDUint32, err := types.ConvIDFromStr(draftID)
	if err != nil {
		return nil, err
	}
	return ABCIQuery("/vote", struct {
		DraftID uint32 `json:"draft_id"`
		Voter   string `json:"voter"`
	}{draftIDUint32, address})
}

func QueryStorage(storageID string) ([]byte, error) {
	storageIDUint32, err := types.ConvIDFromStr(storageID)
	if err != nil {
		return nil, err
	}
	return ABCIQuery("/storage", storageIDUint32)
}

func QueryParcel(parcelID string) ([]byte, error) {
	return ABCIQuery("/parcel", parcelID)
}

func QueryRequest(buyer string, target string) ([]byte, error) {
	return ABCIQuery("/request", struct {
		Buyer  string `json:"buyer"`
		Target string `json:"target"`
	}{buyer, target})
}

func QueryUsage(buyer string, target string) ([]byte, error) {
	return ABCIQuery("/usage", struct {
		Buyer  string `json:"buyer"`
		Target string `json:"target"`
	}{buyer, target})
}

func QueryBlockIncentive(height string) ([]byte, error) {
	return ABCIQuery("/inc_block", height)
}

func QueryAddressIncentive(address string) ([]byte, error) {
	return ABCIQuery("/inc_address", address)
}

func QueryIncentive(height string, address string) ([]byte, error) {
	return ABCIQuery("/inc", struct {
		Height  string `json:"height"`
		Address string `json:"address"`
	}{height, address})
}
