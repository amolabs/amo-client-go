package rpc

import ()

func QueryBalance(address string) ([]byte, error) {
	ret, err := ABCIQuery("/balance", address)
	return ret, err
}

func QueryStake(address string) ([]byte, error) {
	ret, err := ABCIQuery("/stake", address)
	return ret, err
}

func QueryDelegate(address string) ([]byte, error) {
	ret, err := ABCIQuery("/delegate", address)
	return ret, err
}

func QueryParcel(parcelID string) ([]byte, error) {
	ret, err := ABCIQuery("/parcel", parcelID)
	return ret, err
}

func QueryRequest(buyer string, target string) ([]byte, error) {
	ret, err := ABCIQuery("/request", struct {
		Buyer  string `json:"buyer"`
		Target string `json:"target"`
	}{buyer, target})
	return ret, err
}

func QueryUsage(buyer string, target string) ([]byte, error) {
	ret, err := ABCIQuery("/usage", struct {
		Buyer  string `json:"buyer"`
		Target string `json:"target"`
	}{buyer, target})
	return ret, err
}
