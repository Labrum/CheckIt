package main

import (
	"github.com/Labrum/CheckIt"
	"time"
	"path/filepath"
	"io/ioutil"
	"errors"
	"encoding/base64"
	"os"
)

type pic struct{}

func dot(dotStr string,filename string) ([]byte, error) {
	
	dotFilename := filepath.Join("", filename)
	if err := ioutil.WriteFile(dotFilename, []byte(dotStr), 0666); err != nil {
		return nil, err
	}
	dotPicFilename := filepath.Join("", filename)
	return CheckIt.CombinedRun("dot", dotFilename, "-Tpng", "-o", dotPicFilename)
}


func (l *pic) Run(TextAreas []string, directory string, timeout time.Duration) (out []byte, err error) {

	fName := "my"
	fType := "png"
	out,err = dot(TextAreas[1],fName+"."+fType)

	f,_ := os.Open(fName+"."+fType)
	out,err = ioutil.ReadAll(f)

	imageString := base64.StdEncoding.EncodeToString(out)
	image := []byte("data:image/"+fType+";base64, "+imageString)
	err = errors.New("Picture")
	return image,err
}

func (l *pic) Descriptors() (string,string) {
	title := "Picture"
	description :="This textbox uses the command line to list files"
	return description,title
}

func (l *pic) Default() (string,time.Duration) {
	return `digraph G {Hello->World}`, 10000
}

func (l *pic) Syntax() string {
	return "cmd"
}
