package keys

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/amolabs/amo-client-go/cli/util"
)

type KeyRing struct {
	filePath string
	keyList  map[string]KeyEntry // just a cache
}

func GetKeyRing(path string) (*KeyRing, error) {
	kr := new(KeyRing)
	kr.filePath = path
	kr.keyList = make(map[string]KeyEntry)
	err := kr.Load()
	if err != nil {
		return nil, err
	}
	return kr, nil
}

func (kr *KeyRing) Load() error {
	err := util.EnsureFile(kr.filePath)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(kr.filePath)
	if err != nil {
		return err
	}

	newKeyList := make(map[string]KeyEntry)
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
	err := util.EnsureFile(kr.filePath)
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

func (kr *KeyRing) GenerateNewKey(username string, seed string, passphrase []byte, encrypt bool) (*KeyEntry, error) {
	key, err := GenerateKey(seed, passphrase, encrypt)
	if err != nil {
		return nil, errors.New("Fail to generate new key.")
	}

	return key, kr.AddKey(username, key)
}

func (kr *KeyRing) ImportNewKey(username string, keyBytes []byte, passphrase []byte, encrypt bool) (*KeyEntry, error) {
	key, err := ImportKey(keyBytes, passphrase, encrypt)
	if err != nil {
		return nil, errors.New("Fail to import key.")
	}

	return key, kr.AddKey(username, key)
}

func (kr *KeyRing) GetKey(username string) *KeyEntry {
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

func (kr *KeyRing) PrintKeyList(withPubKey bool) {
	sortKey := make([]string, 0, len(kr.keyList))
	maxUsernameSize := 0
	for k := range kr.keyList {
		if len(k) > maxUsernameSize {
			maxUsernameSize = len(k)
		}
		sortKey = append(sortKey, k)
	}

	sort.Strings(sortKey)

	fmt.Printf("%3s %-*s %-3s %-40s", "#", maxUsernameSize, "username", "enc", "address")
	if withPubKey {
		fmt.Printf(" %-65s", "pubkey")
	}
	fmt.Printf("\n")

	i := 0
	for _, username := range sortKey {
		i++
		key := kr.keyList[username]

		enc := "x"
		if key.Encrypted {
			enc = "o"
		}
		fmt.Printf("%3d %-*s %-3s %-40s", i, maxUsernameSize, username, enc, key.Address)
		if withPubKey {
			fmt.Printf(" %-65X", key.PubKey)
		}
		fmt.Printf("\n")
	}
}

func (kr *KeyRing) GetNumKeys() int {
	return len(kr.keyList)
}

func (kr *KeyRing) GetFirstKey() *KeyEntry {
	var key *KeyEntry = nil
	for _, v := range kr.keyList {
		key = &v
		break
	}
	return key
}

func (kr *KeyRing) AddKey(username string, key *KeyEntry) error {
	_, ok := kr.keyList[username]
	if ok {
		return errors.New("Username already exists.")
	}

	kr.keyList[username] = *key
	return kr.Save()
}
