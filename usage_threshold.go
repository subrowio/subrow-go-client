package subrow

import (
	"time"

	"github.com/google/uuid"
)

type UsageThresholdInput struct {
	SubrowId             *uuid.UUID `json:"id,omitempty"`
	ThresholdDisplayName string     `json:"threshold_display_name,omitempty"`
	AmountCents          int        `json:"amount_cents"`
	Recurring            bool       `json:"recurring"`
}

type UsageThreshold struct {
	SubrowID             uuid.UUID `json:"subrow_id"`
	ThresholdDisplayName string    `json:"threshold_display_name,omitempty"`
	AmountCents          int       `json:"amount_cents"`
	Recurring            bool      `json:"recurring"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type AppliedUsageThreshold struct {
	LifetimeUsageAmountCents int            `json:"lifetime_usage_amount_cents"`
	CreatedAt                time.Time      `json:"created_at"`
	UsageThreshold           UsageThreshold `json:"usage_threshold"`
}
