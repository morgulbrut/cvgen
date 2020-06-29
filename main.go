package main

import (
	"flag"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/cvgen/cv"
)

func executeCv(data string, f *os.File, t *template.Template) {
	var d cv.CV
	d.Read(data)
	color256.PrintBgHiYellow("%v", d)

	err := t.Execute(f, d)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
}

func executeLetter(data string, f *os.File, t *template.Template) {
	var d cv.Letter
	d.Read(data)
	color256.PrintBgHiYellow("%v", d)

	err := t.Execute(f, d)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
}

func output(data, templ, out string) {
	color256.PrintBgHiCyan("Rendering output(%s, %s, %s)", data, templ, out)

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

	if strings.Contains(out, "letter.tex") {
		executeLetter(data, f, t)
	} else {
		executeCv(data, f, t)
	}

	f.Close()
}

func main() {
	var templ, data, out string
	var letter bool
	var ltempl, ldata string
	flag.StringVar(&templ, "t", "templates/template.tex", "Path to cv template")
	flag.StringVar(&data, "i", "input/cv.yaml", "Path to cv data yaml file")
	flag.StringVar(&out, "o", "output/cv.tex", "Path to output file")
	flag.BoolVar(&letter, "L", false, "Compile a cover letter, the template should have it enabled too")
	flag.StringVar(&ltempl, "c", "templates/letter_template.tex", "Path to letter template")
	flag.StringVar(&ldata, "l", "input/letter.yaml", "Path to lette data yaml file")
	flag.Parse()

	if letter {
		output(ldata, ltempl, "output/letter.tex")
	}
	output(data, templ, out)
}
