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
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Box interface {
	Desc() (heading string, description string, text string, syntax string)
	Run(textAreas []string, runPath string) ([]byte, error)
}

type Config struct {
	Path      string
	Port      string
	About     string
	AboutSide string
	Heading   string
}

type Page struct {
	Heading string
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

func CombinedRun(timeout time.Duration, runPath string, args ...string) (out []byte, err error) {

	var cmd *exec.Cmd

	cmd = exec.Command(args[0], args[1:]...)
	cmd.Dir = runPath
	kill := make(chan bool, 1)
	completed := make(chan bool, 1)

	go func() {
		time.Sleep(timeout)
		kill <- true
	}()
	go func() {
		out, err = cmd.CombinedOutput()
		completed <- true
	}()

	select {
	case <-completed:
	case <-kill:
		out = []byte("Error Command timed out!")
		err = errors.New("Timed Out")
	}

	return out, err
}

func Run(timeout time.Duration, runPath string, args ...string) (out []byte, stderr []byte, err error) {

	var buf bytes.Buffer
	var errBuf bytes.Buffer

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = &buf
	cmd.Stderr = &errBuf
	cmd.Dir = runPath

	kill := make(chan bool, 1)
	completed := make(chan bool, 1)

	go func() {
		time.Sleep(timeout)
		kill <- true
	}()
	go func() {
		err = cmd.Run()
		completed <- true
	}()

	select {
	case <-completed:
	case <-kill:
		out = []byte("Error Command timed out!")
		err = errors.New("Timed Out")
		return out, nil, err
	}

	return buf.Bytes(), errBuf.Bytes(), err

}
