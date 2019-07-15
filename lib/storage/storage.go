package storage

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/amolabs/amo-client-go/lib/keys"
)

var (
	Endpoint = "http://139.162.111.178"
	c        = elliptic.P256() // move to crypto sub-package
)

type AuthBody struct {
	User      string          `json:"user"`
	Operation json.RawMessage `json:"operation"`
}

func getOp(name, param string) (string, error) {
	var (
		op  []byte
		err error
	)
	switch name {
	case "upload":
		authOp := struct {
			Name string `json:"name"`
			Hash string `json:"hash"`
		}{"upload", param}
		op, err = json.Marshal(authOp)
	case "download":
		authOp := struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		}{"download", param}
		op, err = json.Marshal(authOp)
	case "remove":
		authOp := struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		}{"remove", param}
		op, err = json.Marshal(authOp)
	default:
		return "", errors.New("Unknown operation name")
	}

	return string(op), err
}

func requestToken(account, op string) ([]byte, error) {
	authBody := AuthBody{
		User:      account,
		Operation: []byte(op),
	}
	reqJson, err := json.Marshal(authBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		Endpoint+"/api/v1/auth",
		bytes.NewBuffer(reqJson),
	)
	req.Header.Add("Content-Type", "application/json")

	return doHTTP(client, req)
}

func signToken(key keys.Key, token []byte) ([]byte, error) {
	// do sign
	h := sha256.Sum256(token)
	X, Y := c.ScalarBaseMult(key.PrivKey[:])
	ecdsaPrivKey := ecdsa.PrivateKey{
		D: new(big.Int).SetBytes(key.PrivKey[:]),
		PublicKey: ecdsa.PublicKey{
			Curve: c,
			X:     X,
			Y:     Y,
		},
	}
	r, s, err := ecdsa.Sign(rand.Reader, &ecdsaPrivKey, h[:])
	if err != nil {
		return nil, err
	}
	rb := r.Bytes()
	sb := s.Bytes()
	sigBytes := make([]byte, 64)
	copy(sigBytes[32-len(rb):], rb)
	copy(sigBytes[64-len(sb):], sb)
	// done sign

	return sigBytes, nil
}

func doHTTP(client *http.Client, req *http.Request) ([]byte, error) {
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf(
			"HTTP status code %d, res body: %s",
			rsp.StatusCode,
			body,
		)
	}

	return body, nil
}
