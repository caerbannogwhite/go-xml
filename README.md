This repository contains a collection of Go packages for working
with XML, with the ultimate goal of enabling code generation based
on XML documents.

- The `xmltree` package converts xml documents to a tree data
  structure, and provides convenient methods for manipulating and
  searching through that tree.
- The `xsd` package implements a parser for XML Schema. It takes
  some liberties from the specification, and would need some work for
  use as a validator, but it handles type inheritance and XML namespaces
  in a relatively sane way.
- The `xsdgen` package provides a customizable code generator that
  generates Go type declarations and marshal/unmarshal methods for
  an XML Schema.
<!-- - The `wsdl` package parses Web Service Definition Language (WSDL)
  files, which describe a (usually) SOAP web service.
- The `wsdlgen` package generates Go source code from WSDL files. -->
- The `xsdgen` and `wsdlgen` commands generate Go code with default
  settings and are suitable for use with `go generate`.
