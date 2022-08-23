package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriveItemable
type DriveItemable interface {
	BaseItemable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAnalytics() ItemAnalyticsable
	GetAudio() Audioable
	GetBundle() Bundleable
	GetChildren() []DriveItemable
	GetContent() []byte
	GetCTag() *string
	GetDeleted() Deletedable
	GetFile() Fileable
	GetFileSystemInfo() FileSystemInfoable
	GetFolder() Folderable
	GetImage() Imageable
	GetListItem() ListItemable
	GetLocation() GeoCoordinatesable
	GetMalware() Malwareable
	GetPackage() Package_escapedable
	GetPendingOperations() PendingOperationsable
	GetPermissions() []Permissionable
	GetPhoto() Photoable
	GetPublication() PublicationFacetable
	GetRemoteItem() RemoteItemable
	GetRoot() Rootable
	GetSearchResult() SearchResultable
	GetShared() Sharedable
	GetSharepointIds() SharepointIdsable
	GetSize() *int64
	GetSpecialFolder() SpecialFolderable
	GetSubscriptions() []Subscriptionable
	GetThumbnails() []ThumbnailSetable
	GetVersions() []DriveItemVersionable
	GetVideo() Videoable
	GetWebDavUrl() *string
	GetWorkbook() Workbookable
	SetAnalytics(value ItemAnalyticsable)
	SetAudio(value Audioable)
	SetBundle(value Bundleable)
	SetChildren(value []DriveItemable)
	SetContent(value []byte)
	SetCTag(value *string)
	SetDeleted(value Deletedable)
	SetFile(value Fileable)
	SetFileSystemInfo(value FileSystemInfoable)
	SetFolder(value Folderable)
	SetImage(value Imageable)
	SetListItem(value ListItemable)
	SetLocation(value GeoCoordinatesable)
	SetMalware(value Malwareable)
	SetPackage(value Package_escapedable)
	SetPendingOperations(value PendingOperationsable)
	SetPermissions(value []Permissionable)
	SetPhoto(value Photoable)
	SetPublication(value PublicationFacetable)
	SetRemoteItem(value RemoteItemable)
	SetRoot(value Rootable)
	SetSearchResult(value SearchResultable)
	SetShared(value Sharedable)
	SetSharepointIds(value SharepointIdsable)
	SetSize(value *int64)
	SetSpecialFolder(value SpecialFolderable)
	SetSubscriptions(value []Subscriptionable)
	SetThumbnails(value []ThumbnailSetable)
	SetVersions(value []DriveItemVersionable)
	SetVideo(value Videoable)
	SetWebDavUrl(value *string)
	SetWorkbook(value Workbookable)
}
