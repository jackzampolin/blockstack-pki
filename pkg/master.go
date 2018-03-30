package pkg

// MasterNode implements the HDNode interface
import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/landonia/crypto/bip39"
)

// MasterNode is the HDWallet node that all the other nodes are derived from
type MasterNode struct {
	Node     *hdkeychain.ExtendedKey
	Mnemonic string
}

// GetAddress implements foo
func (mn MasterNode) GetAddress() string {
	addr, _ := mn.Node.Address(&chaincfg.MainNetParams)
	return addr.String()
}

// GetPKHex implements foo
func (mn MasterNode) GetPKHex() string {
	pk, _ := mn.Node.ECPrivKey()
	return fmt.Sprintf("%x01", pk.ToECDSA().D)
}

// String implements foo
func (mn MasterNode) String() string {
	return fmt.Sprintf(`==== Master Node ====
  - Address: %s
  - PrivKey: %s
`, mn.GetAddress(), mn.GetPKHex())
}

// MnemonicWords implements foo
func (mn MasterNode) MnemonicWords() string {
	return fmt.Sprintf(`==== Mnemonic ====
%s
`, mn.Mnemonic)
}

// GetWalletNode returns the wallet node at m/44'/0'/0'/0/0
func (mn MasterNode) GetWalletNode() WalletNode {
	one, _ := mn.Node.Child(hdkeychain.HardenedKeyStart + 44)
	two, _ := one.Child(hdkeychain.HardenedKeyStart + 0)
	three, _ := two.Child(hdkeychain.HardenedKeyStart + 0)
	four, _ := three.Child(0)
	five, _ := four.Child(0)
	return WalletNode{
		Node: five,
	}
}

// GetIdentityNode returns the identiy node at m/888'/0'/{{ index }}'
func (mn MasterNode) GetIdentityNode(index uint32) IdentityNode {
	one, _ := mn.Node.Child(hdkeychain.HardenedKeyStart + 888)
	two, _ := one.Child(hdkeychain.HardenedKeyStart + 0)

	// Generate the salt from the sha256 hash of the hex representation of the m/888'/0' node
	idsk, _ := two.ECPubKey()
	hsh := sha256.New()
	pubKeyHex := fmt.Sprintf("%x", idsk.SerializeCompressed())
	hsh.Write([]byte(pubKeyHex))
	salt := fmt.Sprintf("%x", hsh.Sum(nil))

	three, _ := two.Child(hdkeychain.HardenedKeyStart + index)
	return IdentityNode{
		Node:  three,
		Index: index,
		Salt:  salt,
	}
}

// NewMasterNode returns a MasterNode
func NewMasterNode() MasterNode {
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

	key, err := hdkeychain.NewMaster(mnemonic.GenerateSeed(""), &chaincfg.MainNetParams)
	if err != nil {
		panic(err)
	}

	return MasterNode{
		Node:     key,
		Mnemonic: mnemonic.JoinWords(),
	}
}

// NewMasterFromMnemonic returns a MasterNode
func NewMasterFromMnemonic(mnemonic string) MasterNode {
	key, err := hdkeychain.NewMaster(bip39.GenerateSeed(mnemonic, ""), &chaincfg.MainNetParams)
	if err != nil {
		panic(err)
	}

	return MasterNode{
		Node:     key,
		Mnemonic: mnemonic,
	}
}
