package CheckIt

import(
	"testing"
	"io/ioutil"
	"time"
	"os"
)


type tester struct{}


func (l *tester) Run(TextAreas []string,runPath string) (out []byte, err error) {
		os.Chdir("/"+runPath)
		err = ioutil.WriteFile("hello.go", []byte(TextAreas[0]), 0777)

		out,err = CombinedRun(100000000000,"go","run","hello.go")
		os.Chdir("..")
		return out, err
}

func (l *tester) Desc() (heading string, description string, text string, syntax string) {
	heading = ""
	description = ""
	text = `package main
		import(
		"fmt"
		) 
		func main(){
			fmt.Print("hello")
			}`
	syntax = "go"
	return heading, description, text, syntax
}

func TestConc(t *testing.T){

	texts := []string{`package main
		import(
			"fmt"
		) 
		func main(){
			fmt.Print("hello")
		}`,}

	for i := 0; i< 10;i++{
		go func(){
			out,_ := InterfaceRun(&tester{},texts)
			if string(out) != "hello" {
				t.Errorf("Expected %s received %s",string(out),"hello")
			}
		}()
    }

	fileNames, _ := ioutil.ReadDir(".")
	for _, f := range fileNames {
        if len(f.Name()) == 32{
        	t.Fail()
        }

    }
    time.Sleep(time.Second *2)
}