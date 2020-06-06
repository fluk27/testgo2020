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

// ReadPemFilePrivateKey is fuction read RSA key from pravateKey from RSA key in file .pem
func (RSAKey) ReadPemFilePrivateKey() rsa.PrivateKey {
	pemPrivateKey, err := ioutil.ReadFile("privateKey.pem")
	if err != nil {
		log.Fatal(err)
	}

	rsaKey := PemtoPrivateKeyOfRSAKey(pemPrivateKey)
	return rsaKey
}

//ReadPemFilePublicKey is fuction read RSA key from publicKey from RSA key in file .pem
func (RSAKey) ReadPemFilePublicKey() rsa.PublicKey {
	pemPublicKey, err := ioutil.ReadFile("publicKey.pem")
	if err != nil {
		log.Fatal(err)
	}

	rsaKey := PemtoPublicKeyOfRSAKey(pemPublicKey)
	return *rsaKey
}

//PemtoPrivateKeyOfRSAKey is fuction cover RSA key file pem to RSA key
func PemtoPrivateKeyOfRSAKey(pemPrivateKey []byte) rsa.PrivateKey {

	block, _ := pem.Decode([]byte(pemPrivateKey))
	if block == nil {
		// return nil, errors.New("failed to parse PEM block containing the key")
		log.Fatalln("from block :")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalln("from pemtoPrivateKey", err)
	}
	fmt.Println("private key from pem", priv)
	return *priv
}

//PemtoPublicKeyOfRSAKey is fuction cover RSA key file pem to RSA key
func PemtoPublicKeyOfRSAKey(pemPublicKey []byte) *rsa.PublicKey{

	block, _ := pem.Decode([]byte(pemPublicKey))
	if block == nil {
		// return nil, errors.New("failed to parse PEM block containing the key")
		log.Fatalln("from block :")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
           
    }

    switch pub := pub.(type) {
    case *rsa.PublicKey:
            return pub
    default:
            break // fall through
	}
	return nil
	// fmt.Println("PublicKey key from pem", public)
	// return *public
}

// EncyptData is fuctions encyption RSA key with publicKey
func (RSA RSAKey) EncyptData(data string) {
	PublicKey := RSA.ReadPemFilePublicKey()
	resultCipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &PublicKey, []byte(data), nil)
	if err != nil {
		log.Fatalln(err)
	}
	sEnc := base64.StdEncoding.EncodeToString([]byte(resultCipherText))
	fmt.Println("CipherText:", sEnc)

}

//EncyptDataWithPKC is function encryptData with publicKey of RSA key type PKC
func (RSA RSAKey) EncyptDataWithPKC() {
	PublicKey := RSA.ReadPemFilePublicKey()
	resultEncrypt, err := rsa.EncryptPKCS1v15(rand.Reader, &PublicKey, []byte("test"))
	if err != nil {
		log.Fatalln(err)
	}
	sEnc := base64.StdEncoding.EncodeToString(resultEncrypt)
	fmt.Println("CipherText:", sEnc)
}

// DncyptDataWithPKC is function decrypt cipherText to PlanText with RSA key type PKC#1
func (RSA RSAKey) DncyptDataWithPKC()  {
	PrivateKey := RSA.ReadPemFilePrivateKey()
	resultEncrypt:="XLCkx4SfLPpxC2MPgZxlRDh7tEGWFBR2W88NFW4szey9Kl/MLDpSdBcLZUO8YIOu+tdvzuZxp+V50ibeWdvNj7zEfrid5SWKzdpyQvHPbKGSZ9iAn3jZzrAy1B3QWLZUlWwh10NBnshzP6iYXdfUmxJQ3+DKpqZpufLCSdc2amO4J2qrSoqKUPA0PeIQsBt+iD3bejMTDoRj81+oGwmtlNEyt+wYwCenLgB1sBVTlPQGAQKZD9k8L/M5JTdd/5jGYIKucfz3gsSqJ8ArwJLatkNEFxdFKCHdg6Iq6fuuEj+mjvNUCOWb1EwdFZMxTUrgit6GFQU5dcVsyGkY3wHe7w=="
	sEnc ,err:= base64.StdEncoding.DecodeString(resultEncrypt)
	if err != nil {
		log.Fatalln("err from decrypt base64:",err)
	}
	fmt.Println("CipherText:", sEnc)
resultToMe,err:=rsa.DecryptPKCS1v15(rand.Reader,&PrivateKey,sEnc)
if err != nil {
	log.Fatalln("err from DecryptPKCS1v15:",err)
}
fmt.Println("text is decrypt :",string(resultToMe))
}