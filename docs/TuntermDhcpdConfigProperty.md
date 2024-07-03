# TuntermDhcpdConfigProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Servers** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**TuntermDhcpdType**](TuntermDhcpdType.md) |  | [optional] [default to TUNTERMDHCPDTYPE_RELAY]

## Methods

### NewTuntermDhcpdConfigProperty

`func NewTuntermDhcpdConfigProperty() *TuntermDhcpdConfigProperty`

NewTuntermDhcpdConfigProperty instantiates a new TuntermDhcpdConfigProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTuntermDhcpdConfigPropertyWithDefaults

`func NewTuntermDhcpdConfigPropertyWithDefaults() *TuntermDhcpdConfigProperty`

NewTuntermDhcpdConfigPropertyWithDefaults instantiates a new TuntermDhcpdConfigProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TuntermDhcpdConfigProperty) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TuntermDhcpdConfigProperty) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TuntermDhcpdConfigProperty) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TuntermDhcpdConfigProperty) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServers

`func (o *TuntermDhcpdConfigProperty) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *TuntermDhcpdConfigProperty) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *TuntermDhcpdConfigProperty) SetServers(v []string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *TuntermDhcpdConfigProperty) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetType

`func (o *TuntermDhcpdConfigProperty) GetType() TuntermDhcpdType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TuntermDhcpdConfigProperty) GetTypeOk() (*TuntermDhcpdType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TuntermDhcpdConfigProperty) SetType(v TuntermDhcpdType)`

SetType sets Type field to given value.

### HasType

`func (o *TuntermDhcpdConfigProperty) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


