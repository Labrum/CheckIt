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

package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

var outputs = [10]*CompileOut{}
var boxes = []*BoxStruct{}
var interfaces = []Box{}
var page = Page{}
var aboutPage = Page{}
var about = AboutStruct{}
var configuration = Config{}

var (
	httpListen = flag.String("http", "127.0.0.1:3999", "host:port to listen on")
	htmlOutput = flag.Bool("html", false, "render program output as HTML")
)

func baseCase(w http.ResponseWriter, r *http.Request) {

	//_, boxes = InitDefault()

	headTemp.Execute(w, nil)
	openBodyTemp.Execute(w, nil)
	pageStartTemp.Execute(w, page)

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

	if len(title) < 1 {
		baseCase(w, r)
	} else {
		title := r.URL.Path[len("/"):]

		pageNames, _ := filepath.Glob(title + "/*.page")
		boxNames, _ := filepath.Glob(title + "/*.box")

		fmt.Println(len(pageNames))
		fmt.Println("hello")

		if pageNames == nil || boxNames == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {

			pageName := pageNames[0]
			fmt.Println(boxNames)

			headTemp.Execute(w, nil)
			openBodyTemp.Execute(w, nil)

			p := ReadPage(pageName)
			pageStartTemp.Execute(w, p)

			boxes = []*BoxStruct{}
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
	//aboutPage, about := InitAbout()

	headTemp.Execute(w, nil)
	openBodyTemp.Execute(w, nil)
	pageStartTemp.Execute(w, aboutPage)
	aboutTemp.Execute(w, about)
	pageCloseTemp.Execute(w, nil)
	htmlCloseTemp.Execute(w, nil)

}

var outputText = `<pre>{{printf "%s" . |html}}</pre>`
var output = template.Must(template.New("output").Parse(outputText))
var shareText = `{{printf "%s" . |html}}`
var shareOutput = template.Must(template.New("shareOutput").Parse(shareText))

// Compile is an HTTP handler that reads Source code from the request,
// runs the program (returning any errors),
// and sends the program's output as the HTTP response.
func PipeCompile(w http.ResponseWriter, req *http.Request) {

	title := req.URL.Path[len("/pipeile/"):]

	fmt.Println(title)
	str := strings.Split(title, "/")
	title = str[0]

	position, _ := strconv.Atoi(str[1])

	body := new(bytes.Buffer)

	if _, err := body.ReadFrom(req.Body); err != nil {
		return
	}

	fmt.Println(position)

	/*var in []byte
	if position == 1 {
		in = nil
	} else {
		in = outputs[position-2].Out
	}
	*/
	/*	If you want to use predefine languages, a language must be able to
			run	in the format:

			[Runner] [Filename]

		langName := Lang(boxes, title)
		var lang = getLang(langName)
		out, err := Compile(title, in, body.Bytes(), *lang)
	*/
	/*  Run command takes input from the previous box and an array of strings
	as commands
	*/

	out, err := InterfaceRun(interfaces[position-1], outputs, body.Bytes(), title)
	compOut := CompileOut{Out: out, Error: err}

	outputs[position-1] = &compOut

	updateBody(boxes, title, body.String())

	if err != nil {
		w.WriteHeader(404)
		output.Execute(w, out)
	} else if *htmlOutput {
		w.Write(out)
	} else {
		output.Execute(w, out)
	}
}

func sharHandler(w http.ResponseWriter, r *http.Request) {
	out := Share()
	Save(out)
	shareOutput.Execute(w, out)
}

func initConfig(config *Config) {
	page.Heading = config.Heading
	page.SubHeading = config.SubHeading
	about.Text = config.About
	about.SecondaryText = config.AboutSide
	aboutPage.Heading = "About"
	aboutPage.SubHeading = ""
}

func initBoxes(boxs ...Box) {

	for key := range boxs {

		var box = BoxStruct{}

		box.Id = strconv.Itoa(key)
		box.Position = strconv.Itoa(key + 1)
		box.Total = len(boxs)
		box.Lang = boxs[key].Syntax()
		box.Body = boxs[key].Default()
		box.Head = "heading"
		box.SubHead = "subhead"
		box.Text = boxs[key].Help()

		boxes = append(boxes, &box)

		interfaces = append(interfaces, boxs[key])

	}
}

func Serve(config *Config, boxs ...Box) {
	initConfig(config)
	initBoxes(boxs...)

	fmt.Println("cool beans")
	http.HandleFunc("/share/", sharHandler)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/", FrontPage)
	http.HandleFunc("/compile/", PipeCompile)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("js"))))
	http.ListenAndServe(":8088", nil)
}

var helloWorld = []byte(`package main

import "fmt"

func main() {
	fmt.Println("hello, world")
}
`)
