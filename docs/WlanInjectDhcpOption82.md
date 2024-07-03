# WlanInjectDhcpOption82

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | whether to inject option 82 when forwarding DHCP packets | [optional] [default to false]

## Methods

### NewWlanInjectDhcpOption82

`func NewWlanInjectDhcpOption82() *WlanInjectDhcpOption82`

NewWlanInjectDhcpOption82 instantiates a new WlanInjectDhcpOption82 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanInjectDhcpOption82WithDefaults

`func NewWlanInjectDhcpOption82WithDefaults() *WlanInjectDhcpOption82`

NewWlanInjectDhcpOption82WithDefaults instantiates a new WlanInjectDhcpOption82 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitId

`func (o *WlanInjectDhcpOption82) GetCircuitId() string`

GetCircuitId returns the CircuitId field if non-nil, zero value otherwise.

### GetCircuitIdOk

`func (o *WlanInjectDhcpOption82) GetCircuitIdOk() (*string, bool)`

GetCircuitIdOk returns a tuple with the CircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitId

`func (o *WlanInjectDhcpOption82) SetCircuitId(v string)`

SetCircuitId sets CircuitId field to given value.

### HasCircuitId

`func (o *WlanInjectDhcpOption82) HasCircuitId() bool`

HasCircuitId returns a boolean if a field has been set.

### GetEnabled

`func (o *WlanInjectDhcpOption82) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WlanInjectDhcpOption82) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WlanInjectDhcpOption82) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WlanInjectDhcpOption82) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


