# SwitchPortUsageStormControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoBroadcast** | Pointer to **bool** | whether to disable storm control on broadcast traffic | [optional] [default to false]
**NoMulticast** | Pointer to **bool** | whether to disable storm control on multicast traffic | [optional] [default to false]
**NoRegisteredMulticast** | Pointer to **bool** | whether to disable storm control on registered multicast traffic | [optional] [default to false]
**NoUnknownUnicast** | Pointer to **bool** | whether to disable storm control on unknown unicast traffic | [optional] [default to false]
**Percentage** | Pointer to **int32** | bandwidth-percentage, configures the storm control level as a percentage of the available bandwidth | [optional] [default to 80]

## Methods

### NewSwitchPortUsageStormControl

`func NewSwitchPortUsageStormControl() *SwitchPortUsageStormControl`

NewSwitchPortUsageStormControl instantiates a new SwitchPortUsageStormControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPortUsageStormControlWithDefaults

`func NewSwitchPortUsageStormControlWithDefaults() *SwitchPortUsageStormControl`

NewSwitchPortUsageStormControlWithDefaults instantiates a new SwitchPortUsageStormControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoBroadcast

`func (o *SwitchPortUsageStormControl) GetNoBroadcast() bool`

GetNoBroadcast returns the NoBroadcast field if non-nil, zero value otherwise.

### GetNoBroadcastOk

`func (o *SwitchPortUsageStormControl) GetNoBroadcastOk() (*bool, bool)`

GetNoBroadcastOk returns a tuple with the NoBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBroadcast

`func (o *SwitchPortUsageStormControl) SetNoBroadcast(v bool)`

SetNoBroadcast sets NoBroadcast field to given value.

### HasNoBroadcast

`func (o *SwitchPortUsageStormControl) HasNoBroadcast() bool`

HasNoBroadcast returns a boolean if a field has been set.

### GetNoMulticast

`func (o *SwitchPortUsageStormControl) GetNoMulticast() bool`

GetNoMulticast returns the NoMulticast field if non-nil, zero value otherwise.

### GetNoMulticastOk

`func (o *SwitchPortUsageStormControl) GetNoMulticastOk() (*bool, bool)`

GetNoMulticastOk returns a tuple with the NoMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoMulticast

`func (o *SwitchPortUsageStormControl) SetNoMulticast(v bool)`

SetNoMulticast sets NoMulticast field to given value.

### HasNoMulticast

`func (o *SwitchPortUsageStormControl) HasNoMulticast() bool`

HasNoMulticast returns a boolean if a field has been set.

### GetNoRegisteredMulticast

`func (o *SwitchPortUsageStormControl) GetNoRegisteredMulticast() bool`

GetNoRegisteredMulticast returns the NoRegisteredMulticast field if non-nil, zero value otherwise.

### GetNoRegisteredMulticastOk

`func (o *SwitchPortUsageStormControl) GetNoRegisteredMulticastOk() (*bool, bool)`

GetNoRegisteredMulticastOk returns a tuple with the NoRegisteredMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRegisteredMulticast

`func (o *SwitchPortUsageStormControl) SetNoRegisteredMulticast(v bool)`

SetNoRegisteredMulticast sets NoRegisteredMulticast field to given value.

### HasNoRegisteredMulticast

`func (o *SwitchPortUsageStormControl) HasNoRegisteredMulticast() bool`

HasNoRegisteredMulticast returns a boolean if a field has been set.

### GetNoUnknownUnicast

`func (o *SwitchPortUsageStormControl) GetNoUnknownUnicast() bool`

GetNoUnknownUnicast returns the NoUnknownUnicast field if non-nil, zero value otherwise.

### GetNoUnknownUnicastOk

`func (o *SwitchPortUsageStormControl) GetNoUnknownUnicastOk() (*bool, bool)`

GetNoUnknownUnicastOk returns a tuple with the NoUnknownUnicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoUnknownUnicast

`func (o *SwitchPortUsageStormControl) SetNoUnknownUnicast(v bool)`

SetNoUnknownUnicast sets NoUnknownUnicast field to given value.

### HasNoUnknownUnicast

`func (o *SwitchPortUsageStormControl) HasNoUnknownUnicast() bool`

HasNoUnknownUnicast returns a boolean if a field has been set.

### GetPercentage

`func (o *SwitchPortUsageStormControl) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *SwitchPortUsageStormControl) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *SwitchPortUsageStormControl) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *SwitchPortUsageStormControl) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


