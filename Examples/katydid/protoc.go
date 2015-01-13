package main

import (
	"github.com/Labrum/CheckIt"
	"github.com/gogo/protobuf/parser"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

type protoc struct{}

func (this *protoc) Run(path string, textboxes []string) ([]byte, error) {
	protoBufDef := textboxes[0]
	filename := filepath.Join(path, "my.proto")
	if err := ioutil.WriteFile(filename, protoBufDef, 0666); err != nil {
		return nil, err
	}
	desc, err := parser.ParseFile(filename, ".")
	if err != nil {
		return nil, err
	}
	result := CheckIt.NewBuffer()
	result.WriteParagraph(proto.MarshalTextString(desc) + "\n")

	buf, err = CheckIt.CombinedRun("protoc", "--gogo_out="+path, filename)
	if err != nil {
		return buf, err
	}
	pbgo, err := ioutil.ReadFile(filepath.Join(path, "my.pb.go"))
	if err != nil {
		return nil, err
	}
	result.WriteParagraph(string(pbgo))

	return result.Bytes(), nil
}

func (this *protoc) Syntax() string {
	return "proto"
}

func (this *protoc) Default() string {
	return `package main;

message Hello {
	optional string World = 1;
}`
}

func (this *proto) Help() string {
	return `Protocol Buffers are a way of encoding structured data in an efficient yet extensible format. 
For more information please see <a href="http://code.google.com/p/protobuf/">http://code.google.com/p/protobuf/</a>.`
}
