# SimpleAlertArpFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCount** | Pointer to **int32** |  | [optional] [default to 10]
**Duration** | Pointer to **int32** | failing within minutes | [optional] [default to 20]
**IncidentCount** | Pointer to **int32** |  | [optional] [default to 10]

## Methods

### NewSimpleAlertArpFailure

`func NewSimpleAlertArpFailure() *SimpleAlertArpFailure`

NewSimpleAlertArpFailure instantiates a new SimpleAlertArpFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleAlertArpFailureWithDefaults

`func NewSimpleAlertArpFailureWithDefaults() *SimpleAlertArpFailure`

NewSimpleAlertArpFailureWithDefaults instantiates a new SimpleAlertArpFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCount

`func (o *SimpleAlertArpFailure) GetClientCount() int32`

GetClientCount returns the ClientCount field if non-nil, zero value otherwise.

### GetClientCountOk

`func (o *SimpleAlertArpFailure) GetClientCountOk() (*int32, bool)`

GetClientCountOk returns a tuple with the ClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCount

`func (o *SimpleAlertArpFailure) SetClientCount(v int32)`

SetClientCount sets ClientCount field to given value.

### HasClientCount

`func (o *SimpleAlertArpFailure) HasClientCount() bool`

HasClientCount returns a boolean if a field has been set.

### GetDuration

`func (o *SimpleAlertArpFailure) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SimpleAlertArpFailure) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SimpleAlertArpFailure) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SimpleAlertArpFailure) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetIncidentCount

`func (o *SimpleAlertArpFailure) GetIncidentCount() int32`

GetIncidentCount returns the IncidentCount field if non-nil, zero value otherwise.

### GetIncidentCountOk

`func (o *SimpleAlertArpFailure) GetIncidentCountOk() (*int32, bool)`

GetIncidentCountOk returns a tuple with the IncidentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentCount

`func (o *SimpleAlertArpFailure) SetIncidentCount(v int32)`

SetIncidentCount sets IncidentCount field to given value.

### HasIncidentCount

`func (o *SimpleAlertArpFailure) HasIncidentCount() bool`

HasIncidentCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


