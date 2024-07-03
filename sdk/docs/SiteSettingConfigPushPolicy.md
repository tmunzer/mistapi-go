# SiteSettingConfigPushPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoPush** | Pointer to **bool** | stop any new config from being pushed to the device | [optional] [default to false]
**PushWindow** | Pointer to [**SiteSettingConfigPushPolicyPushWindow**](SiteSettingConfigPushPolicyPushWindow.md) |  | [optional] 

## Methods

### NewSiteSettingConfigPushPolicy

`func NewSiteSettingConfigPushPolicy() *SiteSettingConfigPushPolicy`

NewSiteSettingConfigPushPolicy instantiates a new SiteSettingConfigPushPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingConfigPushPolicyWithDefaults

`func NewSiteSettingConfigPushPolicyWithDefaults() *SiteSettingConfigPushPolicy`

NewSiteSettingConfigPushPolicyWithDefaults instantiates a new SiteSettingConfigPushPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoPush

`func (o *SiteSettingConfigPushPolicy) GetNoPush() bool`

GetNoPush returns the NoPush field if non-nil, zero value otherwise.

### GetNoPushOk

`func (o *SiteSettingConfigPushPolicy) GetNoPushOk() (*bool, bool)`

GetNoPushOk returns a tuple with the NoPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoPush

`func (o *SiteSettingConfigPushPolicy) SetNoPush(v bool)`

SetNoPush sets NoPush field to given value.

### HasNoPush

`func (o *SiteSettingConfigPushPolicy) HasNoPush() bool`

HasNoPush returns a boolean if a field has been set.

### GetPushWindow

`func (o *SiteSettingConfigPushPolicy) GetPushWindow() SiteSettingConfigPushPolicyPushWindow`

GetPushWindow returns the PushWindow field if non-nil, zero value otherwise.

### GetPushWindowOk

`func (o *SiteSettingConfigPushPolicy) GetPushWindowOk() (*SiteSettingConfigPushPolicyPushWindow, bool)`

GetPushWindowOk returns a tuple with the PushWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushWindow

`func (o *SiteSettingConfigPushPolicy) SetPushWindow(v SiteSettingConfigPushPolicyPushWindow)`

SetPushWindow sets PushWindow field to given value.

### HasPushWindow

`func (o *SiteSettingConfigPushPolicy) HasPushWindow() bool`

HasPushWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


