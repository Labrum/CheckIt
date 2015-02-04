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
	"github.com/Labrum/CheckIt"
	"strings"
)

var BOXESONPAGE = 3

type list struct{}

func (l *list) Run(textAreas []string, runPath string) (out []byte, err error) {

	args := strings.Fields(textAreas[0])

	out, _, err = CheckIt.Run(1000000000,runPath,  args...)

	return out, err
}

func (l *list) Desc() (heading string, description string, text string, syntax string) {
	heading = "List files"
	description = "This textbox uses the command line to list files"
	text = `ls -l -a`
	syntax = "cmd"
	return heading, description, text, syntax
}
