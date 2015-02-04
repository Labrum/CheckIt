package CheckIt

import(
	"fmt"
	"testing"
	"io/ioutil"
	"time"
	"os"
	"path/filepath"
)


type tester struct{}


func (l *tester) Run(TextAreas []string,runPath string) (out []byte, err error) {
		filename := filepath.Join(runPath,"hello.go")
		
		err = ioutil.WriteFile(filename, []byte(TextAreas[0]), 0777)

		out,err = CombinedRun(10000000000000000,runPath,"go", "run", "hello.go")

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
			root , _  = os.Getwd()
			out,_ := InterfaceRun(&tester{},texts)
			if string(out) != "hello" {
				t.Errorf("Received %s , expected %s.",string(out),"hello")
			}
		}()
    }

	fileNames, _ := ioutil.ReadDir(".")
	for _, f := range fileNames {
        if len(f.Name()) == 32{
        	fmt.Print(f.Name())
        	t.Fail()
        }

    }
    time.Sleep(time.Second *2)
}