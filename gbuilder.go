package main

import (
	"fmt"
	"github.com/fzft/gbuilder/pkg/filemaker"
	"github.com/fzft/gbuilder/pkg/gmod"
	"log"
	"os"
)

const (
	DOCS    = "docs"
	PKG     = "pkg"
	EXAMPLE = "example"
	IMAGES  = "images"
)

type Builder struct {
	goVersion *gmod.GoVersion
	project   string
	fileMaker *filemaker.FileMaker
}

func NewBuilder(name, author string, mode os.FileMode) *Builder {
	fileMaker := filemaker.NewFileMaker(name, author, mode)
	return &Builder{project: name, fileMaker: fileMaker}
}

func (b *Builder) Run() error {
	if err := b.mkPath(); err != nil {
		fmt.Errorf("gbuilder run failed on [g mkPath] %v\n", err)
		return err
	}

	if err := b.init(); err != nil {
		fmt.Errorf("gbuilder run failed on [g init] %v\n", err)
		return err
	}

	if err := b.mkFile(); err != nil {
		fmt.Errorf("gbuilder run failed on [g mkFile] %v\n", err)
		return err
	}
	return nil
}

func (b *Builder) init() (err error) {
	b.goVersion, err = gmod.GOVNew()
	if err != nil {
		log.Fatalf("gbuilder init failed on [GOVNEW] %v\n", err)
		return err
	}
	if err := b.goVersion.Validate(); err != nil {
		log.Fatalf("gbuilder init failed on [GOValidate] %v\n", err)
		return err
	}
	return nil
}

func (b *Builder) mkPath() error {
	if err := b.fileMaker.MkRootPath(); err != nil {
		log.Fatalf("gbuilder mkPath failed on [MkRootPath] %v", err)
		return err
	}
	return nil
}

func (b *Builder) mkFile() error {
	// go mod init
	if err := b.goVersion.GoMod(b.fileMaker.ProjPath()); err != nil {
		log.Fatalf("gbuilder mkFile failed on [go mod init] %v", err)
		return err
	}
	// .gitignore
	if err := b.fileMaker.GitIgnore(); err != nil {
		log.Fatalf("gbuilder mkFile failed on [GitIgnore] %v", err)
		return err
	}
	// README.md
	if err := b.fileMaker.ReadME(); err != nil {
		log.Fatalf("gbuilder mkFile failed on [ReadME] %v", err)
		return err
	}
	// license
	if err := b.fileMaker.License(); err != nil {
		log.Fatalf("gbuilder mkFile failed on [License] %v", err)
		return err
	}
	// pkg dir
	if err := b.fileMaker.MkPath(PKG); err != nil {
		log.Fatalf("gbuilder mkFile failed on [PKG] %v", err)
		return err
	}
	// example dir
	if err := b.fileMaker.MkPath(EXAMPLE); err != nil {
		log.Fatalf("gbuilder mkFile failed on [EXAMPLE] %v", err)
		return err
	}
	// docs dir
	if err := b.fileMaker.MkPath(DOCS); err != nil {
		log.Fatalf("gbuilder mkFile failed on [DOCS] %v", err)
		return err
	}
	// images dir
	if err := b.fileMaker.MkPath(IMAGES); err != nil {
		log.Fatalf("gbuilder mkFile failed on [IMAGES] %v", err)
		return err
	}
	// subReadme.md
	// go project structure, type: lib
	// [project].go
	goFile := fmt.Sprintf("%v.go", b.project)
	if err := b.fileMaker.MkGo(goFile); err != nil {
		log.Fatalf("gbuilder mkFile failed on [goFile] %v", err)
		return err
	}

	// [project]_test.go
	goTestFile := fmt.Sprintf("%v_test.go", b.project)
	if err := b.fileMaker.MkGo(goTestFile); err != nil {
		log.Fatalf("gbuilder mkFile failed on [goTestFile] %v", err)
		return err
	}
	return nil
}
