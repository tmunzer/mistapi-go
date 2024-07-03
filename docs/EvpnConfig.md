# EvpnConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to [**EvpnConfigRole**](EvpnConfigRole.md) |  | [optional] 

## Methods

### NewEvpnConfig

`func NewEvpnConfig() *EvpnConfig`

NewEvpnConfig instantiates a new EvpnConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvpnConfigWithDefaults

`func NewEvpnConfigWithDefaults() *EvpnConfig`

NewEvpnConfigWithDefaults instantiates a new EvpnConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *EvpnConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EvpnConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EvpnConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EvpnConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRole

`func (o *EvpnConfig) GetRole() EvpnConfigRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EvpnConfig) GetRoleOk() (*EvpnConfigRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EvpnConfig) SetRole(v EvpnConfigRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *EvpnConfig) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


