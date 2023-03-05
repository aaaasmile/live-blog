package live

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/aaaasmile/live-blog/conf"
)

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case "GET":
		err = fmt.Errorf("unsupported GET method in upload")
	case "POST":
		log.Println("POST on ", req.RequestURI)
		err = handleUploadPost(w, req)
	}
	if err != nil {
		log.Println("Error on process request: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handleUploadPost(w http.ResponseWriter, req *http.Request) error {
	start := time.Now()
	var err error
	lastPath := getURLForRoute(req.RequestURI)
	log.Println("Check the last path ", lastPath)

	err = procUpload(w, req)
	if err != nil {
		return err
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Service total call duration: %v\n", elapsed)
	return nil
}

func procUpload(w http.ResponseWriter, req *http.Request) error {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	req.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := req.FormFile("myFile")
	if err != nil {
		return err
	}
	defer file.Close()
	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	now := time.Now()
	tempFile, err := ioutil.TempFile(conf.Current.UploadDir, fmt.Sprintf("%s-*", now.Format("2006-01-02")))
	if err != nil {
		return err
	}
	log.Println("Using the temp file: ", tempFile.Name())
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	log.Println("OK, file uploaded: ", tempFile.Name())
	// return that we have successfully uploaded our file!
	return nil
}
