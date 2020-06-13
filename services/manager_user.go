package services

//UserServices is all fuctions manager user
type UserServices struct {
}

//Register is function recvier data of user
func (US *UserServices)  Register(password string) (string, error) {
	RSAService := &RSAKey{}
	RSAService.FileNamePrivateKey = "privateKey.pem"
	RSAService.FileNamePublicKey = "publicKey.pem"
	RSAService.PathPublicKey = "./"
	cipherText, err := RSAService.EncyptDataWithPKC(password, "")
	if err != nil {
		return "", err
	}
	return cipherText, nil
}
