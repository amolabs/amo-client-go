package storage

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amolabs/amo-client-go/lib/keys"
)

const (
	testToken = `{"token":"token body"}`
	testBody  = "test parcel content"
	testId    = "eeee"
)

// see https://github.com/amolabs/amo-storage#auth-api
func testHandleAuth(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte(`{"error":"Expected POST method"}`))
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"empty request body"}`))
		return
	}

	// same as AuthBody but, change each field as a pointer
	var authBody struct {
		User      *string          `json:"user"`
		Operation *json.RawMessage `json:"operation"`
	}
	err = json.Unmarshal(body, &authBody)
	if err != nil || authBody.User == nil || authBody.Operation == nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"malformed request body"}`))
		return
	}
	var opReq struct {
	}
	err = json.Unmarshal(*authBody.Operation, &opReq)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"malformed request body"}`))
		return
	}

	w.Write([]byte(testToken))
}

func testHandleUpload(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte(`{"error":"Expected POST method"}`))
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"empty request body"}`))
		return
	}

	// same as AuthBody but, change each field as a pointer
	var uploadBody struct {
		Owner    *string          `json:"owner"`
		Metadata *json.RawMessage `json:"metadata"`
		Data     *string          `json:"data"`
	}
	err = json.Unmarshal(body, &uploadBody)
	if err != nil || uploadBody.Owner == nil || uploadBody.Metadata == nil || uploadBody.Data == nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"malformed request body"}`))
		return
	}

	w.Write([]byte(testId))
}

func testHandleDownload(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte(`{"error":"Expected GET method"}`))
		return
	}
	authToken := req.Header.Get("X-Auth-Token")
	pubKey := req.Header.Get("X-Public-Key")
	sig := req.Header.Get("X-Signature")
	if authToken != testToken {
		w.WriteHeader(401)
		w.Write([]byte(`{"error":"X-Auth-Token header missing"}`))
		return
	}
	if len(pubKey) == 0 {
		w.WriteHeader(401)
		w.Write([]byte(`{"error":"X-Public-Key header missing"}`))
		return
	}
	_, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"malformed pubKey"}`))
		return
	}
	if len(sig) == 0 {
		w.WriteHeader(401)
		w.Write([]byte(`{"error":"X-Signature header missing"}`))
		return
	}
	_, err = base64.StdEncoding.DecodeString(sig)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"malformed signature"}`))
		return
	}

	w.Write([]byte(testBody))
}

func setUp() {
	// serve test auth challenge
	http.HandleFunc(
		"/api/v1/auth",
		testHandleAuth,
	)
	// serve parcel upload
	http.HandleFunc(
		"/api/v1/parcels",
		testHandleUpload,
	)
	// serve test parcel data
	http.HandleFunc(
		"/api/v1/parcels/",
		testHandleDownload,
	)
	go http.ListenAndServe("localhost:12345", nil)
	Endpoint = "http://localhost:12345"
}

func tearDown() {
	// TODO: kill the HTTP server launched in setUp()
}

func TestAll(t *testing.T) {
	setUp()
	defer tearDown()

	//Endpoint = "http://139.162.111.178"

	op, err := getOp("unknown", "blah")
	assert.Empty(t, op)
	assert.Error(t, err)

	// download
	op, err = getOp("download", "2f2f")
	assert.NotEmpty(t, op)
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"download","id":"2f2f"}`, op)

	authToken, err := requestToken("tester", `{ fjdska}`)
	assert.Error(t, err)
	assert.Nil(t, authToken)

	key, err := keys.GenerateKey(nil, false, "tester")
	assert.NoError(t, err)

	authToken, err = requestToken(key.Address, op)
	assert.NoError(t, err)
	assert.NotNil(t, authToken)
	assert.Equal(t, testToken, string(authToken))

	sig, err := signToken(*key, authToken)
	assert.NoError(t, err)
	assert.NotNil(t, sig)

	data, err := doDownload("2f2f", authToken, key.PubKey, sig)
	assert.NoError(t, err)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, testBody, string(data))

	// upload
	op, err = getOp("upload", "ffff")
	assert.NotEmpty(t, op)
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"upload","hash":"ffff"}`, op)

	authToken, err = requestToken(key.Address, op)
	assert.NoError(t, err)
	assert.NotNil(t, authToken)
	assert.Equal(t, testToken, string(authToken))

	sig, err = signToken(*key, authToken)
	assert.NoError(t, err)
	assert.NotNil(t, sig)

	id, err := doUpload(key.Address, nil, authToken, key.PubKey, sig)
	assert.NoError(t, err)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, testId, id)
}
