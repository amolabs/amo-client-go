package storage

import (
	"encoding/base64"
	"net/http"

	"github.com/amolabs/amo-client-go/lib/keys"
)

func doDownload(id string, token, pubKey, sig []byte) ([]byte, error) {
	// TODO: check op description in token
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		Endpoint+"/api/v1/parcels/"+id,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Auth-Token", string(token))
	req.Header.Add("X-Public-Key", base64.StdEncoding.EncodeToString(pubKey))
	req.Header.Add("X-Signature", base64.StdEncoding.EncodeToString(sig))

	return doHTTP(client, req)
}

func Download(parcelID string, key keys.KeyEntry) ([]byte, error) {
	op, err := getOp("download", parcelID)
	if err != nil {
		return nil, err
	}
	authToken, err := requestToken(key.Address, op)
	if err != nil {
		return nil, err
	}
	sig, err := signToken(key, authToken)

	return doDownload(parcelID, authToken, key.PubKey, sig)
}
