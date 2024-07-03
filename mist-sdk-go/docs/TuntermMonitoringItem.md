# TuntermMonitoringItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | can be ip, ipv6, hostname | [optional] 
**Port** | Pointer to **int32** | when &#x60;protocol&#x60;&#x3D;&#x3D;&#x60;tcp&#x60; | [optional] 
**Protocol** | Pointer to [**TunternMonitoringProtocol**](TunternMonitoringProtocol.md) |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] [default to 300]

## Methods

### NewTuntermMonitoringItem

`func NewTuntermMonitoringItem() *TuntermMonitoringItem`

NewTuntermMonitoringItem instantiates a new TuntermMonitoringItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTuntermMonitoringItemWithDefaults

`func NewTuntermMonitoringItemWithDefaults() *TuntermMonitoringItem`

NewTuntermMonitoringItemWithDefaults instantiates a new TuntermMonitoringItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *TuntermMonitoringItem) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TuntermMonitoringItem) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TuntermMonitoringItem) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TuntermMonitoringItem) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *TuntermMonitoringItem) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TuntermMonitoringItem) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TuntermMonitoringItem) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *TuntermMonitoringItem) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *TuntermMonitoringItem) GetProtocol() TunternMonitoringProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *TuntermMonitoringItem) GetProtocolOk() (*TunternMonitoringProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *TuntermMonitoringItem) SetProtocol(v TunternMonitoringProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *TuntermMonitoringItem) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTimeout

`func (o *TuntermMonitoringItem) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TuntermMonitoringItem) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TuntermMonitoringItem) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TuntermMonitoringItem) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


