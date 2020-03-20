package crypto

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type UserCred struct {
	UserName     string
	PasswordHash string
	CredFile     string
}

func NewUserCred() *UserCred {
	res := UserCred{
		CredFile: "./cert/cred.json",
	}
	return &res
}

func (uc *UserCred) CreateAdminCredentials() error {
	err := uc.CredFromFile()
	if err != nil {
		log.Println("Missed od malformed credential: ", err)
		if err := uc.credFromPrompt(); err != nil {
			return err
		}
		return uc.saveCredential()
	}
	return fmt.Errorf("Unable to create a new crediantial. File %s already exist. Please delete it if you want a new crendential", uc.CredFile)
}

func (uc *UserCred) saveCredential() error {
	path := uc.CredFile
	log.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	cred := struct {
		UserName     string
		PasswordHash string
	}{
		UserName:     uc.UserName,
		PasswordHash: uc.PasswordHash,
	}

	return json.NewEncoder(f).Encode(cred)
}

func (uc *UserCred) credFromPrompt() error {
	var user, pwd, pwdcfm string
	fmt.Println("Please enter the username")
	fmt.Scanln(&user)
	//fmt.Println("*** user: ", []byte(user))

	fmt.Println("Please enter the password")
	fmt.Scanln(&pwd)
	if len(pwd) < 6 {
		return fmt.Errorf("Password must be al least 8 character lenght")
	}

	fmt.Println("Please confirm the password")
	fmt.Scanln(&pwdcfm)
	if pwd != pwdcfm {
		return fmt.Errorf("Passowrd confirm is different")
	}

	uc.UserName = user
	uc.PasswordHash = pwd

	return nil
}

func (uc *UserCred) CredFromFile() error {
	file := uc.CredFile
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	cred := struct {
		UserName     string
		PasswordHash string
	}{}

	err = json.NewDecoder(f).Decode(&cred)
	uc.UserName = cred.UserName
	uc.PasswordHash = cred.PasswordHash
	return err
}
