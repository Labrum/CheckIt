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
	"strings"
	"github.com/Labrum/CheckIt"
	"time"
)

var BOXESONPAGE = 3

type list struct{}

func (l *list) Run(TextAreas []string, directory string, timeout time.Duration) (out []byte, err error) {

	args := strings.Fields(TextAreas[0])

//	var stderr []byte

	out,_,err = CheckIt.Run(args...)
	
	return out, err
}

func (l *list) Descriptors() (string,string) {
	title := "List files"
	description :="This textbox uses the command line to list files"
	return description,title
}

func (l *list) Default() (string,time.Duration) {
	return `ls -l -a`, 10000
}

func (l *list) Syntax() string {
	return "cmd"
}

/*
func InitDefault() (p Page, boxs []*BoxStruct) {

	p = Page{
		Heading:    "Testing",
		SubHeading: "this is a SubHeading",
	}

	boxOne := BoxStruct{
		Id:       "A",
		Position: "1",
		Total:    BOXESONPAGE,
		Head:     "Hello",
		SubHead:  "My First program",
		Text:     "Lorem ipsum dolor sit amet",
		Lang:     "cmd",
		Body: `cd ..
ls -l -a`,
	}

	boxTwo := BoxStruct{
		Id:       "B",
		Position: "2",
		Total:    BOXESONPAGE,
		Lang:     "cmd",
		Body:     `grep init`,
		Head:     "Hello Again",
		SubHead:  "I'm in java!",
		Text:     "Ipsum dolor sit amet",
	}

	boxThree := BoxStruct{
		Id:       "C",
		Position: "3",
		Total:    BOXESONPAGE,
		Head:     "Hi",
		SubHead:  "Gophers unite",
		Text:     "This text doesn't have to be latin",
		Lang:     "cmd",
		Body:     `cut -d' ' -f10`,
	}

	boxs = append(boxs, &boxOne)
	boxs = append(boxs, &boxTwo)
	boxs = append(boxs, &boxThree)

	return p, boxs
}

func InitAbout() (Page, AboutStruct) {
	p := Page{
		Heading:    "About",
		SubHeading: "The CheckIt Project",
	}

	ab := AboutStruct{
		Text:          "CheckIt is for the demonstration, sharing and storing of code snippets",
		SecondaryText: `<a href="http://www.google.com"> Google </a>`,
	}

	return p, ab
}
*/
