# RrmConsideration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **int32** |  | 
**Noise** | **float32** |  | 
**OtherRssi** | Pointer to **float32** | the avg RSSI heard from other APs (that does NOT belongs to the same site) | [optional] 
**OtherSsid** | Pointer to **string** | SSID from other AP that we heard from with the max RSSI | [optional] 
**UtilScore** | **float32** | utilization score, 0-1, lower means less utilization (cleaner RF) | 
**UtilScoreNonWifi** | **float32** | non-wifi utilization score, 0-1, lower means less utilization (cleaner RF) | 
**UtilScoreOther** | **float32** | other utilization score, 0-1, lower means less utilization (cleaner RF) | 

## Methods

### NewRrmConsideration

`func NewRrmConsideration(channel int32, noise float32, utilScore float32, utilScoreNonWifi float32, utilScoreOther float32, ) *RrmConsideration`

NewRrmConsideration instantiates a new RrmConsideration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRrmConsiderationWithDefaults

`func NewRrmConsiderationWithDefaults() *RrmConsideration`

NewRrmConsiderationWithDefaults instantiates a new RrmConsideration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *RrmConsideration) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *RrmConsideration) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *RrmConsideration) SetChannel(v int32)`

SetChannel sets Channel field to given value.


### GetNoise

`func (o *RrmConsideration) GetNoise() float32`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *RrmConsideration) GetNoiseOk() (*float32, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *RrmConsideration) SetNoise(v float32)`

SetNoise sets Noise field to given value.


### GetOtherRssi

`func (o *RrmConsideration) GetOtherRssi() float32`

GetOtherRssi returns the OtherRssi field if non-nil, zero value otherwise.

### GetOtherRssiOk

`func (o *RrmConsideration) GetOtherRssiOk() (*float32, bool)`

GetOtherRssiOk returns a tuple with the OtherRssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherRssi

`func (o *RrmConsideration) SetOtherRssi(v float32)`

SetOtherRssi sets OtherRssi field to given value.

### HasOtherRssi

`func (o *RrmConsideration) HasOtherRssi() bool`

HasOtherRssi returns a boolean if a field has been set.

### GetOtherSsid

`func (o *RrmConsideration) GetOtherSsid() string`

GetOtherSsid returns the OtherSsid field if non-nil, zero value otherwise.

### GetOtherSsidOk

`func (o *RrmConsideration) GetOtherSsidOk() (*string, bool)`

GetOtherSsidOk returns a tuple with the OtherSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSsid

`func (o *RrmConsideration) SetOtherSsid(v string)`

SetOtherSsid sets OtherSsid field to given value.

### HasOtherSsid

`func (o *RrmConsideration) HasOtherSsid() bool`

HasOtherSsid returns a boolean if a field has been set.

### GetUtilScore

`func (o *RrmConsideration) GetUtilScore() float32`

GetUtilScore returns the UtilScore field if non-nil, zero value otherwise.

### GetUtilScoreOk

`func (o *RrmConsideration) GetUtilScoreOk() (*float32, bool)`

GetUtilScoreOk returns a tuple with the UtilScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilScore

`func (o *RrmConsideration) SetUtilScore(v float32)`

SetUtilScore sets UtilScore field to given value.


### GetUtilScoreNonWifi

`func (o *RrmConsideration) GetUtilScoreNonWifi() float32`

GetUtilScoreNonWifi returns the UtilScoreNonWifi field if non-nil, zero value otherwise.

### GetUtilScoreNonWifiOk

`func (o *RrmConsideration) GetUtilScoreNonWifiOk() (*float32, bool)`

GetUtilScoreNonWifiOk returns a tuple with the UtilScoreNonWifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilScoreNonWifi

`func (o *RrmConsideration) SetUtilScoreNonWifi(v float32)`

SetUtilScoreNonWifi sets UtilScoreNonWifi field to given value.


### GetUtilScoreOther

`func (o *RrmConsideration) GetUtilScoreOther() float32`

GetUtilScoreOther returns the UtilScoreOther field if non-nil, zero value otherwise.

### GetUtilScoreOtherOk

`func (o *RrmConsideration) GetUtilScoreOtherOk() (*float32, bool)`

GetUtilScoreOtherOk returns a tuple with the UtilScoreOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilScoreOther

`func (o *RrmConsideration) SetUtilScoreOther(v float32)`

SetUtilScoreOther sets UtilScoreOther field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


