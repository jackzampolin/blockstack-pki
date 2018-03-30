package pkg

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

// AppNode represents an apps node
type AppNode struct {
	Node   *hdkeychain.ExtendedKey
	Salt   string
	Domain string
	Index  uint32
}

// GetAddress implements foo
func (an AppNode) GetAddress() string {
	addr, _ := an.Node.Address(&chaincfg.MainNetParams)
	return addr.String()
}

// GetPKHex implements foo
func (an AppNode) GetPKHex() string {
	pk, _ := an.Node.ECPrivKey()
	return fmt.Sprintf("%x01", pk.ToECDSA().D)
}

// String implements foo
func (an AppNode) String() string {
	return fmt.Sprintf(`==== Identity Node %d App Node "%s" ====
  - Address: %s
  - PrivKey: %s
`, an.Index, an.Domain, an.GetAddress(), an.GetPKHex())
}
