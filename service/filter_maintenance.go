// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/gsapkal/hap/characteristic"
)

const TypeFilterMaintenance = "BA"

type FilterMaintenance struct {
	*S

	FilterChangeIndication *characteristic.FilterChangeIndication
}

func NewFilterMaintenance() *FilterMaintenance {
	s := FilterMaintenance{}
	s.S = New(TypeFilterMaintenance)

	s.FilterChangeIndication = characteristic.NewFilterChangeIndication()
	s.AddC(s.FilterChangeIndication.C)

	return &s
}
