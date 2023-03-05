package live

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/aaaasmile/live-blog/conf"
)

func handleList(w http.ResponseWriter, req *http.Request) error {
	log.Println("Handle List ", req.Header["Authorization"])
	user := ""
	var err error
	adminCred := conf.Current.AdminCred
	auth_tk := req.Header.Get("Authorization")
	if auth_tk != "" {
		tk := strings.Split(auth_tk, " ")
		//fmt.Println("*** ", tk)
		if len(tk) == 2 {
			if tk[0] == "Bearer" {
				user, err = adminCred.ParseJwtToken(tk[1])
				if err != nil {
					return err
				}
			}
		}
	}

	if user == "" {
		return writeErrorResponse(w, 403, "")
	}

	rawbody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	fmt.Println("*** Request: ", string(rawbody))

	reqPayload := struct {
		Path string
	}{}
	if err := json.Unmarshal(rawbody, &req); err != nil {
		return err
	}
	log.Println("List request for ", reqPayload.Path)

	return listResult(reqPayload.Path, w)
}

type ResType int

const (
	RTFile = iota
	RTDir
)

type Resource struct {
	Name string
	Type ResType
}

func listResult(reqPath string, w http.ResponseWriter) error {
	resp := struct {
		Resources []Resource
	}{
		Resources: make([]Resource, 0),
	}
	pp := path.Join(conf.Current.UploadDir, reqPath)
	targetPath, _ := filepath.Abs(pp)
	fmt.Println("path is: ", targetPath)
	files, err := ioutil.ReadDir(targetPath)
	if err != nil {
		return err
	}
	for _, f := range files {
		ri := Resource{
			Name: f.Name(),
			Type: RTFile,
		}
		itemAbs := path.Join(targetPath, f.Name())
		if info, err := os.Stat(itemAbs); err == nil && info.IsDir() {
			ri.Type = RTDir
		}
		resp.Resources = append(resp.Resources, ri)
	}

	return writeResponse(w, resp)
}
