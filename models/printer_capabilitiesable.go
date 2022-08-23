package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterCapabilitiesable
type PrinterCapabilitiesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBottomMargins() []int32
	GetCollation() *bool
	GetColorModes() []string
	GetContentTypes() []string
	GetCopiesPerJob() IntegerRangeable
	GetDpis() []int32
	GetDuplexModes() []string
	GetFeedOrientations() []string
	GetFinishings() []string
	GetInputBins() []string
	GetIsColorPrintingSupported() *bool
	GetIsPageRangeSupported() *bool
	GetLeftMargins() []int32
	GetMediaColors() []string
	GetMediaSizes() []string
	GetMediaTypes() []string
	GetMultipageLayouts() []string
	GetOdataType() *string
	GetOrientations() []string
	GetOutputBins() []string
	GetPagesPerSheet() []int32
	GetQualities() []string
	GetRightMargins() []int32
	GetScalings() []string
	GetSupportsFitPdfToPage() *bool
	GetTopMargins() []int32
	SetBottomMargins(value []int32)
	SetCollation(value *bool)
	SetColorModes(value []string)
	SetContentTypes(value []string)
	SetCopiesPerJob(value IntegerRangeable)
	SetDpis(value []int32)
	SetDuplexModes(value []string)
	SetFeedOrientations(value []string)
	SetFinishings(value []string)
	SetInputBins(value []string)
	SetIsColorPrintingSupported(value *bool)
	SetIsPageRangeSupported(value *bool)
	SetLeftMargins(value []int32)
	SetMediaColors(value []string)
	SetMediaSizes(value []string)
	SetMediaTypes(value []string)
	SetMultipageLayouts(value []string)
	SetOdataType(value *string)
	SetOrientations(value []string)
	SetOutputBins(value []string)
	SetPagesPerSheet(value []int32)
	SetQualities(value []string)
	SetRightMargins(value []int32)
	SetScalings(value []string)
	SetSupportsFitPdfToPage(value *bool)
	SetTopMargins(value []int32)
}
