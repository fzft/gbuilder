package filemaker

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	GIT_IGNORE = ".gitignore"
	README     = "README.md"
	LICENSE    = "LICENSE"
)

type FileMaker struct {
	project  string
	author   string
	projPath string
	modePerm os.FileMode
}

func NewFileMaker(name, author string, mp os.FileMode) *FileMaker {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return &FileMaker{project: name, author: author, modePerm: mp, projPath: path.Join(dir, name)}
}

func (fm *FileMaker) GitIgnore() error {

	gitignoreTemplate := []byte("# Binaries for programs and plugins\n" +
		"*.exe\n" +
		"*.exe~\n" +
		"*.dll\n" +
		"*.so\n" +
		"*.dylib\n" +
		"\n" +
		"# Test binary, built with `go test -c`\n" +
		"*.test\n" +
		"\n" +
		"# Output of the go coverage tool, specifically when used with LiteIDE\n" +
		"*.out\n" +
		"\n" +
		"# Dependency directories (remove the comment below to include it)\n" +
		"# vendor/\n")

	err := ioutil.WriteFile(path.Join(fm.projPath, GIT_IGNORE), gitignoreTemplate, fm.modePerm)
	if err != nil {
		log.Fatalf("fileMaker failed on [gitignore] %v", err)
		return err
	}
	return nil
}

func (fm *FileMaker) ReadME() error {

	readMeTemplate := []byte(fmt.Sprintf("# %v\n", fm.project) +
		"\n" +
		"[comment]: logo position\n" +
		"\n" +
		"## Installing\n" +
		"\n" +
		fmt.Sprintf("To start using %v, install Go and run `go get`:\n", fm.project) +
		"\n" +
		"```sh\n" +
		fmt.Sprintf("$ go get -u github.com/%v/%v\n", fm.author, fm.project) +
		"```\n" +
		"\n" +
		"## Usage\n" +
		"\n" +
		"<!---go code position--->\n" +
		"```go\n" +
		"\n" +
		"```\n" +
		"\n" +
		"## Operations\n" +
		"\n" +
		"### Basic\n" +
		"\n" +
		"<!--- lib basic usage--->\n" +
		"```\n" +
		"\n" +
		"```\n" +
		"\n" +
		"## Performance\n" +
		"\n" +
		"<!--- The following benchmarks were run on xxx--->\n" +
		"<!--- benchmarks records--->\n" +
		"```\n" +
		"```\n" +
		"\n" +
		"## Contact\n" +
		"\n" +
		fmt.Sprintf("%v \n   github@%v.com", fm.author, fm.author) +
		"\n" +
		"## License\n")

	err := ioutil.WriteFile(path.Join(fm.projPath, README), readMeTemplate, fm.modePerm)
	if err != nil {
		log.Fatalf("fileMaker failed on [README] %v", err)
		return err
	}
	return nil
}

func (fm *FileMaker) License() error {
	year, _, _ := time.Now().Date()

	licenseTemplate := []byte(fmt.Sprintf("Copyright (c) %v %v\n", year, fm.author) +
		"\n" +
		"Permission is hereby granted, free of charge, to any person obtaining a copy of\n" +
		"this software and associated documentation files (the \"Software\"), to deal in\n" +
		"the Software without restriction, including without limitation the rights to\n" +
		"use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of\n" +
		"the Software, and to permit persons to whom the Software is furnished to do so,\n" +
		"subject to the following conditions:\n" +
		"\n" +
		"The above copyright notice and this permission notice shall be included in all\n" +
		"copies or substantial portions of the Software.\n" +
		"\n" +
		"THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n" +
		"IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS\n" +
		"FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR\n" +
		"COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER\n" +
		"IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN\n" +
		"CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.\n")

	err := ioutil.WriteFile(path.Join(fm.projPath, LICENSE), licenseTemplate, fm.modePerm)
	if err != nil {
		log.Fatalf("fileMaker failed on [license] %v", err)
		return err
	}
	return nil
}

func (fm *FileMaker) MkPath(name string) error {
	fullPath := path.Join(fm.projPath, name)
	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		log.Fatalf("fileMaker failed on [mkPath] %v", name)
		return err
	}

	if err := os.Mkdir(fullPath, fm.modePerm); err != nil {
		log.Fatalf("fileMaker failed on [mkPath] %v", err)
		return err
	}
	return nil
}
func (fm *FileMaker) MkRootPath() error {
	if _, err := os.Stat(fm.projPath); !os.IsNotExist(err) {
		log.Fatalf("fileMaker failed on [mkRootPath] %v", err)
		return err
	}

	if err := os.Mkdir(fm.projPath, fm.modePerm); err != nil {
		log.Fatalf("fileMaker failed on [mkRootPath] %v", err)
		return err
	}
	return nil
}

func (fm *FileMaker) MkGo(name string) error {
	fullPath := path.Join(fm.projPath, name)
	err := ioutil.WriteFile(fullPath, []byte("package main\n"), fm.modePerm)
	if err != nil {
		fmt.Printf("fileMaker failed on [MkGo %v] %v", name, err)
		return err
	}
	return nil
}

func (fm *FileMaker) ProjPath() string {
	return fm.projPath
}
