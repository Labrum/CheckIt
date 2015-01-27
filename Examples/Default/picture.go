package main

import (
	"encoding/base64"
	"github.com/Labrum/CheckIt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type pic struct{}

func dot(dotStr string, filename string) ([]byte, error) {

	dotFilename := filepath.Join("", filename)
	if err := ioutil.WriteFile(dotFilename, []byte(dotStr), 0666); err != nil {
		return nil, err
	}
	dotPicFilename := filepath.Join("", filename)
	timeout, _ := time.ParseDuration("2m")
	return CheckIt.CombinedRun(timeout, "dot", dotFilename, "-Tpng", "-o", dotPicFilename)
}

func (l *pic) Run(TextAreas []string) (out []byte, err error) {

	fName := "my"
	fType := "png"
	out, err = dot(TextAreas[1], fName+"."+fType)
	if err != nil {
		return out, err
	}

	f, _ := os.Open(fName + "." + fType)
	out, err = ioutil.ReadAll(f)

	imageString := base64.StdEncoding.EncodeToString(out)
	image := []byte(`<img src="data:image/` + fType + ";base64, " + imageString + `">`)
	return image, err
}

func (l *pic) Desc() (heading string, description string, text string, syntax string) {
	heading = "Picture"
	description = "This textbox uses the command line to list files"
	text = `digraph G {Hello->World}`
	syntax = "cmd"
	return heading, description, text, syntax
}
