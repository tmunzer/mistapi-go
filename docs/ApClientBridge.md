# ApClientBridge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**ApClientBridgeAuth**](ApClientBridgeAuth.md) |  | [optional] 
**Enabled** | Pointer to **bool** | when acted as client bridge: * only 5G radio can be used * will not serve as AP on any radios | [optional] [default to false]
**Ssid** | Pointer to **string** |  | [optional] 

## Methods

### NewApClientBridge

`func NewApClientBridge() *ApClientBridge`

NewApClientBridge instantiates a new ApClientBridge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApClientBridgeWithDefaults

`func NewApClientBridgeWithDefaults() *ApClientBridge`

NewApClientBridgeWithDefaults instantiates a new ApClientBridge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *ApClientBridge) GetAuth() ApClientBridgeAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ApClientBridge) GetAuthOk() (*ApClientBridgeAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ApClientBridge) SetAuth(v ApClientBridgeAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ApClientBridge) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetEnabled

`func (o *ApClientBridge) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApClientBridge) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApClientBridge) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApClientBridge) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSsid

`func (o *ApClientBridge) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ApClientBridge) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ApClientBridge) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ApClientBridge) HasSsid() bool`

HasSsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


