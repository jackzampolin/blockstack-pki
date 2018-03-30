var blockstack = require('blockstack')
const bitcoin = require('bitcoinjs-lib')
const bip39 = require('bip39')
const crypto = require('crypto')

// Simple hashing function
function hashCode(string) {
  let hash = 0
  if (string.length === 0) return hash
  for (let i = 0; i < string.length; i++) {
    const character = string.charCodeAt(i)
    hash = (hash << 5) - hash + character
    hash = hash & hash
  }
  return hash & 0x7fffffff
}

// Base class for HD (hierarchal deterministic) nodes
class HdNode {
  constructor(node) {
    this.node = node
  }
  getAddress(){
    return this.node.keyPair.getAddress()
  }
  getSKHex(){
    return this.node.keyPair.d.toBuffer(32).toString('hex')
  }
  getHex(){
    return this.node.keyPair.d.toHex() + '01'
  }
}

///////////////////////
// Mnemonic
///////////////////////

console.log("Using saved mnemonic to show equality with golang example code")

var mnemonic = "slam cheap sponsor average issue lemon nuclear file gesture snake other seminar"

console.log(`
==== Mnemonic ====
${ mnemonic }
`)

///////////////////////
// Master Node Generation
///////////////////////

class MasterNode extends HdNode {
  constructor(node){
    super(node)
  }
  getWalletNode() {
    // Derive browser wallet node at `44'/0'/0'/0/0`
    return new WalletNode(this.node.deriveHardened(44).deriveHardened(0).deriveHardened(0).derive(0).derive(0))
  }
  getIdentityNode(index) {
    // Derive browser wallet node at `888'/0'/index'`
    var identitiesNode = this.node.deriveHardened(888).deriveHardened(0)
    var pubKeyHex = identitiesNode.keyPair.getPublicKeyBuffer().toString('hex')
    var salt = crypto.createHash('sha256').update(pubKeyHex).digest('hex')
    return new IdentityNode(identitiesNode.deriveHardened(index), salt)
  }
}

var myMasterNode = new MasterNode(bitcoin.HDNode.fromSeedBuffer(bip39.mnemonicToSeed(mnemonic)))

console.log(`==== Master Node ====
  - Address: ${myMasterNode.getAddress()}
  - PrivKey: ${myMasterNode.getHex()}
`)


///////////////////////
// Wallet Node Generation
///////////////////////

class WalletNode extends HdNode {
  constructor(node){
    super(node)
  }
}

const myWalletNode = myMasterNode.getWalletNode()

console.log(`==== Wallet Node ====
  - Address: ${myWalletNode.getAddress()}
  - PrivKey: ${myWalletNode.getHex()}
`)


///////////////////////
// Identity Node Generation
///////////////////////

class IdentityNode extends HdNode {
  constructor(node, salt){
    super(node)
    this.salt = salt
  }
  getAppsNode() {
    return new AppsNode(this.node.deriveHardened(0), this.salt)
  }
  getEncryptionNode() {
    return this.node.deriveHardened(1)
  }
  getSigningNode() {
    return this.node.deriveHardened(2)
  }
}

const myIdentityNode = myMasterNode.getIdentityNode(0)
const myIdentityNode1 = myMasterNode.getIdentityNode(1)

console.log(`==== Identity Node 0 ====
  - Address: ${myIdentityNode.getAddress()}
  - PrivKey: ${myIdentityNode.getHex()}
`)

console.log(`==== Identity Node 1 ====
  - Address: ${myIdentityNode1.getAddress()}
  - PrivKey: ${myIdentityNode1.getHex()}
`)


///////////////////////
// Apps Node Generation
///////////////////////

class AppsNode extends HdNode {
  constructor(node, salt){
    super(node)
    this.salt = salt
  }
  getAppNode(appDomain) {
    const hash = crypto
      .createHash('sha256')
      .update(`${appDomain}${this.salt}`)
      .digest('hex')
    const appIndex = hashCode(hash)
    const appNode = this.node.deriveHardened(appIndex)
    return new AppNode(appNode, appDomain)
  }
}

const myAppsNode = myIdentityNode.getAppsNode()

console.log(`==== Identity Node 0 Apps Node ====
  - Address: ${myAppsNode.getAddress()}
  - PrivKey: ${myAppsNode.getHex()}
`)


///////////////////////
// App Node Generation
///////////////////////

class AppNode extends HdNode {
  constructor(node, appDomain) {
    super(node)
    this.appDomain = appDomain
  }
  getAppDomain() {
    return this.appDomain
  }
}

const myAppNode = myAppsNode.getAppNode('https://www.foo.com')

console.log(`==== Identity Node 0 App Node "https://www.foo.com" ====
  - Address: ${myAppNode.getAddress()}
  - PrivKey: ${myAppNode.getHex()}
`)
