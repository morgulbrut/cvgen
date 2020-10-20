package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"text/template"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/cvgen/cv"
)

func executeCv(data string, f *os.File, t *template.Template) {
	var d cv.CV
	d.Read(data)
	err := t.Execute(f, d)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
}

// func executeLetter(data string, f *os.File, t *template.Template) {
// 	var d cv.Letter
// 	d.Read(data)
// 	err := t.Execute(f, d)
// 	if err != nil {
// 		log.Print("execute: ", err)
// 		return
// 	}
// }

func output(data, templ, out string) {
	os.Remove(out)

	color256.PrintHiCyan("Rendering output(%s, %s, %s)", data, templ, out)
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
	executeCv(data, f, t)

	f.Close()
}

func logo() {
	color256.Init()
	color256.PrintRandom("                                         ██╗	")
	color256.PrintRandom("                  ██████╗ ███████╗███╗   ██║	")
	color256.PrintRandom(" ██████╗██╗   ██╗██╔════╝ ██╔════╝████╗  ██║	")
	color256.PrintRandom("██║     ██║   ██║██║  ███╗█████╗  ██╔██╗ ██║	")
	color256.PrintRandom("██║     ╚██╗ ██╔╝██║   ██║██╔══╝  ██║╚██╗██║  ")
	color256.PrintRandom("╚██████╗ ╚████╔╝ ╚██████╔╝███████╗██║ ╚████║  ")
	color256.PrintRandom(" ╚═════╝  ╚═══╝   ╚═════╝ ╚══════╝██║  ╚═══╝  ")
	color256.PrintRandom(" CV Generator                     ╚═╝         ")
}

func runcmd(cmd string, file string) {
	c := exec.Command(cmd, file)
	st, err := c.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(st))
}

func cleanup() {
	remFile("cv.aux")
	remFile("cv.log")
	remFile("cv.out")
}

func remFile(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	logo()
	var templ, data, out string
	var letter, pdfl, xel bool
	var ltempl, ldata string
	flag.StringVar(&templ, "t", "templates/moderncv.tex", "Path to cv template")
	flag.StringVar(&data, "i", "input/cv.yaml", "Path to cv data yaml file")
	flag.StringVar(&out, "o", "output/cv.tex", "Path to output file")
	flag.BoolVar(&pdfl, "P", false, "Genderate a pdf using pdflatex")
	flag.BoolVar(&xel, "X", false, "Genderate a pdf using xelatex")

	flag.Parse()

	if letter {
		output(ldata, ltempl, "output/letter.tex")
	}
	output(data, templ, out)

	if pdfl {
		runcmd("pdflatex", out)
		cleanup()
	}
	if xel {
		runcmd("xelatex", out)
		cleanup()
	}
}
