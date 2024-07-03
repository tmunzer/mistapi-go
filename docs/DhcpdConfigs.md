# DhcpdConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | if set to &#x60;true&#x60;, disable the DHCP server | [optional] [default to false]

## Methods

### NewDhcpdConfigs

`func NewDhcpdConfigs() *DhcpdConfigs`

NewDhcpdConfigs instantiates a new DhcpdConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpdConfigsWithDefaults

`func NewDhcpdConfigsWithDefaults() *DhcpdConfigs`

NewDhcpdConfigsWithDefaults instantiates a new DhcpdConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DhcpdConfigs) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DhcpdConfigs) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DhcpdConfigs) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DhcpdConfigs) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


