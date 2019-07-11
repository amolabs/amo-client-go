package keys

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/xsalsa20symmetric"

	"github.com/amolabs/amoabci/crypto/p256"
)

type KeyRing struct {
	filePath string
	keyList  map[string]Key // just a cache
}

func EnsureDir(dir string, mode os.FileMode) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, mode)
		if err != nil {
			return fmt.Errorf("Could not create directory %v. %v", dir, err)
		}
	}
	return nil
}

func EnsureFile(path string) error {
	dirPath, _ := filepath.Split(path)

	if len(dirPath) > 0 {
		err := EnsureDir(dirPath, 0775)
		if err != nil {
			return err
		}
	}

	_, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	return err
}

func GetKeyRing(path string) (*KeyRing, error) {
	kr := new(KeyRing)
	kr.filePath = path
	kr.keyList = make(map[string]Key)
	err := kr.Load()
	if err != nil {
		return nil, err
	}
	return kr, nil
}

func (kr *KeyRing) Load() error {
	err := EnsureFile(kr.filePath)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(kr.filePath)
	if err != nil {
		return err
	}

	newKeyList := make(map[string]Key)
	if len(b) > 0 {
		err = json.Unmarshal(b, &newKeyList)
		if err != nil {
			return err
		}
	}

	kr.keyList = newKeyList

	return nil
}

func (kr *KeyRing) Save() error {
	err := EnsureFile(kr.filePath)
	if err != nil {
		return err
	}

	b, err := json.Marshal(kr.keyList)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(kr.filePath, b, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (kr *KeyRing) GenerateNewKey(username string, passphrase []byte, encrypt bool, seed string) (*Key, error) {
	key, err := GenerateKey(passphrase, encrypt, seed)
	if err != nil {
		return nil, errors.New("Fail to generate new key.")
	}

	return key, kr.addKey(username, key)
}

// TODO: use KeyRing.addKey()
func (kr *KeyRing) ImportPrivKey(keyBytes []byte,
	username string, passphrase []byte, encrypt bool) (*Key, error) {
	_, ok := kr.keyList[username]
	if ok {
		return nil, errors.New("Username already exists.")
	}

	if len(keyBytes) != p256.PrivKeyP256Size {
		return nil, errors.New("Input private key size mismatch.")
	}
	var privKey p256.PrivKeyP256
	copy(privKey[:], keyBytes)

	return kr.addNewP256Key(privKey, username, passphrase, encrypt)
}

func (kr *KeyRing) GetKey(username string) *Key {
	key, ok := kr.keyList[username]
	if !ok {
		return nil
	}
	return &key
}

func (kr *KeyRing) RemoveKey(username string) error {
	_, ok := kr.keyList[username]
	if !ok {
		return errors.New("Username not found")
	}

	delete(kr.keyList, username)

	return kr.Save()
}

func (kr *KeyRing) PrintKeyList() {
	sortKey := make([]string, 0, len(kr.keyList))
	for k := range kr.keyList {
		sortKey = append(sortKey, k)
	}

	sort.Strings(sortKey)

	fmt.Printf("%3s %-9s %-20s %-3s %-40s\n",
		"#", "username", "type", "enc", "address")

	i := 0
	for _, username := range sortKey {
		i++
		key := kr.keyList[username]

		enc := "x"
		if key.Encrypted {
			enc = "o"
		}
		fmt.Printf("%3d %-9s %-20s %-3s %-40s\n",
			i, username, key.Type, enc, key.Address)
	}
}

func (kr *KeyRing) GetNumKeys() int {
	return len(kr.keyList)
}

func (kr *KeyRing) GetFirstKey() *Key {
	var key *Key = nil
	for _, v := range kr.keyList {
		key = &v
		break
	}
	return key
}

func (kr *KeyRing) addKey(username string, key *Key) error {
	_, ok := kr.keyList[username]
	if ok {
		return errors.New("Username already exists.")
	}

	kr.keyList[username] = *key
	return kr.Save()
}

func (kr *KeyRing) addNewP256Key(privKey p256.PrivKeyP256,
	username string, passphrase []byte, encrypt bool) (*Key, error) {
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

	kr.keyList[username] = *key
	err := kr.Save()
	if err != nil {
		return nil, err
	}

	return key, nil
}
