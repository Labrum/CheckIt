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
	"strings"
)

type CompileOut struct {
	Out   []byte
	Error error
}

type Page struct {
	Heading    string
	SubHeading string
	Author     string
	Body       []byte
}

type About struct {
	Text          string
	SecondaryText string
}

type Box struct {
	Id       string
	Position string
	Total    int
	Head     string
	SubHead  string
	Text     string
	Lang     string
	Body     string
	Output   string
	ErrorOut string
}

type Boxes []*Box

func (this Boxes) Len() int {
	return len(this)
}

func Lang(this []*Box, name string) string {

	for key := range this {
		if strings.EqualFold(this[key].Id, name) {
			return this[key].Lang
		}
	}
	return ""
}

func updateBody(this []*Box, name string, bod string) {
	for key := range this {
		if strings.EqualFold(this[key].Id, name) {
			this[key].Body = bod
		}
	}
}

func printBoxes(this []*Box) {
	for key := range this {
		fmt.Print(this[key].Body)

	}

}

func (this Boxes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
