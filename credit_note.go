package chargebee

type CreditNote struct {
	ID                 string         `json:"id"`
	CustomerID         string         `json:"customer_id"`
	SubscriptionID     string         `json:"subscription_id"`
	ReferenceInvoiceID string         `json:"reference_invoice_id"`
	CreditNoteType     string         `json:"type"`
	ReasonCode         string         `json:"reason_code"`
	Status             string         `json:"status"`
	VatNumber          string         `json:"vat_number"`
	Date               int64          `json:"date"`
	PriceType          string         `json:"price_type"`
	Total              uint64         `json:"total"`
	AmountAllocated    uint64         `json:"amount_allocated"`
	AmountRefunded     uint64         `json:"amount_refunded"`
	AmountAvailable    uint64         `json:"amount_available"`
	RefundedAt         int64          `json:"refunded_at"`
	SubTotal           uint64         `json:"sub_total"`
	LineItems          []LineItem     `json:"line_items"`
	Discounts          []Discount     `json:"discounts"`
	Taxes              []Tax          `json:"taxes"`
	LineItemTaxes      []LineItemTax  `json:"line_item_taxes"`
	LinkedRefunds      []LinkedRefund `json:"linked_refunds"`
	Allocations        []Allocation   `json:"allocations"`
}

type Allocation struct {
	InvoiceID       string `json:"invoice_id"`
	AllocatedAmount uint64 `json:"allocated_amount"`
	AllocatedAt     int64  `json:"allocated_at"`
	InvoiceDate     int64  `json:"invoice_date"`
	InvoiceStatus   string `json:"invoice_status"`
}
