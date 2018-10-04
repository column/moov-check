// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// CashLetter contains CashLetterHeader, CashLetterControl and Bundle records.
type CashLetter struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// CashLetterHeader is a Cash Letter Header Record
	CashLetterHeader *CashLetterHeader `json:"cashLetterHeader,omitempty"`
	// Bundles is an array of Bundle
	Bundles []*Bundle `json:"bundles,omitempty"`
	// ReturnBundles is an array of ReturnBundle
	//ReturnBundles []*ReturnBundle `json:"returnBundle,omitempty"`
	// RoutingNumberSummary is an X9 RoutingNumberSummary
	RoutingNumberSummary []*RoutingNumberSummary `json:"routingNumberSummary,omitempty"`
	// currentBundle is the currentBundle being parsed
	currentBundle *Bundle
	// currentReturnBundle is the current ReturnBundle being parsed
	//currentReturnBundle *ReturnBundle
	// RoutingNumberSummary is an X9 RoutingNumberSummary
	currentRoutingNumberSummary *RoutingNumberSummary
	// CashLetterControl is a Cash Letter Control Record
	CashLetterControl *CashLetterControl `json:"cashLetterControl,omitempty"`
}

// NewCashLetter takes a CashLetterHeader and returns a CashLetter
func NewCashLetter(clh *CashLetterHeader) CashLetter {
	cl := CashLetter{}
	cl.SetControl(NewCashLetterControl())
	cl.SetHeader(clh)
	return cl
}

// Validate performs X9 validations and format rule checks and returns an error if not Validated
func (cl *CashLetter) Validate() error {
	// ToDo:  If CashLetterRecordTypeIndicator is "N", There should be no bundle, it is an empty cash letter
	return nil
}

// ToDo:  Add verify

// build by building a valid CashLetter by building a CashLetterControl. An error is returned if
// the CashLetter being built has invalid records.
func (cl *CashLetter) build() error {

	// Requires a valid BundleHeader
	if err := cl.CashLetterHeader.Validate(); err != nil {
		return err
	}

	// if len(cl.Bundles) == 0, it is an empty cash letter
	cashLetterBundleCount := len(cl.Bundles)
	cashLetterItemsCount := 0
	cashLetterTotalAmount := 0
	cashLetterImagesCount := 0

	// ToDo: Sequences, research if a bundle should be part of the cash letter item count?

	// Bundles
	for _, b := range cl.Bundles {
		// Validate Bundle
		if err := b.Validate(); err != nil {
			return err
		}

		// Check Items
		for _, cd := range b.Checks {

			if err := b.build(); err != nil {
				return err
			}
			cashLetterItemsCount = cashLetterItemsCount + 1
			cashLetterItemsCount = cashLetterItemsCount + len(cd.CheckDetailAddendumA) + len(cd.CheckDetailAddendumB) + len(cd.CheckDetailAddendumC)
			cashLetterItemsCount = cashLetterItemsCount + len(cd.ImageViewDetail) + len(cd.ImageViewData) + len(cd.ImageViewAnalysis)
			cashLetterTotalAmount = cashLetterTotalAmount + cd.ItemAmount
			cashLetterImagesCount = cashLetterImagesCount + len(cd.ImageViewDetail)
		}

		// Returns Items
		for _, rd := range b.Returns {

			if err := b.build(); err != nil {
				return err
			}
			cashLetterItemsCount = cashLetterItemsCount + 1
			cashLetterItemsCount = cashLetterItemsCount + len(rd.ReturnDetailAddendumA) + len(rd.ReturnDetailAddendumB) + len(rd.ReturnDetailAddendumC) + len(rd.ReturnDetailAddendumD)
			cashLetterItemsCount = cashLetterItemsCount + len(rd.ImageViewDetail) + len(rd.ImageViewData) + len(rd.ImageViewAnalysis)
			cashLetterTotalAmount = cashLetterTotalAmount + rd.ItemAmount
			cashLetterImagesCount = cashLetterImagesCount + len(rd.ImageViewDetail)
		}

	}

	// build a CashLetterControl record
	clc := NewCashLetterControl()
	clc.CashLetterBundleCount = cashLetterBundleCount
	clc.CashLetterItemsCount = cashLetterItemsCount
	clc.CashLetterTotalAmount = cashLetterTotalAmount
	clc.CashLetterImagesCount = cashLetterImagesCount
	clc.ECEInstitutionName = ""
	clc.CreditTotalIndicator = 0
	cl.CashLetterControl = clc
	return nil
}

// Create creates a CashLetter of Bundles containing CheckDetail or ReturnDetail
func (cl *CashLetter) Create() error {
	if err := cl.build(); err != nil {
		return err
	}
	return cl.Validate()
}

// SetHeader appends a CashLetterHeader to the CashLetter
func (cl *CashLetter) SetHeader(cashLetterHeader *CashLetterHeader) {
	cl.CashLetterHeader = cashLetterHeader
}

// GetHeader returns the current CashLetter header
func (cl *CashLetter) GetHeader() *CashLetterHeader {
	return cl.CashLetterHeader
}

// SetControl appends a CashLetterControl to the CashLetter
func (cl *CashLetter) SetControl(cashLetterControl *CashLetterControl) {
	cl.CashLetterControl = cashLetterControl
}

// GetControl returns the current CashLetter Control
func (cl *CashLetter) GetControl() *CashLetterControl {
	return cl.CashLetterControl
}

// AddBundle appends a Bundle to the CashLetter
func (cl *CashLetter) AddBundle(bundle *Bundle) []*Bundle {
	cl.Bundles = append(cl.Bundles, bundle)
	return cl.Bundles
}

// GetBundles returns a slice of Bundles for the CashLetter
func (cl *CashLetter) GetBundles() []*Bundle {
	return cl.Bundles
}

// AddRoutingNumberSummary appends a RoutingNumberSummary to the CashLetter
func (cl *CashLetter) AddRoutingNumberSummary(rns *RoutingNumberSummary) []*RoutingNumberSummary {
	cl.RoutingNumberSummary = append(cl.RoutingNumberSummary, rns)
	return cl.RoutingNumberSummary
}

// GetRoutingNumberSummary returns a slice of RoutingNumberSummary for the CashLetter
func (cl *CashLetter) GetRoutingNumberSummary() []*RoutingNumberSummary {
	return cl.RoutingNumberSummary
}