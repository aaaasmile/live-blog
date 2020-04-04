package live

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aaaasmile/live-blog/conf"
	"github.com/aaaasmile/live-blog/crypto"
	"github.com/aaaasmile/live-blog/web/session"
)

func handleLogin(w http.ResponseWriter, req *http.Request) error {
	rawbody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	if conf.Current.DebugVerbose {
		log.Println("Request: ", string(rawbody))
	}

	credReq := struct {
		Username string
		Password string
	}{}
	if err := json.Unmarshal(rawbody, &credReq); err != nil {
		return err
	}
	log.Println("Login for user ", credReq.Username)

	session, err := session.SessMgr.GetSession(w, req)
	if err != nil {
		return err
	}

	if session.Username != "" {
		if session.Username == credReq.Username { // only one user is supported
			return loginResult(210, credReq.Username, w)
		}
	} else {
		refCred := conf.Current.AdminCred
		if credReq.Username == refCred.UserName {
			log.Println("Check password for user ", credReq.Username)
			//fmt.Println("*** refcred", refCred)
			hash := crypto.GetHashOfSecret(credReq.Password, refCred.Salt)
			//log.Println("Hash is: ", hash)
			if hash == refCred.PasswordHash {
				session.Username = credReq.Username
				return loginResult(200, credReq.Username, w)
			}
		}
	}

	loginResult(403, credReq.Username, w)

	return nil
}

func loginResult(resultCode int, username string, w http.ResponseWriter) error {

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
		resp.Info = fmt.Sprintf("User login OK")
		expires := 3600 * 24
		log.Printf("Create JWT Token for user %s, expires in %d", username, expires)
		refCred := conf.Current.AdminCred
		err := refCred.GetJWTToken(username, expires, &resp.Token)
		if err != nil {
			return err
		}
		return writeResponse(w, &resp)
	case 210:
		resp.Info = fmt.Sprintf("User already logged in")
		return writeResponse(w, &resp)
	case 403:
		resp.Info = fmt.Sprintf("User Unauthorized")
	default:
		resp.Info = fmt.Sprintf("User login ERROR")
	}

	return writeErrorResponse(w, resp.ResultCode, resp)
}

func handleLogout(w http.ResponseWriter, req *http.Request) error {
	rawbody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	if conf.Current.DebugVerbose {
		log.Println("Logut Request: ", string(rawbody))
	}

	credReq := struct {
		Username string
	}{}
	if err := json.Unmarshal(rawbody, &credReq); err != nil {
		return err
	}
	log.Println("Logout for user ", credReq.Username)

	session, err := session.SessMgr.GetSession(w, req)
	if err != nil {
		return err
	}

	resp := struct {
		Info       string
		ResultCode int
		Username   string
	}{
		ResultCode: 403,
		Info:       "Logout error",
	}

	if session.Username != "" {
		if session.Username == credReq.Username { // only one user is supported
			resp.ResultCode = 200
			resp.Info = "Logout OK"
			resp.Username = credReq.Username
			return writeResponse(w, &resp)
		}
	}

	writeErrorResponse(w, resp.ResultCode, resp)
	return nil
}
