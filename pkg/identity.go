package pkg

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

// IdentityNode represents an identiy node
type IdentityNode struct {
	Node  *hdkeychain.ExtendedKey
	Index uint32
	Salt  string
}

// GetAddress implements foo
func (in IdentityNode) GetAddress() string {
	addr, _ := in.Node.Address(&chaincfg.MainNetParams)
	return addr.String()
}

// GetPKHex implements foo
func (in IdentityNode) GetPKHex() string {
	pk, _ := in.Node.ECPrivKey()
	return fmt.Sprintf("%x01", pk.ToECDSA().D)
}

// String implements foo
func (in IdentityNode) String() string {
	return fmt.Sprintf(`==== Identity Node %d ====
  - Address: %s
  - PrivKey: %s
`, in.Index, in.GetAddress(), in.GetPKHex())
}

// GetAppsNode returns the apps node for the Identity
func (in IdentityNode) GetAppsNode() AppsNode {
	node, _ := in.Node.Child(hdkeychain.HardenedKeyStart + 0)
	return AppsNode{
		Node:  node,
		Index: in.Index,
		Salt:  in.Salt,
	}
}

// GetEncryptionNode returns the Encryption node for this identity
func (in IdentityNode) GetEncryptionNode() *hdkeychain.ExtendedKey {
	node, _ := in.Node.Child(hdkeychain.HardenedKeyStart + 1)
	return node
}

// GetSigningNode returns the Signing node for this identity
func (in IdentityNode) GetSigningNode() *hdkeychain.ExtendedKey {
	node, _ := in.Node.Child(hdkeychain.HardenedKeyStart + 1)
	return node
}
