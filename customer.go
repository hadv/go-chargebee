package chargebee

type Result struct {
	Customer Customer `json:"customer"`
	Card     Card     `json:"card"`
	Invoice  Invoice  `json:"invoice"`
}

// CustomerParams is the set of parameters that can be used when creating or updating a customer.
// For more details see https://stripe.com/docs/api#create_customer and https://stripe.com/docs/api#update_customer.
type CustomerParams struct {
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	Company        string
	AutoCollection string
}

// https://apidocs.chargebee.com/docs/api/customers?lang=node#customer_attributes
type Customer struct {
	ID                 string `json:"id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	Company            string `json:"company"`
	VATNumber          string `json:"vat_number"`
	AutoCollection     string `json:"auto_collection"`
	AllowDirectDebit   bool   `json:"allow_direct_debit"`
	CreatedAt          int64  `json:"created_at"`
	CreatedFromIP      string `json:"created_from_ip"`
	Taxability         string `json:"taxability"`
	EntityCode         string `json:"entity_code"`
	ExemptNumber       string `json:"exempt_number"`
	InvoiceNotes       string `json:"invoice_notes"`
	PromotionalCredits uint64 `json:"promotional_credits"`
	RefundableCredits  uint64 `json:"refundable_credits"`
	ExcessPayments     uint64 `json:"excess_payments"`
}
