package invoice

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/hadv/chargebee"
)

// Client is used to invoke /invoices APIs.
// https://apidocs.chargebee.com/docs/api/invoices
type Client struct {
	B   chargebee.Backend
	Key string
}

func New(params *chargebee.InvoiceParams) (*chargebee.Invoice, error) {
	return getC().New(params)
}

func (c Client) New(params *chargebee.InvoiceParams) (*chargebee.Invoice, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.CustomerID) > 0 {
			body.Add("customer_id", params.CustomerID)
		}

		if len(params.Charges) > 0 {
			for i := range params.Charges {
				params.Charges[i].AppendDetails(i, body)
			}
		}
	}
	result := &chargebee.Result{}
	err := c.B.Call("POST", "/invoices", c.Key, body, result)
	if err != nil {
		return nil, err
	}
	invoice := result.Invoice
	return &invoice, nil
}

func Refund(params *chargebee.RefundParams) (*chargebee.Refund, error) {
	return getC().Refund(params)
}

func (c Client) Refund(params *chargebee.RefundParams) (*chargebee.Refund, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.RefundAmount > 0 {
			body.Add("refund_amount", strconv.FormatUint(params.RefundAmount, 10))
		}

		if len(params.Comment) > 0 {
			body.Add("comment", params.Comment)
		}

		if len(params.CustomerNotes) > 0 {
			body.Add("customer_notes", params.CustomerNotes)
		}

		if len(params.ReasonCode) > 0 {
			body.Add("credit_note[reason_code]", (string)(params.ReasonCode))
		}
	}

	refund := &chargebee.Refund{}
	err := c.B.Call("POST", fmt.Sprintf("/invoices/%v/refund", params.InvoiceID), c.Key, body, refund)
	return refund, err
}

func getC() Client {
	return Client{
		B:   chargebee.NewBackend(chargebee.SiteName),
		Key: chargebee.Key,
	}
}
