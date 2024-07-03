# MxedgeTuntermDhcpdConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Servers** | Pointer to **[]string** | list of DHCP servers; required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;relay&#x60; | [optional] 
**Type** | Pointer to [**MxedgeTuntermDhcpdConfigType**](MxedgeTuntermDhcpdConfigType.md) |  | [optional] [default to MXEDGETUNTERMDHCPDCONFIGTYPE_RELAY]

## Methods

### NewMxedgeTuntermDhcpdConfig

`func NewMxedgeTuntermDhcpdConfig() *MxedgeTuntermDhcpdConfig`

NewMxedgeTuntermDhcpdConfig instantiates a new MxedgeTuntermDhcpdConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeTuntermDhcpdConfigWithDefaults

`func NewMxedgeTuntermDhcpdConfigWithDefaults() *MxedgeTuntermDhcpdConfig`

NewMxedgeTuntermDhcpdConfigWithDefaults instantiates a new MxedgeTuntermDhcpdConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MxedgeTuntermDhcpdConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MxedgeTuntermDhcpdConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MxedgeTuntermDhcpdConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MxedgeTuntermDhcpdConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServers

`func (o *MxedgeTuntermDhcpdConfig) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *MxedgeTuntermDhcpdConfig) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *MxedgeTuntermDhcpdConfig) SetServers(v []string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *MxedgeTuntermDhcpdConfig) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetType

`func (o *MxedgeTuntermDhcpdConfig) GetType() MxedgeTuntermDhcpdConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MxedgeTuntermDhcpdConfig) GetTypeOk() (*MxedgeTuntermDhcpdConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MxedgeTuntermDhcpdConfig) SetType(v MxedgeTuntermDhcpdConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *MxedgeTuntermDhcpdConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


