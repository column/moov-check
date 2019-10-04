/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// FileHeader struct for FileHeader
type FileHeader struct {
	// FileHeader ID
	ID string `json:"ID,omitempty"`
	// StandardLevel identifies the standard level of the file.  * `03` - DSTU X9.37 - 2003 * `30` - X9.100-187-2008 * `35` - X9.100-187-2013 and 2016
	StandardLevel string `json:"standardLevel"`
	// TestFileIndicator dentifies whether the file is a test or production file.  * `T` - Test File * `P` - Production File
	TestFileIndicator string `json:"testFileIndicator"`
	// ImmediateDestination is the routing and transit number of the Federal Reserve Bank (FRB) or receiver to which the file is being sent.
	ImmediateDestination string `json:"immediateDestination"`
	// ImmediateOrigin is the routing and transit number of the Federal Reserve Bank (FRB) or originator from which the file is being sent.
	ImmediateOrigin string `json:"immediateOrigin"`
	// FileCreationDate is the date that the immediate origin institution creates the file.
	FileCreationDate time.Time `json:"fileCreationDate"`
	// FileCreationTime is the time the immediate origin institution creates the file. (Format - hhmm, where - hh hour, mm minute)
	FileCreationTime string `json:"fileCreationTime"`
	// ResendIndicator Indicates whether the file has been previously transmitted. (Y - Yes, N - No)
	ResendIndicator string `json:"resendIndicator"`
	// ImmediateDestinationName Identifies the short name of the institution that receives the file.
	ImmediateDestinationName string `json:"immediateDestinationName,omitempty"`
	// immediateOriginName identifies the short name of the institution that sends the file.
	ImmediateOriginName string `json:"immediateOriginName,omitempty"`
	// FileIDModifier is a code that permits multiple files, created on the same date, same time and between the same institutions, to be distinguished one from another. If all of the following fields in a previous file are equal to the same fields in this file: FileHeader ImmediateDestination, ImmediateOrigin, FileCreationDate, and FileCreationTime, it must be defined.
	FileIDModifier string `json:"fileIDModifier,omitempty"`
	// CountryCode is a 2-character code as approved by the International Organization for Standardization (ISO) used to identify the country in which the payer bank is located.
	CountryCode string `json:"countryCode,omitempty"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField,omitempty"`
	// CompanionDocumentIndicator identifies a field used to indicate the Companion Document being used. Shall be present only under clearing arrangements. Companion Document usage and values defined by clearing arrangements. Values: 0–9 Reserved for United States use A–J Reserved for Canadian use Other - as defined by clearing arrangements.
	CompanionDocumentIndicator string `json:"companionDocumentIndicator,omitempty"`
}
