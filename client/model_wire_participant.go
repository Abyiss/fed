/*
 * FED API
 *
 * FED API is designed to create FEDACH and FEDWIRE dictionaries.  The FEDACH dictionary contains receiving depository financial institutions (RDFI’s) which are qualified to receive ACH entries.  The FEDWIRE dictionary contains receiving depository financial institutions (RDFI’s) which are qualified to receive WIRE entries.  This project implements a modern REST HTTP API for FEDACH Dictionary and FEDWIRE Dictionary.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// WireParticipant WIREParticipant holds a FedWIRE dir routing record as defined by Fed WIRE Format.  https://frbservices.org/EPaymentsDirectory/fedwireFormat.html
type WireParticipant struct {
	// The institution's routing number
	RoutingNumber string `json:"routingNumber,omitempty"`
	// Short name of financial institution
	TelegraphicName string `json:"telegraphicName,omitempty"`
	// Financial Institution Name
	CustomerName string       `json:"customerName,omitempty"`
	WireLocation WireLocation `json:"wireLocation,omitempty"`
	// Designates funds transfer status  * `Y` - Eligible * `N` - Ineligible
	FundsTransferStatus string `json:"fundsTransferStatus,omitempty"`
	// Designates funds settlement only status   * `S` - Settlement-Only
	FundsSettlementOnlyStatus string `json:"fundsSettlementOnlyStatus,omitempty"`
	// Designates book entry securities transfer status  * `Y` - Eligible * `N` - Ineligible
	BookEntrySecuritiesTransferStatus string `json:"bookEntrySecuritiesTransferStatus,omitempty"`
	// Date of last revision  * YYYYMMDD * Blank
	Date string `json:"date,omitempty"`
}
