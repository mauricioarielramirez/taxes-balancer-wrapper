package models

import "taxes-balancer-wrapper/pkg/business/models/types"

type Product struct {
	ID               string
	Name             string
	ShortDescription string
	Currency         types.Currency
	CurrencyType     types.CurrencyType
	Amount           float32
	UseImmediatly    bool
}
