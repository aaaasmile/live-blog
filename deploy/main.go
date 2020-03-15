package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/aaaasmile/live-blog/deploy/depl"
)

var (
	defOutDir = "..\\..\\Deployed"
)

func main() {
	const (
		invido = "invido"
	)
	var outdir = flag.String("outdir", "",
		fmt.Sprintf("Output zip directory. If empty use the hardcoded one: %s\n", defOutDir))

	var target = flag.String("target", "",
		fmt.Sprintf("Target of deployment: %s", invido))

	flag.Parse()

	rootDirRel := ".."
	pathItems := []string{"live-blog.exe", "static", "templates", "cert"}
	switch *target {
	case invido:
		pathItems = append(pathItems, "deploy/config_files/invido_config.toml")
		//pathItems[0] = "live-blog.bin"
	default:
		log.Fatalf("Deployment target %s is not recognized or not specified", *target)
	}
	log.Printf("Create the zip package for target %s", *target)

	outFn := getOutFileName(*outdir, *target)
	depl.CreateDeployZip(rootDirRel, pathItems, outFn, func(pathItem string) string {
		if strings.HasPrefix(pathItem, "deploy/config_files") {
			return "config.toml"
		}
		return pathItem
	})
}

func getOutFileName(outdir string, tgt string) string {
	if outdir == "" {
		outdir = defOutDir
	}
	vn := depl.GetVersionNrFromFile("../web/idl/idl.go", "")
	log.Println("Version is ", vn)

	currentTime := time.Now()
	s := fmt.Sprintf("LiveBlog_%s_%s_%s.zip", strings.Replace(vn, ".", "-", -1), currentTime.Format("02012006-150405"), tgt) // current date-time stamp using 2006 date time format template
	s = filepath.Join(outdir, s)
	return s
}

func testGetVersion() {
	buf, err := ioutil.ReadFile("../web/idl/idl.go")
	if err != nil {
		log.Fatalln("Cannot read input file", err)
	}
	s := string(buf)
	fmt.Println(s)
	vn := depl.GetBuildVersionNr(s, "")
	if vn == "" {
		log.Fatalln("Version not found")
	}
	fmt.Println("Version is ", vn)
	//depl.TestLexer()
}
