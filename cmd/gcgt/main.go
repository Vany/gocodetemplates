package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/pkg/errors"
)

type Stash struct {
	GOFILE    string
	GOLINE    int
	GOPACKAGE string

	SUBJECT string
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("gcgt <identificator> <template> <output fname>")
		return
	}

	templateFNmae := os.Args[2]
	outputFName := os.Args[3]

	line, err := strconv.Atoi(os.Getenv("GOLINE"))
	if err != nil {
		fmt.Printf("GOLINE is not numeric : %s", err.Error())
		return
	}

	stash := Stash{
		GOFILE:    os.Getenv("GOFILE"),
		GOLINE:    line,
		GOPACKAGE: os.Getenv("GOPACKAGE"),

		SUBJECT: os.Args[1],
	}

	T, err := preparetemplate(templateFNmae)
	if err != nil {
		fmt.Printf("Template OOPS: %s", err.Error())
		return
	}

	f, err := os.OpenFile(outputFName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("File Write oops: %s", err.Error())
		return
	}

	err = T.Execute(f, stash)
	if err != nil {
		fmt.Printf("Template execute oops: %s", err.Error())
		return
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func preparetemplate(fname string) (*template.Template, error) {
	getwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	actualFname := filepath.Join(getwd, fname)

	T, err := template.ParseFiles(actualFname)
	if err != nil {
		return nil, errors.Wrapf(err, "Can't parse template '%s'", fname)
	}

	return T, nil
}
