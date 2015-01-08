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

func InitDefault() (p Page,boxs []*Box){
	

	p = Page{
		Title : "",
		Heading : "Testing",
		SubHeading : "this is a SubHeading",
		Author :"",
		Body: nil,
	}	

	boxOne := Box{
		Id : "A" ,
		Head : "Hello",
		SubHead :"My First program",
		Text : "Lorem ipsum dolor sit amet",
		Lang : "python",
		Body : `print "Hello World!" `,
		Output : "",
		ErrorOut :  "",
	}

	boxTwo := Box{
		Id : "B" ,
		Lang : "java",
		Body : `public class B{
   public static void main(String [] args){
        System.out.println("hello");
    }
} `,
		Head : "Hello Again",
		SubHead :"I'm in java!",
		Text : "Ipsum dolor sit amet",
		Output : "",
		ErrorOut :  "",
	}

	boxThree := Box{
		Id : "C" ,
		Head : "Hi",
		SubHead :"Gophers unite",
		Text : "This text doesn't have to be latin",
		Lang : "go",
		Body : `package main

import( "fmt")
func main(){
	fmt.Println("Hi")
}`,
		Output : "",
		ErrorOut :  "",
	}

	boxs = append(boxs, &boxOne)
	boxs = append(boxs, &boxTwo)
	boxs = append(boxs,&boxThree)

	return p,boxs
}

func InitContact()(Page,Contact){
	p := Page{
		Title : "",
		Heading : "Contact",
		SubHeading : "",
		Author :"",
		Body: nil,
	}	
	
	con := Contact{
		TelNum : "Tel : 076 111 1111",
		Author :"Author : Steven Labrum",
		Text : "CheckIt is for the demonstration, sharing and storing of code snippets",
		Email : `Email : labrumsteven@gmail.com`,
	}

	return p,con
}


func InitAbout()(Page,About){
	p := Page{
		Title : "",
		Heading : "About",
		SubHeading : "The CheckIt Project",
		Author :"",
		Body: nil,
	}

	ab := About{
		Text : "CheckIt is for the demonstration, sharing and storing of code snippets",
		SecondaryText : "This is the secondary text",
	}

	return p,ab
}