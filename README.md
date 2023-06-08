# Read Me

Original GitHub [repository](https://github.com/droyo/go-xml).

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

## How to use

```powershell
go run .\main.go C:\dev\qubasschemas\analysis\6\ C:\dev\qubas-site\QBAS\Go\analysis\schema\ analysisSchema
```

## Future Features/Nice To Have

- Do something about the import paths.
- Might be a good idea to convert optional children and attributes into Nullable types.

## Change Log

- Added `InnerXML` type.
- Optional children and attributes are converted into pointers.