/*
Copyright 2015 Steven Labrum

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package CheckIt

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"text/template"
)

var interfaces = []Box{}
var configuration = &Config{}
var boxes = []*BoxStruct{}

var (
	httpListen = flag.String("http", "127.0.0.1:3999", "host:port to listen on")
	htmlOutput = flag.Bool("html", false, "render program output as HTML")
)

func baseCase(w http.ResponseWriter, r *http.Request, page Page) {

	headTemp.Execute(w, nil)
	openBodyTemp.Execute(w, nil)
	pageStartTemp.Execute(w, page)

	boxes = initBoxes(interfaces)

	for key := range boxes {
		boxTemp.Execute(w, boxes[key])
	}

	pageCloseTemp.Execute(w, nil)
	htmlCloseTemp.Execute(w, nil)

}

/*  FrontPage is an HTTP handler that displays the basecase
unless a stored page is being loaded.
*/
func FrontPage(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]

	page, _ := initConfig(configuration)

	if len(title) < 1 {
		baseCase(w, r, page)
	} else {
		title := r.URL.Path[len("/"):]
		title = configuration.Path + "/" + title

		pageNames, _ := filepath.Glob(title + "/*.config")
		boxNames, _ := filepath.Glob(title + "/*.box")

		fmt.Println("Loaded shared page")

		if pageNames == nil || boxNames == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {

			pageName := pageNames[0]

			headTemp.Execute(w, nil)
			openBodyTemp.Execute(w, nil)

			configuration = ReadConfig(pageName)

			initConfig(configuration)

			pageStartTemp.Execute(w, page)

			for key := range boxNames {
				boxP := ReadBox(boxNames[key])
				boxes = append(boxes, boxP)
				boxTemp.Execute(w, boxP)
			}

			pageCloseTemp.Execute(w, nil)
			htmlCloseTemp.Execute(w, nil)
		}
	}
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	var aboutPage = Page{Heading: "About"}

	_, about := initConfig(configuration)

	headTemp.Execute(w, nil)
	openBodyTemp.Execute(w, nil)
	pageStartTemp.Execute(w, aboutPage)
	aboutTemp.Execute(w, about)
	pageCloseTemp.Execute(w, nil)
	htmlCloseTemp.Execute(w, nil)

}

var outputText = `<pre>{{printf "%s" .}}</pre>`
var output = template.Must(template.New("output").Parse(outputText))
var shareText = `{{printf "%s" . |html}}`
var shareOutput = template.Must(template.New("shareOutput").Parse(shareText))

// Compile is an HTTP handler that reads Source code from the request,
// runs the program (returning any errors),
// and sends the program's output as the HTTP response.
func PipeCompile(w http.ResponseWriter, req *http.Request) {

	title := req.URL.Path[len("/pipeile/"):]

	str := strings.Split(title, "/")

	position, _ := strconv.Atoi(str[1])

	body := new(bytes.Buffer)

	if _, err := body.ReadFrom(req.Body); err != nil {
		return
	}

	var textboxes []string

	if err := json.Unmarshal(body.Bytes(), &textboxes); err != nil {
		panic(err)
	}

	updateBody(boxes, textboxes)

	out, err := InterfaceRun(interfaces[position-1], textboxes)

	if err != nil {
		w.WriteHeader(404)
		output.Execute(w, out)
	} else if *htmlOutput {
		w.Write(out)
	} else {
		output.Execute(w, out)
	}
}

func sharHandler(w http.ResponseWriter, req *http.Request) {

	body := new(bytes.Buffer)

	if _, err := body.ReadFrom(req.Body); err != nil {
		return
	}

	var textboxes []string

	if err := json.Unmarshal(body.Bytes(), &textboxes); err != nil {
		panic(err)
	}

	page, _ := initConfig(configuration)

	out := Share(page)

	mux := &sync.Mutex{}
	mux.Lock()
	updateBody(boxes, textboxes)
	Save(configuration.Path+"/", out)
	mux.Unlock()
	shareOutput.Execute(w, out)
}

func initConfig(config *Config) (Page, AboutStruct) {
	var about = AboutStruct{}
	var page = Page{}

	configuration = config
	page.Heading = config.Heading
	about.Text = config.About
	about.SecondaryText = config.AboutSide

	return page, about
}

func initBoxes(boxs []Box) (boxes []*BoxStruct) {

	for key := range boxs {

		var box = BoxStruct{}
		heading, description, text, syntax := boxs[key].Desc()

		box.Id = strconv.Itoa(key)
		box.Position = strconv.Itoa(key + 1)
		box.Total = len(boxs)
		box.Lang = syntax

		box.Body = text
		box.Text = description
		box.Head = heading
		boxes = append(boxes, &box)
	}
	return boxes
}

var root string

func Serve(config *Config, boxs ...Box) (err error) {
	configuration = config
	interfaces = boxs

	root, err = os.Getwd()

	fmt.Println("Server Hosted")
	http.HandleFunc("/share/", sharHandler)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/", FrontPage)
	http.HandleFunc("/compile/", PipeCompile)
	http.ListenAndServe(configuration.Port, nil)

	err = errors.New("Server crashed")
	return err
}
