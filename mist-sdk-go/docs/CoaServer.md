# CoaServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableEventTimestampCheck** | Pointer to **bool** | whether to disable Event-Timestamp Check | [optional] [default to false]
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Ip** | **string** |  | 
**Port** | Pointer to **int32** |  | [optional] [default to 3799]
**Secret** | **string** |  | 

## Methods

### NewCoaServer

`func NewCoaServer(ip string, secret string, ) *CoaServer`

NewCoaServer instantiates a new CoaServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoaServerWithDefaults

`func NewCoaServerWithDefaults() *CoaServer`

NewCoaServerWithDefaults instantiates a new CoaServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableEventTimestampCheck

`func (o *CoaServer) GetDisableEventTimestampCheck() bool`

GetDisableEventTimestampCheck returns the DisableEventTimestampCheck field if non-nil, zero value otherwise.

### GetDisableEventTimestampCheckOk

`func (o *CoaServer) GetDisableEventTimestampCheckOk() (*bool, bool)`

GetDisableEventTimestampCheckOk returns a tuple with the DisableEventTimestampCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEventTimestampCheck

`func (o *CoaServer) SetDisableEventTimestampCheck(v bool)`

SetDisableEventTimestampCheck sets DisableEventTimestampCheck field to given value.

### HasDisableEventTimestampCheck

`func (o *CoaServer) HasDisableEventTimestampCheck() bool`

HasDisableEventTimestampCheck returns a boolean if a field has been set.

### GetEnabled

`func (o *CoaServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CoaServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CoaServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CoaServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIp

`func (o *CoaServer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CoaServer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CoaServer) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPort

`func (o *CoaServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CoaServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CoaServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CoaServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *CoaServer) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CoaServer) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CoaServer) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


