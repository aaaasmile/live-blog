package live

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/aaaasmile/live-blog/conf"
	"github.com/aaaasmile/live-blog/crypto"
)

func handleToken(w http.ResponseWriter, req *http.Request) error {
	rawbody, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	//fmt.Println("*** Request: ", string(rawbody))

	credReq := struct {
		Username string
		Password string
		UserHash string
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
		hash := ""
		if credReq.Password != "" && credReq.UserHash == "" {
			hash = crypto.GetHashOfSecret(credReq.Password, refCred.Salt)
		} else {
			hash = credReq.UserHash
		}
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
		resp.Info = "user credential OK"
		expires := 3600
		log.Printf("Create JWT Token for user %s, expires in %d", username, expires)
		refCred := conf.Current.AdminCred
		err := refCred.GetJWTToken(username, expires, &resp.Token)
		if err != nil {
			return err
		}
		return writeResponse(w, &resp)
	case 403:
		resp.Info = "user Unauthorized"
	default:
		resp.Info = "user credential ERROR"
	}

	return writeErrorResponse(w, resp.ResultCode, resp)
}

func checkRefreshToken(w http.ResponseWriter, refrTk string) error {

	if refrTk == "" {
		return fmt.Errorf("refresh token is empty")
	}
	if len(refrTk) > 10 {
		b := len(refrTk) - 1
		a := b - 10
		log.Println("Check for refresh token ", refrTk[a:b])
	}
	refCred := conf.Current.AdminCred
	user, err := refCred.ParseJwtTokenForRefresh(refrTk)
	if err != nil {
		return err
	}
	if user != "" {
		return tokenResult(200, user, w)
	}
	return tokenResult(403, user, w)
}

func isRequestAuthorized(req *http.Request) (string, error) {
	adminCred := conf.Current.AdminCred
	user := ""
	var err error
	auth_tk := req.Header.Get("Authorization")
	if auth_tk != "" {
		tk := strings.Split(auth_tk, " ")
		//fmt.Println("*** ", tk)
		if len(tk) == 2 {
			if tk[0] == "Bearer" {
				user, err = adminCred.ParseJwtTokenForAuth(tk[1])
				if err != nil {
					return "", err
				}
			}
		}
	}
	return user, nil
}
