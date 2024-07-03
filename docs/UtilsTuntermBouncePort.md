# UtilsTuntermBouncePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoldTime** | Pointer to **int32** | in milli seconds, hold time between multiple port bounces | [optional] 
**Ports** | **[]string** | list of ports to bounce | 

## Methods

### NewUtilsTuntermBouncePort

`func NewUtilsTuntermBouncePort(ports []string, ) *UtilsTuntermBouncePort`

NewUtilsTuntermBouncePort instantiates a new UtilsTuntermBouncePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsTuntermBouncePortWithDefaults

`func NewUtilsTuntermBouncePortWithDefaults() *UtilsTuntermBouncePort`

NewUtilsTuntermBouncePortWithDefaults instantiates a new UtilsTuntermBouncePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoldTime

`func (o *UtilsTuntermBouncePort) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *UtilsTuntermBouncePort) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *UtilsTuntermBouncePort) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *UtilsTuntermBouncePort) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.

### GetPorts

`func (o *UtilsTuntermBouncePort) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *UtilsTuntermBouncePort) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *UtilsTuntermBouncePort) SetPorts(v []string)`

SetPorts sets Ports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


