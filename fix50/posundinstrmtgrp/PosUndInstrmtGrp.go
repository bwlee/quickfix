package posundinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50/underlyingamount"
	"github.com/quickfixgo/quickfix/fix50/underlyinginstrument"
)

func New() *PosUndInstrmtGrp {
	var m PosUndInstrmtGrp
	return &m
}

//NoUnderlyings is a repeating group in PosUndInstrmtGrp
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
	//UnderlyingSettlPrice is a non-required field for NoUnderlyings.
	UnderlyingSettlPrice *float64 `fix:"732"`
	//UnderlyingSettlPriceType is a non-required field for NoUnderlyings.
	UnderlyingSettlPriceType *int `fix:"733"`
	//UnderlyingAmount is a non-required component for NoUnderlyings.
	UnderlyingAmount *underlyingamount.UnderlyingAmount
	//UnderlyingDeliveryAmount is a non-required field for NoUnderlyings.
	UnderlyingDeliveryAmount *float64 `fix:"1037"`
}

//PosUndInstrmtGrp is a fix50 Component
type PosUndInstrmtGrp struct {
	//NoUnderlyings is a non-required field for PosUndInstrmtGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func (m *PosUndInstrmtGrp) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
