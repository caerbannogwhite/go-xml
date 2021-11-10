module xsd

go 1.17

replace xmltree v0.0.0 => ../xmltree

replace internal/dependency v0.0.0 => ../internal/dependency

require (
	internal/dependency v0.0.0
	xmltree v0.0.0
)

require (
	golang.org/x/net v0.0.0-20211108170745-6635138e15ea // indirect
	golang.org/x/text v0.3.6 // indirect
)
