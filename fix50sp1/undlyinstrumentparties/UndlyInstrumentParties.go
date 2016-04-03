package undlyinstrumentparties

import (
	"github.com/quickfixgo/quickfix/fix50sp1/undlyinstrumentptyssubgrp"
)

func New() *UndlyInstrumentParties {
	var m UndlyInstrumentParties
	return &m
}

//NoUndlyInstrumentParties is a repeating group in UndlyInstrumentParties
type NoUndlyInstrumentParties struct {
	//UndlyInstrumentPartyID is a non-required field for NoUndlyInstrumentParties.
	UndlyInstrumentPartyID *string `fix:"1059"`
	//UndlyInstrumentPartyIDSource is a non-required field for NoUndlyInstrumentParties.
	UndlyInstrumentPartyIDSource *string `fix:"1060"`
	//UndlyInstrumentPartyRole is a non-required field for NoUndlyInstrumentParties.
	UndlyInstrumentPartyRole *int `fix:"1061"`
	//UndlyInstrumentPtysSubGrp is a non-required component for NoUndlyInstrumentParties.
	UndlyInstrumentPtysSubGrp *undlyinstrumentptyssubgrp.UndlyInstrumentPtysSubGrp
}

//UndlyInstrumentParties is a fix50sp1 Component
type UndlyInstrumentParties struct {
	//NoUndlyInstrumentParties is a non-required field for UndlyInstrumentParties.
	NoUndlyInstrumentParties []NoUndlyInstrumentParties `fix:"1058,omitempty"`
}

func (m *UndlyInstrumentParties) SetNoUndlyInstrumentParties(v []NoUndlyInstrumentParties) {
	m.NoUndlyInstrumentParties = v
}
