package server

import (
	"context"

	protos "github.com/Xpectuer/micro-service/currency/protos/currency"

	hclog "github.com/hashicorp/go-hclog"
)

// Currency something
type Currency struct {
	log hclog.Logger
}

// NewCurrency something
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

// GetRate something
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
