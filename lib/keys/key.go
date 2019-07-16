package keys

import (
	"crypto/sha256"
	"errors"

	"github.com/tendermint/tendermint/crypto/xsalsa20symmetric"

	"github.com/amolabs/amoabci/crypto/p256"
)

type KeyEntry struct {
	Type      string `json:"type"`
	Address   string `json:"address"`
	PubKey    []byte `json:"pub_key"`
	PrivKey   []byte `json:"priv_key"`
	Encrypted bool   `json:"encrypted"`
}

func GenerateKey(seed string, passphrase []byte, encrypt bool) (*KeyEntry, error) {
	var privKey p256.PrivKeyP256
	if len(seed) > 0 {
		privKey = p256.GenPrivKeyFromSecret([]byte(seed))
	} else {
		privKey = p256.GenPrivKey()
	}

	pubKey, ok := privKey.PubKey().(p256.PubKeyP256)
	if !ok {
		return nil, errors.New("Error when deriving pubkey from privkey.")
	}

	key := new(KeyEntry)
	key.Type = p256.PrivKeyAminoName
	key.Address = pubKey.Address().String()
	key.PubKey = pubKey.RawBytes()
	if encrypt {
		encKey := sha256.Sum256(passphrase)
		key.PrivKey = xsalsa20symmetric.EncryptSymmetric(
			privKey.RawBytes(), encKey[:])
	} else {
		key.PrivKey = privKey.RawBytes()
	}
	key.Encrypted = encrypt

	return key, nil
}

func ImportKey(keyBytes []byte, passphrase []byte, encrypt bool) (*KeyEntry, error) {
	if len(keyBytes) != p256.PrivKeyP256Size {
		return nil, errors.New("Input private key size mismatch.")
	}
	var privKey p256.PrivKeyP256
	copy(privKey[:], keyBytes)

	pubKey, ok := privKey.PubKey().(p256.PubKeyP256)
	if !ok {
		return nil, errors.New("Error when deriving pubkey from privkey.")
	}

	key := new(KeyEntry)
	key.Type = p256.PrivKeyAminoName
	key.Address = pubKey.Address().String()
	key.PubKey = pubKey.RawBytes()
	if encrypt {
		encKey := sha256.Sum256(passphrase)
		key.PrivKey = xsalsa20symmetric.EncryptSymmetric(
			privKey.RawBytes(), encKey[:])
	} else {
		key.PrivKey = privKey.RawBytes()
	}
	key.Encrypted = encrypt

	return key, nil
}

func (key *KeyEntry) Decrypt(passphrase []byte) error {
	if !key.Encrypted {
		return errors.New("The key is not encrypted")
	}

	encKey := sha256.Sum256(passphrase)
	plainKey, err := xsalsa20symmetric.DecryptSymmetric(key.PrivKey, encKey[:])
	if err != nil {
		return err
	}

	key.Encrypted = false
	key.PrivKey = plainKey

	return nil
}
