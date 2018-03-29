var blockstack = require('blockstack')
const keychains = require('blockstack-keychains')
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

console.log("Generating a new mnemonic and then deriving keys with explanations")

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
    return new IdentityNode(this.node.deriveHardened(888).deriveHardened(0).deriveHardened(index))
  }
}

const createMasterNode = (mnemonic) => {
  return new MasterNode(
    bitcoin.HDNode.fromSeedBuffer(bip39.mnemonicToSeed(mnemonic))
  )
}

var myMasterNode = createMasterNode(mnemonic)

console.log(`
==== Master Node ====
- Address:
${myMasterNode.getAddress()}
- PrivateKey:
${myMasterNode.getSKHex()}
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

console.log(`
==== Bitcoin Wallet Node ====
- Address:
${myWalletNode.getAddress()}
- PrivateKey:
${myWalletNode.getSKHex()}
`)


///////////////////////
// Identity Node Generation
///////////////////////

class IdentityNode extends HdNode {
  constructor(node){
    super(node)
  }
  getAppsNode() {
    return new AppsNode(this.node.deriveHardened(0))
  }
}

const createIdentityNode = index => (
  myMasterNode.getIdentityNode(index)
)

const myIdentityNode = createIdentityNode(0)

console.log(`
==== Identity 1 Node (You can make many) ====
- Address:
${myIdentityNode.getAddress()}
- PrivateKey:
${myIdentityNode.getSKHex()}
`)


///////////////////////
// Apps Node Generation
///////////////////////

class AppsNode extends HdNode {
  constructor(node){
    super(node)
  }
  getAppNode(appDomain) {
    const hash = crypto
      .createHash('sha256')
      .update(appDomain)
      .digest('hex')
    const appIndex = hashCode(hash)
    const appNode = this.node.deriveHardened(appIndex)
    return new AppNode(appNode, appDomain)
  }
}

const myAppsNode = myIdentityNode.getAppsNode()

console.log(`
==== Identity 1's Apps Node  ====
- Address:
${myAppsNode.getAddress()}
- PrivateKey:
${myAppsNode.getSKHex()}
`)


///////////////////////
// App Node Generation
///////////////////////

class AppNode extends HdNode {
  constructor(node, appDomain) {
    super(node)
  }
  getAppDomain() {
    return this.appDomain
  }
}

const createAppNode = domain => (
  myAppsNode.getAppNode(domain)
)

const myAppNode = createAppNode('https://www.foo.com')

console.log(`
==== Identity 1's "https://www.foo.com" App Node  ====
- Address:
${myAppsNode.getAddress()}
- PrivateKey:
${myAppsNode.getSKHex()}
`)
