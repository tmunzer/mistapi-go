# SiteSettingVna

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | enable Virtual Network Assistant (using SUB-VNA license). This applied to AP / Switch / Gateway | [optional] [default to false]

## Methods

### NewSiteSettingVna

`func NewSiteSettingVna() *SiteSettingVna`

NewSiteSettingVna instantiates a new SiteSettingVna object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingVnaWithDefaults

`func NewSiteSettingVnaWithDefaults() *SiteSettingVna`

NewSiteSettingVnaWithDefaults instantiates a new SiteSettingVna object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SiteSettingVna) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SiteSettingVna) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SiteSettingVna) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SiteSettingVna) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


