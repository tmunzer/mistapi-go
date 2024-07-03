# InsightRogueAp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | **string** | mac of the device that had strongest signal strength for ssid/bssid pair | 
**AvgRssi** | **float32** | average signal strength of ap_mac for ssid/bssid pair | 
**Bssid** | **string** | bssid of the network detected as threat | 
**Channel** | **string** | channel over which ap_mac heard ssid/bssid pair | 
**DeltaX** | Pointer to **float32** | X position relative to the reporting AP (&#x60;ap_mac&#x60;) | [optional] 
**DeltaY** | Pointer to **float32** | Y position relative to the reporting AP (&#x60;ap_mac&#x60;) | [optional] 
**NumAps** | **int32** | num of aps that heard the ssid/bssid pair | 
**SeenOnLan** | Pointer to **bool** | whether the reporting AP see a wireless client (on LAN) connecting to it | [optional] 
**Ssid** | Pointer to **string** | ssid of the network detected as threat | [optional] 
**TimesHeard** | Pointer to **int32** | represents number of times the pair was heard in the interval. Each count roughly corresponds to a minute. | [optional] 

## Methods

### NewInsightRogueAp

`func NewInsightRogueAp(apMac string, avgRssi float32, bssid string, channel string, numAps int32, ) *InsightRogueAp`

NewInsightRogueAp instantiates a new InsightRogueAp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightRogueApWithDefaults

`func NewInsightRogueApWithDefaults() *InsightRogueAp`

NewInsightRogueApWithDefaults instantiates a new InsightRogueAp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *InsightRogueAp) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *InsightRogueAp) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *InsightRogueAp) SetApMac(v string)`

SetApMac sets ApMac field to given value.


### GetAvgRssi

`func (o *InsightRogueAp) GetAvgRssi() float32`

GetAvgRssi returns the AvgRssi field if non-nil, zero value otherwise.

### GetAvgRssiOk

`func (o *InsightRogueAp) GetAvgRssiOk() (*float32, bool)`

GetAvgRssiOk returns a tuple with the AvgRssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgRssi

`func (o *InsightRogueAp) SetAvgRssi(v float32)`

SetAvgRssi sets AvgRssi field to given value.


### GetBssid

`func (o *InsightRogueAp) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *InsightRogueAp) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *InsightRogueAp) SetBssid(v string)`

SetBssid sets Bssid field to given value.


### GetChannel

`func (o *InsightRogueAp) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InsightRogueAp) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InsightRogueAp) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetDeltaX

`func (o *InsightRogueAp) GetDeltaX() float32`

GetDeltaX returns the DeltaX field if non-nil, zero value otherwise.

### GetDeltaXOk

`func (o *InsightRogueAp) GetDeltaXOk() (*float32, bool)`

GetDeltaXOk returns a tuple with the DeltaX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaX

`func (o *InsightRogueAp) SetDeltaX(v float32)`

SetDeltaX sets DeltaX field to given value.

### HasDeltaX

`func (o *InsightRogueAp) HasDeltaX() bool`

HasDeltaX returns a boolean if a field has been set.

### GetDeltaY

`func (o *InsightRogueAp) GetDeltaY() float32`

GetDeltaY returns the DeltaY field if non-nil, zero value otherwise.

### GetDeltaYOk

`func (o *InsightRogueAp) GetDeltaYOk() (*float32, bool)`

GetDeltaYOk returns a tuple with the DeltaY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaY

`func (o *InsightRogueAp) SetDeltaY(v float32)`

SetDeltaY sets DeltaY field to given value.

### HasDeltaY

`func (o *InsightRogueAp) HasDeltaY() bool`

HasDeltaY returns a boolean if a field has been set.

### GetNumAps

`func (o *InsightRogueAp) GetNumAps() int32`

GetNumAps returns the NumAps field if non-nil, zero value otherwise.

### GetNumApsOk

`func (o *InsightRogueAp) GetNumApsOk() (*int32, bool)`

GetNumApsOk returns a tuple with the NumAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAps

`func (o *InsightRogueAp) SetNumAps(v int32)`

SetNumAps sets NumAps field to given value.


### GetSeenOnLan

`func (o *InsightRogueAp) GetSeenOnLan() bool`

GetSeenOnLan returns the SeenOnLan field if non-nil, zero value otherwise.

### GetSeenOnLanOk

`func (o *InsightRogueAp) GetSeenOnLanOk() (*bool, bool)`

GetSeenOnLanOk returns a tuple with the SeenOnLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenOnLan

`func (o *InsightRogueAp) SetSeenOnLan(v bool)`

SetSeenOnLan sets SeenOnLan field to given value.

### HasSeenOnLan

`func (o *InsightRogueAp) HasSeenOnLan() bool`

HasSeenOnLan returns a boolean if a field has been set.

### GetSsid

`func (o *InsightRogueAp) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InsightRogueAp) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InsightRogueAp) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *InsightRogueAp) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetTimesHeard

`func (o *InsightRogueAp) GetTimesHeard() int32`

GetTimesHeard returns the TimesHeard field if non-nil, zero value otherwise.

### GetTimesHeardOk

`func (o *InsightRogueAp) GetTimesHeardOk() (*int32, bool)`

GetTimesHeardOk returns a tuple with the TimesHeard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesHeard

`func (o *InsightRogueAp) SetTimesHeard(v int32)`

SetTimesHeard sets TimesHeard field to given value.

### HasTimesHeard

`func (o *InsightRogueAp) HasTimesHeard() bool`

HasTimesHeard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


