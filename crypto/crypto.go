package crypto

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"

	"golang.org/x/crypto/argon2"
)

type UserCred struct {
	UserName     string
	PasswordHash string
	Salt         string
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
		Salt         string
	}{
		UserName:     uc.UserName,
		PasswordHash: uc.PasswordHash,
		Salt:         uc.Salt,
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

	salt := make([]byte, 32)
	io.ReadFull(rand.Reader, salt)

	uc.UserName = user
	uc.PasswordHash = hashPassword(pwd, salt)
	uc.Salt = fmt.Sprintf("%x", salt)

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
		Salt         string
	}{}

	err = json.NewDecoder(f).Decode(&cred)
	uc.UserName = cred.UserName
	uc.PasswordHash = cred.PasswordHash
	return err
}

func hashPassword(pwd string, salt []byte) string {
	ram := 512 * 1024
	t0 := time.Now()
	key := argon2.IDKey([]byte(pwd), salt, 1, uint32(ram), uint8(runtime.NumCPU()<<1), 32)
	log.Printf("hash time: %v, key: %x\n", time.Since(t0), key)
	return fmt.Sprintf("%x", key)
}
