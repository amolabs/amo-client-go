package keys

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testfile = "test_keyring.json"
)

func _tearDown() {
	os.RemoveAll(testfile)
}

func TestGetKeyRing(t *testing.T) {
	kr, err := GetKeyRing(testfile)
	assert.NoError(t, err)
	assert.NotNil(t, kr)

	_tearDown()
}

func TestGenKey(t *testing.T) {
	kr, err := GetKeyRing(testfile)
	assert.NoError(t, err)
	assert.NotNil(t, kr)
	assert.Equal(t, 0, len(kr.keyList))

	key, err := kr.GenerateNewKey("test", "test", []byte("pass"), true)
	assert.NoError(t, err)
	assert.NotNil(t, key)
	assert.Equal(t, 40, len(key.Address))
	assert.Equal(t, 65, len(key.PubKey)) // XXX: really?
	assert.True(t, key.Encrypted)
	key2 := kr.GetKey("test")
	assert.NotNil(t, key2)
	assert.Equal(t, key, key2)

	// check if the actual file was updated
	err = kr.Load()
	assert.NoError(t, err)
	key2 = kr.GetKey("test")
	assert.NotNil(t, key2)
	assert.Equal(t, key, key2)

	// test remove
	err = kr.RemoveKey("test")
	assert.NoError(t, err)
	key2 = kr.GetKey("test")
	assert.Nil(t, key2)

	err = kr.Load()
	assert.NoError(t, err)
	key2 = kr.GetKey("test")
	assert.Nil(t, key2)

	// test genkey without enc
	key, err = kr.GenerateNewKey("test", "test", nil, false)
	assert.NoError(t, err)
	assert.NotNil(t, key)
	assert.Equal(t, 40, len(key.Address))
	assert.Equal(t, 65, len(key.PubKey)) // XXX: really?
	assert.False(t, key.Encrypted)

	_tearDown()
}

func TestImportKey(t *testing.T) {
	kr, err := GetKeyRing(testfile)
	assert.NoError(t, err)
	assert.NotNil(t, kr)
	assert.Equal(t, 0, len(kr.keyList))

	// test import
	testKey, err := GenerateKey("seed", nil, false)
	assert.NotNil(t, testKey)
	assert.NoError(t, err)
	testKeyBytes := testKey.PrivKey

	key, err := kr.ImportNewKey("bob", testKeyBytes[:31], []byte("pass"), true)
	assert.Nil(t, key)
	assert.Error(t, err)

	key, err = kr.ImportNewKey("bob", testKeyBytes, []byte("pass"), true)
	assert.NotNil(t, key)
	assert.NoError(t, err)
	assert.Equal(t, 40, len(key.Address))
	assert.Equal(t, testKey.Address, key.Address)
	assert.Equal(t, testKey.PubKey, key.PubKey)
	assert.True(t, key.Encrypted)
	key2 := kr.GetKey("bob")
	assert.NotNil(t, key2)
	assert.Equal(t, key, key2)

	// check if the actual file was updated
	err = kr.Load()
	assert.NoError(t, err)
	key2 = kr.GetKey("bob")
	assert.NotNil(t, key2)
	assert.Equal(t, key, key2)

	_tearDown()
}
