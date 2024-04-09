package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOrder

type CreateOrderRequest struct {
	date          string `json:"date"`
	customer      string `json:"customer"`
	address       string `json:"address"`
	city          string `json:"city"`
	postcode      string `json:"postcode"`
	country       string `json:"country"`
	amount        int64  `json:"amount"`
	status        string `json:"status"`
	deleted       string `json:"deleted"`
	last_modified string `json:"last_modified"`
}

func (r *CreateOrderRequest) Validate() error {
	if r.date == "" && r.customer == "" && r.address == "" &&
		r.city == "" && r.postcode == "" && r.country == "" &&
		r.amount == 0 && r.status == "" && r.deleted == "" &&
		r.last_modified == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.date == "" {
		return errParamIsRequired("role", "string")
	}
	if r.customer == "" {
		return errParamIsRequired("customer", "string")
	}
	if r.address == "" {
		return errParamIsRequired("address", "string")
	}
	if r.city == "" {
		return errParamIsRequired("city", "string")
	}
	if r.postcode == "" {
		return errParamIsRequired("postcode", "string")
	}
	if r.amount <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	if r.deleted == "" {
		return errParamIsRequired("deleted", "string")
	}

	if r.status == "" {
		return errParamIsRequired("status", "string")
	}
	if r.last_modified == "" {
		return errParamIsRequired("last_modified", "string")
	}
	return nil
}

// UpdateOrder

type UpdateOrderRequest struct {
	date          string `json:"date"`
	customer      string `json:"customer"`
	address       string `json:"address"`
	city          string `json:"city"`
	postcode      string `json:"postcode"`
	country       string `json:"country"`
	amount        int64  `json:"amount"`
	status        string `json:"status"`
	deleted       string `json:"deleted"`
	last_modified string `json:"last_modified"`
}

func (r *UpdateOrderRequest) Validate() error {
	// If any field is provided, validation is truthy

	if r.date != "" || r.customer != "" || r.address != "" || r.city != "" ||
		r.postcode != "" || r.country != "" || r.amount > 0 || r.status != "" ||
		r.deleted != "" || r.last_modified != "" {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
