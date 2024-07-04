# OrgSettingVpnOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsBase** | Pointer to **int32** |  | [optional] 
**StSubnet** | Pointer to **string** | equiring /12 or bigger to support 16 private IPs for 65535 gateways | [optional] [default to "10.224.0.0/12"]

## Methods

### NewOrgSettingVpnOptions

`func NewOrgSettingVpnOptions() *OrgSettingVpnOptions`

NewOrgSettingVpnOptions instantiates a new OrgSettingVpnOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingVpnOptionsWithDefaults

`func NewOrgSettingVpnOptionsWithDefaults() *OrgSettingVpnOptions`

NewOrgSettingVpnOptionsWithDefaults instantiates a new OrgSettingVpnOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsBase

`func (o *OrgSettingVpnOptions) GetAsBase() int32`

GetAsBase returns the AsBase field if non-nil, zero value otherwise.

### GetAsBaseOk

`func (o *OrgSettingVpnOptions) GetAsBaseOk() (*int32, bool)`

GetAsBaseOk returns a tuple with the AsBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsBase

`func (o *OrgSettingVpnOptions) SetAsBase(v int32)`

SetAsBase sets AsBase field to given value.

### HasAsBase

`func (o *OrgSettingVpnOptions) HasAsBase() bool`

HasAsBase returns a boolean if a field has been set.

### GetStSubnet

`func (o *OrgSettingVpnOptions) GetStSubnet() string`

GetStSubnet returns the StSubnet field if non-nil, zero value otherwise.

### GetStSubnetOk

`func (o *OrgSettingVpnOptions) GetStSubnetOk() (*string, bool)`

GetStSubnetOk returns a tuple with the StSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStSubnet

`func (o *OrgSettingVpnOptions) SetStSubnet(v string)`

SetStSubnet sets StSubnet field to given value.

### HasStSubnet

`func (o *OrgSettingVpnOptions) HasStSubnet() bool`

HasStSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


