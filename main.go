package main

import (
	"flag"
	"fmt"
	"github.com/kpawlik/gofpdf"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

const (
	// ScaleFactor is default scale factor to fit svg into PDF
	ScaleFactor = 0.1
)

var (
	format      string
	orientation string
	out         string
	dir         string
	lineW       float64
	scale       float64
)

func init() {
	flag.StringVar(&format, "format", "A4", "Page format (A5|A4|A3)")
	flag.StringVar(&orientation, "orientation", "P", "Page orientation (P|L)")
	flag.StringVar(&out, "out", "result.pdf", "output file path")
	flag.StringVar(&dir, "dir", "", "directory wiith SVG files")
	flag.Float64Var(&lineW, "linew", 0, "line width")
	flag.Float64Var(&scale, "scale", ScaleFactor, "Scale")

}

func svg(format, orientation string, lineW float64, scale float64, files []string) {
	var (
		layerName string
	)
	pdf := gofpdf.New(orientation, "mm", format, "")
	if lineW > 0 {
		pdf.SetLineWidth(lineW)
	}
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 1)
	for _, file := range files {
		_, fileName := filepath.Split(file)
		if strs := strings.Split(fileName, "."); len(strs) > 0 {
			layerName = strs[0]
		} else {
			layerName = fileName
		}
		layer := pdf.AddLayer(layerName, true)
		pdf.BeginLayer(layer)
		writeSvg(file, scale, pdf)
		pdf.EndLayer()
	}
	pdf.OpenLayerPane()
	if err := pdf.OutputFileAndClose(out); err != nil {
		fmt.Printf("Error generating PDF: %v\n", err)
	}
}

func writeSvg(file string, scale float64, pdf *gofpdf.Fpdf) (err error) {
	var (
		sig gofpdf.SVGBasicType
	)
	t1 := time.Now()
	if sig, err = gofpdf.SVGBasicFileParse(file); err != nil {
		fmt.Println(err)
		pdf.SetError(err)
		return
	}
	t2 := time.Now()
	pdf.SetLineCapStyle("round")
	pdf.SetXY(0, 0)
	pdf.SVGBasicWrite(&sig, scale)
	t3 := time.Now()
	pdf.SVGWriteTexts(&sig, scale)
	t4 := time.Now()
	fmt.Printf("Parse: %v\nWrite geoms: %v\nWrite texts: %v\n", t2.Sub(t1), t3.Sub(t2), t4.Sub(t3))
	return
}

//
// Returns list of paths to SVG files from directory
//
func readFilesFromDir(dir string) []string {
	var files []string
	filesInfo, _ := ioutil.ReadDir(dir)
	for _, fi := range filesInfo {
		if name := fi.Name(); strings.HasSuffix(name, ".svg") {
			files = append(files, filepath.Join(dir, name))
		}
	}
	return files
}

func main() {
	var (
		files []string
	)
	t := time.Now()
	flag.Parse()
	if flag.NArg() == 0 && len(dir) == 0 {
		fmt.Printf("App uasage:\n\tpdfl.exe -f=A3|A4|A5 -o=P|L file1.svg file2.svg ...\n\nParams:\n")
		flag.PrintDefaults()
		return
	}
	if files = flag.Args(); len(files) == 0 && len(dir) != 0 {
		files = readFilesFromDir(dir)
	}
	orientation = strings.ToUpper(orientation)
	format = strings.ToUpper(format)
	svg(format, orientation, lineW, scale, files)
	fmt.Printf("Total time: %v\n", time.Now().Sub(t))
}
