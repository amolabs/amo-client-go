package rpc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64" // when decoding rpc response
	"encoding/hex"    // when encoding rpc request
	"encoding/json"
	"math/big"
	"strings"

	"github.com/amolabs/amo-client-go/lib/keys"
	"github.com/ybbus/jsonrpc"
)

var (
	RpcRemote       = "http://0.0.0.0:26657"
	rpcWsEndpoint   = "/websocket"
	AddressByteSize = 20
	NonceByteSize   = 4
	c               = elliptic.P256() // move to crypto sub-package
)

// generic ABCI query in Tendermint context

type ABCIQueryParams struct {
	Path   string `json:"path"`
	Data   string `json:"data"`
	Height string `json:"height"`
	Prove  bool   `json:"prove"`
}

type ABCIQueryResponse struct {
	Log   string `json:"log"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// XXX: Weired, but tendermint does this anyway
type TmQueryResult struct {
	Response ABCIQueryResponse `json:"response"`
}

func ABCIQuery(path string, queryData interface{}) ([]byte, error) {
	queryJson, err := json.Marshal(queryData)
	if err != nil {
		return nil, err
	}
	params := ABCIQueryParams{
		Path:   path,
		Data:   hex.EncodeToString([]byte(queryJson)),
		Height: "0",
		Prove:  false,
	}

	c := jsonrpc.NewClient(RpcRemote)
	rsp, err := c.Call("abci_query", params)
	if err != nil { // call error
		return nil, err
	}
	if rsp.Error != nil { // rpc error
		return nil, err
	}
	var res TmQueryResult
	err = rsp.GetObject(&res)
	if err != nil { // conversion error
		return nil, err
	}
	// TODO: check ABCI error
	// XXX: Do we need to do something with Log and Key?
	ret, err := base64.StdEncoding.DecodeString(string(res.Response.Value))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// generic Tx broadcast in Tendermint context

type TxToSign struct {
	Type    string          `json:"type"`
	Sender  string          `json:"sender"`
	Nonce   string          `json:"nonce"`
	Payload json.RawMessage `json:"payload"`
	Fee     string          `json:"fee"`
}

type TxSig struct {
	Pubkey   string `json:"pubkey"`
	SigBytes string `json:"sig_bytes"`
}

type TxToSend struct {
	Type      string          `json:"type"`
	Sender    string          `json:"sender"`
	Nonce     string          `json:"nonce"`
	Payload   json.RawMessage `json:"payload"`
	Fee       string          `json:"fee"`
	Signature TxSig           `json:"signature"`
}

type BroadcastParams struct {
	Tx []byte `json:"tx"`
}

type TmTxResult struct {
	CheckTx struct {
		Code int64  `json:"code,omitempty"`
		Info string `json:"info,omitempty"`
	} `json:"check_tx"`
	DeliverTx struct {
		Code int64  `json:"code,omitempty"`
		Info string `json:"info,omitempty"`
	} `json:"deliver_tx"`
	Hash   string `json:"hash"`
	Height string `json:"height"` // number as a string
}

func getAddressBytes(pubkey []byte) []byte {
	hash := sha256.Sum256(pubkey)
	return hash[:AddressByteSize]
}

func SignSendTx(txType string, payload interface{}, key keys.KeyEntry, fee string) (TmTxResult, error) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return TmTxResult{}, err
	}

	nonceBytes := make([]byte, NonceByteSize)
	_, err = rand.Read(nonceBytes)
	sender := strings.ToUpper(key.Address)
	nonce := strings.ToUpper(hex.EncodeToString(nonceBytes))
	txToSign := TxToSign{
		Type:    txType,
		Sender:  sender,
		Nonce:   nonce,
		Payload: payloadJson,
		Fee:     fee,
	}
	msg, err := json.Marshal(txToSign)
	if err != nil {
		return TmTxResult{}, err
	}
	// do sign
	h := sha256.Sum256(msg)
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
		return TmTxResult{}, err
	}
	rb := r.Bytes()
	sb := s.Bytes()
	sigBytes := make([]byte, 64)
	copy(sigBytes[32-len(rb):], rb)
	copy(sigBytes[64-len(sb):], sb)
	// done sign
	txSig := TxSig{
		Pubkey:   hex.EncodeToString(key.PubKey),
		SigBytes: hex.EncodeToString(sigBytes),
	}
	tx := TxToSend{
		Type:      txToSign.Type,    // forward
		Sender:    txToSign.Sender,  // forward
		Nonce:     txToSign.Nonce,   // forward
		Payload:   txToSign.Payload, // forward
		Fee:       txToSign.Fee,     // forward
		Signature: txSig,            // signature appendix
	}
	b, err := json.Marshal(tx)
	if err != nil {
		return TmTxResult{}, err
	}

	return BroadcastTx(b)
}

func BroadcastTx(tx []byte) (TmTxResult, error) {
	params := BroadcastParams{
		Tx: tx,
	}
	c := jsonrpc.NewClient(RpcRemote)
	rsp, err := c.Call("broadcast_tx_commit", params)
	if err != nil { // call error
		return TmTxResult{}, err
	}
	if rsp.Error != nil { // rpc error
		return TmTxResult{}, err
	}
	var res TmTxResult
	err = rsp.GetObject(&res)
	if err != nil { // conversion error
		return TmTxResult{}, err
	}
	return res, nil
}

// misc rpcs

type TmStatusResult struct {
	NodeInfo      json.RawMessage `json:"node_info"`
	SyncInfo      json.RawMessage `json:"sync_info"`
	ValidatorInfo json.RawMessage `json:"validator_info"`
}

func NodeStatus() (TmStatusResult, error) {
	c := jsonrpc.NewClient(RpcRemote)
	rsp, err := c.Call("status")
	if err != nil { // call error
		return TmStatusResult{}, err
	}
	if rsp.Error != nil { // rpc error
		return TmStatusResult{}, err
	}
	var res TmStatusResult
	err = rsp.GetObject(&res)
	if err != nil { // conversion error
		return TmStatusResult{}, err
	}
	return res, nil
}
