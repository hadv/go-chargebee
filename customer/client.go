package customer

import (
	"net/url"

	"github.com/hadv/chargebee"
)

// Client is used to invoke /customers APIs.
type Client struct {
	B   chargebee.Backend
	Key string
}

func New(params *chargebee.CustomerParams) (*chargebee.Customer, error) {
	return getC().New(params)
}

func (c Client) New(params *chargebee.CustomerParams) (*chargebee.Customer, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.FirstName) > 0 {
			body.Add("first_name", params.FirstName)
		}

		if len(params.LastName) > 0 {
			body.Add("last_name", params.LastName)
		}

		if len(params.Email) > 0 {
			body.Add("email", params.Email)
		}

		if len(params.Phone) > 0 {
			body.Add("phone", params.Phone)
		}

		if len(params.Company) > 0 {
			body.Add("company", params.Company)
		}

		if len(params.AutoCollection) > 0 {
			body.Add("auto_collection", params.AutoCollection)
		}
	}

	result := &chargebee.Result{}
	err := c.B.Call("POST", "/customers", c.Key, body, result)
	cust := result.Customer
	return &cust, err
}

func getC() Client {
	return Client{
		B:   chargebee.NewBackend(chargebee.SiteName),
		Key: chargebee.Key,
	}
}
