package main

import (
	"flag"
	"html/template"
	"log"
	"os"

	"github.com/morgulbrut/cvgen/cv"
)

func output(data, templ, out string) {
	var cv cv.CV
	cv.Read(data)

	t, err := template.ParseFiles(templ)
	if err != nil {
		log.Print(err)
		return
	}

	f, err := os.Create(out)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(f, cv)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()
}

func main() {
	var templ, data, out string
	var letter bool
	var ltempl, ldata string
	flag.StringVar(&templ, "t", "template.tex", "Path to cv template")
	flag.StringVar(&data, "i", "cv.yaml", "Path to cv data yaml file")
	flag.StringVar(&out, "o", "cv.tex", "Path to output file")
	flag.BoolVar(&letter, "L", false, "Compile a cover letter, the template should have it enabled too")
	flag.StringVar(&ltempl, "c", "letter_template.tex", "Path to letter template")
	flag.StringVar(&ldata, "l", "letter.yaml", "Path to lette data yaml file")
	flag.Parse()

	if letter {
		output(ldata, ltempl, "letter.tex")
	}
	output(data, templ, out)
}
