package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

// EducationUserable
type EducationUserable interface {
	Entityable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAccountEnabled() *bool
	GetAssignedLicenses() []AssignedLicenseable
	GetAssignedPlans() []AssignedPlanable
	GetAssignments() []EducationAssignmentable
	GetBusinessPhones() []string
	GetClasses() []EducationClassable
	GetCreatedBy() IdentitySetable
	GetDepartment() *string
	GetDisplayName() *string
	GetExternalSource() *EducationExternalSource
	GetExternalSourceDetail() *string
	GetGivenName() *string
	GetMail() *string
	GetMailingAddress() PhysicalAddressable
	GetMailNickname() *string
	GetMiddleName() *string
	GetMobilePhone() *string
	GetOfficeLocation() *string
	GetOnPremisesInfo() EducationOnPremisesInfoable
	GetPasswordPolicies() *string
	GetPasswordProfile() PasswordProfileable
	GetPreferredLanguage() *string
	GetPrimaryRole() *EducationUserRole
	GetProvisionedPlans() []ProvisionedPlanable
	GetRefreshTokensValidFromDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetRelatedContacts() []RelatedContactable
	GetResidenceAddress() PhysicalAddressable
	GetRubrics() []EducationRubricable
	GetSchools() []EducationSchoolable
	GetShowInAddressList() *bool
	GetStudent() EducationStudentable
	GetSurname() *string
	GetTaughtClasses() []EducationClassable
	GetTeacher() EducationTeacherable
	GetUsageLocation() *string
	GetUser() Userable
	GetUserPrincipalName() *string
	GetUserType() *string
	SetAccountEnabled(value *bool)
	SetAssignedLicenses(value []AssignedLicenseable)
	SetAssignedPlans(value []AssignedPlanable)
	SetAssignments(value []EducationAssignmentable)
	SetBusinessPhones(value []string)
	SetClasses(value []EducationClassable)
	SetCreatedBy(value IdentitySetable)
	SetDepartment(value *string)
	SetDisplayName(value *string)
	SetExternalSource(value *EducationExternalSource)
	SetExternalSourceDetail(value *string)
	SetGivenName(value *string)
	SetMail(value *string)
	SetMailingAddress(value PhysicalAddressable)
	SetMailNickname(value *string)
	SetMiddleName(value *string)
	SetMobilePhone(value *string)
	SetOfficeLocation(value *string)
	SetOnPremisesInfo(value EducationOnPremisesInfoable)
	SetPasswordPolicies(value *string)
	SetPasswordProfile(value PasswordProfileable)
	SetPreferredLanguage(value *string)
	SetPrimaryRole(value *EducationUserRole)
	SetProvisionedPlans(value []ProvisionedPlanable)
	SetRefreshTokensValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetRelatedContacts(value []RelatedContactable)
	SetResidenceAddress(value PhysicalAddressable)
	SetRubrics(value []EducationRubricable)
	SetSchools(value []EducationSchoolable)
	SetShowInAddressList(value *bool)
	SetStudent(value EducationStudentable)
	SetSurname(value *string)
	SetTaughtClasses(value []EducationClassable)
	SetTeacher(value EducationTeacherable)
	SetUsageLocation(value *string)
	SetUser(value Userable)
	SetUserPrincipalName(value *string)
	SetUserType(value *string)
}
