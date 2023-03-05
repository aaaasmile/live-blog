package live

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/aaaasmile/live-blog/conf"
)

func handleList(w http.ResponseWriter, req *http.Request) error {
	log.Println("Handle List ", req.Header["Authorization"])
	user, err := isRequestAuthorized(req)
	if err != nil {
		return err
	}
	if user == "" {
		return writeErrorResponse(w, 403, "")
	}
	log.Println("User is autorized", user)

	rawbody, err := io.ReadAll(req.Body)
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
	files, err := os.ReadDir(targetPath)
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
