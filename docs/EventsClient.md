# EventsClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to **string** |  | [optional] 
**Band** | [**Dot11Band**](Dot11Band.md) |  | 
**Bssid** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **int32** |  | [optional] 
**Proto** | [**Dot11Proto**](Dot11Proto.md) |  | 
**Ssid** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Timestamp** | **float32** |  | 
**Type** | Pointer to **string** | event type, e.g. MARVIS_EVENT_CLIENT_FBT_FAILURE | [optional] 
**TypeCode** | Pointer to **int32** | for assoc/disassoc events | [optional] 
**WlanId** | Pointer to **string** |  | [optional] 

## Methods

### NewEventsClient

`func NewEventsClient(band Dot11Band, proto Dot11Proto, timestamp float32, ) *EventsClient`

NewEventsClient instantiates a new EventsClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsClientWithDefaults

`func NewEventsClientWithDefaults() *EventsClient`

NewEventsClientWithDefaults instantiates a new EventsClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *EventsClient) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *EventsClient) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *EventsClient) SetAp(v string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *EventsClient) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetBand

`func (o *EventsClient) GetBand() Dot11Band`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *EventsClient) GetBandOk() (*Dot11Band, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *EventsClient) SetBand(v Dot11Band)`

SetBand sets Band field to given value.


### GetBssid

`func (o *EventsClient) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *EventsClient) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *EventsClient) SetBssid(v string)`

SetBssid sets Bssid field to given value.

### HasBssid

`func (o *EventsClient) HasBssid() bool`

HasBssid returns a boolean if a field has been set.

### GetChannel

`func (o *EventsClient) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *EventsClient) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *EventsClient) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *EventsClient) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetProto

`func (o *EventsClient) GetProto() Dot11Proto`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *EventsClient) GetProtoOk() (*Dot11Proto, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *EventsClient) SetProto(v Dot11Proto)`

SetProto sets Proto field to given value.


### GetSsid

`func (o *EventsClient) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *EventsClient) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *EventsClient) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *EventsClient) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetText

`func (o *EventsClient) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EventsClient) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EventsClient) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *EventsClient) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTimestamp

`func (o *EventsClient) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventsClient) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventsClient) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *EventsClient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventsClient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventsClient) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventsClient) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeCode

`func (o *EventsClient) GetTypeCode() int32`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *EventsClient) GetTypeCodeOk() (*int32, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *EventsClient) SetTypeCode(v int32)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *EventsClient) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetWlanId

`func (o *EventsClient) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *EventsClient) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *EventsClient) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *EventsClient) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


