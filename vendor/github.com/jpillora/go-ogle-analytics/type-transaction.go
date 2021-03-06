package ga

import "net/url"

//WARNING: This file was generated. Do not edit.

//Transaction Hit Type
type Transaction struct {
	iD              string
	affiliation     string
	affiliationSet  bool
	revenue         float64
	revenueSet      bool
	shipping        float64
	shippingSet     bool
	tax             float64
	taxSet          bool
	currencyCode    string
	currencyCodeSet bool
}

// NewTransaction creates a new Transaction Hit Type.
// A unique identifier for the transaction. This value should
// be the same for both the Transaction hit and Items hits
// associated to the particular transaction.

func NewTransaction(iD string) *Transaction {
	h := &Transaction{
		iD: iD,
	}
	return h
}

func (h *Transaction) addFields(v url.Values) error {
	v.Add("ti", h.iD)
	if h.affiliationSet {
		v.Add("ta", h.affiliation)
	}
	if h.revenueSet {
		v.Add("tr", float2str(h.revenue))
	}
	if h.shippingSet {
		v.Add("ts", float2str(h.shipping))
	}
	if h.taxSet {
		v.Add("tt", float2str(h.tax))
	}
	if h.currencyCodeSet {
		v.Add("cu", h.currencyCode)
	}
	return nil
}

// Specifies the affiliation or store name.
func (h *Transaction) Affiliation(affiliation string) *Transaction {
	h.affiliation = affiliation
	h.affiliationSet = true
	return h
}

// Specifies the total revenue associated with the transaction.
// This value should include any shipping or tax costs.
func (h *Transaction) Revenue(revenue float64) *Transaction {
	h.revenue = revenue
	h.revenueSet = true
	return h
}

// Specifies the total shipping cost of the transaction.
func (h *Transaction) Shipping(shipping float64) *Transaction {
	h.shipping = shipping
	h.shippingSet = true
	return h
}

// Specifies the total tax of the transaction.
func (h *Transaction) Tax(tax float64) *Transaction {
	h.tax = tax
	h.taxSet = true
	return h
}

// When present indicates the local currency for all transaction
// currency values. Value should be a valid ISO 4217 currency
// code.
func (h *Transaction) CurrencyCode(currencyCode string) *Transaction {
	h.currencyCode = currencyCode
	h.currencyCodeSet = true
	return h
}

func (h *Transaction) Copy() *Transaction {
	c := *h
	return &c
}
