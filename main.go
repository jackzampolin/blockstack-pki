package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/landonia/crypto/bip39"
)

func genMnemonic() bip39.Mnemonics {
	token := make([]byte, 16)
	rand.Read(token)

	entropy, err := bip39.GenerateRandomEntropy(128)
	if err != nil {
		log.Fatal(err)
	}

	mnemonic, err := entropy.GenerateMnemonics(bip39.English)
	if err != nil {
		log.Fatal(err)
	}

	return mnemonic
}

func main() {
	fmt.Println("Generating a new mnomonic and then deriving keys with explainations")
	mnemonic := genMnemonic()

	fmt.Println("  mnemonic:", mnemonic.JoinWords())

	// Generate a new master node using the seed.
	key, err := hdkeychain.NewMaster(mnemonic.GenerateSeed(""), &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}

	// Generate the ECPrivKey for pulling out the hardened hex representation

	fmt.Println("  Root Private Key hardened:", privKeyHex(key))

	btc := getBTC(key)

	fmt.Println("  Browser Bitcoin Wallet Address:", getAddress(btc))
	fmt.Println("  Browser Bitcoin Wallet PrivateKey:", privKeyHex(btc))

	id0 := getIdentity(key, 0)

	fmt.Println("  ID0 Identity Address:", getAddress(id0))
	fmt.Println("  ID0 Identity PKHex:", privKeyHex(id0))

	id1 := getIdentity(key, 1)

	fmt.Println("  ID1 Identity Address:", getAddress(id1))
	fmt.Println("  ID1 Identity PKHex:", privKeyHex(id1))
}

func getAddress(key *hdkeychain.ExtendedKey) string {
	addr, err := key.Address(&chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	return addr.String()
}

func privKeyHex(key *hdkeychain.ExtendedKey) string {
	pk, err := key.ECPrivKey()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x01", pk.ToECDSA().D)
}

func getBTC(key *hdkeychain.ExtendedKey) *hdkeychain.ExtendedKey {
	one, err := key.Child(hdkeychain.HardenedKeyStart + 44)
	if err != nil {
		log.Fatal(err)
	}
	two, err := one.Child(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		log.Fatal(err)
	}
	three, err := two.Child(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		log.Fatal(err)
	}
	four, err := three.Child(0)
	if err != nil {
		log.Fatal(err)
	}
	five, err := four.Child(0)
	if err != nil {
		log.Fatal(err)
	}
	return five
}

func getIdentity(key *hdkeychain.ExtendedKey, index uint32) *hdkeychain.ExtendedKey {
	one, err := key.Child(hdkeychain.HardenedKeyStart + 888)
	if err != nil {
		log.Fatal(err)
	}
	two, err := one.Child(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		log.Fatal(err)
	}
	three, err := two.Child(hdkeychain.HardenedKeyStart + index)
	if err != nil {
		log.Fatal(err)
	}
	return three
}
