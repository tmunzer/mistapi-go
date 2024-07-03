# ResponseConfigHistorySearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel24** | **int32** |  | 
**Channel5** | **int32** |  | 
**RadioMacs** | Pointer to **[]string** |  | [optional] 
**Radios** | Pointer to [**[]ResponseConfigHistorySearchItemRadio**](ResponseConfigHistorySearchItemRadio.md) |  | [optional] 
**SecpolicyViolated** | **bool** |  | 
**Ssids** | Pointer to **[]string** |  | [optional] 
**Ssids24** | Pointer to **[]string** |  | [optional] 
**Ssids5** | Pointer to **[]string** |  | [optional] 
**Timestamp** | **float32** |  | 
**Version** | **string** |  | 
**Wlans** | Pointer to [**[]ResponseConfigHistorySearchItemWlan**](ResponseConfigHistorySearchItemWlan.md) |  | [optional] 

## Methods

### NewResponseConfigHistorySearchItem

`func NewResponseConfigHistorySearchItem(channel24 int32, channel5 int32, secpolicyViolated bool, timestamp float32, version string, ) *ResponseConfigHistorySearchItem`

NewResponseConfigHistorySearchItem instantiates a new ResponseConfigHistorySearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseConfigHistorySearchItemWithDefaults

`func NewResponseConfigHistorySearchItemWithDefaults() *ResponseConfigHistorySearchItem`

NewResponseConfigHistorySearchItemWithDefaults instantiates a new ResponseConfigHistorySearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel24

`func (o *ResponseConfigHistorySearchItem) GetChannel24() int32`

GetChannel24 returns the Channel24 field if non-nil, zero value otherwise.

### GetChannel24Ok

`func (o *ResponseConfigHistorySearchItem) GetChannel24Ok() (*int32, bool)`

GetChannel24Ok returns a tuple with the Channel24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel24

`func (o *ResponseConfigHistorySearchItem) SetChannel24(v int32)`

SetChannel24 sets Channel24 field to given value.


### GetChannel5

`func (o *ResponseConfigHistorySearchItem) GetChannel5() int32`

GetChannel5 returns the Channel5 field if non-nil, zero value otherwise.

### GetChannel5Ok

`func (o *ResponseConfigHistorySearchItem) GetChannel5Ok() (*int32, bool)`

GetChannel5Ok returns a tuple with the Channel5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel5

`func (o *ResponseConfigHistorySearchItem) SetChannel5(v int32)`

SetChannel5 sets Channel5 field to given value.


### GetRadioMacs

`func (o *ResponseConfigHistorySearchItem) GetRadioMacs() []string`

GetRadioMacs returns the RadioMacs field if non-nil, zero value otherwise.

### GetRadioMacsOk

`func (o *ResponseConfigHistorySearchItem) GetRadioMacsOk() (*[]string, bool)`

GetRadioMacsOk returns a tuple with the RadioMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioMacs

`func (o *ResponseConfigHistorySearchItem) SetRadioMacs(v []string)`

SetRadioMacs sets RadioMacs field to given value.

### HasRadioMacs

`func (o *ResponseConfigHistorySearchItem) HasRadioMacs() bool`

HasRadioMacs returns a boolean if a field has been set.

### GetRadios

`func (o *ResponseConfigHistorySearchItem) GetRadios() []ResponseConfigHistorySearchItemRadio`

GetRadios returns the Radios field if non-nil, zero value otherwise.

### GetRadiosOk

`func (o *ResponseConfigHistorySearchItem) GetRadiosOk() (*[]ResponseConfigHistorySearchItemRadio, bool)`

GetRadiosOk returns a tuple with the Radios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadios

`func (o *ResponseConfigHistorySearchItem) SetRadios(v []ResponseConfigHistorySearchItemRadio)`

SetRadios sets Radios field to given value.

### HasRadios

`func (o *ResponseConfigHistorySearchItem) HasRadios() bool`

HasRadios returns a boolean if a field has been set.

### GetSecpolicyViolated

`func (o *ResponseConfigHistorySearchItem) GetSecpolicyViolated() bool`

GetSecpolicyViolated returns the SecpolicyViolated field if non-nil, zero value otherwise.

### GetSecpolicyViolatedOk

`func (o *ResponseConfigHistorySearchItem) GetSecpolicyViolatedOk() (*bool, bool)`

GetSecpolicyViolatedOk returns a tuple with the SecpolicyViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecpolicyViolated

`func (o *ResponseConfigHistorySearchItem) SetSecpolicyViolated(v bool)`

SetSecpolicyViolated sets SecpolicyViolated field to given value.


### GetSsids

`func (o *ResponseConfigHistorySearchItem) GetSsids() []string`

GetSsids returns the Ssids field if non-nil, zero value otherwise.

### GetSsidsOk

`func (o *ResponseConfigHistorySearchItem) GetSsidsOk() (*[]string, bool)`

GetSsidsOk returns a tuple with the Ssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids

`func (o *ResponseConfigHistorySearchItem) SetSsids(v []string)`

SetSsids sets Ssids field to given value.

### HasSsids

`func (o *ResponseConfigHistorySearchItem) HasSsids() bool`

HasSsids returns a boolean if a field has been set.

### GetSsids24

`func (o *ResponseConfigHistorySearchItem) GetSsids24() []string`

GetSsids24 returns the Ssids24 field if non-nil, zero value otherwise.

### GetSsids24Ok

`func (o *ResponseConfigHistorySearchItem) GetSsids24Ok() (*[]string, bool)`

GetSsids24Ok returns a tuple with the Ssids24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids24

`func (o *ResponseConfigHistorySearchItem) SetSsids24(v []string)`

SetSsids24 sets Ssids24 field to given value.

### HasSsids24

`func (o *ResponseConfigHistorySearchItem) HasSsids24() bool`

HasSsids24 returns a boolean if a field has been set.

### GetSsids5

`func (o *ResponseConfigHistorySearchItem) GetSsids5() []string`

GetSsids5 returns the Ssids5 field if non-nil, zero value otherwise.

### GetSsids5Ok

`func (o *ResponseConfigHistorySearchItem) GetSsids5Ok() (*[]string, bool)`

GetSsids5Ok returns a tuple with the Ssids5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids5

`func (o *ResponseConfigHistorySearchItem) SetSsids5(v []string)`

SetSsids5 sets Ssids5 field to given value.

### HasSsids5

`func (o *ResponseConfigHistorySearchItem) HasSsids5() bool`

HasSsids5 returns a boolean if a field has been set.

### GetTimestamp

`func (o *ResponseConfigHistorySearchItem) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponseConfigHistorySearchItem) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponseConfigHistorySearchItem) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetVersion

`func (o *ResponseConfigHistorySearchItem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponseConfigHistorySearchItem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponseConfigHistorySearchItem) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetWlans

`func (o *ResponseConfigHistorySearchItem) GetWlans() []ResponseConfigHistorySearchItemWlan`

GetWlans returns the Wlans field if non-nil, zero value otherwise.

### GetWlansOk

`func (o *ResponseConfigHistorySearchItem) GetWlansOk() (*[]ResponseConfigHistorySearchItemWlan, bool)`

GetWlansOk returns a tuple with the Wlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlans

`func (o *ResponseConfigHistorySearchItem) SetWlans(v []ResponseConfigHistorySearchItemWlan)`

SetWlans sets Wlans field to given value.

### HasWlans

`func (o *ResponseConfigHistorySearchItem) HasWlans() bool`

HasWlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


