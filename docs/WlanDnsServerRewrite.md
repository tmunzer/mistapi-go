# WlanDnsServerRewrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**RadiusGroups** | Pointer to **map[string]string** | map between radius_group and the desired DNS server (IPv4 only) Property key is the RADIUS group, property value is the desired DNS Server | [optional] 

## Methods

### NewWlanDnsServerRewrite

`func NewWlanDnsServerRewrite() *WlanDnsServerRewrite`

NewWlanDnsServerRewrite instantiates a new WlanDnsServerRewrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanDnsServerRewriteWithDefaults

`func NewWlanDnsServerRewriteWithDefaults() *WlanDnsServerRewrite`

NewWlanDnsServerRewriteWithDefaults instantiates a new WlanDnsServerRewrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WlanDnsServerRewrite) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WlanDnsServerRewrite) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WlanDnsServerRewrite) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WlanDnsServerRewrite) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRadiusGroups

`func (o *WlanDnsServerRewrite) GetRadiusGroups() map[string]string`

GetRadiusGroups returns the RadiusGroups field if non-nil, zero value otherwise.

### GetRadiusGroupsOk

`func (o *WlanDnsServerRewrite) GetRadiusGroupsOk() (*map[string]string, bool)`

GetRadiusGroupsOk returns a tuple with the RadiusGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroups

`func (o *WlanDnsServerRewrite) SetRadiusGroups(v map[string]string)`

SetRadiusGroups sets RadiusGroups field to given value.

### HasRadiusGroups

`func (o *WlanDnsServerRewrite) HasRadiusGroups() bool`

HasRadiusGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


