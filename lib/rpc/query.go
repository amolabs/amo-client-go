package rpc

// ABCI queries in AMO context

func QueryAppConfig() ([]byte, error) {
	ret, err := ABCIQuery("/app_config", nil)
	return ret, err
}

func QueryBalance(address string) ([]byte, error) {
	ret, err := ABCIQuery("/balance", address)
	if ret == nil {
		ret = []byte("0")
	}
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

func QueryBlockIncentive(height string) ([]byte, error) {
	ret, err := ABCIQuery("/inc_block", height)
	return ret, err
}

func QueryAddressIncentive(address string) ([]byte, error) {
	ret, err := ABCIQuery("/inc_address", address)
	return ret, err
}

func QueryIncentive(height string, address string) ([]byte, error) {
	ret, err := ABCIQuery("/inc", struct {
		Height  string `json:"height"`
		Address string `json:"address"`
	}{height, address})
	return ret, err
}
