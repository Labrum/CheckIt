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

func textAreas() []string {
	var texts []string

	for key := range boxes {
		texts = append(texts, boxes[key].Body)
	}
	return texts
}

/*
	Runs commands specified in args using input as Stdin
*/
func InterfaceRun(box Box, body []byte, args ...string) (out []byte, err error) {

	tempDirectory = randFolder()

	os.Mkdir("./"+tempDirectory, 0777)

	writefile("./"+tempDirectory+"/"+args[0], body, "")
	args = args[1:]
	textareas := textAreas()

	killer := time.NewTimer(10000000)
	var timer bool

	go func() {
		killer = time.AfterFunc(configuration.Timeout, TimeOut)
	}()

	go func() {
		out, err = box.Run(textareas, tempDirectory, args...)
		timer = killer.Stop()
	}()

	if timer {
		out = []byte("TIME OUT")
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
