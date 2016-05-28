package chargebee

// https://apidocs.chargebee.com/docs/api/transactions#transaction_attributes
type Transaction struct {
	ID                     string             `json:"id"`
	CustomerID             string             `json:"customer_id"`
	SubscripionID          string             `json:"subscription_id"`
	PaymentMethod          string             `json:"payment_method"`
	ReferenceNumber        string             `json:"reference_number"`
	Gateway                string             `json:"gateway"`
	Type                   string             `json:"type"`
	Date                   int64              `json:"date"`
	Amount                 uint64             `json:"amount"`
	IDAtGateway            string             `json:"id_at_gateway"`
	Status                 string             `json:"status"`
	ErrorCode              string             `json:"error_code"`
	ErrorText              string             `json:"error_text"`
	VoidedAt               int64              `json:"voided_at"`
	AmountUnused           uint64             `json:"amount_unused"`
	MaskedCardNumber       string             `json:"masked_card_number"`
	ReferenceTransactionID string             `json:"reference_transaction_id"`
	RefundedTxnID          string             `json:"refunded_txn_id"`
	ReversalTransactionID  string             `json:"reversal_transaction_id"`
	CurrencyCode           string             `json:"currency_code"`
	LinkedInvoices         []LinkedInvoice    `json:"linked_invoices"`
	LinkedCreditNotes      []LinkedCreditNote `json:"linked_credit_notes"`
	LinkedRefunds          []LinkedRefund     `json:"linked_refunds"`
}

type LinkedInvoice struct {
	InvoiceID     string `json:"invoice_id"`
	AppliedAmount uint64 `json:"applied_amount"`
	AppliedAt     int64  `json:"applied_at"`
	InvoiceDate   int64  `json:"invoice_date"`
	InvoiceTotal  uint64 `json:"invoice_total"`
	InvoiceStatus string `json:"invoice_status"`
}

type LinkedCreditNote struct {
	CNID                 string `json:"cn_id"`
	CNReasonCode         string `json:"cn_reason_code"`
	CNDate               int64  `json:"cn_date"`
	CNTotal              uint64 `json:"cn_total"`
	CNStatus             string `json:"cn_status"`
	CNReferenceInvoiceID string `json:"cn_reference_invoice_id"`
}

type LinkedRefund struct {
	TnxID         string `json:"txn_id"`
	AppliedAmount uint64 `json:"applied_amount"`
	AppliedAt     int64  `json:"applied_at"`
	TxnStatus     string `json:"txn_status"`
	TxnDate       int64  `json:"txn_date"`
	TxnAmount     uint64 `json:"txn_amount"`
}
