package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"git.sr.ht/~mekyt/latex2mathml"
)

func must[T any](d T, e error) T {
	if e != nil {
       panic(e)
   }
   return d
}

func build(t *template.Template, in string, out string) {
	must(t.ParseFiles(in))
	f := must(os.Create(out))
    defer f.Close()
    t.Execute(f, nil)
}

func texToMathML(tex string) template.HTML {
	return template.HTML(latex2mathml.Convert(
		tex,
		"http://www.w3.org/1998/Math/MathML",
		"inline",
		0,
	))
}

func texToDisplayMathML(tex string) template.HTML {
	return template.HTML(latex2mathml.Convert(
		tex,
		"http://www.w3.org/1998/Math/MathML",
		"display",
		0,
	))
}

func main() {
	funcs := template.FuncMap{
		"math": texToMathML,
		"displayMath": texToDisplayMathML,
	}
	t := must(template.ParseFiles("layout.html")).Funcs(funcs)

	path := flag.String("d", "**/template.html", "file to build")
	flag.Parse()
	files := must(filepath.Glob(*path))
	for _, f := range files {
		d := filepath.Dir(f)
		o := filepath.Join(d, "index.html")
		fmt.Printf("processing %s", d)
		build(t, f, o)
	}
}
