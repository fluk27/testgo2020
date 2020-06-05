package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

// RSAKey is all fuction use RSA
type RSAKey struct{}

// GenerateRSAKey is GenerateRSAKey
func (rsak RSAKey) GenerateRSAKey(bits int) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		// return nil, err
	}
	rsak.exportRSAKey(privateKey)
}

// exportRSAKey is function export privatekey and publickey
func (RSAKey) exportRSAKey(PrivateKey *rsa.PrivateKey) (string, error) {

	err := rsaKeyTopemFile(PrivateKey)
	if err != nil {
		return "", err
	}
	return "ceate rsa key sccessfully", nil
}

func rsaKeyTopemFile(PrivateKey *rsa.PrivateKey) error {
	privkeyBytes := x509.MarshalPKCS1PrivateKey(PrivateKey)
	privkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkeyBytes,
		},
	)
	err := ioutil.WriteFile("privateKey.pem", privkeyPem, 0644)
	if err != nil {
		return errors.New("err from PrivateKey.pem")
	}
	// publicKey
	publickeyBytes := x509.MarshalPKCS1PublicKey(&PrivateKey.PublicKey)
	publickeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publickeyBytes,
		},
	)
	err = ioutil.WriteFile("publicKey.pem", publickeyPem, 0644)
	return errors.New("err from publickey.pem")
}

// ReadPemFile is fuction read RSA key from pravateKey from RSA key in file .pem
func (RSAKey) ReadPemFile() rsa.PrivateKey {
	pemPrivateKey, err := ioutil.ReadFile("privateKey.pem")
	if err != nil {
		log.Fatal(err)
	}

	rsaKey := PemtoRSAKey(pemPrivateKey)
	return rsaKey
}

func (RSAKey) ReadPemFilePublicKey() rsa.PrivateKey {
	pemPrivateKey, err := ioutil.ReadFile("public.pem.pem")
	if err != nil {
		log.Fatal(err)
	}

	rsaKey := PemtoRSAKey(pemPrivateKey)
	return rsaKey
}

//PemtoRSAKey is fuction cover RSA key file pem to RSA key
func PemtoRSAKey(pemPrivateKey []byte) rsa.PrivateKey {

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
	return *priv
}

// EncyptData is fuctions encyption RSA key with publicKey
func (RSA RSAKey) EncyptData(data string) {
	PublicKey := RSA.ReadPemFile()
	resultCipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &PublicKey.PublicKey, []byte(data), nil)
	if err != nil {
		log.Fatalln(err)
	}
	sEnc := base64.StdEncoding.EncodeToString([]byte(resultCipherText))
	fmt.Println("CipherText:", sEnc)

}

func (RSA RSAKey) EncyptDataWithPKC() {
	PublicKey := RSA.ReadPemFile()
	resultEncrypt, err := rsa.EncryptPKCS1v15(rand.Reader, &PublicKey.PublicKey, []byte("test"))
	if err != nil {
		log.Fatalln(err)
	}
	sEnc := base64.StdEncoding.EncodeToString(resultEncrypt)
	fmt.Println("CipherText:", sEnc)
}
