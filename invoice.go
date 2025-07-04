package subrow

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type InvoiceType string
type InvoiceStatus string
type InvoicePaymentStatus string
type InvoiceCreditItemType string

const (
	SubscriptionInvoiceType       InvoiceType = "subscription"
	AddOnInvoiceType              InvoiceType = "add_on"
	CreditInvoiceType             InvoiceType = "credit"
	OneOffInvoiceType             InvoiceType = "one_off"
	ProgressiveBillingInvoiceType InvoiceType = "progressive_billing"
)

const (
	InvoiceStatusDraft     InvoiceStatus = "draft"
	InvoiceStatusFinalized InvoiceStatus = "finalized"
	InvoiceStatusFailed    InvoiceStatus = "failed"
	InvoiceStatusVoided    InvoiceStatus = "voided"
	InvoiceStatusPending   InvoiceStatus = "pending"
)

const (
	InvoicePaymentStatusPending   InvoicePaymentStatus = "pending"
	InvoicePaymentStatusSucceeded InvoicePaymentStatus = "succeeded"
	InvoicePaymentStatusFailed    InvoicePaymentStatus = "failed"
)

const (
	InvoiceCreditItemCoupon     InvoiceCreditItemType = "coupon"
	InvoiceCreditItemCreditNote InvoiceCreditItemType = "credit_note"
	InvoiceCreditItemInvoice    InvoiceCreditItemType = "invoice"
)

type InvoiceRequest struct {
	client *Client
}

type InvoiceResult struct {
	Invoice  *Invoice  `json:"invoice,omitempty"`
	Invoices []Invoice `json:"invoices,omitempty"`
	Meta     Metadata  `json:"meta,omitempty"`
}

type InvoicePaymentUrlResult struct {
	InvoicePaymentUrl *InvoicePaymentUrl `json:"invoice_payment_url"`
}

type InvoiceParams struct {
	Invoice *InvoiceInput `json:"invoice"`
}

type InvoiceOneOffParams struct {
	Invoice *InvoiceOneOffInput `json:"invoice"`
}

type InvoiceMetadataInput struct {
	SubrowID *uuid.UUID `json:"id,omitempty"`
	Key      string     `json:"key,omitempty"`
	Value    string     `json:"value,omitempty"`
}

type InvoiceFeesInput struct {
	AddOnCode          string   `json:"add_on_code,omitempty"`
	InvoiceDisplayName string   `json:"invoice_display_name,omitempty"`
	UnitAmountCents    int      `json:"unit_amount_cents,omitempty"`
	Description        string   `json:"description,omitempty"`
	Units              float32  `json:"units,omitempty"`
	TaxCodes           []string `json:"tax_codes,omitempty"`
}

type InvoiceMetadataResponse struct {
	SubrowID  uuid.UUID `json:"subrow_id,omitempty"`
	Key       string    `json:"key,omitempty"`
	Value     string    `json:"value,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type InvoiceInput struct {
	SubrowID      uuid.UUID              `json:"subrow_id,omitempty"`
	PaymentStatus InvoicePaymentStatus   `json:"payment_status,omitempty"`
	Metadata      []InvoiceMetadataInput `json:"metadata,omitempty"`
}

type InvoiceOneOffInput struct {
	ExternalCustomerId string             `json:"external_customer_id,omitempty"`
	Currency           string             `json:"currency,omitempty"`
	Fees               []InvoiceFeesInput `json:"fees,omitempty"`
	SkipPsp            bool               `json:"skip_psp,omitempty"`
}

type InvoicePreviewInput struct {
	PlanCode          string              `json:"plan_code,omitempty"`
	BillingTime       string              `json:"billing_time,omitempty"`
	SubscriptionAt    string              `json:"subscription_at,omitempty"`
	Coupons           []CouponInput       `json:"coupons,omitempty"`
	Customer          *CustomerInput      `json:"customer,omitempty"`
	Subscriptions     *SubscriptionsInput `json:"subscriptions,omitempty"`
	BillingEntityCode string              `json:"billing_entity_code,omitempty"`
}

type InvoiceListInput struct {
	PerPage int `json:"per_page,omitempty,string"`
	Page    int `json:"page,omitempty,string"`

	IssuingDateFrom string `json:"issuing_date_from,omitempty"`
	IssuingDateTo   string `json:"issuing_date_to,omitempty"`

	ExternalCustomerID string               `json:"external_customer_id,omitempty"`
	Status             InvoiceStatus        `json:"status,omitempty"`
	PaymentStatus      InvoicePaymentStatus `json:"payment_status,omitempty"`
	PaymentOverdue     bool                 `json:"payment_overdue,omitempty"`

	AmountFrom int `json:"amount_from,omitempty"`
	AmountTo   int `json:"amount_to,omitempty"`
}

type InvoiceCreditItem struct {
	SubrowID uuid.UUID             `json:"subrow_id,omitempty"`
	Type     InvoiceCreditItemType `json:"type,omitempty"`
	Code     string                `json:"code,omitempty"`
	Name     string                `json:"name,omitempty"`
}

type InvoiceSummary struct {
	SubrowID      uuid.UUID            `json:"subrow_id,omitempty"`
	PaymentStatus InvoicePaymentStatus `json:"payment_status,omitempty"`
}

type InvoiceCredit struct {
	Item InvoiceCreditItem `json:"item,omitempty"`

	Invoice InvoiceSummary `json:"invoice,omitempty"`

	SubrowItemID   uuid.UUID `json:"subrow_item_id,omitempty"`
	AmountCents    int       `json:"amount_cents,omitempty"`
	AmountCurrency Currency  `json:"amount_currency,omitempty"`
	BeforeTaxes    bool      `json:"before_taxes,omitempty"`
}

type InvoiceAppliedInvoiceCustomSection struct {
	SubrowId        uuid.UUID `json:"subrow_id,omitempty"`
	SubrowInvoiceId uuid.UUID `json:"subrow_invoice_id,omitempty"`
	Code            string    `json:"code,omitempty"`
	Details         string    `json:"details,omitempty"`
	DisplayName     string    `json:"display_name,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

type InvoiceAppliedTax struct {
	SubrowId        uuid.UUID `json:"subrow_id,omitempty"`
	SubrowInvoiceId uuid.UUID `json:"subrow_invoice_id,omitempty"`
	SubrowTaxId     uuid.UUID `json:"subrow_tax_id,omitempty"`
	TaxName         string    `json:"tax_name,omitempty"`
	TaxCode         string    `json:"tax_code,omitempty"`
	TaxRate         float32   `json:"tax_rate,omitempty"`
	TaxDescription  string    `json:"tax_description,omitempty"`
	AmountCents     int       `json:"amount_cents,omitempty"`
	AmountCurrency  Currency  `json:"amount_currency,omitempty"`
	FeesAmountCents int       `json:"fees_amount_cents,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

type InvoiceErrorDetail struct {
	SubrowId  uuid.UUID      `json:"subrow_id,omitempty"`
	ErrorCode string         `json:"error_code,omitempty"`
	Details   map[string]any `json:"details,omitempty"`
}

type Invoice struct {
	SubrowID          uuid.UUID `json:"subrow_id,omitempty"`
	SequentialID      int       `json:"sequential_id,omitempty"`
	BillingEntityCode string    `json:"billing_entity_code,omitempty"`
	Number            string    `json:"number,omitempty"`

	IssuingDate          string    `json:"issuing_date,omitempty"`
	PaymentDisputeLostAt time.Time `json:"payment_dispute_lost_at,omitempty"`
	PaymentDueDate       string    `json:"payment_due_date,omitempty"`
	PaymentOverdue       bool      `json:"payment_overdue,omitempty"`

	InvoiceType   InvoiceType          `json:"invoice_type,omitempty"`
	Status        InvoiceStatus        `json:"status,omitempty"`
	PaymentStatus InvoicePaymentStatus `json:"payment_status,omitempty"`

	Currency Currency `json:"currency,omitempty"`

	FeesAmountCents                     int `json:"fees_amount_cents,omitempty"`
	TaxesAmountCents                    int `json:"taxes_amount_cents,omitempty"`
	CouponsAmountCents                  int `json:"coupons_amount_cents,omitempty"`
	CreditNotesAmountCents              int `json:"credit_notes_amount_cents,omitempty"`
	SubTotalExcludingTaxesAmountCents   int `json:"sub_total_excluding_taxes_amount_cents,omitempty"`
	SubTotalIncludingTaxesAmountCents   int `json:"sub_total_including_taxes_amount_cents,omitempty"`
	TotalAmountCents                    int `json:"total_amount_cents,omitempty"`
	TotalDueAmountCents                 int `json:"total_due_amount_cents,omitempty"`
	PrepaidCreditAmountCents            int `json:"prepaid_credit_amount_cents,omitempty"`
	ProgressiveBillingCreditAmountCents int `json:"progressive_billing_credit_amount_cents"`
	NetPaymentTerm                      int `json:"net_payment_term,omitempty"`

	FileURL       string                    `json:"file_url,omitempty"`
	Metadata      []InvoiceMetadataResponse `json:"metadata,omitempty"`
	VersionNumber int                       `json:"version_number,omitempty"`

	Customer       *Customer       `json:"customer,omitempty"`
	BillingPeriods []BillingPeriod `json:"billing_periods,omitempty"`
	Subscriptions  []Subscription  `json:"subscriptions,omitempty"`

	Fees                         []Fee                                `json:"fees,omitempty"`
	Credits                      []InvoiceCredit                      `json:"credits,omitempty"`
	AppliedInvoiceCustomSections []InvoiceAppliedInvoiceCustomSection `json:"applied_invoice_custom_sections,omitempty"`
	AppliedTaxes                 []InvoiceAppliedTax                  `json:"applied_taxes,omitempty"`
	ErrorDetails                 []InvoiceErrorDetail                 `json:"error_details,omitempty"`
	AppliedUsageThreshold        []AppliedUsageThreshold              `json:"applied_usage_threshold,omitempty"`
}

type InvoicePaymentUrl struct {
	PaymentUrl string `json:"payment_url,omitempty"`
}

func (c *Client) Invoice() *InvoiceRequest {
	return &InvoiceRequest{
		client: c,
	}
}

func (ir *InvoiceRequest) Get(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	subPath := fmt.Sprintf("%s/%s", "invoices", invoiceID)
	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &InvoiceResult{},
	}

	result, err := ir.client.Get(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	invoiceResult, ok := result.(*InvoiceResult)
	if !ok {
		return nil, &ErrorTypeAssert
	}

	return invoiceResult.Invoice, nil
}

func (ir *InvoiceRequest) GetList(ctx context.Context, invoiceListInput *InvoiceListInput) (*InvoiceResult, *Error) {
	jsonQueryParams, err := json.Marshal(invoiceListInput)
	if err != nil {
		return nil, &Error{Err: err}
	}

	queryParams := make(map[string]string)
	if err = json.Unmarshal(jsonQueryParams, &queryParams); err != nil {
		return nil, &Error{Err: err}
	}

	clientRequest := &ClientRequest{
		Path:        "invoices",
		QueryParams: queryParams,
		Result:      &InvoiceResult{},
	}

	result, clientErr := ir.client.Get(ctx, clientRequest)
	if clientErr != nil {
		return nil, clientErr
	}

	invoiceResult, ok := result.(*InvoiceResult)
	if !ok {
		return nil, &ErrorTypeAssert
	}

	return invoiceResult, nil
}

func (ir *InvoiceRequest) Create(ctx context.Context, oneOffInput *InvoiceOneOffInput) (*Invoice, *Error) {
	invoiceOneOffParams := &InvoiceOneOffParams{
		Invoice: oneOffInput,
	}

	clientRequest := &ClientRequest{
		Path:   "invoices",
		Result: &InvoiceResult{},
		Body:   invoiceOneOffParams,
	}

	result, err := ir.client.Post(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	invoiceResult, ok := result.(*InvoiceResult)
	if !ok {
		return nil, &ErrorTypeAssert
	}

	return invoiceResult.Invoice, nil
}

func (ir *InvoiceRequest) Preview(ctx context.Context, invoicePreviewInput *InvoicePreviewInput) (*Invoice, *Error) {
	clientRequest := &ClientRequest{
		Path:   "invoices/preview",
		Result: &InvoiceResult{},
		Body:   invoicePreviewInput,
	}

	result, err := ir.client.Post(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	invoiceResult, ok := result.(*InvoiceResult)
	if !ok {
		return nil, &ErrorTypeAssert
	}

	return invoiceResult.Invoice, nil
}

func (ir *InvoiceRequest) Update(ctx context.Context, invoiceInput *InvoiceInput) (*Invoice, *Error) {
	subPath := fmt.Sprintf("%s/%s", "invoices", invoiceInput.SubrowID)
	invoiceParams := &InvoiceParams{
		Invoice: invoiceInput,
	}

	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &InvoiceResult{},
		Body:   invoiceParams,
	}

	result, err := ir.client.Put(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	invoiceResult, ok := result.(*InvoiceResult)
	if !ok {
		return nil, &ErrorTypeAssert
	}

	return invoiceResult.Invoice, nil
}

func (ir *InvoiceRequest) Download(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	subPath := fmt.Sprintf("%s/%s/%s", "invoices", invoiceID, "download")
	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &InvoiceResult{},
	}

	result, err := ir.client.PostWithoutBody(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	if result != nil {
		invoiceResult, ok := result.(*InvoiceResult)
		if !ok {
			return nil, &ErrorTypeAssert
		}

		return invoiceResult.Invoice, nil
	}

	return nil, nil
}

func (ir *InvoiceRequest) Refresh(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	subPath := fmt.Sprintf("%s/%s/%s", "invoices", invoiceID, "refresh")
	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &InvoiceResult{},
	}

	result, err := ir.client.Put(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	if result != nil {
		invoiceResult, ok := result.(*InvoiceResult)
		if !ok {
			return nil, &ErrorTypeAssert
		}

		return invoiceResult.Invoice, nil
	}

	return nil, nil
}

func (ir *InvoiceRequest) Retry(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	subPath := fmt.Sprintf("%s/%s/%s", "invoices", invoiceID, "retry")
	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &InvoiceResult{},
	}

	result, err := ir.client.Post(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	if result != nil {
		invoiceResult, ok := result.(*InvoiceResult)
		if !ok {
			return nil, &ErrorTypeAssert
		}

		return invoiceResult.Invoice, nil
	}

	return nil, nil
}

func (ir *InvoiceRequest) Finalize(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	subPath := fmt.Sprintf("%s/%s/%s", "invoices", invoiceID, "finalize")
	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &InvoiceResult{},
	}

	result, err := ir.client.Put(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	if result != nil {
		invoiceResult, ok := result.(*InvoiceResult)
		if !ok {
			return nil, &ErrorTypeAssert
		}

		return invoiceResult.Invoice, nil
	}

	return nil, nil
}

type VoidInvoiceOptions struct {
	GenerateCreditNote bool `json:"generate_credit_note,omitempty"`
	RefundAmount       int  `json:"refund_amount,omitempty"`
	CreditAmount       int  `json:"credit_amount,omitempty"`
}

func (ir *InvoiceRequest) Void(ctx context.Context, invoiceID string, opts *VoidInvoiceOptions) (*Invoice, *Error) {
	subPath := fmt.Sprintf("%s/%s/%s", "invoices", invoiceID, "void")
	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &InvoiceResult{},
	}

	if opts != nil {
		clientRequest.Body = opts
	}

	result, err := ir.client.Post(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	if result != nil {
		invoiceResult, ok := result.(*InvoiceResult)
		if !ok {
			return nil, &ErrorTypeAssert
		}

		return invoiceResult.Invoice, nil
	}

	return nil, nil
}

func (ir *InvoiceRequest) LoseDispute(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	subPath := fmt.Sprintf("%s/%s/%s", "invoices", invoiceID, "lose_dispute")
	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &InvoiceResult{},
	}

	result, err := ir.client.Put(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	if result != nil {
		invoiceResult, ok := result.(*InvoiceResult)
		if !ok {
			return nil, &ErrorTypeAssert
		}

		return invoiceResult.Invoice, nil
	}

	return nil, nil
}

// We have Invoice as a possible return to be consitent with other endpoints, but no Invoice will be returned.
func (ir *InvoiceRequest) RetryPayment(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	subPath := fmt.Sprintf("%s/%s/%s", "invoices", invoiceID, "retry_payment")
	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &InvoiceResult{},
	}

	// We don't return an invoice here due to async retry payment processing
	_, err := ir.client.PostWithoutBody(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (ir *InvoiceRequest) PaymentUrl(ctx context.Context, invoiceID string) (*InvoicePaymentUrl, *Error) {
	subPath := fmt.Sprintf("%s/%s/%s", "invoices", invoiceID, "payment_url")

	clientRequest := &ClientRequest{
		Path:   subPath,
		Result: &InvoicePaymentUrlResult{},
	}

	result, err := ir.client.PostWithoutBody(ctx, clientRequest)
	if err != nil {
		return nil, err
	}

	if result != nil {
		paymentUrlResult, ok := result.(*InvoicePaymentUrlResult)
		if !ok {
			return nil, &ErrorTypeAssert
		}

		return paymentUrlResult.InvoicePaymentUrl, nil
	}

	return nil, nil
}
