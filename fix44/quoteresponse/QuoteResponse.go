//Package quoteresponse msg type = AJ.
package quoteresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/legstipulations"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/orderqtydata"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fix44/yielddata"
	"time"
)

//NoQuoteQualifiers is a repeating group in QuoteResponse
type NoQuoteQualifiers struct {
	//QuoteQualifier is a non-required field for NoQuoteQualifiers.
	QuoteQualifier *string `fix:"695"`
}

//NewNoQuoteQualifiers returns an initialized NoQuoteQualifiers instance
func NewNoQuoteQualifiers() *NoQuoteQualifiers {
	var m NoQuoteQualifiers
	return &m
}

func (m *NoQuoteQualifiers) SetQuoteQualifier(v string) { m.QuoteQualifier = &v }

//NoUnderlyings is a repeating group in QuoteResponse
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
}

//NewNoUnderlyings returns an initialized NoUnderlyings instance
func NewNoUnderlyings() *NoUnderlyings {
	var m NoUnderlyings
	return &m
}

func (m *NoUnderlyings) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}

//NoLegs is a repeating group in QuoteResponse
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegStipulations is a non-required component for NoLegs.
	LegStipulations *legstipulations.LegStipulations
	//NestedParties is a non-required component for NoLegs.
	NestedParties *nestedparties.NestedParties
	//LegPriceType is a non-required field for NoLegs.
	LegPriceType *int `fix:"686"`
	//LegBidPx is a non-required field for NoLegs.
	LegBidPx *float64 `fix:"681"`
	//LegOfferPx is a non-required field for NoLegs.
	LegOfferPx *float64 `fix:"684"`
	//LegBenchmarkCurveData is a non-required component for NoLegs.
	LegBenchmarkCurveData *legbenchmarkcurvedata.LegBenchmarkCurveData
}

//NewNoLegs returns an initialized NoLegs instance
func NewNoLegs() *NoLegs {
	var m NoLegs
	return &m
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg)       { m.InstrumentLeg = &v }
func (m *NoLegs) SetLegQty(v float64)                                  { m.LegQty = &v }
func (m *NoLegs) SetLegSwapType(v int)                                 { m.LegSwapType = &v }
func (m *NoLegs) SetLegSettlType(v string)                             { m.LegSettlType = &v }
func (m *NoLegs) SetLegSettlDate(v string)                             { m.LegSettlDate = &v }
func (m *NoLegs) SetLegStipulations(v legstipulations.LegStipulations) { m.LegStipulations = &v }
func (m *NoLegs) SetNestedParties(v nestedparties.NestedParties)       { m.NestedParties = &v }
func (m *NoLegs) SetLegPriceType(v int)                                { m.LegPriceType = &v }
func (m *NoLegs) SetLegBidPx(v float64)                                { m.LegBidPx = &v }
func (m *NoLegs) SetLegOfferPx(v float64)                              { m.LegOfferPx = &v }
func (m *NoLegs) SetLegBenchmarkCurveData(v legbenchmarkcurvedata.LegBenchmarkCurveData) {
	m.LegBenchmarkCurveData = &v
}

//Message is a QuoteResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"AJ"`
	fix44.Header
	//QuoteRespID is a required field for QuoteResponse.
	QuoteRespID string `fix:"693"`
	//QuoteID is a non-required field for QuoteResponse.
	QuoteID *string `fix:"117"`
	//QuoteRespType is a required field for QuoteResponse.
	QuoteRespType int `fix:"694"`
	//ClOrdID is a non-required field for QuoteResponse.
	ClOrdID *string `fix:"11"`
	//OrderCapacity is a non-required field for QuoteResponse.
	OrderCapacity *string `fix:"528"`
	//IOIID is a non-required field for QuoteResponse.
	IOIID *string `fix:"23"`
	//QuoteType is a non-required field for QuoteResponse.
	QuoteType *int `fix:"537"`
	//NoQuoteQualifiers is a non-required field for QuoteResponse.
	NoQuoteQualifiers []NoQuoteQualifiers `fix:"735,omitempty"`
	//Parties is a non-required component for QuoteResponse.
	Parties *parties.Parties
	//TradingSessionID is a non-required field for QuoteResponse.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for QuoteResponse.
	TradingSessionSubID *string `fix:"625"`
	//Instrument is a required component for QuoteResponse.
	instrument.Instrument
	//FinancingDetails is a non-required component for QuoteResponse.
	FinancingDetails *financingdetails.FinancingDetails
	//NoUnderlyings is a non-required field for QuoteResponse.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//Side is a non-required field for QuoteResponse.
	Side *string `fix:"54"`
	//OrderQtyData is a non-required component for QuoteResponse.
	OrderQtyData *orderqtydata.OrderQtyData
	//SettlType is a non-required field for QuoteResponse.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for QuoteResponse.
	SettlDate *string `fix:"64"`
	//SettlDate2 is a non-required field for QuoteResponse.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for QuoteResponse.
	OrderQty2 *float64 `fix:"192"`
	//Currency is a non-required field for QuoteResponse.
	Currency *string `fix:"15"`
	//Stipulations is a non-required component for QuoteResponse.
	Stipulations *stipulations.Stipulations
	//Account is a non-required field for QuoteResponse.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for QuoteResponse.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for QuoteResponse.
	AccountType *int `fix:"581"`
	//NoLegs is a non-required field for QuoteResponse.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//BidPx is a non-required field for QuoteResponse.
	BidPx *float64 `fix:"132"`
	//OfferPx is a non-required field for QuoteResponse.
	OfferPx *float64 `fix:"133"`
	//MktBidPx is a non-required field for QuoteResponse.
	MktBidPx *float64 `fix:"645"`
	//MktOfferPx is a non-required field for QuoteResponse.
	MktOfferPx *float64 `fix:"646"`
	//MinBidSize is a non-required field for QuoteResponse.
	MinBidSize *float64 `fix:"647"`
	//BidSize is a non-required field for QuoteResponse.
	BidSize *float64 `fix:"134"`
	//MinOfferSize is a non-required field for QuoteResponse.
	MinOfferSize *float64 `fix:"648"`
	//OfferSize is a non-required field for QuoteResponse.
	OfferSize *float64 `fix:"135"`
	//ValidUntilTime is a non-required field for QuoteResponse.
	ValidUntilTime *time.Time `fix:"62"`
	//BidSpotRate is a non-required field for QuoteResponse.
	BidSpotRate *float64 `fix:"188"`
	//OfferSpotRate is a non-required field for QuoteResponse.
	OfferSpotRate *float64 `fix:"190"`
	//BidForwardPoints is a non-required field for QuoteResponse.
	BidForwardPoints *float64 `fix:"189"`
	//OfferForwardPoints is a non-required field for QuoteResponse.
	OfferForwardPoints *float64 `fix:"191"`
	//MidPx is a non-required field for QuoteResponse.
	MidPx *float64 `fix:"631"`
	//BidYield is a non-required field for QuoteResponse.
	BidYield *float64 `fix:"632"`
	//MidYield is a non-required field for QuoteResponse.
	MidYield *float64 `fix:"633"`
	//OfferYield is a non-required field for QuoteResponse.
	OfferYield *float64 `fix:"634"`
	//TransactTime is a non-required field for QuoteResponse.
	TransactTime *time.Time `fix:"60"`
	//OrdType is a non-required field for QuoteResponse.
	OrdType *string `fix:"40"`
	//BidForwardPoints2 is a non-required field for QuoteResponse.
	BidForwardPoints2 *float64 `fix:"642"`
	//OfferForwardPoints2 is a non-required field for QuoteResponse.
	OfferForwardPoints2 *float64 `fix:"643"`
	//SettlCurrBidFxRate is a non-required field for QuoteResponse.
	SettlCurrBidFxRate *float64 `fix:"656"`
	//SettlCurrOfferFxRate is a non-required field for QuoteResponse.
	SettlCurrOfferFxRate *float64 `fix:"657"`
	//SettlCurrFxRateCalc is a non-required field for QuoteResponse.
	SettlCurrFxRateCalc *string `fix:"156"`
	//Commission is a non-required field for QuoteResponse.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for QuoteResponse.
	CommType *string `fix:"13"`
	//CustOrderCapacity is a non-required field for QuoteResponse.
	CustOrderCapacity *int `fix:"582"`
	//ExDestination is a non-required field for QuoteResponse.
	ExDestination *string `fix:"100"`
	//Text is a non-required field for QuoteResponse.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for QuoteResponse.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for QuoteResponse.
	EncodedText *string `fix:"355"`
	//Price is a non-required field for QuoteResponse.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for QuoteResponse.
	PriceType *int `fix:"423"`
	//SpreadOrBenchmarkCurveData is a non-required component for QuoteResponse.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData is a non-required component for QuoteResponse.
	YieldData *yielddata.YieldData
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized QuoteResponse instance
func New(quoterespid string, quoteresptype int, instrument instrument.Instrument) *Message {
	var m Message
	m.SetQuoteRespID(quoterespid)
	m.SetQuoteRespType(quoteresptype)
	m.SetInstrument(instrument)
	return &m
}

func (m *Message) SetQuoteRespID(v string)                                 { m.QuoteRespID = v }
func (m *Message) SetQuoteID(v string)                                     { m.QuoteID = &v }
func (m *Message) SetQuoteRespType(v int)                                  { m.QuoteRespType = v }
func (m *Message) SetClOrdID(v string)                                     { m.ClOrdID = &v }
func (m *Message) SetOrderCapacity(v string)                               { m.OrderCapacity = &v }
func (m *Message) SetIOIID(v string)                                       { m.IOIID = &v }
func (m *Message) SetQuoteType(v int)                                      { m.QuoteType = &v }
func (m *Message) SetNoQuoteQualifiers(v []NoQuoteQualifiers)              { m.NoQuoteQualifiers = v }
func (m *Message) SetParties(v parties.Parties)                            { m.Parties = &v }
func (m *Message) SetTradingSessionID(v string)                            { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)                         { m.TradingSessionSubID = &v }
func (m *Message) SetInstrument(v instrument.Instrument)                   { m.Instrument = v }
func (m *Message) SetFinancingDetails(v financingdetails.FinancingDetails) { m.FinancingDetails = &v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)                      { m.NoUnderlyings = v }
func (m *Message) SetSide(v string)                                        { m.Side = &v }
func (m *Message) SetOrderQtyData(v orderqtydata.OrderQtyData)             { m.OrderQtyData = &v }
func (m *Message) SetSettlType(v string)                                   { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)                                   { m.SettlDate = &v }
func (m *Message) SetSettlDate2(v string)                                  { m.SettlDate2 = &v }
func (m *Message) SetOrderQty2(v float64)                                  { m.OrderQty2 = &v }
func (m *Message) SetCurrency(v string)                                    { m.Currency = &v }
func (m *Message) SetStipulations(v stipulations.Stipulations)             { m.Stipulations = &v }
func (m *Message) SetAccount(v string)                                     { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)                                   { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                                    { m.AccountType = &v }
func (m *Message) SetNoLegs(v []NoLegs)                                    { m.NoLegs = v }
func (m *Message) SetBidPx(v float64)                                      { m.BidPx = &v }
func (m *Message) SetOfferPx(v float64)                                    { m.OfferPx = &v }
func (m *Message) SetMktBidPx(v float64)                                   { m.MktBidPx = &v }
func (m *Message) SetMktOfferPx(v float64)                                 { m.MktOfferPx = &v }
func (m *Message) SetMinBidSize(v float64)                                 { m.MinBidSize = &v }
func (m *Message) SetBidSize(v float64)                                    { m.BidSize = &v }
func (m *Message) SetMinOfferSize(v float64)                               { m.MinOfferSize = &v }
func (m *Message) SetOfferSize(v float64)                                  { m.OfferSize = &v }
func (m *Message) SetValidUntilTime(v time.Time)                           { m.ValidUntilTime = &v }
func (m *Message) SetBidSpotRate(v float64)                                { m.BidSpotRate = &v }
func (m *Message) SetOfferSpotRate(v float64)                              { m.OfferSpotRate = &v }
func (m *Message) SetBidForwardPoints(v float64)                           { m.BidForwardPoints = &v }
func (m *Message) SetOfferForwardPoints(v float64)                         { m.OfferForwardPoints = &v }
func (m *Message) SetMidPx(v float64)                                      { m.MidPx = &v }
func (m *Message) SetBidYield(v float64)                                   { m.BidYield = &v }
func (m *Message) SetMidYield(v float64)                                   { m.MidYield = &v }
func (m *Message) SetOfferYield(v float64)                                 { m.OfferYield = &v }
func (m *Message) SetTransactTime(v time.Time)                             { m.TransactTime = &v }
func (m *Message) SetOrdType(v string)                                     { m.OrdType = &v }
func (m *Message) SetBidForwardPoints2(v float64)                          { m.BidForwardPoints2 = &v }
func (m *Message) SetOfferForwardPoints2(v float64)                        { m.OfferForwardPoints2 = &v }
func (m *Message) SetSettlCurrBidFxRate(v float64)                         { m.SettlCurrBidFxRate = &v }
func (m *Message) SetSettlCurrOfferFxRate(v float64)                       { m.SettlCurrOfferFxRate = &v }
func (m *Message) SetSettlCurrFxRateCalc(v string)                         { m.SettlCurrFxRateCalc = &v }
func (m *Message) SetCommission(v float64)                                 { m.Commission = &v }
func (m *Message) SetCommType(v string)                                    { m.CommType = &v }
func (m *Message) SetCustOrderCapacity(v int)                              { m.CustOrderCapacity = &v }
func (m *Message) SetExDestination(v string)                               { m.ExDestination = &v }
func (m *Message) SetText(v string)                                        { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                                 { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                                 { m.EncodedText = &v }
func (m *Message) SetPrice(v float64)                                      { m.Price = &v }
func (m *Message) SetPriceType(v int)                                      { m.PriceType = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetYieldData(v yielddata.YieldData) { m.YieldData = &v }

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
	return enum.BeginStringFIX44, "AJ", r
}
