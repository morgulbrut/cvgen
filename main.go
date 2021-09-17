package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/cvgen/cv"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func output(data, templ, out string, d cv.CV) {

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
	color256.PrintHiCyan("Compiling %s using %s", file, cmd)

	_, err := exec.Command(cmd, "--interaction=nonstopmode", file).Output()
	if err != nil {
		color256.PrintHiOrange("Compiling PDF: %s", err)
	}
}

func cleanup() {
	remFile("cv.aux")
	remFile("cv.log")
	remFile("cv.out")
}

func remFile(path string) {
	err := os.Remove(path)
	if err != nil {
		color256.PrintHiOrange("Removing file: %s", err)
	}
}

func appendAttachments(cv string, atts []string, outfile string) {
	inFiles := []string{cv}
	inFiles = append(inFiles, atts...)
	color256.PrintHiCyan("Merging %q to %s", inFiles, outfile)
	err := api.MergeCreateFile(inFiles, outfile, nil)
	if err != nil {
		color256.PrintHiOrange("Merging PDF: %s", err)
	}
}

func generateDocument(templ, data, out string, pdfl, xel, att bool, d cv.CV) {

	output(data, templ, out, d)

	if pdfl {
		compPDF("pdflatex", out)
		cleanup()
	}
	if xel {
		compPDF("xelatex", out)
		cleanup()
	}

	if att {
		tn := filepath.Base(data)
		outname := strings.Split(tn, ".")[0] + ".pdf"
		appendAttachments("cv.pdf", d.Settings.Attachments, outname)
	}
}

func main() {
	logo()

	var templ, data, out string
	var pdfl, xel, att, all bool
	flag.StringVar(&templ, "t", "templates/moderncv.tex", "Path to cv template")
	flag.StringVar(&data, "i", "input/cv.yaml", "Path to cv data yaml file")
	flag.StringVar(&out, "o", "output/cv.tex", "Path to output file")
	flag.BoolVar(&pdfl, "P", false, "Generate a pdf using pdflatex")
	flag.BoolVar(&xel, "X", false, "Generate a pdf using xelatex")
	flag.BoolVar(&att, "A", false, "Attach documents to the cv (only works with -P or -X")
	flag.BoolVar(&all, "ba", false, "Build all input files with all outputfiles")

	flag.Parse()

	var d cv.CV
	d.Read(data)

	err := os.Setenv("TEXINPUTS", "classfiles:")
	if err != nil {
		log.Fatalf("err %v", err)
	}

	if all {
		pdfl = true

		ts, err := ioutil.ReadDir("templates")
		if err != nil {
			log.Fatal(err)
		}

		is, err := ioutil.ReadDir("input")
		if err != nil {
			log.Fatal(err)
		}

		for _, i := range is {
			for _, t := range ts {
				var d cv.CV
				d.Read("input/" + i.Name())
				generateDocument("templates/"+t.Name(), "input/"+i.Name(), out, true, false, false, d)

				outname := d.Personalinfo.Prename + "_" + d.Personalinfo.Name + ".pdf"
				path := "pdf/" + strings.Split(filepath.Base(i.Name()), ".")[0] + "/" + strings.Split(filepath.Base(t.Name()), ".")[0]
				err := os.MkdirAll(path, os.ModePerm)
				if err != nil {
					color256.PrintHiOrange("Mkdir: %s", err)
				}
				appendAttachments("cv.pdf", d.Settings.Attachments, path+"/"+outname)
			}
		}
		remFile("cv.pdf")
	} else {
		generateDocument(templ, data, out, pdfl, xel, att, d)
	}

}
