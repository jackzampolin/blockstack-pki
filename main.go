package main

import (
	"fmt"

	"github.com/jackzampolin/blockstack-pki/pkg"
)

func main() {
	fmt.Println("Using saved mnemonic to show equality with node.js example code")
	fmt.Println()
	// NOTE: To generate a new mnemonic use the NewMasterNode() function instead of the NewMasterFromMnemonic
	mn := pkg.NewMasterFromMnemonic("slam cheap sponsor average issue lemon nuclear file gesture snake other seminar")
	fmt.Println(mn.MnemonicWords())
	fmt.Println(mn)
	wn := mn.GetWalletNode()
	fmt.Println(wn)
	in0 := mn.GetIdentityNode(0)
	in1 := mn.GetIdentityNode(1)
	fmt.Println(in0)
	fmt.Println(in1)
	asn0 := in0.GetAppsNode()
	fmt.Println(asn0)
	an0 := asn0.GetAppNode("https://www.foo.com")
	fmt.Println(an0)
}
