package chargebee

type Status string
type Gateway string
type CardType string

const (
	Valid    Status = "valid"
	Expiring Status = "expiring"
	Expired  Status = "expired"
)

const (
	Chargebee    Gateway = "chargebee"
	Stripe       Gateway = "stripe"
	Braintree    Gateway = "braintree"
	AuthorizeNet Gateway = "authorize_net"
)

const (
	Visa            CardType = "visa"
	Mastercard      CardType = "mastercard"
	AmericanExpress CardType = "american_express"
	Discover        CardType = "discover"
	JCB             CardType = "jcb"
	DinersClub      CardType = "diners_club"
	Others          CardType = "other"
	NotApplicable   CardType = "not_applicable"
)

type CardParams struct {
	Gateway   Gateway
	FirstName string
	LastName  string
	TmpToken  string
}

// https://apidocs.chargebee.com/docs/api/cards#card_attributes
type Card struct {
	CustomerID       string   `json:"customer_id"`
	Status           Status   `json:"status"`
	Gateway          Gateway  `json:"gateway"`
	FirstName        string   `json:"first_name"`
	LastName         string   `json:"last_name"`
	IIN              string   `json:"iin"`
	LastFour         string   `json:"last4"`
	CardType         CardType `json:"card_type"`
	ExpiryMonth      uint8    `json:"expiry_month"`
	ExpiryYear       uint16   `json:"expiry_year"`
	BillingAddr1     string   `json:"billing_addr1"`
	BillingAddr2     string   `json:"billing_addr2"`
	BillingCity      string   `json:"billing_city"`
	BillingStateCode string   `json:"billing_state_code"`
	BillingState     string   `json:"billing_state"`
	BillingCountry   string   `json:"billing_country"`
	BillingZip       string   `json:"billing_zip"`
	IPAddress        string   `json:"ip_address"`
	MaskedNumber     string   `json:"masked_number"`
}
