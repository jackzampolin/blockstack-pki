# Blockstack PKI

This repo contains `js` and `go` code to help understand Blockstack key derivation and provide example code for working with the keychain in multiple languages.

To run the golang code, have golang installed and go get the following packages:

```
go get github.com/btcsuite/btcd/
go get github.com/btcsuite/btcutil/
go get github.com/landonia/crypto/bip39
```

Then you can run `go run main.go`. 

> NOTE: This run uses a preset mnemonic to show that the two approaches are equivalent

```
$ go run main.go 
Using saved mnemonic to show equality with node.js example code

==== Mnemonic ====
slam cheap sponsor average issue lemon nuclear file gesture snake other seminar

==== Master Node ====
  - Address: 1AMD7hLvWVLwTJeCcSzMB7zLSLsC2NiV6V
  - PrivKey: fc54af64efaac5513ae1541796b395473908c6120e5062e80ea5b8c086c1758701

==== Wallet Node ====
  - Address: 1KeCXWDYHhpVZ9yatkDZnVg6qh2gnbAtKJ
  - PrivKey: c83d3062d97e3e66d0b3899534777fbea289612e334e028e2b1a14a937727da201

==== Identity Node 0 ====
  - Address: 1EWreS2xiBAB4BnZnFMo6StdTjGXsanDgf
  - PrivKey: 13b8777d2787e8fa666b260eafb36d7808277ed9fc855087b40c4f34430c8e001

==== Identity Node 1 ====
  - Address: 138orRYDsumLdchNmxgA6BnzYXZa6XMr9v
  - PrivKey: e54dfe9ede902f4905a32d3e5660d8adb36060180104717e4de826a76a22445a01

==== Identity Node 0 Apps Node ====
  - Address: 16Nru7xoU9yBhJVxeN7WjreuogT7rBBHXB
  - PrivKey: 8d0f4e7287659b26f830e89c7f8840c43cbae1d68aa7d6926778ba5dfa002baa01

==== Identity Node 0 App Node "https://www.foo.com" ====
  - Address: 1HzcgSVZEzeb7nPPZmgwFnij1ypsahT54w
  - PrivKey: 6b330358f01eda7121ba3f2d7e280212ae6fec9eb11350b4d49151760a1a418001
```

To run the `node.js` code first run `npm install`, then `node index.js`:

```
$ node index.js
Using saved mnemonic to show equality with golang example code

==== Mnemonic ====
slam cheap sponsor average issue lemon nuclear file gesture snake other seminar

==== Master Node ====
  - Address: 1AMD7hLvWVLwTJeCcSzMB7zLSLsC2NiV6V
  - PrivKey: fc54af64efaac5513ae1541796b395473908c6120e5062e80ea5b8c086c1758701

==== Wallet Node ====
  - Address: 1KeCXWDYHhpVZ9yatkDZnVg6qh2gnbAtKJ
  - PrivKey: c83d3062d97e3e66d0b3899534777fbea289612e334e028e2b1a14a937727da201

==== Identity Node 0 ====
  - Address: 1EWreS2xiBAB4BnZnFMo6StdTjGXsanDgf
  - PrivKey: 013b8777d2787e8fa666b260eafb36d7808277ed9fc855087b40c4f34430c8e001

==== Identity Node 1 ====
  - Address: 138orRYDsumLdchNmxgA6BnzYXZa6XMr9v
  - PrivKey: e54dfe9ede902f4905a32d3e5660d8adb36060180104717e4de826a76a22445a01

==== Identity Node 0 Apps Node ====
  - Address: 16Nru7xoU9yBhJVxeN7WjreuogT7rBBHXB
  - PrivKey: 8d0f4e7287659b26f830e89c7f8840c43cbae1d68aa7d6926778ba5dfa002baa01

==== Identity Node 0 App Node "https://www.foo.com" ====
  - Address: 1HzcgSVZEzeb7nPPZmgwFnij1ypsahT54w
  - PrivKey: 6b330358f01eda7121ba3f2d7e280212ae6fec9eb11350b4d49151760a1a418001
```

### Derivation Paths:

> NOTE: `'` denotes a hardened derivation:

For the Browser Bitcoin wallet the following derivation paths are followed.

```
Master                    m
Browser Wallet:           m/44'/0'/0'/0/0
Identity:                 m/888'/0'/{{ id_index }}'
Identity Encryption Node: m/888'/0'/{{ id_index }}'/1'
Identity Signing Node:    m/888'/0'/{{ id_index }}'/2'
Identity Apps Key:        m/888'/0'/{{ id_index }}'/0'
Identity App Key:         m/888'/0'/{{ id_index }}'/0'/{{ app_index }}'
```

### Open Questions

1. The javascript program generates 24 word mnemonics. How do we generate the 12 word mnemonics?
2. The 0th identity in the Golang program omits leading `0`s. This may need to be fixed. Ask @kantai
