package main

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/m0thm4n/currency/currencypb"
	"google.golang.org/grpc"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, req *currencypb.RateRequest) (*currencypb.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", req.GetBase(), "destination", req.GetDestination())

	return &currencypb.RateResponse{Rate: 0.5}, nil
}

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := NewCurrency()

	currencypb.RegisterCurrencyService(gs, cs)

}
