# UtilsShowArp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value. | [optional] [default to 0]
**Interval** | Pointer to **int32** | rate at which output will refresh | [optional] [default to 0]
**Ip** | Pointer to **string** | IP Address | [optional] 
**PortId** | Pointer to **string** | device Port ID | [optional] 
**Vrf** | Pointer to **string** | VRF Name | [optional] 

## Methods

### NewUtilsShowArp

`func NewUtilsShowArp() *UtilsShowArp`

NewUtilsShowArp instantiates a new UtilsShowArp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsShowArpWithDefaults

`func NewUtilsShowArpWithDefaults() *UtilsShowArp`

NewUtilsShowArpWithDefaults instantiates a new UtilsShowArp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *UtilsShowArp) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UtilsShowArp) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UtilsShowArp) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UtilsShowArp) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetInterval

`func (o *UtilsShowArp) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UtilsShowArp) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UtilsShowArp) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UtilsShowArp) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetIp

`func (o *UtilsShowArp) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *UtilsShowArp) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *UtilsShowArp) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *UtilsShowArp) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPortId

`func (o *UtilsShowArp) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *UtilsShowArp) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *UtilsShowArp) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *UtilsShowArp) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetVrf

`func (o *UtilsShowArp) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *UtilsShowArp) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *UtilsShowArp) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *UtilsShowArp) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


