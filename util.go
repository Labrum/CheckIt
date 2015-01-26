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
	"fmt"
	"os/exec"
	"strings"
	"bytes"
	"time"
)

type Box interface {
	Descriptors() (string,string)
	Default() (string,time.Duration)
	Syntax() string
	Run(TextAreas []string, directory string, timeout time.Duration) ([]byte, error)
}

type Config struct {
	Path       string
	Port       string
	About      string
	AboutSide  string
	Heading    string
}

type Page struct {
	Heading    string
}

type AboutStruct struct {
	Text          string
	SecondaryText string
}

type BoxStruct struct {
	Id       string
	Position string
	Total    int
	Head     string
	Text     string
	Lang     string
	Body     string
}

type Boxes []*BoxStruct

func (this Boxes) Len() int {
	return len(this)
}

func Lang(this []*BoxStruct, name string) string {

	for key := range this {
		if strings.EqualFold(this[key].Id, name) {
			return this[key].Lang
		}
	}
	return ""
}

func updateBody(this []*BoxStruct, text []string) {
	for key := range this {
			this[key].Body = text[key]
	}
}

func printBoxes(this []*BoxStruct) {
	for key := range this {
		fmt.Print(this[key].Body)
	}
}

func (this Boxes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func CombinedRun(args ...string) (out []byte, err error) {

	var cmd *exec.Cmd

	cmd = exec.Command(args[0], args[1:]...)

	out, err = cmd.CombinedOutput()

	return out, err
}


func Run(args ...string) (out []byte, stderr []byte, err error) {

	var buf bytes.Buffer
	var errBuf bytes.Buffer

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = &buf
	cmd.Stderr = &errBuf

	err = cmd.Run()

	return buf.Bytes(), errBuf.Bytes() , err
	
}