package main

type populate struct{}

func getPop(path string, protoBufDef string, jsonStr string) ([]byte, string, error) {
	filename := filepath.Join(path, "my.proto")
	if err := ioutil.WriteFile(filename, protoBufDef, 0666); err != nil {
		return nil, "", err
	}
	buf, err := CheckIt.CombinedRun("protoc", "--gogo_out="+path, filename)
	if err != nil {
		return nil, "", err
	}
	desc, err := parser.ParseFile(filename, ".")
	if err != nil {
		return nil, "", err
	}
	msgs := desc.File[0].GetMessageType()
	if len(msgs) == 0 {
		return nil, "", fmt.Errorf("please define at least one message")
	}
	msg := msgs[0].GetName()
	goJson := `
	package main

	import "io/ioutil"
	import "fmt"
	import "os"
	import "encoding/json"
	import "github.com/gogo/protobuf/proto"

	func main() {
		m := &` + msg + `{}
		if err := json.Unmarshal([]byte("` + jsonStr + `"), m); err != nil {
			panic(err)
		}
		data, err := proto.Marshal(m)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(data)
		fmt.Fprintf(os.Stderr, proto.MarshalTextString(m)+"\n")
	}`
	goFile := filepath.Join(path, "main.go")
	if err := ioutil.WriteFile(goFile, goJson, 0666); err != nil {
		return nil, "", err
	}
	stdout, stderr, err := CheckIt.Run("go", "run", "main.go", "my.pb.go")
	return stdout, string(stderr), err
}

func (this *populate) Run(path string, textboxes []string, args ...string) ([]byte, error) {
	_, text, err := getPop(path, textboxes[0], string(textboxes[1]))
	return []byte(text), err
}

func (this *populate) Syntax() string {
	return "json"
}

func (this *populate) Default() string {
	return `{"World": "World"}`
}

func (this *populate) Help() string {
	return `populate the protocol buffer using json`
}
