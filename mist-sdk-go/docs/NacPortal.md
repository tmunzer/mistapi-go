# NacPortal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to [**NacPortalAccessType**](NacPortalAccessType.md) |  | [optional] [default to NACPORTALACCESSTYPE_WIRELESS]
**BgImageUrl** | Pointer to **string** | background image | [optional] 
**CertExpireTime** | Pointer to **int32** | in days | [optional] 
**EnableTelemetry** | Pointer to **bool** | model, version, fingering, events (connecting, disconnect, roaming), which ap | [optional] 
**ExpiryNotificationTime** | Pointer to **int32** | in days | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotifyExpiry** | Pointer to **bool** | phase 2 | [optional] 
**Ssid** | Pointer to **string** |  | [optional] 
**Sso** | Pointer to [**NacPortalSso**](NacPortalSso.md) |  | [optional] 
**TemplateUrl** | Pointer to **string** |  | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] [readonly] 
**Tos** | Pointer to **string** |  | [optional] 

## Methods

### NewNacPortal

`func NewNacPortal() *NacPortal`

NewNacPortal instantiates a new NacPortal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNacPortalWithDefaults

`func NewNacPortalWithDefaults() *NacPortal`

NewNacPortalWithDefaults instantiates a new NacPortal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *NacPortal) GetAccessType() NacPortalAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *NacPortal) GetAccessTypeOk() (*NacPortalAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *NacPortal) SetAccessType(v NacPortalAccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *NacPortal) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetBgImageUrl

`func (o *NacPortal) GetBgImageUrl() string`

GetBgImageUrl returns the BgImageUrl field if non-nil, zero value otherwise.

### GetBgImageUrlOk

`func (o *NacPortal) GetBgImageUrlOk() (*string, bool)`

GetBgImageUrlOk returns a tuple with the BgImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgImageUrl

`func (o *NacPortal) SetBgImageUrl(v string)`

SetBgImageUrl sets BgImageUrl field to given value.

### HasBgImageUrl

`func (o *NacPortal) HasBgImageUrl() bool`

HasBgImageUrl returns a boolean if a field has been set.

### GetCertExpireTime

`func (o *NacPortal) GetCertExpireTime() int32`

GetCertExpireTime returns the CertExpireTime field if non-nil, zero value otherwise.

### GetCertExpireTimeOk

`func (o *NacPortal) GetCertExpireTimeOk() (*int32, bool)`

GetCertExpireTimeOk returns a tuple with the CertExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpireTime

`func (o *NacPortal) SetCertExpireTime(v int32)`

SetCertExpireTime sets CertExpireTime field to given value.

### HasCertExpireTime

`func (o *NacPortal) HasCertExpireTime() bool`

HasCertExpireTime returns a boolean if a field has been set.

### GetEnableTelemetry

`func (o *NacPortal) GetEnableTelemetry() bool`

GetEnableTelemetry returns the EnableTelemetry field if non-nil, zero value otherwise.

### GetEnableTelemetryOk

`func (o *NacPortal) GetEnableTelemetryOk() (*bool, bool)`

GetEnableTelemetryOk returns a tuple with the EnableTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTelemetry

`func (o *NacPortal) SetEnableTelemetry(v bool)`

SetEnableTelemetry sets EnableTelemetry field to given value.

### HasEnableTelemetry

`func (o *NacPortal) HasEnableTelemetry() bool`

HasEnableTelemetry returns a boolean if a field has been set.

### GetExpiryNotificationTime

`func (o *NacPortal) GetExpiryNotificationTime() int32`

GetExpiryNotificationTime returns the ExpiryNotificationTime field if non-nil, zero value otherwise.

### GetExpiryNotificationTimeOk

`func (o *NacPortal) GetExpiryNotificationTimeOk() (*int32, bool)`

GetExpiryNotificationTimeOk returns a tuple with the ExpiryNotificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryNotificationTime

`func (o *NacPortal) SetExpiryNotificationTime(v int32)`

SetExpiryNotificationTime sets ExpiryNotificationTime field to given value.

### HasExpiryNotificationTime

`func (o *NacPortal) HasExpiryNotificationTime() bool`

HasExpiryNotificationTime returns a boolean if a field has been set.

### GetName

`func (o *NacPortal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NacPortal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NacPortal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NacPortal) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyExpiry

`func (o *NacPortal) GetNotifyExpiry() bool`

GetNotifyExpiry returns the NotifyExpiry field if non-nil, zero value otherwise.

### GetNotifyExpiryOk

`func (o *NacPortal) GetNotifyExpiryOk() (*bool, bool)`

GetNotifyExpiryOk returns a tuple with the NotifyExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyExpiry

`func (o *NacPortal) SetNotifyExpiry(v bool)`

SetNotifyExpiry sets NotifyExpiry field to given value.

### HasNotifyExpiry

`func (o *NacPortal) HasNotifyExpiry() bool`

HasNotifyExpiry returns a boolean if a field has been set.

### GetSsid

`func (o *NacPortal) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *NacPortal) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *NacPortal) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *NacPortal) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSso

`func (o *NacPortal) GetSso() NacPortalSso`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *NacPortal) GetSsoOk() (*NacPortalSso, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *NacPortal) SetSso(v NacPortalSso)`

SetSso sets Sso field to given value.

### HasSso

`func (o *NacPortal) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetTemplateUrl

`func (o *NacPortal) GetTemplateUrl() string`

GetTemplateUrl returns the TemplateUrl field if non-nil, zero value otherwise.

### GetTemplateUrlOk

`func (o *NacPortal) GetTemplateUrlOk() (*string, bool)`

GetTemplateUrlOk returns a tuple with the TemplateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUrl

`func (o *NacPortal) SetTemplateUrl(v string)`

SetTemplateUrl sets TemplateUrl field to given value.

### HasTemplateUrl

`func (o *NacPortal) HasTemplateUrl() bool`

HasTemplateUrl returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *NacPortal) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *NacPortal) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *NacPortal) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *NacPortal) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetTos

`func (o *NacPortal) GetTos() string`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *NacPortal) GetTosOk() (*string, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *NacPortal) SetTos(v string)`

SetTos sets Tos field to given value.

### HasTos

`func (o *NacPortal) HasTos() bool`

HasTos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


