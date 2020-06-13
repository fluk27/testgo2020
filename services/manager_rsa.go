package services

import (
	"crypto"
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
	"os"
)

// RSAKey is all fuction use RSA key
type RSAKey struct {
	PathPrivateKey     string
	FileNamePrivateKey string
	PathPublicKey      string
	FileNamePublicKey  string
}

// GenerateRSAKey is GenerateRSAKey
func (rsak *RSAKey) GenerateRSAKey(bits int) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		// return nil, err
	}
	exportRSAKey(privateKey)
}

// exportRSAKey is function export privatekey and publickey
func exportRSAKey(PrivateKey *rsa.PrivateKey) (string, error) {

	err := RSAKeyTopemFile(PrivateKey)
	if err != nil {
		return "", err
	}
	return "ceate rsa key sccessfully", nil
}

// RSAKeyTopemFile is fuction cover RSA key to pem file
func RSAKeyTopemFile(PrivateKey *rsa.PrivateKey) error {
	privkeyBytes := x509.MarshalPKCS1PrivateKey(PrivateKey)
	privkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkeyBytes,
		},
	)

	err := ioutil.WriteFile(RSAKey{}.PathPrivateKey+"privateKey.pem", privkeyPem, 0644)
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
	err = ioutil.WriteFile(RSAKey{}.PathPublicKey+"publicKey.pem", publickeyPem, 0644)
	return errors.New("err from publickey.pem")
}

// ReadPemFilePrivateKey is fuction read RSA key from pravateKey from RSA key in file .pem
func (rsak *RSAKey) ReadPemFilePrivateKey() (*rsa.PrivateKey, error) {
	fmt.Println("path", rsak.PathPrivateKey+rsak.FileNamePrivateKey)
	pemPrivateKey, err := ioutil.ReadFile(rsak.PathPrivateKey + rsak.FileNamePrivateKey)
	if err != nil {
		// log.Fatalln("error read:")
		return nil, errors.New(err.Error())
	}

	rsaKey, err := PemtoPrivateKeyOfRSAKey(pemPrivateKey)
	if err != nil {
		// log.Fatalln("error covert:")
		return nil, errors.New(err.Error())
	}
	return rsaKey, nil
}

//ReadPemFilePublicKey is fuction read RSA key from publicKey from RSA key in file .pem
func (rsak *RSAKey) ReadPemFilePublicKey() (*rsa.PublicKey, error) {
	fmt.Println("file and path:", rsak.PathPublicKey+rsak.FileNamePublicKey)
	pemPublicKey, err := ioutil.ReadFile(rsak.PathPublicKey + rsak.FileNamePublicKey)
	if err != nil {
		// log.Fatal(err)
		return nil, errors.New("error ReadFile publicKey:" + err.Error())
	}

	rsaKey, err := PemtoPublicKeyOfRSAKey(pemPublicKey)
	if err != nil {
		return nil, errors.New("error ReadFile PemtoPublicKeyOfRSAKey:" + err.Error())
	}
	return rsaKey, nil
}

//PemtoPrivateKeyOfRSAKey is fuction cover RSA key file pem to RSA key
func PemtoPrivateKeyOfRSAKey(pemPrivateKey []byte) (*rsa.PrivateKey, error) {

	block, _ := pem.Decode(pemPrivateKey)
	if block == nil {
		// return nil, errors.New("failed to parse PEM block containing the key")
		return nil, errors.New("decode pem error")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalln("x509.ParsePKCS1PrivateKey:")
		return nil, errors.New(err.Error())
	}
	return priv, nil
}

//PemtoPublicKeyOfRSAKey is fuction cover RSA key file pem to RSA key
func PemtoPublicKeyOfRSAKey(pemPublicKey []byte) (*rsa.PublicKey, error) {

	block, _ := pem.Decode([]byte(pemPublicKey))
	if block == nil {
		//log.Fatalln("from block :")
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, errors.New("error ParsePKCS1PublicKey:" + err.Error())
		}
		return pub, nil
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall throughc

	}
	return nil, nil
}

// EncyptDataWithOAEP is fuctions encyption RSA key with publicKey of RSA key type OAEP
func (rsak *RSAKey) EncyptDataWithOAEP(data string) (string, error) {
	PublicKey, err := rsak.ReadPemFilePublicKey()
	if err != nil {
		return "", err
	}
	resultCipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, PublicKey, []byte(data), nil)
	if err != nil {
		log.Fatalln(err)
		return "", errors.New(err.Error())
	}
	sEnc := base64.StdEncoding.EncodeToString([]byte(resultCipherText))
	// fmt.Println("CipherText:", sEnc)
	return sEnc, nil

}

//EncyptDataWithPKC is function encryptData with publicKey of RSA key type PKC
func (rsak *RSAKey) EncyptDataWithPKC(password string, sign string) (string, error) {
	PublicKey, err := rsak.ReadPemFilePublicKey()
	if err != nil {
		return " ", err
	}

	resultEncrypt, err := rsa.EncryptPKCS1v15(rand.Reader, PublicKey, []byte(sign))
	if err != nil {
		log.Fatalln(err)
	}
	if sign != "" {
		rsak.SignatureMassage(sign)
	}

	sEnc := base64.StdEncoding.EncodeToString(resultEncrypt)
	// fmt.Println("CipherText:", sEnc)
	return sEnc, nil
}

// DncyptDataWithPKC is function decrypt cipherText to PlanText with RSA key type PKC#1
func (rsak *RSAKey) DncyptDataWithPKC(cipherText string) (string, error) {
	PrivateKey, err := rsak.ReadPemFilePrivateKey()
	sEnc, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", errors.New("err from decrypt base64 555:" + err.Error())
	}
	fmt.Println("CipherText:", sEnc)
	resultToMe, err := rsa.DecryptPKCS1v15(rand.Reader, PrivateKey, sEnc)
	if err != nil {
		return "", errors.New("err from DecryptPKCS1v15:" + err.Error())
	}
	fmt.Println("text is decrypt :", string(resultToMe))
	// rsak.verifyCipherText(resultToMe)
	return string(resultToMe), nil
}

// SignatureMassage is function sign data messsage
func (rsak *RSAKey) SignatureMassage(message string) (string, error) {
	PrivateKey, err := rsak.ReadPemFilePrivateKey()
	if err != nil {
		return "", errors.New(err.Error())
	}
	hashed := sha256.Sum256([]byte(message))

	signature, err := rsa.SignPKCS1v15(rand.Reader, PrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return "", errors.New("SignPKCS1v15 :" + err.Error())
	}

	sEnc := base64.StdEncoding.EncodeToString(signature)
	fmt.Println("Signature:", sEnc)
	return sEnc, nil
}

// verifyCipherText is function verify betaween data and sign(hash form sha 256)
func (rsak *RSAKey) verifyCipherText(sEnc string, massage string) error {
	//verify
	PublicKey, err := rsak.ReadPemFilePublicKey()
	if err != nil {
		log.Fatalln("ReadPemFilePublicKey from verify :", err)
		// return " ", err
	}

	signature, err := base64.StdEncoding.DecodeString(sEnc)
	if err != nil {

		// log.Fatalln("err from decrypt base64 :" + err.Error())
		return errors.New("err from decrypt base64 :" + err.Error())
	}

	hashedData := sha256.Sum256([]byte(massage))

	err = rsa.VerifyPKCS1v15(PublicKey, crypto.SHA256, hashedData[:], signature)
	if err != nil {
		// log.Fatalln("Error from verification:", err.Error())
		return errors.New("VerifyPKCS1v15:" + err.Error())
	}

	fmt.Println("verify successfully")
	return nil
}
