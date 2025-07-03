package subrow

import (
	"github.com/google/uuid"
)

type InvoiceCustomSection struct {
	SubrowId    uuid.UUID `json:"subrow_id,omitempty"`
	Code        string    `json:"code,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Details     string    `json:"details,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
}
