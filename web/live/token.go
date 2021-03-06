package live

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aaaasmile/live-blog/conf"
	"github.com/aaaasmile/live-blog/crypto"
)

func handleToken(w http.ResponseWriter, req *http.Request) error {
	rawbody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	fmt.Println("*** Request: ", string(rawbody)) //do not log passwords

	credReq := struct {
		Username string
		Password string
	}{}
	if err := json.Unmarshal(rawbody, &credReq); err != nil {
		return err
	}
	log.Println("Token for user ", credReq.Username)

	if credReq.Username == "" {
		log.Println("Username is empty")
		refrToken := struct {
			Token string
		}{}
		if err := json.Unmarshal(rawbody, &refrToken); err != nil {
			return err
		}
		return checkRefreshToken(w, refrToken.Token)
	}

	refCred := conf.Current.AdminCred
	if credReq.Username == refCred.UserName {
		log.Println("Check password for user ", credReq.Username)
		//fmt.Println("*** refcred", refCred)
		hash := crypto.GetHashOfSecret(credReq.Password, refCred.Salt)
		//log.Println("Hash is: ", hash)
		if hash == refCred.PasswordHash {
			return tokenResult(200, credReq.Username, w)
		}
	}

	return tokenResult(403, credReq.Username, w)
}

func tokenResult(resultCode int, username string, w http.ResponseWriter) error {

	resp := struct {
		Info       string
		ResultCode int
		Username   string
		Token      crypto.Token
	}{
		ResultCode: resultCode,
		Username:   username,
	}

	switch resultCode {
	case 200:
		resp.Info = fmt.Sprintf("User credential OK")
		expires := 3600
		log.Printf("Create JWT Token for user %s, expires in %d", username, expires)
		refCred := conf.Current.AdminCred
		err := refCred.GetJWTToken(username, expires, &resp.Token)
		if err != nil {
			return err
		}
		return writeResponse(w, &resp)
	case 403:
		resp.Info = fmt.Sprintf("User Unauthorized")
	default:
		resp.Info = fmt.Sprintf("User credential ERROR")
	}

	return writeErrorResponse(w, resp.ResultCode, resp)
}

func checkRefreshToken(w http.ResponseWriter, refrTk string) error {

	if refrTk == "" {
		return fmt.Errorf("Refresh token is empty")
	}
	if len(refrTk) > 10 {
		b := len(refrTk) - 1
		a := b - 10
		log.Println("Check for refresh token ", refrTk[a:b])
	}
	refCred := conf.Current.AdminCred
	user, err := refCred.ParseJwtToken(refrTk)
	if err != nil {
		return err
	}
	if user != "" {
		return tokenResult(200, user, w)
	}
	return tokenResult(403, user, w)
}
