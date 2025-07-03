package subrow

import (
	"context"
	"fmt"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (c *Client) HealthCheck(ctx context.Context) (*HealthCheckResponse, error) {
	result, err := c.Get(ctx, &ClientRequest{
		Path:   "health",
		Result: &HealthCheckResponse{},
	})

	if err != nil {
		return nil, fmt.Errorf("SubRow API health check failed: %w", err)
	}

	healthResponse, ok := result.(*HealthCheckResponse)
	if !ok {
		return nil, fmt.Errorf("SubRow API health check failed: unexpected response type")
	}

	return healthResponse, nil
}
