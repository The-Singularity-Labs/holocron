package main

import (
    "fmt"
	"io/ioutil"
	"flag"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
	"github.com/mikesmitty/edkey"
    "github.com/algorand/go-algorand-sdk/crypto"
    "github.com/algorand/go-algorand-sdk/mnemonic"
)

func main() {
	namePtr := flag.String("name", "foo", "any name")
	seedPhrase := flag.String("seed", "", "A 24 word mnemonic Algorand compatiable phrase")
	
	flag.Parse()
	name := *namePtr

	fmt.Printf("Generating key '%s', has_seed: %t\n", name, *seedPhrase != "")

    var account crypto.Account

	if *seedPhrase == "" {
		account = crypto.GenerateAccount()
	} else {
		m, err := mnemonic.ToPrivateKey(*seedPhrase)
		if err != nil {
			fmt.Printf("Error creating transaction: %v\n", err)
		}

		account, _ = crypto.AccountFromPrivateKey(m)
	}

	passphrase, err := mnemonic.FromPrivateKey(account.PrivateKey)
    if err != nil {
        fmt.Printf("Error creating mnemonic: %v\n", err)
    }

	publicKey, err := ssh.NewPublicKey(account.PublicKey)
    if err != nil {
        fmt.Printf("Error creating public key: %v\n", err)
    }

	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(account.PrivateKey),  // <- marshals ed25519 correctly
	}
	privateKey := pem.EncodeToMemory(pemKey)
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	_ = ioutil.WriteFile(fmt.Sprintf(".build/%s", name), privateKey, 0600)
	_ = ioutil.WriteFile(fmt.Sprintf(".build/%s.pub", name), authorizedKey, 0644)
	_ = ioutil.WriteFile(fmt.Sprintf(".build/%s.key", name), []byte(passphrase), 0644)
	fmt.Printf("New address: %s\n", account.Address)
}