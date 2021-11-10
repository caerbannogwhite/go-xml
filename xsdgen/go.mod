module xsdgen

go 1.17

replace xsd v0.0.0 => ../xsd

replace xmltree v0.0.0 => ../xmltree

replace internal/commandline v0.0.0 => ../internal/commandline

replace internal/dependency v0.0.0 => ../internal/dependency

replace internal/gen v0.0.0 => ../internal/gen

require (
	internal/commandline v0.0.0
	internal/dependency v0.0.0
	internal/gen v0.0.0
	xmltree v0.0.0
	xsd v0.0.0
)

require (
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/net v0.0.0-20211108170745-6635138e15ea // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/tools v0.1.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)
