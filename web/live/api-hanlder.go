package live

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aaaasmile/live-blog/conf"
	"github.com/aaaasmile/live-blog/web/idl"
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
	was_auth := false
	switch req.Method {
	case "GET":
		err = handleGet(w, req)
	case "POST":
		log.Println("POST on ", req.RequestURI)
		was_auth, err = handlePost(w, req)
	}
	if err != nil {
		log.Println("Error on process request: ", err)
		err_info := "Internal Server Error"
		if was_auth {
			err_info = fmt.Sprintf("Internal Server Error: %s", err.Error())
		}
		http.Error(w, err_info, http.StatusInternalServerError)
	}
}

func handlePost(w http.ResponseWriter, req *http.Request) (bool, error) {
	start := time.Now()
	var err error
	was_auth := false
	lastPath := getURLForRoute(req.RequestURI)
	log.Println("Check the last path ", lastPath)
	switch lastPath {
	case "Token":
		err = handleToken(w, req)
	case "List":
		was_auth, err = handleList(w, req)
	default:
		return was_auth, fmt.Errorf("%s method is not supported", lastPath)
	}

	if err != nil {
		return was_auth, err
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Service total call duration: %v\n", elapsed)
	return was_auth, nil
}

func writeResponse(w http.ResponseWriter, resp interface{}) error {
	blobresp, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	w.Write(blobresp)
	return nil
}

func writeErrorResponse(w http.ResponseWriter, errorcode int, resp interface{}) error {
	blobresp, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	http.Error(w, string(blobresp), errorcode)
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
