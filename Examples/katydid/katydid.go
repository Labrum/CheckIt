package main

import (
	"github.com/Labrum/CheckIt"
	"github.com/katydid/katydid/asm/ast"
	"github.com/katydid/katydid/asm/lexer"
	"github.com/katydid/katydid/asm/parser"
	"io/ioutil"
	"path/filepath"
)

func dot(path, katydidStr string) ([]byte, error) {
	p := parser.NewParser()
	r, err := p.Parse(lexer.NewLexer(katydidStr))
	if err != nil {
		return nil, err
	}
	rules := r.(*ast.Rules)
	dotStr := rules.Dot()
	dotFilename := filepath.Join(path, "my.dot")
	if err := ioutil.WriteFile(dotFilename, []byte(dotStr), 0666); err != nil {
		return nil, err
	}
	dotPicFilename := filepath.Join(path, "my.png")
	return CheckIt.CombinedRun("dot", dotFilename, "-Tpng", "-o", dotPicFilename)
}

type katydid struct{}

func (this *katydid) Run(path string, textboxes []string) ([]byte, error) {
	popBuf, _, err := getPop(path, textboxes[0], textboxes[1]))
	if err != nil {
		return nil, err
	}
	desc, err := parser.ParseFile(filename, ".")
	if err != nil {
		return nil, err
	}
	descBuf, err := proto.Marshal(desc)
	if err != nil {
		return nil, err
	}
	katydidBuf := textboxes[2]
	prepStr := prepStr(fmt.Sprintf("%#v", popBuf), fmt.Sprintf("%#v", katydidBuf), fmt.Sprintf("%#v", descBuf))
	result := CheckIt.NewBuffer()

	evalBuf, err := eval(path, prepStr)
	if err != nil {
		return nil, err
	}
	result.WriteParagraph(string(evalBuf))

	benchOut, err := bench(path, prepStr)
	if err != nil {
		return nil, err
	}
	result.WriteParagraph(string(benchOut))

	dotBuf, err := dot(path, string(katydidBuf))
	if err != nil {
		return nil, err
	}
	result.Write(dotBuf)

	return result.Bytes(), nil
}

func prepStr(popBytes, katydidBytes, descBytes string) string {
	return `buf := ` + popBytes + `
	p := parser.NewParser()
	r, err := p.Parse(lexer.NewLexer(` + katydidBytes + `))
	if err != nil {
		panic(err)
	}
	rules := r.(*ast.Rules)
	desc := &descriptor.FileDescriptorSet{}
	err = proto.Unmarshal(` + descBytes + `, desc)
	if err != nil {
		panic(err)
	}
	protoTokens, err := tokens.NewZipped(rules, desc)
	if err != nil {
		panic(err)
	}
	exec, rootToken, err := compiler.Compile(rules, protoTokens)
	if err != nil {
		panic(err)
	}
	s := scanner.NewProtoScanner(protoTokens, rootToken)`
}

func eval(path string, prepStr string) ([]byte, error) {
	filename := filepath.Join(path, "main.go")
	if err := ioutil.WriteFile(filename, mainStr(prepStr), 0666); err != nil {
		return nil, err
	}
	return CheckIt.CombinedRun("go", "run", "main.go", "my.pb.go")
}

func mainStr(prepStr string) string {
	return `
	package main

	import (
		"github.com/katydid/katydid/asm/ast"
		"github.com/katydid/katydid/asm/compiler"
		"github.com/katydid/katydid/asm/lexer"
		"github.com/katydid/katydid/asm/parser"
		"github.com/katydid/katydid/serialize/proto/scanner"
		"github.com/katydid/katydid/serialize/proto/tokens"
		"github.com/gogo/protobuf/proto"
		descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	)

	func main() {
		` + prepStr + `
		if err := s.Init(buf); err != nil {
			panic(err)
		}
		if succ, err := exec.Eval(s); err != nil {
			panic(err)
		}
		if succ {
			fmt.Printf("ACCEPT")
			return
		}
		fmt.Printf("REJECT")
	}`
}

func bench(path string, prepStr string) ([]byte, error) {

}

func benchStr(prepStr string) string {
	return `
	package main

	import (
		"testing"
		"github.com/katydid/katydid/asm/ast"
		"github.com/katydid/katydid/asm/compiler"
		"github.com/katydid/katydid/asm/lexer"
		"github.com/katydid/katydid/asm/parser"
		"github.com/katydid/katydid/serialize/proto/scanner"
		"github.com/katydid/katydid/serialize/proto/tokens"
		"github.com/gogo/protobuf/proto"
		descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	)

	func BenchmarkKatydid(b *testing.B) {
		` + prepStr + `
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := s.Init(buf); err != nil {
				panic(err)
			}
			if _, err := exec.Eval(s); err != nil {
				panic(err)
			}
		}
	}`
}

func (this *katydid) Default() string {
	return `root = main.Hello
main.Hello = start
start world = accept
start _ = start
accept _ = accept

if contains($string(main.Hello.World), "World") then world else noworld`
}

func (this *katydid) Help() string {
	return ""
}

func (this *katydid) Syntax() string {
	return ""
}
