package live

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/aaaasmile/live-blog/conf"
)

func handleList(w http.ResponseWriter, req *http.Request) (bool, error) {
	log.Println("Handle List ", req.Header["Authorization"])
	was_auth := false
	user, err := isRequestAuthorized(req)
	if err != nil {
		return was_auth, err
	}
	if user == "" {
		return was_auth, writeErrorResponse(w, 403, "")
	}
	was_auth = true
	log.Println("User is autorized", user)

	rawbody, err := io.ReadAll(req.Body)
	if err != nil {
		return was_auth, err
	}
	//fmt.Println("*** Request: ", string(rawbody))

	reqPayload := struct {
		Path string
	}{}
	if err := json.Unmarshal(rawbody, &reqPayload); err != nil {
		return was_auth, err
	}
	log.Println("List request for ", reqPayload.Path)
	err = listResult(reqPayload.Path, w)
	return was_auth, err
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
	log.Println("Path to read is: ", targetPath)
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
