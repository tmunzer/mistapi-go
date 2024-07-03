# ResponseDeviceRadioChannels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band2440mhzAllowed** | **bool** |  | 
**Band24Channels** | **map[string][]int32** | Property key is the channel width | 
**Band24Enabled** | **bool** |  | 
**Band5Channels** | **map[string][]int32** | Property key is the channel width | 
**Band5Enabled** | **bool** |  | 
**Band6Channels** | Pointer to **map[string][]int32** | Property key is the channel width | [optional] 
**Band6Enabled** | Pointer to **bool** |  | [optional] 
**Certified** | **bool** |  | 
**Code** | **int32** |  | 
**DfsOk** | **bool** |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**Uses** | **string** |  | 

## Methods

### NewResponseDeviceRadioChannels

`func NewResponseDeviceRadioChannels(band2440mhzAllowed bool, band24Channels map[string][]int32, band24Enabled bool, band5Channels map[string][]int32, band5Enabled bool, certified bool, code int32, dfsOk bool, key string, name string, uses string, ) *ResponseDeviceRadioChannels`

NewResponseDeviceRadioChannels instantiates a new ResponseDeviceRadioChannels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeviceRadioChannelsWithDefaults

`func NewResponseDeviceRadioChannelsWithDefaults() *ResponseDeviceRadioChannels`

NewResponseDeviceRadioChannelsWithDefaults instantiates a new ResponseDeviceRadioChannels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand2440mhzAllowed

`func (o *ResponseDeviceRadioChannels) GetBand2440mhzAllowed() bool`

GetBand2440mhzAllowed returns the Band2440mhzAllowed field if non-nil, zero value otherwise.

### GetBand2440mhzAllowedOk

`func (o *ResponseDeviceRadioChannels) GetBand2440mhzAllowedOk() (*bool, bool)`

GetBand2440mhzAllowedOk returns a tuple with the Band2440mhzAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand2440mhzAllowed

`func (o *ResponseDeviceRadioChannels) SetBand2440mhzAllowed(v bool)`

SetBand2440mhzAllowed sets Band2440mhzAllowed field to given value.


### GetBand24Channels

`func (o *ResponseDeviceRadioChannels) GetBand24Channels() map[string][]int32`

GetBand24Channels returns the Band24Channels field if non-nil, zero value otherwise.

### GetBand24ChannelsOk

`func (o *ResponseDeviceRadioChannels) GetBand24ChannelsOk() (*map[string][]int32, bool)`

GetBand24ChannelsOk returns a tuple with the Band24Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Channels

`func (o *ResponseDeviceRadioChannels) SetBand24Channels(v map[string][]int32)`

SetBand24Channels sets Band24Channels field to given value.


### GetBand24Enabled

`func (o *ResponseDeviceRadioChannels) GetBand24Enabled() bool`

GetBand24Enabled returns the Band24Enabled field if non-nil, zero value otherwise.

### GetBand24EnabledOk

`func (o *ResponseDeviceRadioChannels) GetBand24EnabledOk() (*bool, bool)`

GetBand24EnabledOk returns a tuple with the Band24Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Enabled

`func (o *ResponseDeviceRadioChannels) SetBand24Enabled(v bool)`

SetBand24Enabled sets Band24Enabled field to given value.


### GetBand5Channels

`func (o *ResponseDeviceRadioChannels) GetBand5Channels() map[string][]int32`

GetBand5Channels returns the Band5Channels field if non-nil, zero value otherwise.

### GetBand5ChannelsOk

`func (o *ResponseDeviceRadioChannels) GetBand5ChannelsOk() (*map[string][]int32, bool)`

GetBand5ChannelsOk returns a tuple with the Band5Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5Channels

`func (o *ResponseDeviceRadioChannels) SetBand5Channels(v map[string][]int32)`

SetBand5Channels sets Band5Channels field to given value.


### GetBand5Enabled

`func (o *ResponseDeviceRadioChannels) GetBand5Enabled() bool`

GetBand5Enabled returns the Band5Enabled field if non-nil, zero value otherwise.

### GetBand5EnabledOk

`func (o *ResponseDeviceRadioChannels) GetBand5EnabledOk() (*bool, bool)`

GetBand5EnabledOk returns a tuple with the Band5Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5Enabled

`func (o *ResponseDeviceRadioChannels) SetBand5Enabled(v bool)`

SetBand5Enabled sets Band5Enabled field to given value.


### GetBand6Channels

`func (o *ResponseDeviceRadioChannels) GetBand6Channels() map[string][]int32`

GetBand6Channels returns the Band6Channels field if non-nil, zero value otherwise.

### GetBand6ChannelsOk

`func (o *ResponseDeviceRadioChannels) GetBand6ChannelsOk() (*map[string][]int32, bool)`

GetBand6ChannelsOk returns a tuple with the Band6Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6Channels

`func (o *ResponseDeviceRadioChannels) SetBand6Channels(v map[string][]int32)`

SetBand6Channels sets Band6Channels field to given value.

### HasBand6Channels

`func (o *ResponseDeviceRadioChannels) HasBand6Channels() bool`

HasBand6Channels returns a boolean if a field has been set.

### GetBand6Enabled

`func (o *ResponseDeviceRadioChannels) GetBand6Enabled() bool`

GetBand6Enabled returns the Band6Enabled field if non-nil, zero value otherwise.

### GetBand6EnabledOk

`func (o *ResponseDeviceRadioChannels) GetBand6EnabledOk() (*bool, bool)`

GetBand6EnabledOk returns a tuple with the Band6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6Enabled

`func (o *ResponseDeviceRadioChannels) SetBand6Enabled(v bool)`

SetBand6Enabled sets Band6Enabled field to given value.

### HasBand6Enabled

`func (o *ResponseDeviceRadioChannels) HasBand6Enabled() bool`

HasBand6Enabled returns a boolean if a field has been set.

### GetCertified

`func (o *ResponseDeviceRadioChannels) GetCertified() bool`

GetCertified returns the Certified field if non-nil, zero value otherwise.

### GetCertifiedOk

`func (o *ResponseDeviceRadioChannels) GetCertifiedOk() (*bool, bool)`

GetCertifiedOk returns a tuple with the Certified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertified

`func (o *ResponseDeviceRadioChannels) SetCertified(v bool)`

SetCertified sets Certified field to given value.


### GetCode

`func (o *ResponseDeviceRadioChannels) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ResponseDeviceRadioChannels) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ResponseDeviceRadioChannels) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDfsOk

`func (o *ResponseDeviceRadioChannels) GetDfsOk() bool`

GetDfsOk returns the DfsOk field if non-nil, zero value otherwise.

### GetDfsOkOk

`func (o *ResponseDeviceRadioChannels) GetDfsOkOk() (*bool, bool)`

GetDfsOkOk returns a tuple with the DfsOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsOk

`func (o *ResponseDeviceRadioChannels) SetDfsOk(v bool)`

SetDfsOk sets DfsOk field to given value.


### GetKey

`func (o *ResponseDeviceRadioChannels) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ResponseDeviceRadioChannels) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ResponseDeviceRadioChannels) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ResponseDeviceRadioChannels) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseDeviceRadioChannels) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseDeviceRadioChannels) SetName(v string)`

SetName sets Name field to given value.


### GetUses

`func (o *ResponseDeviceRadioChannels) GetUses() string`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *ResponseDeviceRadioChannels) GetUsesOk() (*string, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *ResponseDeviceRadioChannels) SetUses(v string)`

SetUses sets Uses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


