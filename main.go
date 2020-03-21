package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aaaasmile/live-blog/crypto"
	"github.com/aaaasmile/live-blog/web"
	"github.com/aaaasmile/live-blog/web/idl"
)

func main() {
	var ver = flag.Bool("ver", false, "Prints the current version")
	var configfile = flag.String("config", "config.toml", "Configuration file path")
	var initAccount = flag.Bool("init-account", false, "Initialize the Admin account credentials")
	flag.Parse()

	if *ver {
		fmt.Printf("%s, version: %s", idl.Appname, idl.Buildnr)
		os.Exit(0)
	}

	uc := crypto.NewUserCred()
	if *initAccount {
		if err := uc.CreateAdminCredentials(); err != nil {
			log.Fatal("Error: ", err)
		}
		log.Println("Credential successfully created. Please restart.")
		os.Exit(0)
	}
	if err := uc.CredFromFile(); err != nil {
		log.Fatal("Credential error. Please make sure that an account has been initiated. Error is: ", err)
	}

	web.RunService(*configfile, uc)
}
