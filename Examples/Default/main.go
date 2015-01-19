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
)

func main() {
	var con = CheckIt.Config{
		Path:       "",
		Port:       "",
		Timeout:    100000,
		Heading:    "Testing",
		SubHeading: "this is a SubHeading",
		About:      "CheckIt is for the demonstration, sharing and storing of code snippets",
		AboutSide:  `<a href="http://www.google.com"> Google </a>`,
	}
	CheckIt.Serve(&con, &list{}, &list{}, &list{})
}
