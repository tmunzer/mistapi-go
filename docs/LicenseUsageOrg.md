# LicenseUsageOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**FullyLoaded** | Pointer to **map[string]int32** | Property key is the service name (e.g. \&quot;SUB-MAN\&quot;) | [optional] [readonly] 
**NumDevices** | **int32** |  | [readonly] 
**SiteId** | **string** |  | [readonly] 
**Usages** | **map[string]int32** | subscriptions and their quantities. Property key is the service name (e.g. \&quot;SUB-MAN\&quot;) | [readonly] 

## Methods

### NewLicenseUsageOrg

`func NewLicenseUsageOrg(numDevices int32, siteId string, usages map[string]int32, ) *LicenseUsageOrg`

NewLicenseUsageOrg instantiates a new LicenseUsageOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseUsageOrgWithDefaults

`func NewLicenseUsageOrgWithDefaults() *LicenseUsageOrg`

NewLicenseUsageOrgWithDefaults instantiates a new LicenseUsageOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForSite

`func (o *LicenseUsageOrg) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *LicenseUsageOrg) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *LicenseUsageOrg) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *LicenseUsageOrg) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetFullyLoaded

`func (o *LicenseUsageOrg) GetFullyLoaded() map[string]int32`

GetFullyLoaded returns the FullyLoaded field if non-nil, zero value otherwise.

### GetFullyLoadedOk

`func (o *LicenseUsageOrg) GetFullyLoadedOk() (*map[string]int32, bool)`

GetFullyLoadedOk returns a tuple with the FullyLoaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyLoaded

`func (o *LicenseUsageOrg) SetFullyLoaded(v map[string]int32)`

SetFullyLoaded sets FullyLoaded field to given value.

### HasFullyLoaded

`func (o *LicenseUsageOrg) HasFullyLoaded() bool`

HasFullyLoaded returns a boolean if a field has been set.

### GetNumDevices

`func (o *LicenseUsageOrg) GetNumDevices() int32`

GetNumDevices returns the NumDevices field if non-nil, zero value otherwise.

### GetNumDevicesOk

`func (o *LicenseUsageOrg) GetNumDevicesOk() (*int32, bool)`

GetNumDevicesOk returns a tuple with the NumDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDevices

`func (o *LicenseUsageOrg) SetNumDevices(v int32)`

SetNumDevices sets NumDevices field to given value.


### GetSiteId

`func (o *LicenseUsageOrg) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *LicenseUsageOrg) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *LicenseUsageOrg) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetUsages

`func (o *LicenseUsageOrg) GetUsages() map[string]int32`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *LicenseUsageOrg) GetUsagesOk() (*map[string]int32, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *LicenseUsageOrg) SetUsages(v map[string]int32)`

SetUsages sets Usages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


