package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81GeneralConfigurationable
type WindowsPhone81GeneralConfigurationable interface {
	DeviceConfigurationable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetApplyOnlyToWindowsPhone81() *bool
	GetAppsBlockCopyPaste() *bool
	GetBluetoothBlocked() *bool
	GetCameraBlocked() *bool
	GetCellularBlockWifiTethering() *bool
	GetCompliantAppListType() *AppListType
	GetCompliantAppsList() []AppListItemable
	GetDiagnosticDataBlockSubmission() *bool
	GetEmailBlockAddingAccounts() *bool
	GetLocationServicesBlocked() *bool
	GetMicrosoftAccountBlocked() *bool
	GetNfcBlocked() *bool
	GetPasswordBlockSimple() *bool
	GetPasswordExpirationDays() *int32
	GetPasswordMinimumCharacterSetCount() *int32
	GetPasswordMinimumLength() *int32
	GetPasswordMinutesOfInactivityBeforeScreenTimeout() *int32
	GetPasswordPreviousPasswordBlockCount() *int32
	GetPasswordRequired() *bool
	GetPasswordRequiredType() *RequiredPasswordType
	GetPasswordSignInFailureCountBeforeFactoryReset() *int32
	GetScreenCaptureBlocked() *bool
	GetStorageBlockRemovableStorage() *bool
	GetStorageRequireEncryption() *bool
	GetWebBrowserBlocked() *bool
	GetWifiBlockAutomaticConnectHotspots() *bool
	GetWifiBlocked() *bool
	GetWifiBlockHotspotReporting() *bool
	GetWindowsStoreBlocked() *bool
	SetApplyOnlyToWindowsPhone81(value *bool)
	SetAppsBlockCopyPaste(value *bool)
	SetBluetoothBlocked(value *bool)
	SetCameraBlocked(value *bool)
	SetCellularBlockWifiTethering(value *bool)
	SetCompliantAppListType(value *AppListType)
	SetCompliantAppsList(value []AppListItemable)
	SetDiagnosticDataBlockSubmission(value *bool)
	SetEmailBlockAddingAccounts(value *bool)
	SetLocationServicesBlocked(value *bool)
	SetMicrosoftAccountBlocked(value *bool)
	SetNfcBlocked(value *bool)
	SetPasswordBlockSimple(value *bool)
	SetPasswordExpirationDays(value *int32)
	SetPasswordMinimumCharacterSetCount(value *int32)
	SetPasswordMinimumLength(value *int32)
	SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)
	SetPasswordPreviousPasswordBlockCount(value *int32)
	SetPasswordRequired(value *bool)
	SetPasswordRequiredType(value *RequiredPasswordType)
	SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)
	SetScreenCaptureBlocked(value *bool)
	SetStorageBlockRemovableStorage(value *bool)
	SetStorageRequireEncryption(value *bool)
	SetWebBrowserBlocked(value *bool)
	SetWifiBlockAutomaticConnectHotspots(value *bool)
	SetWifiBlocked(value *bool)
	SetWifiBlockHotspotReporting(value *bool)
	SetWindowsStoreBlocked(value *bool)
}
