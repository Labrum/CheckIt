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
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Share(page Page) string {
	var buff bytes.Buffer

	buff.WriteString(page.Heading)
	buff.WriteString(page.SubHeading)

	for key := range boxes {
		buff.WriteString(boxes[key].Id)
		buff.WriteString(boxes[key].Head)
		buff.WriteString(boxes[key].SubHead)
		buff.WriteString(boxes[key].Text)
		buff.WriteString(boxes[key].Lang)
		buff.WriteString(boxes[key].Body)
	}

	hash := md5.Sum(buff.Bytes())

	return fmt.Sprintf("%x", hash)

}

func ReadConfig(filename string) *Config {

	f, _ := os.Open(filename)

	file, _ := ioutil.ReadAll(f)

	c := Config{}

	if err := json.Unmarshal(file, &c); err != nil {
		panic(err)
	}

	return &c

}

func ReadBox(filename string) *BoxStruct {

	f, _ := os.Open(filename)

	file, _ := ioutil.ReadAll(f)

	b := BoxStruct{}

	if err := json.Unmarshal(file, &b); err != nil {
		panic(err)
	}

	return &b

}

func writeSave(dir string, filename string, body []byte, ext string) {

	err := ioutil.WriteFile(dir+"/"+filename+ext, body, 0777)
	if err != nil {
		return
	}
}

func Save(path string, folderName string) {
	var buff bytes.Buffer
	os.Mkdir(path+folderName, 0777)

	temp, _ := json.Marshal(configuration)
	buff.WriteString(string(temp))
	writeSave(path+folderName, "conf", buff.Bytes(), ".config")
	buff.Reset()

	for key := range boxes {
		tmp, _ := json.Marshal(boxes[key])
		buff.WriteString(string(tmp))
		writeSave(path+folderName, boxes[key].Id, buff.Bytes(), ".box")
		buff.Reset()
	}

}
