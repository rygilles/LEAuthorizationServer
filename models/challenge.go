package model

import "net/url"

// Challenge description.
// swagger:model challenge
type Challenge struct {
	// Token of the challenge
	//
	// required: true
	Token string `json:"token"`
	// Value of the challenge
	//
	// required: true
	Value string `json:"value"`
}

func (c *Challenge) Validate() url.Values {
	errs := url.Values{}

	// check if the token empty
	if c.Token == "" {
		errs.Add("token", "The token field is required!")
	}

	// check if the value empty
	if c.Value == "" {
		errs.Add("value", "The value field is required!")
	}

	return errs
}