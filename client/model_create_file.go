/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateFile struct for CreateFile
type CreateFile struct {
	// File ID
	ID          string       `json:"ID,omitempty"`
	FileHeader  FileHeader   `json:"fileHeader"`
	CashLetters []CashLetter `json:"cashLetters,omitempty"`
	Bundles     []Bundle     `json:"bundles,omitempty"`
	FileControl FileControl  `json:"fileControl,omitempty"`
}
