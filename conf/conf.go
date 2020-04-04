package conf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/aaaasmile/live-blog/crypto"
)

type Config struct {
	ServiceURL        string
	RootURLPattern    string
	UseRelativeRoot   bool
	DebugVerbose      bool
	UploadDir         string
	PrivatSecretFname string
	AdminCred         *crypto.UserCred
}

var Current = &Config{}

func ReadConfig(configfile string, ac *crypto.UserCred) *Config {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := toml.DecodeFile(configfile, &Current); err != nil {
		log.Fatal(err)
	}
	Current.AdminCred = ac
	log.Println("User configured: ", Current.AdminCred.String())
	return Current
}
