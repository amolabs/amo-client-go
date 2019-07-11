package keys

import (
	"errors"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/xsalsa20symmetric"

	"github.com/amolabs/amoabci/crypto/p256"
)

type Key struct {
	Type      string `json:"type"`
	Address   string `json:"address"`
	PubKey    []byte `json:"pub_key"`
	PrivKey   []byte `json:"priv_key"`
	Encrypted bool   `json:"encrypted"`
}

func GenerateKey(passphrase []byte, encrypt bool, seed string) (*Key, error) {
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

	key := new(Key)
	key.Type = p256.PrivKeyAminoName
	key.Address = pubKey.Address().String()
	key.PubKey = pubKey.RawBytes()
	if encrypt {
		key.PrivKey = xsalsa20symmetric.EncryptSymmetric(
			privKey.RawBytes(), crypto.Sha256(passphrase))
	} else {
		key.PrivKey = privKey.RawBytes()
	}
	key.Encrypted = encrypt

	return key, nil
}

func (key *Key) Decrypt(passphrase []byte) error {
	if !key.Encrypted {
		return errors.New("The key is not encrypted")
	}

	plainKey, err := xsalsa20symmetric.DecryptSymmetric(key.PrivKey, crypto.Sha256(passphrase))
	if err != nil {
		return err
	}

	key.Encrypted = false
	key.PrivKey = plainKey

	return nil
}
