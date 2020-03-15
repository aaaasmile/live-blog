package live

import (
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/aaaasmile/live-blog/conf"
	"github.com/aaaasmile/live-blog/web/idl"
)

type PageCtx struct {
	RootUrl string
	Buildnr string
}

func APiHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		handleGet(w, req)
	case "POST":
		log.Println("POST on ", req.RequestURI)
		handlePost(w, req)
	}
}

func handlePost(w http.ResponseWriter, req *http.Request) {
}

func handleGet(w http.ResponseWriter, req *http.Request) {
	u, _ := url.Parse(req.RequestURI)

	pagectx := PageCtx{
		RootUrl: conf.Current.RootURLPattern,
		Buildnr: idl.Buildnr,
	}
	templName := "templates/vue/index.html"

	tmplIndex := template.Must(template.New("AppIndex").ParseFiles(templName))

	err := tmplIndex.ExecuteTemplate(w, "base", pagectx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GET requested ", u)
}
