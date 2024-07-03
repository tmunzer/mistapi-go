# ApAeroscout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | whether to enable aeroscout config | [optional] [default to false]
**Host** | Pointer to **NullableString** | required if enabled, aeroscout server host | [optional] 
**LocateConnected** | Pointer to **bool** | whether to enable the feature to allow wireless clients data received and sent to AES server for location calculation | [optional] [default to true]

## Methods

### NewApAeroscout

`func NewApAeroscout() *ApAeroscout`

NewApAeroscout instantiates a new ApAeroscout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApAeroscoutWithDefaults

`func NewApAeroscoutWithDefaults() *ApAeroscout`

NewApAeroscoutWithDefaults instantiates a new ApAeroscout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApAeroscout) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApAeroscout) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApAeroscout) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApAeroscout) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *ApAeroscout) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ApAeroscout) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ApAeroscout) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ApAeroscout) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *ApAeroscout) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *ApAeroscout) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetLocateConnected

`func (o *ApAeroscout) GetLocateConnected() bool`

GetLocateConnected returns the LocateConnected field if non-nil, zero value otherwise.

### GetLocateConnectedOk

`func (o *ApAeroscout) GetLocateConnectedOk() (*bool, bool)`

GetLocateConnectedOk returns a tuple with the LocateConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateConnected

`func (o *ApAeroscout) SetLocateConnected(v bool)`

SetLocateConnected sets LocateConnected field to given value.

### HasLocateConnected

`func (o *ApAeroscout) HasLocateConnected() bool`

HasLocateConnected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


