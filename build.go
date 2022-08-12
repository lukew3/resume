package main

import (
	"text/template"
	"log"
	"os"

	resumeValidator "github.com/lukew3/json-resume-validator"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Hello, World!")
	rv := new(resumeValidator.ResumeValidator).WithFile("resume.json")
	err := rv.Validate()
	if err != nil {
		log.Println(err)
	}
	// log.Println(rv.Resume)

	t, err := template.New("").Delims("<%", "%>").ParseFiles("template.tex")
	checkErr(err)
	err = t.Execute(os.Stdout, rv.Resume)
	checkErr(err)
}
