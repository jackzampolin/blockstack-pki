var blockstack = require('blockstack')
const keychains = require('blockstack-keychains')
const bitcoin = require('bitcoinjs-lib')
const bip39 = require('bip39')

class IdentityNode{
    constructor(key){
        this.key = key
    }
    getAddress(){
        return this.key.keyPair.getAddress()
    }
    getSKHex(){
        return this.key.keyPair.d.toBuffer(32).toString('hex')
    }
    getHex(){
      return this.key.keyPair.d.toHex() + '01'
    }
}

console.log("Generating a new mnomonic and then deriving keys with explanations")

let keychain = new keychains.PrivateKeychain()

var mnemonic = "slam cheap sponsor average issue lemon nuclear file gesture snake other seminar"

// console.log("  mnemonic:", keychain.mnemonic())
console.log("  mnemonic:", mnemonic)

var getMaster = (mnemonic) => { return bitcoin.HDNode.fromSeedBuffer(bip39.mnemonicToSeed(mnemonic)) }
 
// var master = getMaster(keychain.mnemonic())
var master = getMaster(mnemonic)

console.log("  Root Private Key hardened:", master.keyPair.d.toHex() + '01')

// Derive browser wallet key at `44/0/0`
var getBTC = (master) => { return master.deriveHardened(44).deriveHardened(0).deriveHardened(0).derive(0).derive(0) }

let btcKey = getBTC(master)

console.log("  Browser Bitcoin Wallet Address:", btcKey.keyPair.getAddress())
console.log("  Browser Bitcoin Wallet PrivateKey:", btcKey.keyPair.d.toHex() + '01')

// Derive Identity keys at `888/0/index`
var getIdentityKeyCurrent = (master, index) => { return new IdentityNode(master.deriveHardened(888).deriveHardened(0).deriveHardened(index)) }

let id0 = getIdentityKeyCurrent(master, 0)

console.log("  ID0 Identity Address:", id0.getAddress())
console.log("  ID0 Identity PKHex:", id0.getHex())

let id1 = getIdentityKeyCurrent(master, 1)

console.log("  ID1 Identity Address:", id1.getAddress())
console.log("  ID1 Identity PKHex:", id1.getHex())
