# EventsRogue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | **string** |  | 
**Bssid** | **string** |  | 
**Channel** | **int32** |  | 
**Rssi** | **int32** |  | 
**Ssid** | **string** |  | 
**Timestamp** | **float32** |  | 

## Methods

### NewEventsRogue

`func NewEventsRogue(ap string, bssid string, channel int32, rssi int32, ssid string, timestamp float32, ) *EventsRogue`

NewEventsRogue instantiates a new EventsRogue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsRogueWithDefaults

`func NewEventsRogueWithDefaults() *EventsRogue`

NewEventsRogueWithDefaults instantiates a new EventsRogue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *EventsRogue) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *EventsRogue) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *EventsRogue) SetAp(v string)`

SetAp sets Ap field to given value.


### GetBssid

`func (o *EventsRogue) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *EventsRogue) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *EventsRogue) SetBssid(v string)`

SetBssid sets Bssid field to given value.


### GetChannel

`func (o *EventsRogue) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *EventsRogue) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *EventsRogue) SetChannel(v int32)`

SetChannel sets Channel field to given value.


### GetRssi

`func (o *EventsRogue) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *EventsRogue) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *EventsRogue) SetRssi(v int32)`

SetRssi sets Rssi field to given value.


### GetSsid

`func (o *EventsRogue) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *EventsRogue) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *EventsRogue) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetTimestamp

`func (o *EventsRogue) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventsRogue) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventsRogue) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


