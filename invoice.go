package chargebee

import (
	"fmt"
	"net/url"
	"strconv"
)

type InvoiceStatus string

const (
	Paid       InvoiceStatus = "paid"
	PaymentDue InvoiceStatus = "payment_due"
	NotPaid    InvoiceStatus = "not_paid"
	Voided     InvoiceStatus = "voided"
	Pending    InvoiceStatus = "pending"
)

type PriceType string

const (
	TaxExclusive PriceType = "tax_exclusive"
	TaxInclusive PriceType = "tax_inclusive"
)

type Refund struct {
	Invoice     Invoice     `json:"invoice"`
	Transaction Transaction `json:"transaction"`
	CreditNote  CreditNote  `json:"credit_note"`
}

type InvoiceParams struct {
	CustomerID string
	Charges    []*ChargeDetail
}

type ChargeDetail struct {
	Amount      uint64
	Description string
}

func (c *ChargeDetail) AppendDetails(i int, values *url.Values) {
	if c.Amount >= 1 {
		values.Add(fmt.Sprintf("charges[amount][%d]", i), strconv.FormatUint(c.Amount, 10))
	}

	if len(c.Description) > 0 {
		values.Add(fmt.Sprintf("charges[description][%d]", i), c.Description)
	}
}

type ReasonCode string

const (
	ProductUnsatisfactory ReasonCode = "product_unsatisfactory"
	ServiceUnsatisfactory ReasonCode = "service_unsatisfactory"
	OrderChange           ReasonCode = "order_change"
	OrderCancellation     ReasonCode = "order_cancellation"
	Waiver                ReasonCode = "waiver"
	Other                 ReasonCode = "other"
)

type RefundParams struct {
	InvoiceID     string
	RefundAmount  uint64
	Comment       string
	CustomerNotes string
	ReasonCode    ReasonCode
}

//https://apidocs.chargebee.com/docs/api/invoices
type Invoice struct {
	ID                    string              `json:"id"`
	PONumber              string              `json:"po_number"`
	CustomerID            string              `json:"customer_id"`
	SubscriptionID        string              `json:"subscription_id"`
	Recurring             bool                `json:"recurring"`
	Status                InvoiceStatus       `json:"status"`
	VatNumber             string              `json:"vat_number"`
	PriceType             PriceType           `json:"price_type"`
	Date                  int64               `json:"date"`
	Total                 uint64              `json:"total"`
	AmountPaid            uint64              `json:"amount_paid"`
	AmountAdjusted        uint64              `json:"amount_adjusted"`
	WriteOffAmount        uint64              `json:"write_off_amount"`
	CreditsApplied        uint64              `json:"credits_applied"`
	AmountDue             uint64              `json:"amount_due"`
	PaidAt                int64               `json:"paid_at"`
	DunningStatus         string              `json:"dunning_status"`
	NextRetryAt           int64               `json:"next_retry_at"`
	SubTotal              uint64              `json:"sub_total"`
	Tax                   uint64              `json:"tax"`
	FirstInvoice          bool                `json:"first_invoice"`
	CurrencyCode          string              `json:"currency_code"`
	ShippingAddress       Address             `json:"shipping_address"`
	BillingAddress        Address             `json:"billing_address"`
	LineItems             []LineItem          `json:"line_items"`
	LinkedPayments        []LinkedPayment     `json:"linked_payments"`
	AppliedCredits        []AppliedCredit     `json:"applied_credits"`
	AdjustmentCreditNotes []InvoiceCreditNote `json:"adjustment_credit_notes"`
	IssuedCreditNotes     []InvoiceCreditNote `json:"issued_credit_notes"`
	LinkedOrders          []LinkedOrder       `json:"linked_orders"`
	Discounts             []Discount          `json:"discounts"`
	Taxes                 []Tax               `json:"taxes"`
	LineItemTaxes         []LineItemTax       `json:"line_item_taxes"`
	Notes                 []Note              `json:"notes"`
}

type LineItem struct {
	ID                      string  `json:"id"`
	DateFrom                int64   `json:"date_from"`
	DateTo                  int64   `json:"date_to"`
	UnitMmount              uint64  `json:"unit_amount"`
	Quantity                int64   `json:quantity`
	IsTaxed                 bool    `json:"is_taxed"`
	TaxAmount               uint64  `json:"tax_amount"`
	TaxRate                 float32 `json:"tax_rate"`
	Amount                  uint64  `json:"amount"`
	DiscountAmount          uint64  `json:"discount_amount"`
	ItemLevelDiscountAmount uint64  `json:"item_level_discount_amount"`
	Description             string  `json:"description"`
	EntityType              string  `json:"entity_type"`
	EntityID                string  `json:"entity_id"`
}

type Discount struct {
	Amount      uint64 `json:"amount"`
	Description string `json:"description"`
	EntityType  string `json:"entity_type"`
	EntityID    string `json:"entity_id"`
}

type Tax struct {
	Name        string `json:"name"`
	Amount      uint64 `json:"amount"`
	Description string `json:"description"`
}

type LineItemTax struct {
	LineItemID   string  `json:"line_item_id"`
	TaxName      string  `json:"tax_name"`
	TaxRate      float32 `json:"tax_rate"`
	TaxAmount    uint64  `json:"tax_amount"`
	TaxJurisType string  `json:"tax_juris_type"`
	TaxJurisName string  `json:"tax_juris_name"`
	TaxJurisCode string  `json:"tax_juris_code"`
}

type LinkedPayment struct {
	TxnID         string `json:"txn_id"`
	AppliedAmount uint64 `json:"applied_amount"`
	AppliedAt     int64  `json:"applied_at"`
	TxnStatus     string `json:"txn_status"`
	TxnDate       int64  `json:"txn_date"`
	TxnAmount     uint64 `json:"txn_amount"`
}

type AppliedCredit struct {
	CNID          string `json:"cn_id"`
	AppliedAmount uint64 `json:"applied_amount"`
	AppliedAt     int64  `json:"applied_at"`
	CNReasonCode  string `json:"cn_reason_code"`
	CNDate        int64  `json:"cn_date"`
	CNStatus      string `json:"cn_status"`
}

type InvoiceCreditNote struct {
	CNID         string `json:"cn_id"`
	CNReasonCode string `json:"cn_reason_code"`
	CNDate       int64  `json:"cn_date"`
	CNTotal      uint64 `json:"cn_total"`
	CNStatus     string `json:"cn_status"`
}

type LinkedOrder struct {
	ID                string `json:"id"`
	Status            string `json:"status"`
	ReferenceID       string `json:"reference_id"`
	FulfillmentStatus string `json:"fulfillment_status"`
	BatchID           string `json:"batch_id"`
	CreatedAt         int64  `json:"created_at"`
}

type Note struct {
	EntityType string `json:"entity_type"`
	Note       string `json:"note"`
	EntityID   string `json:"entity_id"`
}

type Address struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Company   string `json:"company"`
	Phone     string `json:"phone"`
	Line1     string `json:"line1"`
	Line2     string `json:"line2"`
	Line3     string `json:"line3"`
	City      string `json:"city"`
	StateCode string `json:"state_code"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Zip       string `json:"zip"`
}
