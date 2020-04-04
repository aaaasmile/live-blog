package live

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aaaasmile/live-blog/conf"
	"github.com/aaaasmile/live-blog/crypto"
	"github.com/aaaasmile/live-blog/web/idl"
	"github.com/aaaasmile/live-blog/web/session"
)

type PageCtx struct {
	RootUrl string
	Buildnr string
}

func getURLForRoute(uri string) string {
	arr := strings.Split(uri, "/")
	//fmt.Println("split: ", arr, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		ss := arr[i]
		if ss != "" {
			if !strings.HasPrefix(ss, "?") {
				//fmt.Printf("Url for route is %s\n", ss)
				return ss
			}
		}
	}
	return uri
}

func APiHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case "GET":
		err = handleGet(w, req)
	case "POST":
		log.Println("POST on ", req.RequestURI)
		err = handlePost(w, req)
	}
	if err != nil {
		log.Println("Error on process request: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handlePost(w http.ResponseWriter, req *http.Request) error {
	start := time.Now()
	var err error
	lastPath := getURLForRoute(req.RequestURI)
	log.Println("Check the last path ", lastPath)
	switch lastPath {
	case "Login":
		err = handleLogin(w, req)
	default:
		return fmt.Errorf("%s method is not supported", lastPath)
	}

	if err != nil {
		return err
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Service total call duration: %v\n", elapsed)
	return nil
}

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
	}{
		ResultCode: resultCode,
		Username:   username,
	}

	switch resultCode {
	case 200:
		resp.Info = fmt.Sprintf("User login OK")
	case 210:
		resp.Info = fmt.Sprintf("User already logged in")
	case 403:
		resp.Info = fmt.Sprintf("User Unauthorized")
	default:
		resp.Info = fmt.Sprintf("User login ERROR")
	}

	return writeResponse(w, &resp)
}

func writeResponse(w http.ResponseWriter, resp interface{}) error {
	blobresp, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	w.Write(blobresp)
	return nil
}

func handleGet(w http.ResponseWriter, req *http.Request) error {
	u, _ := url.Parse(req.RequestURI)
	log.Println("GET requested ", u)

	pagectx := PageCtx{
		RootUrl: conf.Current.RootURLPattern,
		Buildnr: idl.Buildnr,
	}
	templName := "templates/vue/index.html"

	tmplIndex := template.Must(template.New("AppIndex").ParseFiles(templName))

	err := tmplIndex.ExecuteTemplate(w, "base", pagectx)
	if err != nil {
		return err
	}

	return nil
}
