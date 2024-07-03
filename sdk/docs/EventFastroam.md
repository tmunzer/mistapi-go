# EventFastroam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | **string** |  | 
**ClientMac** | **string** |  | 
**Fromap** | **string** |  | 
**Latency** | **float32** |  | 
**Ssid** | **string** |  | 
**Subtype** | Pointer to **string** |  | [optional] 
**Timestamp** | **float32** | timestamp of the event in nsec | 
**Type** | Pointer to [**EventFastroamType**](EventFastroamType.md) |  | [optional] 

## Methods

### NewEventFastroam

`func NewEventFastroam(apMac string, clientMac string, fromap string, latency float32, ssid string, timestamp float32, ) *EventFastroam`

NewEventFastroam instantiates a new EventFastroam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventFastroamWithDefaults

`func NewEventFastroamWithDefaults() *EventFastroam`

NewEventFastroamWithDefaults instantiates a new EventFastroam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *EventFastroam) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *EventFastroam) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *EventFastroam) SetApMac(v string)`

SetApMac sets ApMac field to given value.


### GetClientMac

`func (o *EventFastroam) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *EventFastroam) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *EventFastroam) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.


### GetFromap

`func (o *EventFastroam) GetFromap() string`

GetFromap returns the Fromap field if non-nil, zero value otherwise.

### GetFromapOk

`func (o *EventFastroam) GetFromapOk() (*string, bool)`

GetFromapOk returns a tuple with the Fromap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromap

`func (o *EventFastroam) SetFromap(v string)`

SetFromap sets Fromap field to given value.


### GetLatency

`func (o *EventFastroam) GetLatency() float32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *EventFastroam) GetLatencyOk() (*float32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *EventFastroam) SetLatency(v float32)`

SetLatency sets Latency field to given value.


### GetSsid

`func (o *EventFastroam) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *EventFastroam) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *EventFastroam) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetSubtype

`func (o *EventFastroam) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *EventFastroam) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *EventFastroam) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *EventFastroam) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetTimestamp

`func (o *EventFastroam) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventFastroam) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventFastroam) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *EventFastroam) GetType() EventFastroamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventFastroam) GetTypeOk() (*EventFastroamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventFastroam) SetType(v EventFastroamType)`

SetType sets Type field to given value.

### HasType

`func (o *EventFastroam) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


