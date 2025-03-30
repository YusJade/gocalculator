module github.com/YusJade/gocalculator

go 1.23.0

toolchain go1.24.1

replace github.com/YusJade/gocalculator => ./

require (
	connectrpc.com/connect v1.18.1
	golang.org/x/net v0.38.0
	google.golang.org/protobuf v1.34.2
)

require golang.org/x/text v0.23.0 // indirect
