package subrow

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CreditNoteCreditStatus string
type CreditNoteRefundStatus string
type CreditNoteReason string

const (
	CreditNoteCreditStatusAvailable CreditNoteCreditStatus = "available"
	CreditNoteCreditStatusConsumed  CreditNoteCreditStatus = "consumed"
)

const (
	CreditNoteRefundStatusPending  CreditNoteRefundStatus = "pending"
	CreditNoteRefundStatusRefunded CreditNoteRefundStatus = "refunded"
)

const (
	CreditNoteReasonDuplicatedCharge      CreditNoteReason = "duplicated_charge"
	CreditNoteReasonProductUnsatisfactory CreditNoteReason = "product_unsatisfactory"
	CreditNoteReasonOrderChange           CreditNoteReason = "order_change"
	CreditNoteReasonOrderCancellation     CreditNoteReason = "order_cancellation"
	CreditNoteReasonFraudulentCharge      CreditNoteReason = "fraudulent_charge"
	CreditNoteReasonOther                 CreditNoteReason = "other"
)

type CreditNoteRequest struct {
	client *Client
}

type CreditNoteParams struct {
	CreditNote *CreditNoteInput `json:"credit_note"`
}

type CreditNoteResult struct {
	CreditNote  *CreditNote  `json:"credit_note,omitempty"`
	CreditNotes []CreditNote `json:"credit_notes,omitempty"`
	Meta        Metadata     `json:"meta,omitempty"`
}

type CreditNoteEstimatedResult struct {
	CreditNoteEstimated *CreditNoteEstimated `json:"credit_note_estimated"`
}

type CreditListInput struct {
	PerPage            int    `json:"per_page,omitempty,string"`
	Page               int    `json:"page,omitempty,string"`
	ExternalCustomerID string `json:"external_customer_id,omitempty"`
}

type CreditNoteItem struct {
	SubrowID       uuid.UUID `json:"subrow_id,omitempty"`
	AmountCents    int       `json:"amount_cents,omitempty"`
	AmountCurrency Currency  `json:"amount_currency,omitempty"`
	Fee            Fee       `json:"fee,omitempty"`
}

type CreditNoteAppliedTax struct {
	SubrowId           uuid.UUID `json:"subrow_id,omitempty"`
	SubrowCreditNoteId uuid.UUID `json:"subrow_credit_note_id,omitempty"`
	SubrowTaxId        uuid.UUID `json:"subrow_tax_id,omitempty"`
	TaxName            string    `json:"tax_name,omitempty"`
	TaxCode            string    `json:"tax_code,omitempty"`
	TaxRate            float32   `json:"tax_rate,omitempty"`
	TaxDescription     string    `json:"tax_description,omitempty"`
	AmountCents        int       `json:"amount_cents,omitempty"`
	AmountCurrency     Currency  `json:"amount_currency,omitempty"`
	BaseAmountCents    int       `json:"base_amount_cents,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
}

type CreditNote struct {
	SubrowID          uuid.UUID        `json:"subrow_id,omitempty"`
	SequentialID      int              `json:"sequential_id,omitempty"`
	BillingEntityCode string           `json:"billing_entity_code,omitempty"`
	Number            string           `json:"number,omitempty"`
	SubrowInvoiceID   uuid.UUID        `json:"subrow_invoice_id,omitempty"`
	InvoiceNumber     string           `json:"invoice_number,omitempty"`
	Reason            CreditNoteReason `json:"reason,omitempty"`

	CreditStatus CreditNoteCreditStatus `json:"credit_status,omitempty"`
	RefundStatus CreditNoteRefundStatus `json:"refund_status,omitempty"`

	Currency                          Currency `json:"currency,omitempty"`
	TotalAmountCents                  int      `json:"total_amount_cents,omitempty"`
	CreditAmountCents                 int      `json:"credit_amount_cents,omitempty"`
	BalanceAmountCents                int      `json:"balance_amount_cents,omitempty"`
	RefundAmountCents                 int      `json:"refund_amount_cents,omitempty"`
	TaxesAmountCents                  int      `json:"taxes_amount_cents,omitempty"`
	TaxesRate                         float32  `json:"taxes_rate,omitempty"`
	SubTotalExcludingTaxesAmountCents int      `json:"sub_total_excluding_taxes_amount_cents,omitempty"`
	CouponsAdjustmentAmountCents      int      `json:"coupons_adjustment_amount_cents,omitempty"`

	FileURL string `json:"file_url,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Items []CreditNoteItem `json:"items,omitempty"`
}

type CreditNoteEstimated struct {
	SubrowInvoiceID uuid.UUID `json:"subrow_invoice_id,omitempty"`
	InvoiceNumber   string    `json:"invoice_number,omitempty"`

	Currency                          Currency `json:"currency,omitempty"`
	MaxCreditableAmountCents          int      `json:"max_creditable_amount_cents,omitempty"`
	MaxRefundableAmountCents          int      `json:"max_refundable_amount_cents,omitempty"`
	TaxesAmountCents                  int      `json:"taxes_amount_cents,omitempty"`
	TaxesRate                         float32  `json:"taxes_rate,omitempty"`
	SubTotalExcludingTaxesAmountCents int      `json:"sub_total_excluding_taxes_amount_cents,omitempty"`
	CouponsAdjustmentAmountCents      int      `json:"coupons_adjustment_amount_cents,omitempty"`

	Items []CreditNoteEstimatedItem `json:"items,omitempty"`

	AppliedTaxes []CreditNoteEstimatedAppliedTax `json:"applied_taxes,omitempty"`
}

type CreditNoteEstimatedItem struct {
	AmountCents int       `json:"amount_cents,omitempty"`
	SubrowFeeID uuid.UUID `json:"subrow_fee_id,omitempty"`
}

type CreditNoteEstimatedAppliedTax struct {
	SubrowTaxId    uuid.UUID `json:"subrow_tax_id,omitempty"`
	TaxName        string    `json:"tax_name,omitempty"`
	TaxCode        string    `json:"tax_code,omitempty"`
	TaxRate        float32   `json:"tax_rate,omitempty"`
	TaxDescription string    `json:"tax_description,omitempty"`
	AmountCents    int       `json:"amount_cents,omitempty"`
	AmountCurrency Currency  `json:"amount_currency,omitempty"`
}

type CreditNoteItemInput struct {
	SubrowFeeID uuid.UUID `json:"fee_id,omitempty"`
	AmountCents int       `json:"amount_cents,omitempty"`
}

type CreditNoteInput struct {
	SubrowInvoiceID   uuid.UUID             `json:"invoice_id,omitempty"`
	Reason            CreditNoteReason      `json:"reason,omitempty"`
	Items             []CreditNoteItemInput `json:"items,omitempty"`
	CreditAmountCents int                   `json:"refund_amount_cents,omitempty"`
	RefundAmountCents int                   `json:"credit_amount_cents,omitempty"`
}

type CreditNoteUpdateInput struct {
	SubrowID     string                 `json:"id,omitempty"`
	RefundStatus CreditNoteRefundStatus `json:"refund_status,omitempty"`
}

type CreditNoteUpdateParams struct {
	CreditNote *CreditNoteUpdateInput `json:"credit_note"`
}

type CreditNoteEstimateInput struct {
	SubrowInvoiceID uuid.UUID             `json:"invoice_id,omitempty"`
	Items           []CreditNoteItemInput `json:"items,omitempty"`
}

type CreditNoteEstimateParams struct {
	CreditNote *CreditNoteEstimateInput `json:"credit_note"`
}

func (c *Client) CreditNote() *CreditNoteRequest {
	return &CreditNoteRequest{
		client: c,
	}
}

func (cr *CreditNoteRequest) Get(ctx context.Context, creditNoteID uuid.UUID) (*CreditNote, *Error) {
	subPath := fmt.Sprintf("%s/%s", "credit_notes", creditNoteID)

	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &CreditNoteResult{},
	}

	result, err := cr.client.Get(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	creditNoteResult, ok := result.(*CreditNoteResult)
	if !ok {
		return nil, &ErrorTypeAssert
	}

	return creditNoteResult.CreditNote, nil
}

func (cr *CreditNoteRequest) Download(ctx context.Context, creditNoteID string) (*CreditNote, *Error) {
	subPath := fmt.Sprintf("%s/%s/%s", "credit_notes", creditNoteID, "download")
	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &CreditNoteResult{},
	}

	result, err := cr.client.PostWithoutBody(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	if result != nil {
		creditNoteResult, ok := result.(*CreditNoteResult)
		if !ok {
			return nil, &ErrorTypeAssert
		}

		return creditNoteResult.CreditNote, nil
	}

	return nil, nil
}

func (cr *CreditNoteRequest) GetList(ctx context.Context, creditNoteListInput *CreditListInput) (*CreditNoteResult, *Error) {
	jsonQueryParams, err := json.Marshal(creditNoteListInput)
	if err != nil {
		return nil, &Error{Err: err}
	}

	queryParams := make(map[string]string)
	if err = json.Unmarshal(jsonQueryParams, &queryParams); err != nil {
		return nil, &Error{Err: err}
	}

	clientRequest := &ClientRequest{
		Path:        "credit_notes",
		QueryParams: queryParams,
		Result:      &CreditNoteResult{},
	}

	result, clientErr := cr.client.Get(ctx, clientRequest)
	if clientErr != nil {
		return nil, clientErr
	}

	creditNoteResult, ok := result.(*CreditNoteResult)
	if !ok {
		return nil, &ErrorTypeAssert
	}

	return creditNoteResult, nil
}

func (cr *CreditNoteRequest) Create(ctx context.Context, creditNoteInput *CreditNoteInput) (*CreditNote, *Error) {
	creditNoteParams := &CreditNoteParams{
		CreditNote: creditNoteInput,
	}

	clientRequest := &ClientRequest{
		Path:   "credit_notes",
		Result: &CreditNoteResult{},
		Body:   creditNoteParams,
	}

	result, err := cr.client.Post(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	creditNoteResult, ok := result.(*CreditNoteResult)
	if !ok {
		return nil, &ErrorTypeAssert
	}

	return creditNoteResult.CreditNote, nil
}

func (cr *CreditNoteRequest) Update(ctx context.Context, creditNoteUpdateInput *CreditNoteUpdateInput) (*CreditNote, *Error) {
	subPath := fmt.Sprintf("%s/%s", "credit_notes", creditNoteUpdateInput.SubrowID)
	creditNoteParams := &CreditNoteUpdateParams{
		CreditNote: creditNoteUpdateInput,
	}

	ClientRequest := &ClientRequest{
		Path:   subPath,
		Result: &PlanResult{},
		Body:   creditNoteParams,
	}

	result, err := cr.client.Put(ctx, ClientRequest)
	if err != nil {
		return nil, err
	}

	creditNoteResult, ok := result.(*CreditNoteResult)
	if !ok {
		return nil, &ErrorTypeAssert
	}

	return creditNoteResult.CreditNote, nil
}

func (cr *CreditNoteRequest) Void(ctx context.Context, creditNoteID string) (*CreditNote, *Error) {
	subPath := fmt.Sprintf("%s/%s/%s", "credit_notes", creditNoteID, "void")

	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &CreditNoteResult{},
	}

	result, err := cr.client.Put(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	creditNoteResult, ok := result.(*CreditNoteResult)
	if !ok {
		return nil, &ErrorTypeAssert
	}

	return creditNoteResult.CreditNote, nil
}

func (cr *CreditNoteRequest) Estimate(ctx context.Context, creditNoteEstimateInput *CreditNoteEstimateInput) (*CreditNoteEstimated, *Error) {
	creditNoteEstimateParams := &CreditNoteEstimateParams{
		CreditNote: creditNoteEstimateInput,
	}

	clientRequest := &ClientRequest{
		Path:   "credit_notes/estimate",
		Result: &CreditNoteEstimatedResult{},
		Body:   creditNoteEstimateParams,
	}

	result, err := cr.client.Post(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	creditNoteEstimatedResult, ok := result.(*CreditNoteEstimatedResult)
	if !ok {
		return nil, &ErrorTypeAssert
	}

	return creditNoteEstimatedResult.CreditNoteEstimated, nil
}
