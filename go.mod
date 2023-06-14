module test

go 1.17

replace xsd v0.0.0 => ./xsd

replace xsdgen v0.0.0 => ./xsdgen

replace xmltree v0.0.0 => ./xmltree

replace internal/dependency v0.0.0 => ./internal/dependency

replace internal/gen v0.0.0 => ./internal/gen

replace internal/commandline v0.0.0 => ./internal/commandline

require (
	aqwari.net/xml v0.0.0-20210331023308-d9421b293817
	github.com/alexflint/go-arg v1.4.3
	xsd v0.0.0
	xsdgen v0.0.0
)

require (
	github.com/alexflint/go-scalar v1.1.0 // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/net v0.0.0-20211108170745-6635138e15ea // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/tools v0.1.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	internal/commandline v0.0.0 // indirect
	internal/dependency v0.0.0 // indirect
	internal/gen v0.0.0 // indirect
	xmltree v0.0.0 // indirect
)
