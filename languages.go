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

import( 
	//"fmt"
	)

type language struct{
	Compiler string
	Runner string
	RunWithExtension bool
	Extension string
}

var languages map[string]language = make(map[string]language)

func declareLanguages(){
	languages["go"] = language{Compiler : "go build", Runner : "go run", RunWithExtension : true , Extension: ".go"}
	languages["python"] = language{Compiler : "", Runner : "python",RunWithExtension : true , Extension: ".py"}
	languages["java"] = language{Compiler : "javac", Runner : "java",RunWithExtension : false , Extension: ".java"}
	languages["cmd"] = language{Compiler : "", Runner : "bash" ,RunWithExtension : true , Extension: ".sh"}
}

func getLang(name string) *language{
	declareLanguages()

	if value, ok := languages[name]; ok{
		return &value
	}
	return nil
}