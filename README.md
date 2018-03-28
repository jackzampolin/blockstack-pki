# Blockstack PKI

This repo contains `js` code to help understand blockstack key derivation and hopefully eventually code examples in multiple languages.

> NOTE: This is experimental software. Use at your own risk!

To run the golang code, have golang installed and go get the following packages:

```
go get github.com/btcsuite/btcd/chaincfg
go get github.com/btcsuite/btcutil/hdkeychain
go get github.com/landonia/crypto/bip39
```

Then you can run `go run main.go`

```
$ go run main.go 
Generating a new mnomonic and then deriving keys with explanations
  mnemonic: slam cheap sponsor average issue lemon nuclear file gesture snake other seminar
  Root Private Key hardened: fc54af64efaac5513ae1541796b395473908c6120e5062e80ea5b8c086c1758701
  Browser Bitcoin Wallet Address: 1KeCXWDYHhpVZ9yatkDZnVg6qh2gnbAtKJ
  Browser Bitcoin Wallet PrivateKey: c83d3062d97e3e66d0b3899534777fbea289612e334e028e2b1a14a937727da201
  ID0 Identity Address: 1EWreS2xiBAB4BnZnFMo6StdTjGXsanDgf
  ID0 Identity PKHex: 13b8777d2787e8fa666b260eafb36d7808277ed9fc855087b40c4f34430c8e001
  ID1 Identity Address: 138orRYDsumLdchNmxgA6BnzYXZa6XMr9v
  ID1 Identity PKHex: e54dfe9ede902f4905a32d3e5660d8adb36060180104717e4de826a76a22445a01
```

Paste the mnemonic into line 25 of `index.js` to ensure that we are getting the same keys:

To run, `npm install` and then `node index.js`

```
$ node index.js 
Generating a new mnomonic and then deriving keys with explanations
  mnemonic: slam cheap sponsor average issue lemon nuclear file gesture snake other seminar
  Root Private Key hardened: fc54af64efaac5513ae1541796b395473908c6120e5062e80ea5b8c086c1758701
  Browser Bitcoin Wallet Address: 1KeCXWDYHhpVZ9yatkDZnVg6qh2gnbAtKJ
  Browser Bitcoin Wallet PrivateKey: c83d3062d97e3e66d0b3899534777fbea289612e334e028e2b1a14a937727da201
  ID0 Identity Address: 1EWreS2xiBAB4BnZnFMo6StdTjGXsanDgf
  ID0 Identity PKHex: 013b8777d2787e8fa666b260eafb36d7808277ed9fc855087b40c4f34430c8e001
  ID1 Identity Address: 138orRYDsumLdchNmxgA6BnzYXZa6XMr9v
  ID1 Identity PKHex: e54dfe9ede902f4905a32d3e5660d8adb36060180104717e4de826a76a22445a01
```

### Derivation Paths:

> NOTE: `'` denotes a hardened derivation:

For the Browser Bitcoin wallet the following derivation path is followed. 

```
Browser Wallet: 44'/0'/0'/0/0
Identity 0: 888'/0'/0'
```

### Open Questions

1. The javascript program generates 24 word mnemonics. How do we generate the 12 word mnemonics?
2. How are app keys derived? Add examples.

### Potential applications

1. Write a program that generates a seed, derives the first ID, issues a subdomain registrar call to register a subdomain to the first identity address and then sends the seed back to the user all the while not exposing the seed. This could be very useful for use on-boarding.
2. Write a program that derives the first `n` identity addresses, btc wallet and discovers all names associated as well as any BTC on the keychainPhrase
3. A utility for recovering BTC sent to an identity address
4. A utility for generating identities for headless devices, signing into applications, and then writing data to that application's storage. 
