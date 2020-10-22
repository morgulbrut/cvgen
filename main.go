package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/cvgen/cv"
)

func output(data, templ, out string) {
	var d cv.CV
	d.Read(data)

	os.Remove(out)
	color256.PrintHiCyan("Rendering output(%s, %s, %s)", data, templ, out)

	tn := filepath.Base(templ)
	t, err := template.New(tn).Delims("#(", ")#").ParseFiles(templ)

	if err != nil {
		color256.PrintHiRed("generating template: ", err)
		return
	}

	f, err := os.Create(out)
	if err != nil {
		color256.PrintHiRed("creating file: ", err)
		return
	}

	err = t.Execute(f, d)
	if err != nil {
		color256.PrintHiRed("executing template: ", err)
		return
	}
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

func compPDF(cmd string, file string) {
	c := exec.Command(cmd, "--interaction=nonstopmode", file)
	st, err := c.Output()
	if err != nil {
		log.Fatal("compiling PDF: ", err)
	}
	color256.PrintHiRed(string(st))
}

func cleanup() {
	remFile("cv.aux")
	remFile("cv.log")
	remFile("cv.out")
}

func remFile(path string) {
	err := os.Remove(path)
	if err != nil {
		color256.PrintHiRed("Removing file: ", err)
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
		color256.PrintHiCyan("Compiling %s using pdflatex", out)
		compPDF("pdflatex", out)
		cleanup()
	}
	if xel {
		color256.PrintHiCyan("Compiling %s using xelatex", out)
		compPDF("xelatex", out)
		cleanup()
	}
}
