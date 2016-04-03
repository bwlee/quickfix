package mdreqgrp

func New(nomdentrytypes []NoMDEntryTypes) *MDReqGrp {
	var m MDReqGrp
	m.SetNoMDEntryTypes(nomdentrytypes)
	return &m
}

//NoMDEntryTypes is a repeating group in MDReqGrp
type NoMDEntryTypes struct {
	//MDEntryType is a required field for NoMDEntryTypes.
	MDEntryType string `fix:"269"`
}

//MDReqGrp is a fix50 Component
type MDReqGrp struct {
	//NoMDEntryTypes is a required field for MDReqGrp.
	NoMDEntryTypes []NoMDEntryTypes `fix:"267"`
}

func (m *MDReqGrp) SetNoMDEntryTypes(v []NoMDEntryTypes) { m.NoMDEntryTypes = v }
