package conf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/aaaasmile/live-blog/crypto"
)

type Config struct {
	ServiceURL      string
	RootURLPattern  string
	UseRelativeRoot bool
	DebugVerbose    bool
	UserCred        *crypto.UserCred
}

var Current = &Config{}

func ReadConfig(configfile string, uc *crypto.UserCred) *Config {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := toml.DecodeFile(configfile, &Current); err != nil {
		log.Fatal(err)
	}
	Current.UserCred = uc
	log.Println("User configured: ", Current.UserCred.String())
	return Current
}
