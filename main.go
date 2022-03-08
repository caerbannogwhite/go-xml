package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"xsd"
	"xsdgen"
)

var OUTPUT_FILE_NAME string = "qbas_schema.go"
var QBAS_API_FIXED_VALUES map[string]string

func main() {
	if len(os.Args) != 3 {
		fmt.Println("\nExample: go run main.go path/to/schema output/path\n")
		os.Exit(0)
	}

	schemaPath := os.Args[1]
	outputPaht := os.Args[2]

	// generate a new map
	QBAS_API_FIXED_VALUES = make(map[string]string)

	files, err := ioutil.ReadDir(schemaPath)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	fileNames := []string{}
	mainPath := ""
	for _, file := range files {
		path := filepath.Join(schemaPath, file.Name())
		if strings.Contains(file.Name(), "main.xsd") {
			mainPath = path
			continue
		}
		fileNames = append(fileNames, path) // on Windows
	}

	// Add the following attributes to the main schema
	attributes := `xmlns:mn="tnsMain" targetNamespace="tnsMain"`
	mainBytes, err := ioutil.ReadFile(mainPath)
	if err != nil {
		fmt.Println(err)
	}

	main := string(mainBytes)
	idx := strings.Index(main, "<xs:schema")
	schema := main[idx:]
	idx = strings.Index(schema, ">")
	schema = schema[:idx+1]
	spt := strings.Split(schema, " ")
	last := spt[len(spt)-1]
	spt[len(spt)-1] = attributes
	schemaUpdated := strings.Join(append(spt, last), " ")

	tempPath := filepath.Join(os.TempDir(), "main.xsd")
	fmt.Println("Temp main.xsd file:", tempPath)
	tempFile, err := os.Create(tempPath)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	tempFile.Write([]byte(strings.Replace(main, schema, schemaUpdated, 1)))

	fileNames = append(fileNames, tempPath)

	var cfg xsdgen.Config
	cfg.Option(xsdgen.DefaultOptions...)
	cfg.Option(
		xsdgen.PackageName("schema"),
		xsdgen.OptionalAsNillable(true),
		xsdgen.IncludeNameSpaceInTags(false),
		xsdgen.ProcessTypes(processTypesCallback),
		xsdgen.StringAsInnerXML(map[string]bool{
			"content": true,
		}),
	)

	source, err := cfg.GenSource(fileNames...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fh, err := os.Create(path.Join(outputPaht, OUTPUT_FILE_NAME))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer fh.Close()

	// write source code
	fmt.Fprint(fh, string(source))

	// add default values to qbas api source code as variables
	fmt.Fprintf(fh, "\n// QBAS API FIXED VALUES: [Type name]_[Element name]_[Attribute name]\n")
	for k, v := range QBAS_API_FIXED_VALUES {
		fmt.Fprintf(fh, "var %s %T = %#v\n", k, v, v)
	}
}

func processTypesCallback(schema xsd.Schema, t xsd.Type) xsd.Type {
	if ct, ok := t.(*xsd.ComplexType); ok {
		for _, element := range ct.Elements {
			for _, attr := range element.Attr {
				name := fmt.Sprintf("%s_%s_%s", ct.Name.Local, element.Name.Local, attr.Name.Local)
				if attr.Name.Local == "fixed" {
					QBAS_API_FIXED_VALUES[name] = attr.Value
				}
			}
		}
		return ct
	}
	return t
}
