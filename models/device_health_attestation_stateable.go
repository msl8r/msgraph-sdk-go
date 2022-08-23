package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

// DeviceHealthAttestationStateable
type DeviceHealthAttestationStateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAttestationIdentityKey() *string
	GetBitLockerStatus() *string
	GetBootAppSecurityVersion() *string
	GetBootDebugging() *string
	GetBootManagerSecurityVersion() *string
	GetBootManagerVersion() *string
	GetBootRevisionListInfo() *string
	GetCodeIntegrity() *string
	GetCodeIntegrityCheckVersion() *string
	GetCodeIntegrityPolicy() *string
	GetContentNamespaceUrl() *string
	GetContentVersion() *string
	GetDataExcutionPolicy() *string
	GetDeviceHealthAttestationStatus() *string
	GetEarlyLaunchAntiMalwareDriverProtection() *string
	GetHealthAttestationSupportedStatus() *string
	GetHealthStatusMismatchInfo() *string
	GetIssuedDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetLastUpdateDateTime() *string
	GetOdataType() *string
	GetOperatingSystemKernelDebugging() *string
	GetOperatingSystemRevListInfo() *string
	GetPcr0() *string
	GetPcrHashAlgorithm() *string
	GetResetCount() *int64
	GetRestartCount() *int64
	GetSafeMode() *string
	GetSecureBoot() *string
	GetSecureBootConfigurationPolicyFingerPrint() *string
	GetTestSigning() *string
	GetTpmVersion() *string
	GetVirtualSecureMode() *string
	GetWindowsPE() *string
	SetAttestationIdentityKey(value *string)
	SetBitLockerStatus(value *string)
	SetBootAppSecurityVersion(value *string)
	SetBootDebugging(value *string)
	SetBootManagerSecurityVersion(value *string)
	SetBootManagerVersion(value *string)
	SetBootRevisionListInfo(value *string)
	SetCodeIntegrity(value *string)
	SetCodeIntegrityCheckVersion(value *string)
	SetCodeIntegrityPolicy(value *string)
	SetContentNamespaceUrl(value *string)
	SetContentVersion(value *string)
	SetDataExcutionPolicy(value *string)
	SetDeviceHealthAttestationStatus(value *string)
	SetEarlyLaunchAntiMalwareDriverProtection(value *string)
	SetHealthAttestationSupportedStatus(value *string)
	SetHealthStatusMismatchInfo(value *string)
	SetIssuedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetLastUpdateDateTime(value *string)
	SetOdataType(value *string)
	SetOperatingSystemKernelDebugging(value *string)
	SetOperatingSystemRevListInfo(value *string)
	SetPcr0(value *string)
	SetPcrHashAlgorithm(value *string)
	SetResetCount(value *int64)
	SetRestartCount(value *int64)
	SetSafeMode(value *string)
	SetSecureBoot(value *string)
	SetSecureBootConfigurationPolicyFingerPrint(value *string)
	SetTestSigning(value *string)
	SetTpmVersion(value *string)
	SetVirtualSecureMode(value *string)
	SetWindowsPE(value *string)
}
