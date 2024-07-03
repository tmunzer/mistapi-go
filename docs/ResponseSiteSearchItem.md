# ResponseSiteSearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpgradeEnabled** | **bool** |  | 
**AutoUpgradeVersion** | **string** |  | [readonly] 
**CountryCode** | **NullableString** |  | [readonly] 
**HoneypotEnabled** | **bool** |  | 
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**OrgId** | **string** |  | [readonly] 
**SiteId** | **string** |  | [readonly] 
**Timestamp** | **float32** |  | [readonly] 
**Timezone** | **string** |  | [readonly] 
**VnaEnabled** | **bool** |  | 
**WifiEnabled** | **bool** |  | 

## Methods

### NewResponseSiteSearchItem

`func NewResponseSiteSearchItem(autoUpgradeEnabled bool, autoUpgradeVersion string, countryCode NullableString, honeypotEnabled bool, id string, name string, orgId string, siteId string, timestamp float32, timezone string, vnaEnabled bool, wifiEnabled bool, ) *ResponseSiteSearchItem`

NewResponseSiteSearchItem instantiates a new ResponseSiteSearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSiteSearchItemWithDefaults

`func NewResponseSiteSearchItemWithDefaults() *ResponseSiteSearchItem`

NewResponseSiteSearchItemWithDefaults instantiates a new ResponseSiteSearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUpgradeEnabled

`func (o *ResponseSiteSearchItem) GetAutoUpgradeEnabled() bool`

GetAutoUpgradeEnabled returns the AutoUpgradeEnabled field if non-nil, zero value otherwise.

### GetAutoUpgradeEnabledOk

`func (o *ResponseSiteSearchItem) GetAutoUpgradeEnabledOk() (*bool, bool)`

GetAutoUpgradeEnabledOk returns a tuple with the AutoUpgradeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgradeEnabled

`func (o *ResponseSiteSearchItem) SetAutoUpgradeEnabled(v bool)`

SetAutoUpgradeEnabled sets AutoUpgradeEnabled field to given value.


### GetAutoUpgradeVersion

`func (o *ResponseSiteSearchItem) GetAutoUpgradeVersion() string`

GetAutoUpgradeVersion returns the AutoUpgradeVersion field if non-nil, zero value otherwise.

### GetAutoUpgradeVersionOk

`func (o *ResponseSiteSearchItem) GetAutoUpgradeVersionOk() (*string, bool)`

GetAutoUpgradeVersionOk returns a tuple with the AutoUpgradeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgradeVersion

`func (o *ResponseSiteSearchItem) SetAutoUpgradeVersion(v string)`

SetAutoUpgradeVersion sets AutoUpgradeVersion field to given value.


### GetCountryCode

`func (o *ResponseSiteSearchItem) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ResponseSiteSearchItem) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ResponseSiteSearchItem) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *ResponseSiteSearchItem) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *ResponseSiteSearchItem) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetHoneypotEnabled

`func (o *ResponseSiteSearchItem) GetHoneypotEnabled() bool`

GetHoneypotEnabled returns the HoneypotEnabled field if non-nil, zero value otherwise.

### GetHoneypotEnabledOk

`func (o *ResponseSiteSearchItem) GetHoneypotEnabledOk() (*bool, bool)`

GetHoneypotEnabledOk returns a tuple with the HoneypotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoneypotEnabled

`func (o *ResponseSiteSearchItem) SetHoneypotEnabled(v bool)`

SetHoneypotEnabled sets HoneypotEnabled field to given value.


### GetId

`func (o *ResponseSiteSearchItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseSiteSearchItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseSiteSearchItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseSiteSearchItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseSiteSearchItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseSiteSearchItem) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *ResponseSiteSearchItem) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResponseSiteSearchItem) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResponseSiteSearchItem) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetSiteId

`func (o *ResponseSiteSearchItem) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ResponseSiteSearchItem) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ResponseSiteSearchItem) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetTimestamp

`func (o *ResponseSiteSearchItem) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponseSiteSearchItem) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponseSiteSearchItem) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetTimezone

`func (o *ResponseSiteSearchItem) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ResponseSiteSearchItem) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ResponseSiteSearchItem) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetVnaEnabled

`func (o *ResponseSiteSearchItem) GetVnaEnabled() bool`

GetVnaEnabled returns the VnaEnabled field if non-nil, zero value otherwise.

### GetVnaEnabledOk

`func (o *ResponseSiteSearchItem) GetVnaEnabledOk() (*bool, bool)`

GetVnaEnabledOk returns a tuple with the VnaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnaEnabled

`func (o *ResponseSiteSearchItem) SetVnaEnabled(v bool)`

SetVnaEnabled sets VnaEnabled field to given value.


### GetWifiEnabled

`func (o *ResponseSiteSearchItem) GetWifiEnabled() bool`

GetWifiEnabled returns the WifiEnabled field if non-nil, zero value otherwise.

### GetWifiEnabledOk

`func (o *ResponseSiteSearchItem) GetWifiEnabledOk() (*bool, bool)`

GetWifiEnabledOk returns a tuple with the WifiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiEnabled

`func (o *ResponseSiteSearchItem) SetWifiEnabled(v bool)`

SetWifiEnabled sets WifiEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


