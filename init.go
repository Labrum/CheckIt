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
	"fmt"
	"os/exec"
	"strings"
)

var BOXESONPAGE = 3

type list struct{}

func (l *list) Run(input [10]*CompileOut, code []byte, directory string, args ...string) (out []byte, err error) {

	var buff bytes.Buffer
	var cmd *exec.Cmd

	commands := bytes.NewBuffer(code)
	args = strings.Fields(commands.String())

	fmt.Println(args[0])
	cmd = exec.Command(args[0], args[1:]...)

	cmd.Stdout = &buff
	cmd.Stderr = cmd.Stdout
	cmd.Dir = "./" + directory + "/"

	err = cmd.Run()
	out = buff.Bytes()

	return out, err
}

func (l *list) Help() string {
	return "This textbox uses the command line to list files"
}

func (l *list) Default() string {
	return `ls -l -a`
}

func (l *list) Syntax() string {
	return "cmd"
}

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
