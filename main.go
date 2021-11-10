package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"xsd"

	"xsdgen"
)

var OUTPUT_FILE_NAME string = "qbas_api.go"
var QBAS_API_FIXED_VALUES map[string]string

func main() {
	schemaPath := os.Args[1]

	// files, err := os.ReadDir(schemaPath)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// data := make([]byte, 0)
	// for _, file := range files {
	// 	d, _ := os.ReadFile(path.Join(schemaPath, file.Name()))

	// 	data = append(data, d...)
	// }

	// var cfg xsdgen.Config
	// code, err := cfg.GenCode(data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(code)

	// schema, err := xsd.Parse(data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(schema)

	files, err := ioutil.ReadDir(schemaPath)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, path.Join(schemaPath, file.Name()))
	}

	var cfg xsdgen.Config
	cfg.Option(xsdgen.DefaultOptions...)
	cfg.Option(
		xsdgen.PackageName("main"),
		xsdgen.ProcessTypes(processTypesCallback),
	)

	source, err := cfg.GenSource(fileNames...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fh, err := os.Create(path.Join("api", OUTPUT_FILE_NAME))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer fh.Close()

	// write source code
	fmt.Fprint(fh, string(source))
}

func processTypesCallback(schema xsd.Schema, t xsd.Type) xsd.Type {
	if ct, ok := t.(*xsd.ComplexType); ok {
		for elementIdx, element := range ct.Elements {
			for attrIdx, _ := range element.Attr {
				// name := fmt.Sprintf("%s_%s_%s", ct.Name.Local, element.Name.Local, attr.Name.Local)
				// if attr.Name.Local == "fixed" {
				// 	QBAS_API_FIXED_VALUES[name] = attr.Value
				// }
				element.Attr[attrIdx].Name.Space = ""
			}
			ct.Elements[elementIdx].Name.Space = ""
			ct.Elements[elementIdx].Nillable = ct.Elements[elementIdx].Optional
		}
		return ct
	}
	return t
}
