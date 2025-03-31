package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/YusJade/gocalculator/app"
	genproto "github.com/YusJade/gocalculator/genproto"
	"github.com/YusJade/gocalculator/genproto/genprotoconnect"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type CalculatorServer struct {
	parser app.ExpressionParser
}

func (c *CalculatorServer) Calculate(
	_ context.Context,
	req *connect.Request[genproto.CalculateRequest]) (*connect.Response[genproto.CalculateResponse], error) {
	logrus.Infof("request in: %v", req.Msg)
	result, err := c.parser.Calculate(req.Msg.GetExpression())
	if err != nil {
		logrus.Error(err)
		return connect.NewResponse(&genproto.CalculateResponse{
			Result:  0,
			Message: err.Error(),
			Code:    -1,
		}), nil
	}
	return connect.NewResponse(&genproto.CalculateResponse{
		Result:  float32(result),
		Message: "ok",
		Code:    0,
	}), nil
}

func main() {
	calculator := &CalculatorServer{
		parser: app.NewExpressionParser(),
	}
	mux := http.NewServeMux()
	path, handler := genprotoconnect.NewCalculatorServiceHandler(calculator)
	mux.Handle(path, handler)

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

}
