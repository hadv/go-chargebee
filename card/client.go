package card

import (
	"fmt"
	"net/url"

	"github.com/hadv/chargebee"
)

// Client is used to invoke /cards APIs.
// https://apidocs.chargebee.com/docs/api/cards
type Client struct {
	B   chargebee.Backend
	Key string
}

func Update(customerID string, params *chargebee.CardParams) (*chargebee.Card, error) {
	return getC().Update(customerID, params)
}

func (c Client) Update(customerID string, params *chargebee.CardParams) (*chargebee.Card, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.FirstName) > 0 {
			body.Add("first_name", params.FirstName)
		}

		if len(params.LastName) > 0 {
			body.Add("last_name", params.LastName)
		}

		if len(params.Gateway) >= 0 {
			body.Add("gateway", (string)(params.Gateway))
		}

		if len(params.TmpToken) > 0 {
			fmt.Println(params.TmpToken)
			body.Add("tmp_token", params.TmpToken)
		}
	}

	result := &chargebee.Result{}
	err := c.B.Call("POST", "/customers/"+customerID+"/credit_card", c.Key, body, result)
	if err != nil {
		return nil, err
	}
	card := result.Card
	return &card, nil
}

func Get(customerID string) (*chargebee.Card, error) {
	return getC().Get(customerID)
}

func (c Client) Get(customerID string) (*chargebee.Card, error) {
	var body *url.Values

	result := &chargebee.Result{}
	err := c.B.Call("GET", "/cards/"+customerID, c.Key, body, result)
	if err != nil {
		return nil, err
	}
	card := result.Card
	return &card, nil
}

func Delete(customerID string) (*chargebee.Customer, error) {
	return getC().Delete(customerID)
}

func (c Client) Delete(customerID string) (*chargebee.Customer, error) {
	var body *url.Values

	cust := &chargebee.Customer{}
	err := c.B.Call("POST", "/customers/"+customerID+"/delete_card", c.Key, body, cust)
	if err != nil {
		return nil, err
	}
	return cust, nil
}

func getC() Client {
	return Client{
		B:   chargebee.NewBackend(chargebee.SiteName),
		Key: chargebee.Key,
	}
}
