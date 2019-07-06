package rpc

import (
	"encoding/base64" // when decoding query response
	"encoding/hex"    // when encoding query request
	"encoding/json"

	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"github.com/ybbus/jsonrpc"

	"github.com/amolabs/amo-client-go/lib/keys"
	"github.com/amolabs/amoabci/amo/tx"
	"github.com/amolabs/amoabci/crypto/p256"
)

var (
	RpcRemote     = "http://0.0.0.0:26657"
	rpcWsEndpoint = "/websocket"
)

type ABCIParams struct {
	Path   string `json:"path"`
	Data   string `json:"data"`
	Height string `json:"height"`
	Prove  bool   `json:"prove"`
}

type ABCIResponse struct {
	Log   string `json:"log"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TmResponse struct {
	Response ABCIResponse `json:"response"`
}

func ABCIQuery(path string, queryData interface{}) ([]byte, error) {
	queryJson, err := json.Marshal(queryData)
	if err != nil {
		return nil, err
	}
	params := ABCIParams{
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
	var res TmResponse
	err = rsp.GetObject(&res)
	if err != nil { // conversion error
		return nil, err
	}
	// TODO: check ABCI error
	// XXX: need to do something with Log and Key?
	ret, err := base64.StdEncoding.DecodeString(string(res.Response.Value))
	return ret, nil
}

// MakeTx handles making tx message
func MakeTx(t string, nonce uint32, payload interface{}, key keys.Key) (types.Tx, error) {
	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var privKey p256.PrivKeyP256
	copy(privKey[:], key.PrivKey)

	msg := tx.Tx{
		Type:    t,
		Payload: raw,
		Sender:  privKey.PubKey().Address(),
		Nonce:   cmn.RandBytes(tx.NonceSize),
	}

	err = msg.Sign(privKey)
	if err != nil {
		return nil, err
	}

	tx, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// RPCBroadcastTxCommit handles sending transactions
func RPCBroadcastTxCommit(tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	cli := client.NewHTTP(RpcRemote, rpcWsEndpoint)
	return cli.BroadcastTxCommit(tx)
}

// RPCStatus handle querying the status
func RPCStatus() (*ctypes.ResultStatus, error) {
	cli := client.NewHTTP(RpcRemote, rpcWsEndpoint)
	return cli.Status()
}

func RPCBlock(height int64) (*ctypes.ResultBlock, error) {
	cli := client.NewHTTP(RpcRemote, rpcWsEndpoint)
	return cli.Block(&height)
}
