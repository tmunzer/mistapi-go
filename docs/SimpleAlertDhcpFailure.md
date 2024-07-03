# SimpleAlertDhcpFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCount** | Pointer to **int32** |  | [optional] [default to 10]
**Duration** | Pointer to **int32** | failing within minutes | [optional] [default to 10]
**IncidentCount** | Pointer to **int32** |  | [optional] [default to 20]

## Methods

### NewSimpleAlertDhcpFailure

`func NewSimpleAlertDhcpFailure() *SimpleAlertDhcpFailure`

NewSimpleAlertDhcpFailure instantiates a new SimpleAlertDhcpFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleAlertDhcpFailureWithDefaults

`func NewSimpleAlertDhcpFailureWithDefaults() *SimpleAlertDhcpFailure`

NewSimpleAlertDhcpFailureWithDefaults instantiates a new SimpleAlertDhcpFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCount

`func (o *SimpleAlertDhcpFailure) GetClientCount() int32`

GetClientCount returns the ClientCount field if non-nil, zero value otherwise.

### GetClientCountOk

`func (o *SimpleAlertDhcpFailure) GetClientCountOk() (*int32, bool)`

GetClientCountOk returns a tuple with the ClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCount

`func (o *SimpleAlertDhcpFailure) SetClientCount(v int32)`

SetClientCount sets ClientCount field to given value.

### HasClientCount

`func (o *SimpleAlertDhcpFailure) HasClientCount() bool`

HasClientCount returns a boolean if a field has been set.

### GetDuration

`func (o *SimpleAlertDhcpFailure) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SimpleAlertDhcpFailure) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SimpleAlertDhcpFailure) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SimpleAlertDhcpFailure) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetIncidentCount

`func (o *SimpleAlertDhcpFailure) GetIncidentCount() int32`

GetIncidentCount returns the IncidentCount field if non-nil, zero value otherwise.

### GetIncidentCountOk

`func (o *SimpleAlertDhcpFailure) GetIncidentCountOk() (*int32, bool)`

GetIncidentCountOk returns a tuple with the IncidentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentCount

`func (o *SimpleAlertDhcpFailure) SetIncidentCount(v int32)`

SetIncidentCount sets IncidentCount field to given value.

### HasIncidentCount

`func (o *SimpleAlertDhcpFailure) HasIncidentCount() bool`

HasIncidentCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


