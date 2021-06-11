package model

import (
	"reflect"
)

type SellingOptionMatchMode string

type Ticket struct {
	Id            int64                  `json:"id,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Price         float64                `json:"price,omitempty"`
	SellingOption SellingOptionMatchMode `json:"sellingOption,omitempty"`
}

var SellingOption = [...]SellingOptionMatchMode{
	"none",
	"allTogether",
	"even",
	"avoidOne",
}

func (ticketSellOption SellingOptionMatchMode) IsValid() bool {
	return contains(ticketSellOption, SellingOption)
}

func contains(val interface{}, array interface{}) bool {
	s := reflect.ValueOf(array)
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
			return true
		}
	}
	return false
}
