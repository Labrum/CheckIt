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
	"fmt"
	"text/template"
	"io/ioutil"
	"net/http"
	"os"
	"flag"
	"bytes"
	"strings"
	"path/filepath"
)

type Page struct {
	Title string
	Heading string
	SubHeading string
	Author string
	Body []byte
}

type Box struct{
	Id string
	Head string
	SubHead string
	Text string
	Lang string
	Body string
	Output string
	ErrorOut string
}

type Boxes []*Box

func (this Boxes) Len() int {
	return len(this)
}

func Lang(this []*Box,name string) string{

	for key := range this {
		if strings.EqualFold(this[key].Id,name){
			return this[key].Lang
		}
	}
	return ""
}

func updateBody(this []*Box,name string, bod string ) {
	for key := range this {
		if strings.EqualFold(this[key].Id,name){
			this[key].Body = bod
		}
	}
}

func printBoxes(this []*Box){
	for key := range this {
		fmt.Print(this[key].Body);
		
	}

}

func (this Boxes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

var boxes = []*Box{}
var tempDirectory string = "./tmp/"
var page = Page{}

var (
	httpListen = flag.String("http", "127.0.0.1:3999", "host:port to listen on")
	htmlOutput = flag.Bool("html", false, "render program output as HTML")
	snipDir = "./snippets/"
	templateDir = "./templates"
	retDir = ".."
)


func (p *Page) save() error {
	filename := snipDir+p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := snipDir + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	} 
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmp string, p *Page){
	os.Chdir(templateDir)
    t, _ := template.ParseFiles(tmp + ".html")
    os.Chdir(retDir)
    t.Execute(w,p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	pageNames,_ := filepath.Glob(title+"/*.page")
	boxNames,_ := filepath.Glob(title+"/*.box")

	pageName := pageNames[0]
	fmt.Println(boxNames)

	head.Execute(w,nil)
	openBody.Execute(w,nil)

	p := ReadPage(pageName)
	pageStart.Execute(w,p)

	boxes = []*Box{}
	for key := range boxNames {
		boxP :=  ReadBox(boxNames[key])
		boxes = append(boxes,boxP)
		box.Execute(w,boxP)

	fmt.Println(boxP.Id)
	fmt.Println(boxP.Head)
	fmt.Println(boxP.SubHead)
	fmt.Println(boxP.Text)
	fmt.Println(boxP.Lang)
	fmt.Println(boxP.Body)
	fmt.Println(boxP.Output)
	fmt.Println(boxP.ErrorOut)
	}
	pageClose.Execute(w,nil)
	htmlClose.Execute(w,nil)
}

func baseCase(w http.ResponseWriter, r *http.Request){
	page = Page{
		Title : "",
		Heading : "Testing",
		SubHeading : "this is a SubHeading",
		Author :"",
		Body: nil,
	}	

	boxOne := Box{
		Id : "box" ,
		Head : "Hello",
		SubHead :"My First program",
		Text : "Lorem ipsum dolor sit amet",
		Lang : "python",
		Body : `print "Hello World" `,
		Output : "",
		ErrorOut :  "",
	}
	boxTwo := Box{
		Id : "pox" ,
		Lang : "java",
		Body : `public class pox{
   public static void main(String [] args){
        System.out.println("hello");
    }
} `,
		Head : "Hello Again",
		SubHead :"I'm in java!",
		Text : "Ipsum dolor sit amet",
		Output : "",
		ErrorOut :  "",
	}
	boxThree := Box{
		Id : "boxThree" ,
		Head : "Hi",
		SubHead :"Gophers unite",
		Text : "This text doesn't have to be latin",
		Lang : "go",
		Body : `package main

import( "fmt")
func main(){
	fmt.Println("Hi")
}`,
		Output : "",
		ErrorOut :  "",
	}

	boxes = append(boxes, &boxOne)
	boxes = append(boxes, &boxTwo)
	boxes = append(boxes,&boxThree)

	head.Execute(w,nil)
	openBody.Execute(w,nil)
	pageStart.Execute(w,page)
	box.Execute(w,boxOne)
	box.Execute(w,boxTwo)
	box.Execute(w,boxThree)
	pageClose.Execute(w,nil)
	htmlClose.Execute(w,nil)

}

// FrontPage is an HTTP handler that renders the goplay interface.
// If a filename is supplied in the path component of the URI,
// its contents will be put in the interface's text area.
// Otherwise, the default "hello, world" program is displayed.
func FrontPage(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]

	if len(title) < 1 {
		baseCase(w,r)
	}else{
		title := r.URL.Path[len("/"):]

		pageNames,_ := filepath.Glob(title+"/*.page")
		boxNames,_ := filepath.Glob(title+"/*.box")

		fmt.Println(len(pageNames))
		fmt.Println("hello")

		if pageNames == nil ||  boxNames  == nil{
			http.Redirect(w, r, "/", http.StatusFound)
		}else{

			pageName := pageNames[0]
			fmt.Println(boxNames)

			head.Execute(w,nil)
			openBody.Execute(w,nil)

			p := ReadPage(pageName)
			pageStart.Execute(w,p)

			boxes = []*Box{}
			for key := range boxNames {
				boxP :=  ReadBox(boxNames[key])
				boxes = append(boxes,boxP)
				box.Execute(w,boxP)

			fmt.Println(boxP.Id)
			fmt.Println(boxP.Head)
			fmt.Println(boxP.SubHead)
			fmt.Println(boxP.Text)
			fmt.Println(boxP.Lang)
			fmt.Println(boxP.Body)
			fmt.Println(boxP.Output)
			fmt.Println(boxP.ErrorOut)
			}
			pageClose.Execute(w,nil)
			htmlClose.Execute(w,nil)
		}	
	}
	
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

var outputText = `<pre>{{printf "%s" . |html}}</pre>`
var output = template.Must(template.New("output").Parse(outputText)) 
var shareText = `{{printf "%s" . |html}}`
var shareOutput = template.Must(template.New("shareOutput").Parse(shareText))
// Compile is an HTTP handler that reads Go source code from the request,
// runs the program (returning any errors),
// and sends the program's output as the HTTP response.
func cmpile(w http.ResponseWriter, req *http.Request) {

	//dir string, filename string,body []byte,lang language
	title := req.URL.Path[len("/compile/"):]
	fmt.Println(title+" This is the title")

	body := new(bytes.Buffer)
	if _, err := body.ReadFrom(req.Body); err != nil {
		return
	}	

	langName := Lang(boxes,title)

	var lang = getLang(langName)
	p,_ := Compile("",title,body.Bytes(), *lang)
	
	fmt.Println(string(p))
	out, err := Compile("",title,body.Bytes(), *lang)

	updateBody(boxes,title,body.String())

	//printBoxes(boxes)

	if err!= nil{
		w.WriteHeader(404)
		output.Execute(w,out)
	}else if *htmlOutput {
		w.Write(out)
	} else {
		output.Execute(w, out)
	}
}

func sharHandler(w http.ResponseWriter, r *http.Request) {
    out := Share()
    Save(out)
    shareOutput.Execute(w,out)
}

func main() {
	fmt.Println("cool beans")
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/share/", sharHandler)
	http.HandleFunc("/", FrontPage)
	http.HandleFunc("/compile/", cmpile)
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