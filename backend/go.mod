module github.com/YusJade/gocalculator

go 1.23.0

toolchain go1.24.1

replace github.com/YusJade/gocalculator => ./

require (
	connectrpc.com/connect v1.18.1
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.38.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
