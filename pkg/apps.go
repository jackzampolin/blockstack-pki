package pkg

import (
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

// AppsNode represents an apps node
type AppsNode struct {
	Node  *hdkeychain.ExtendedKey
	Salt  string
	Index uint32
}

// GetAddress implements foo
func (asn AppsNode) GetAddress() string {
	addr, _ := asn.Node.Address(&chaincfg.MainNetParams)
	return addr.String()
}

// GetPKHex implements foo
func (asn AppsNode) GetPKHex() string {
	pk, _ := asn.Node.ECPrivKey()
	return fmt.Sprintf("%x01", pk.ToECDSA().D)
}

// String implements foo
func (asn AppsNode) String() string {
	return fmt.Sprintf(`==== Identity Node %d Apps Node ====
  - Address: %s
  - PrivKey: %s
`, asn.Index, asn.GetAddress(), asn.GetPKHex())
}

// GetAppNode returns the app node for a given an appDomain
func (asn AppsNode) GetAppNode(appDomain string) AppNode {
	hsh := sha256.New()
	hsh.Write([]byte(fmt.Sprintf("%s%s", appDomain, asn.Salt)))
	appIndex := hashCode(fmt.Sprintf("%x", hsh.Sum(nil)))
	appNode, _ := asn.Node.Child(hdkeychain.HardenedKeyStart + uint32(appIndex))
	return AppNode{
		Node:   appNode,
		Domain: appDomain,
	}
}

func hashCode(str string) int {
	hash := 0
	if len(str) == 0 {
		return hash
	}
	for i := 0; i < len(str); i++ {
		char := []rune(str)[i]
		hash = (hash << 5) - hash + int(char)
		hash = hash & hash
	}
	return hash & 0x7fffffff
}
