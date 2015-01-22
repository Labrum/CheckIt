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
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var tempDirectory = ""

func writefile(filename string, body []byte, ext string) {

	err := ioutil.WriteFile(filename+ext, body, 0777)
	if err != nil {
		return
	}
}

func randFolder() string {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(rand.Intn(1000)))
	hash := md5.Sum(b)

	return fmt.Sprintf("%x", hash)
}

/*
	Runs commands specified in args using input as Stdin
*/
func InterfaceRun(box Box, textareas []string, title string) (out []byte, err error) {

	tempDirectory = randFolder()

	os.Mkdir("./"+tempDirectory, 0777)

//	writefile("./"+tempDirectory+"/"+title, body, "")

	timeout := make(chan bool, 1)
	inTime := make(chan bool, 1)

	go func() {
		fmt.Print(configuration.Timeout)
		time.Sleep(configuration.Timeout * time.Second)
		timeout <- true
	}()

	go func() {
		out, err = box.Run(textareas, tempDirectory)
		inTime <- true
	}()

	select {
	case <-inTime:

	case <-timeout:
		if out == nil {
			out = []byte("ERROR: Execution of code took too long.")
		}
	}

	if err != nil {
		fmt.Println(string(out))
		return
	}

	defer os.RemoveAll("./" + tempDirectory + "/")

	return out, err
}

func TimeOut() {
	fmt.Print("TIME OUT")
}
