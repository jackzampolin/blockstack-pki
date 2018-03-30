package pkg

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

// WalletNode represents the Bitcoin Wallet Node
type WalletNode struct {
	Node *hdkeychain.ExtendedKey
}

// GetAddress implements foo
func (wn WalletNode) GetAddress() string {
	addr, _ := wn.Node.Address(&chaincfg.MainNetParams)
	return addr.String()
}

// GetPKHex implements foo
func (wn WalletNode) GetPKHex() string {
	pk, _ := wn.Node.ECPrivKey()
	return fmt.Sprintf("%x01", pk.ToECDSA().D)
}

// String implements foo
func (wn WalletNode) String() string {
	return fmt.Sprintf(`==== Wallet Node ====
  - Address: %s
  - PrivKey: %s
`, wn.GetAddress(), wn.GetPKHex())
}
