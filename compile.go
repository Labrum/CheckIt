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
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

func writefile(filename string, body []byte, ext string) {

	err := ioutil.WriteFile(filename+ext, body, 0777)
	if err != nil {
		return
	}
}

func compile(args ...string) (out []byte, err error) {

	var buff bytes.Buffer
	var cmd *exec.Cmd

	compiler := strings.Fields(args[0])
	compiler = append(compiler, args[1:]...)
	cmd = exec.Command(compiler[0], compiler[1:]...)

	cmd.Stdout = &buff
	cmd.Stderr = cmd.Stdout
	cmd.Dir = tempDirectory
	err = cmd.Run()
	out = buff.Bytes()
	if err != nil {
		fmt.Println(string(out))
		return
	}
	return out, err
}

func run(args ...string) (out []byte, err error) {
	var buff bytes.Buffer

	var cmd *exec.Cmd

	if run := strings.EqualFold("./", args[0]); run {
		fmt.Println("./" + args[1])
		cmd = exec.Command("./"+args[1], args[2:]...)
	} else {
		runner := strings.Fields(args[0])
		runner = append(runner, args[1:]...)
		cmd = exec.Command(runner[0], runner[1:]...)
	}

	f, err := os.Open("./" + tempDirectory + "/input.txt")

	//f, err := os.Open(os.Stdin.Name())

	butt, err := ioutil.ReadAll(f)
	//butt, err := os.Stdin

	reader := bytes.NewReader(butt)

	cmd.Stdin = reader
	cmd.Stdout = &buff
	cmd.Stderr = cmd.Stdout
	cmd.Dir = "./" + tempDirectory + "/"

	err = cmd.Run()
	out = buff.Bytes()
	if err != nil {
		fmt.Println(string(out))
		return
	}
	//os.Stdout.Write(out)
	return out, err
}

var tempDirectory = ""

func randFolder() string {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(rand.Intn(1000)))
	hash := md5.Sum(b)

	return fmt.Sprintf("%x", hash)
}

/*
	Runs commands specified in args using input as Stdin
*/
func Run(input []byte, title string, body []byte, args []string) (out []byte, err error) {

	var buff bytes.Buffer
	var cmd *exec.Cmd

	tempDirectory = randFolder()

	os.Mkdir("./"+tempDirectory, 0777)

	fmt.Println(args)

	writefile("./"+tempDirectory+"/input", input, ".txt")
	writefile("./"+tempDirectory+"/"+title, body, "")

	cmd = exec.Command(args[0], args[1:]...)

	f, err := os.Open("./" + tempDirectory + "/input.txt")

	butt, err := ioutil.ReadAll(f)

	reader := bytes.NewReader(butt)

	cmd.Stdin = reader
	cmd.Stdout = &buff
	cmd.Stderr = cmd.Stdout
	cmd.Dir = "./" + tempDirectory + "/"

	err = cmd.Run()
	out = buff.Bytes()

	if err != nil {
		fmt.Println(string(out))
		return
	}

	defer os.RemoveAll("./" + tempDirectory)

	return out, err
}

func Compile(filename string, input []byte, body []byte, lang language) (out []byte, err error) {

	tempDirectory = randFolder()

	os.Mkdir("./"+tempDirectory, 0777)
	writefile("./"+tempDirectory+"/"+filename, body, lang.Extension)

	if noCompiler := strings.EqualFold("", lang.Compiler); !noCompiler {
		out, err = compile(lang.Compiler, filename+lang.Extension)
	}

	if err != nil {
		return
	}

	if lang.RunWithExtension {
		if input != nil {
			writefile("./"+tempDirectory+"/input", input, ".txt")
			out, err = run(lang.Runner, filename+lang.Extension, "input.txt")
		} else {
			out, err = run(lang.Runner, filename+lang.Extension)
		}
	} else {
		if input != nil {
			writefile("./"+tempDirectory+"/input", input, ".txt")
			out, err = run(lang.Runner, filename, "input.txt")
		} else {
			out, err = run(lang.Runner, filename)
		}
	}

	defer os.RemoveAll("./" + tempDirectory)

	return
}
