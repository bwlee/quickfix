//Package collateralrequest msg type = AX.
package collateralrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/execcollgrp"
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/miscfeesgrp"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/trdcollgrp"
	"github.com/quickfixgo/quickfix/fix50/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtcollgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a CollateralRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AX"`
	fixt11.Header
	//CollReqID is a required field for CollateralRequest.
	CollReqID string `fix:"894"`
	//CollAsgnReason is a required field for CollateralRequest.
	CollAsgnReason int `fix:"895"`
	//TransactTime is a required field for CollateralRequest.
	TransactTime time.Time `fix:"60"`
	//ExpireTime is a non-required field for CollateralRequest.
	ExpireTime *time.Time `fix:"126"`
	//Parties is a non-required component for CollateralRequest.
	Parties *parties.Parties
	//Account is a non-required field for CollateralRequest.
	Account *string `fix:"1"`
	//AccountType is a non-required field for CollateralRequest.
	AccountType *int `fix:"581"`
	//ClOrdID is a non-required field for CollateralRequest.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for CollateralRequest.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for CollateralRequest.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for CollateralRequest.
	SecondaryClOrdID *string `fix:"526"`
	//ExecCollGrp is a non-required component for CollateralRequest.
	ExecCollGrp *execcollgrp.ExecCollGrp
	//TrdCollGrp is a non-required component for CollateralRequest.
	TrdCollGrp *trdcollgrp.TrdCollGrp
	//Instrument is a non-required component for CollateralRequest.
	Instrument *instrument.Instrument
	//FinancingDetails is a non-required component for CollateralRequest.
	FinancingDetails *financingdetails.FinancingDetails
	//SettlDate is a non-required field for CollateralRequest.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralRequest.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralRequest.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralRequest.
	Currency *string `fix:"15"`
	//InstrmtLegGrp is a non-required component for CollateralRequest.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//UndInstrmtCollGrp is a non-required component for CollateralRequest.
	UndInstrmtCollGrp *undinstrmtcollgrp.UndInstrmtCollGrp
	//MarginExcess is a non-required field for CollateralRequest.
	MarginExcess *float64 `fix:"899"`
	//TotalNetValue is a non-required field for CollateralRequest.
	TotalNetValue *float64 `fix:"900"`
	//CashOutstanding is a non-required field for CollateralRequest.
	CashOutstanding *float64 `fix:"901"`
	//TrdRegTimestamps is a non-required component for CollateralRequest.
	TrdRegTimestamps *trdregtimestamps.TrdRegTimestamps
	//Side is a non-required field for CollateralRequest.
	Side *string `fix:"54"`
	//MiscFeesGrp is a non-required component for CollateralRequest.
	MiscFeesGrp *miscfeesgrp.MiscFeesGrp
	//Price is a non-required field for CollateralRequest.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for CollateralRequest.
	PriceType *int `fix:"423"`
	//AccruedInterestAmt is a non-required field for CollateralRequest.
	AccruedInterestAmt *float64 `fix:"159"`
	//EndAccruedInterestAmt is a non-required field for CollateralRequest.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for CollateralRequest.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for CollateralRequest.
	EndCash *float64 `fix:"922"`
	//SpreadOrBenchmarkCurveData is a non-required component for CollateralRequest.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//Stipulations is a non-required component for CollateralRequest.
	Stipulations *stipulations.Stipulations
	//TradingSessionID is a non-required field for CollateralRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for CollateralRequest.
	TradingSessionSubID *string `fix:"625"`
	//SettlSessID is a non-required field for CollateralRequest.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for CollateralRequest.
	SettlSessSubID *string `fix:"717"`
	//ClearingBusinessDate is a non-required field for CollateralRequest.
	ClearingBusinessDate *string `fix:"715"`
	//Text is a non-required field for CollateralRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for CollateralRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for CollateralRequest.
	EncodedText *string `fix:"355"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized CollateralRequest instance
func New(collreqid string, collasgnreason int, transacttime time.Time) *Message {
	var m Message
	m.SetCollReqID(collreqid)
	m.SetCollAsgnReason(collasgnreason)
	m.SetTransactTime(transacttime)
	return &m
}

func (m *Message) SetCollReqID(v string)                                   { m.CollReqID = v }
func (m *Message) SetCollAsgnReason(v int)                                 { m.CollAsgnReason = v }
func (m *Message) SetTransactTime(v time.Time)                             { m.TransactTime = v }
func (m *Message) SetExpireTime(v time.Time)                               { m.ExpireTime = &v }
func (m *Message) SetParties(v parties.Parties)                            { m.Parties = &v }
func (m *Message) SetAccount(v string)                                     { m.Account = &v }
func (m *Message) SetAccountType(v int)                                    { m.AccountType = &v }
func (m *Message) SetClOrdID(v string)                                     { m.ClOrdID = &v }
func (m *Message) SetOrderID(v string)                                     { m.OrderID = &v }
func (m *Message) SetSecondaryOrderID(v string)                            { m.SecondaryOrderID = &v }
func (m *Message) SetSecondaryClOrdID(v string)                            { m.SecondaryClOrdID = &v }
func (m *Message) SetExecCollGrp(v execcollgrp.ExecCollGrp)                { m.ExecCollGrp = &v }
func (m *Message) SetTrdCollGrp(v trdcollgrp.TrdCollGrp)                   { m.TrdCollGrp = &v }
func (m *Message) SetInstrument(v instrument.Instrument)                   { m.Instrument = &v }
func (m *Message) SetFinancingDetails(v financingdetails.FinancingDetails) { m.FinancingDetails = &v }
func (m *Message) SetSettlDate(v string)                                   { m.SettlDate = &v }
func (m *Message) SetQuantity(v float64)                                   { m.Quantity = &v }
func (m *Message) SetQtyType(v int)                                        { m.QtyType = &v }
func (m *Message) SetCurrency(v string)                                    { m.Currency = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp)          { m.InstrmtLegGrp = &v }
func (m *Message) SetUndInstrmtCollGrp(v undinstrmtcollgrp.UndInstrmtCollGrp) {
	m.UndInstrmtCollGrp = &v
}
func (m *Message) SetMarginExcess(v float64)                               { m.MarginExcess = &v }
func (m *Message) SetTotalNetValue(v float64)                              { m.TotalNetValue = &v }
func (m *Message) SetCashOutstanding(v float64)                            { m.CashOutstanding = &v }
func (m *Message) SetTrdRegTimestamps(v trdregtimestamps.TrdRegTimestamps) { m.TrdRegTimestamps = &v }
func (m *Message) SetSide(v string)                                        { m.Side = &v }
func (m *Message) SetMiscFeesGrp(v miscfeesgrp.MiscFeesGrp)                { m.MiscFeesGrp = &v }
func (m *Message) SetPrice(v float64)                                      { m.Price = &v }
func (m *Message) SetPriceType(v int)                                      { m.PriceType = &v }
func (m *Message) SetAccruedInterestAmt(v float64)                         { m.AccruedInterestAmt = &v }
func (m *Message) SetEndAccruedInterestAmt(v float64)                      { m.EndAccruedInterestAmt = &v }
func (m *Message) SetStartCash(v float64)                                  { m.StartCash = &v }
func (m *Message) SetEndCash(v float64)                                    { m.EndCash = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetStipulations(v stipulations.Stipulations) { m.Stipulations = &v }
func (m *Message) SetTradingSessionID(v string)                { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)             { m.TradingSessionSubID = &v }
func (m *Message) SetSettlSessID(v string)                     { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)                  { m.SettlSessSubID = &v }
func (m *Message) SetClearingBusinessDate(v string)            { m.ClearingBusinessDate = &v }
func (m *Message) SetText(v string)                            { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                     { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                     { m.EncodedText = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.ApplVerID_FIX50, "AX", r
}
