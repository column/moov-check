/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ImageViewAnalysis struct for ImageViewAnalysis
type ImageViewAnalysis struct {
	// ImageViewAnalysis ID
	ID string `json:"ID,omitempty"`
	// GlobalImageQuality is a code that indicates whether the image view was tested for any of the conditions related to image quality defined in the Image Quality Information.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	GlobalImageQuality string `json:"globalImageQuality,omitempty"`
	// GlobalImageUsability is a code that indicates whether the image view was tested for any of the conditions related to image usability defined in the Image Usability Information.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	GlobalImageUsability string `json:"globalImageUsability,omitempty"`
	// ImagingBankSpecificTest designates the capture institution may be able to perform specific tests that can indicate a potentially problematic image view caused by conditions other than those listed in the Image Quality and Image Usability Information fields. By mutual agreement, clearing partners can use the UserField to report the presence or absence of additional image conditions found through tests that are particular to the specific imaging institution. The meaning and interpretation of the User Field data must be understood and agreed upon between participants.   * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	ImagingBankSpecificTest string `json:"imagingBankSpecificTest,omitempty"`
	// PartialImage is a code that indicates if only a portion of the image view is represented digitally while the other portion is suspected to be missing or corrupt.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	PartialImage string `json:"partialImage,omitempty"`
	// ExcessiveImageSkew is a code that indicates if the image view skew exceeds an acceptable value. This value is specific to the imaging institution’s own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	ExcessiveImageSkew string `json:"excessiveImageSkew,omitempty"`
	// PiggybackImage is a code that indicates if a “piggyback” condition has been detected. With a “piggyback” condition, the intended image view may be extended, obscured, or replaced by image(s) of additional document(s). A piggyback occurs when two or more documents are fed together and captured as one document when only a single document should have been fed and captured.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	PiggybackImage string `json:"piggybackImage,omitempty"`
	// TooLightOrTooDark is a code that indicates if the image view is too light or too dark. The value is specific to the imaging institution’s own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	TooLightOrTooDark string `json:"tooLightOrTooDark,omitempty"`
	// StreaksAndOrBands is a A code that indicates if the image view is likely corrupted due to streaks and/or bands. Streaks and bands can be caused by such problems as dirt, dust, ink, or debris on a lens or in the optical path, and failures in the imaging equipment scanner.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	StreaksAndOrBands string `json:"streaksAndOrBands,omitempty"`
	// BelowMinimumImageSize is a code that indicates if the size of the compressed image view is below an acceptable value. The value is specific to the imaging institution’s own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	BelowMinimumImageSize string `json:"belowMinimumImageSize,omitempty"`
	// ExceedsMaximumImageSize is a code that indicates if the size of the compressed image view is above an acceptable value. The value is specific to the imaging institution’s own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	ExceedsMaximumImageSize string `json:"exceedsMaximumImageSize,omitempty"`
	// ImageEnabledPOD is a code that indicates if the image view was used within an image-enabled POD (Proof of Deposit) application.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	ImageEnabledPOD string `json:"imageEnabledPOD,omitempty"`
	// SourceDocumentBad is a code that indicates if it is possible to obtain a better image from the source document when it is known that the current image of the document is unusable.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	SourceDocumentBad string `json:"sourceDocumentBad,omitempty"`
	// DateUsability is a code that indicates if the date Area of Interest is usable and readable from the image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's own defined requirements and/or constraint  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	DateUsability string `json:"dateUsability,omitempty"`
	// PayeeUsability is a code that indicates if the payee name Area of Interest is usable and readable from the image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	PayeeUsability string `json:"payeeUsability,omitempty"`
	// AmountInWordsUsability is a code that indicates if the amount in words (legal amount) Area of Interest is usable and readable from the image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	ConvenienceAmountUsability string `json:"convenienceAmountUsability,omitempty"`
	// SignatureUsability is a code that indicates if the signature Area of Interest is usable and readable from the image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	AmountInWordsUsability string `json:"amountInWordsUsability,omitempty"`
	//  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	SignatureUsability string `json:"signatureUsability,omitempty"`
	// PayorNameAddressUsability is a code that indicates if the payor name and address Area of Interest is usable and readable from the image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	PayorNameAddressUsability string `json:"payorNameAddressUsability,omitempty"`
	// MICRLineUsability is a code that indicates if the MICR line Area of Interest is usable and readable from the image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	MICRLineUsability string `json:"mICRLineUsability,omitempty"`
	// MemoLineUsability is code that indicates if the memo line Area of Interest is usable and readable from the image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's own defined requirements and/or constraints. * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	MemoLineUsability string `json:"memoLineUsability,omitempty"`
	// PayorBankNameAddressUsability is a code that indicates if the payor bank name and address Area of Interest is usable and readable from the image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	PayorBankNameAddressUsability string `json:"payorBankNameAddressUsability,omitempty"`
	// PayeeEndorsementUsability is a code that indicates if the payee endorsement Area of Interest is usable and readable from the image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	PayeeEndorsementUsability string `json:"payeeEndorsementUsability,omitempty"`
	// BOFDEndorsementUsability is a code that indicates if the Bank of First Deposit (BOFD) endorsement Area of Interest is usable and readable from the image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	BOFDEndorsementUsability string `json:"bOFDEndorsementUsability,omitempty"`
	// TransitEndorsementUsability is a code that indicates if the transit endorsement Area of Interest is usable and readable from the image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's own defined requirements and/or constraints.  * `0` -  The image was not tested for any of the image quality conditions * `1` -  The image was tested and one or more image quality conditions were reported * `2` -  The image was tested and no image quality conditions were reported
	TransitEndorsementUsability string `json:"transitEndorsementUsability,omitempty"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField,omitempty"`
}
