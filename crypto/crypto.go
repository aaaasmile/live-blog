package crypto

import (
	"crypto/rand"
	"encoding/base64"
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

func (uc *UserCred) CreateAccountCredentials() error {
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
	//fmt.Println("*** user: ", user)

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
	//fmt.Println("*** pwd: ", pwd)
	uc.PasswordHash = hashPassword(pwd, salt)
	uc.Salt = base64.StdEncoding.EncodeToString(salt)

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
	uc.Salt = cred.Salt
	return err
}

func (uc *UserCred) String() string {
	return fmt.Sprintf("Username: %s, Hash: %s, Salt %s", uc.UserName, uc.PasswordHash, uc.Salt)
}

func hashPassword(pwd string, salt []byte) string {
	ram := 512 * 1024
	t0 := time.Now()
	//fmt.Printf("*** salt is %x\n", salt)
	//fmt.Printf("*** password is %s\n", pwd)
	key := argon2.IDKey([]byte(pwd), salt, 1, uint32(ram), uint8(runtime.NumCPU()<<1), 32)
	log.Printf("hash time: %v, key: %x, salt: %x\n", time.Since(t0), key, salt)
	return fmt.Sprintf("%x", key)
}

func GetHashOfSecret(pwd, salt string) string {
	ss, _ := base64.StdEncoding.DecodeString(salt)
	return hashPassword(pwd, ss)
}
