package clrinstgrp

//New returns an initialized ClrInstGrp instance
func New() *ClrInstGrp {
	var m ClrInstGrp
	return &m
}

//NoClearingInstructions is a repeating group in ClrInstGrp
type NoClearingInstructions struct {
	//ClearingInstruction is a non-required field for NoClearingInstructions.
	ClearingInstruction *int `fix:"577"`
}

//NewNoClearingInstructions returns an initialized NoClearingInstructions instance
func NewNoClearingInstructions() *NoClearingInstructions {
	var m NoClearingInstructions
	return &m
}

func (m *NoClearingInstructions) SetClearingInstruction(v int) { m.ClearingInstruction = &v }

//ClrInstGrp is a fix50sp2 Component
type ClrInstGrp struct {
	//NoClearingInstructions is a non-required field for ClrInstGrp.
	NoClearingInstructions []NoClearingInstructions `fix:"576,omitempty"`
}

func (m *ClrInstGrp) SetNoClearingInstructions(v []NoClearingInstructions) {
	m.NoClearingInstructions = v
}
