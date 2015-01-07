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
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
	"strings"
)

func Share() string {
	var buff bytes.Buffer

	buff.WriteString(page.Title)
	buff.WriteString(page.Heading)
	buff.WriteString(page.SubHeading)
	buff.WriteString(page.Author)

	for key := range boxes {
		buff.WriteString(boxes[key].Id)
		buff.WriteString(boxes[key].Head)
		buff.WriteString(boxes[key].SubHead)
		buff.WriteString(boxes[key].Text)
		buff.WriteString(boxes[key].Lang)
		buff.WriteString(boxes[key].Body)
		buff.WriteString(boxes[key].Output)
		buff.WriteString(boxes[key].ErrorOut)		
	}

	hash := md5.Sum(buff.Bytes())

	return fmt.Sprintf("%x",hash)

}

func ReadPage(filename string) *Page{

	f, _ := os.Open(filename)

	rd := bufio.NewReader(f)

	p := Page{}

	title, _ := rd.ReadString(byte(','))
	heading, _ := rd.ReadString(byte(','))
	subHeading, _ := rd.ReadString(byte(','))
	author, _ := rd.ReadString(byte(','))

	p.Title= strings.TrimSuffix(title,",")
	p.Heading = strings.TrimSuffix(heading,",")
	p.SubHeading = strings.TrimSuffix(subHeading,",")
	p.Author = strings.TrimSuffix(author,",")


	fmt.Println(p.Title)
	fmt.Println(p.Heading)
	fmt.Println(p.SubHeading)
	fmt.Println(p.Author)

	return &p
}


func ReadBox(filename string) *Box{

	f, _ := os.Open(filename)
	box := Box{}

	rd := bufio.NewReader(f)

	box.Id,_ = rd.ReadString(byte(','))
	box.Head,_ = rd.ReadString(byte(','))
	box.SubHead,_ = rd.ReadString(byte(','))
	box.Text,_ = rd.ReadString(byte(','))
	box.Lang,_ = rd.ReadString(byte(','))
	box.Body,_ = rd.ReadString(byte(','))
	box.Output,_ = rd.ReadString(byte(','))
	box.ErrorOut,_ = rd.ReadString(byte(','))

    box.Id = strings.TrimSuffix(box.Id,",")
	box.Head = strings.TrimSuffix(box.Head,",")
	box.SubHead = strings.TrimSuffix(box.SubHead,",")
	box.Text = strings.TrimSuffix(box.Text,",")
	box.Lang = strings.TrimSuffix(box.Lang,",")
	box.Body = strings.TrimSuffix(box.Body,",")
	box.Output = strings.TrimSuffix(box.Output,",")
	box.ErrorOut = strings.TrimSuffix(box.ErrorOut,",")



	return &box
	
}

func writeSave(dir string,filename string, body []byte, ext string){
	
	err := ioutil.WriteFile(dir+"/"+filename+ext, body,0777)
	if(err!= nil){
		return
	}
}

func Save(folderName string){
	var buff bytes.Buffer

	os.Mkdir(folderName,0777)
	buff.WriteString(page.Title)
	buff.WriteString(",")
	buff.WriteString(page.Heading)
	buff.WriteString(",")
	buff.WriteString(page.SubHeading)
	buff.WriteString(",")
	buff.WriteString(page.Author)
	writeSave(folderName,"page",buff.Bytes(),".page")
	buff.Reset()

	for key := range boxes {
		buff.WriteString(boxes[key].Id)
		buff.WriteString(",")
		buff.WriteString(boxes[key].Head)
		buff.WriteString(",")
		buff.WriteString(boxes[key].SubHead)
		buff.WriteString(",")
		buff.WriteString(boxes[key].Text)
		buff.WriteString(",")
		buff.WriteString(boxes[key].Lang)
		buff.WriteString(",")
		buff.WriteString(boxes[key].Body)
		buff.WriteString(",")
		buff.WriteString(boxes[key].Output)
		buff.WriteString(",")
		buff.WriteString(boxes[key].ErrorOut)		

		writeSave(folderName,boxes[key].Id,buff.Bytes(),".box")
		buff.Reset()
	}

}