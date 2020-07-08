/*
 * METRO2 API
 *
 * Moov Metro2 ([Automated Clearing House](https://en.wikipedia.org/wiki/Automated_Clearing_House)) implements an HTTP API for creating, parsing and validating Metro2 files. Metro2 is an open-source consumer credit history report for credit report file creation and validation.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// K3Segment struct for K3Segment
type K3Segment struct {
	SegmentIdentifier            string `json:"segmentIdentifier"`
	AgencyIdentifier             int32  `json:"agencyIdentifier,omitempty"`
	AccountNumber                string `json:"accountNumber,omitempty"`
	MortgageIdentificationNumber string `json:"mortgageIdentificationNumber,omitempty"`
}