package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type RSAKey struct{}

// GenerateRSAKey
func (rsak RSAKey) GenerateRSAKey(bits int) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		// return nil, err
	}

	unixTimeUTC := time.Unix(privateKey.D.Int64(), 0)
	fmt.Println("time:",unixTimeUTC)
	rsak.exportRSAKey(privateKey)
}

func (RSAKey) exportRSAKey(PrivateKey *rsa.PrivateKey) (string, error) {

	err := rsaKeyTopemFile(PrivateKey)
	if err != nil {
		return "", err
	}
	return "ceate rsa key sccessfully", nil
}

func rsaKeyTopemFile(PrivateKey *rsa.PrivateKey) error {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(PrivateKey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	err := ioutil.WriteFile("privateKey.pem", privkey_pem, 0644)
	if err != nil {
		return errors.New("err from PrivateKey.pem")
	}
	// publicKey
	publickey_bytes := x509.MarshalPKCS1PublicKey(&PrivateKey.PublicKey)
	publickey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publickey_bytes,
		},
	)
	err = ioutil.WriteFile("publicKey.pem", publickey_pem, 0644)
	return errors.New("err from publickey.pem")
}

func (RSAKey) ReadPemFile() {
	pemPrivateKey, err := ioutil.ReadFile("privateKey.pem")
	if err != nil {
		log.Fatal(err)
	}

	PemtoRSAKey(pemPrivateKey)
}

func PemtoRSAKey(pemPrivateKey []byte) {

	block, _ := pem.Decode([]byte(pemPrivateKey))
	if block == nil {
		// return nil, errors.New("failed to parse PEM block containing the key")
		log.Fatalln("from block :")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalln("from pemtokey", err)
	}

	fmt.Println("private key from pem", priv)
}
