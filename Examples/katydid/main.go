package main

import (
	"github.com/Labrum/CheckIt"
	"github.com/alecthomas/kingpin"
	"log"
)

var path = kingpin.Flag("path", "path to all temporary and shared files").Default(".").Short('p').String()
var port = kingpin.Flag("port", "http port").Default("8080").String()
var timeout = kingpin.Flag("timeout", "timeout for each forked process").Default("5s").Short('t').Duration()

/*
type Box interface {
	Help() string
	//used for syntax highlighting
	Syntax() string
	//starting text in text box if none is provided
	Default() string
	Run(path string, textboxes []string) ([]byte, error)
}

func NewBuffer() Buffer {
	return &buffer{
		bytes.NewBuffer(nil),
	}
}

type Buffer interface {
	// <p>text</p> or <pre>text</pre>
	WriteParagraph(text string)
	Write([]byte) (n int, err error)
}

func CombinedRun(cmd string, args ...string) ([]byte, error)
func Run(cmd string, args ...string) ([]byte, []byte, error)

*/

func main() {
	kingpin.Parse()
	config := &CheckIt.Config{
		Path:      *path,
		Port:      *port,
		Timeout:   *timeout,
		About:     "this is my example about",
		AboutSide: "my contact details",
		Title:     "My Demo",
	}
	if err := CheckIt.Serve(config, &protoc{}, &populate{}, &katydid{}); err != nil {
		log.Fatal(err)
	}
	//Serve function signature
	//func Serve(config *Config, boxes ...Box) error
}
