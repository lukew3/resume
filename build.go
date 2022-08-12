package main

import (
	//"html/template"
	"log"
	"reflect"

	resumeValidator "github.com/cinarmert/json-resume-validator"
)

func main() {
	log.Println("Hello, World!")
	rv := new(resumeValidator.ResumeValidator).WithFile("resume.json")
	rv.IsValid()
	r := reflect.ValueOf(rv)
	f := reflect.Indirect(r).FieldByName("resume")
	log.Println(f)
}
