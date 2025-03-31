package tests

import (
	"context"
	"net/http"
	"testing"

	"connectrpc.com/connect"
	"github.com/YusJade/gocalculator/genproto"
	"github.com/YusJade/gocalculator/genproto/genprotoconnect"
	"github.com/stretchr/testify/assert"
)

var (
	ctx    = context.TODO()
	client = genprotoconnect.NewCalculatorServiceClient(
		http.DefaultClient,
		"http://localhost:8080/",
	)
)

func TestMain(m *testing.M) {
	before()
	m.Run()
}

func before() {

}

func TestUnmatchedParentheses(t *testing.T) {
	resp, err := client.Calculate(ctx, connect.NewRequest(&genproto.CalculateRequest{
		Expression: "(1+3)/(2+3",
	}))
	t.Log(resp, err)
	assert.Equal(t, 0, resp.Msg.Code)
}

func TestInvalidCharacters(t *testing.T) {
	resp, err := client.Calculate(ctx, connect.NewRequest(&genproto.CalculateRequest{
		Expression: "1K2",
	}))
	t.Log(resp, err)
	assert.Equal(t, 0, resp.Msg.Code)
}

func TestDivisionByZero(t *testing.T) {
	resp, err := client.Calculate(ctx, connect.NewRequest(&genproto.CalculateRequest{
		Expression: "(1+3)/0",
	}))
	t.Log(resp, err)
	assert.Equal(t, 0, resp.Msg.Code)
}

func TestNilNumerator(t *testing.T) {
	resp, err := client.Calculate(ctx, connect.NewRequest(&genproto.CalculateRequest{
		Expression: "/1",
	}))
	t.Log(resp, err)
	assert.Equal(t, 0, resp.Msg.Code)
}
