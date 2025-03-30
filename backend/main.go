package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	genproto "github.com/YusJade/gocalculator/genproto"
	"github.com/YusJade/gocalculator/genproto/genprotoconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type CalculatorServer struct {
}

func (c *CalculatorServer) Calculate(
	_ context.Context,
	req *connect.Request[genproto.CalculateRequest]) (*connect.Response[genproto.CalculateResponse], error) {
	panic("no implement")
}

func main() {
	calculator := &CalculatorServer{}
	mux := http.NewServeMux()
	path, handler := genprotoconnect.NewCalculatorServiceHandler(calculator)
	mux.Handle(path, handler)

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

}
