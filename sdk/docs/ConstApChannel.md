# ConstApChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band2440mhzAllowed** | Pointer to **bool** |  | [optional] 
**Band24Channels** | Pointer to **map[string][]int32** | Property key is the channel width | [optional] 
**Band24Enabled** | Pointer to **bool** |  | [optional] 
**Band5Channels** | Pointer to **map[string][]int32** | Property key is the channel width | [optional] 
**Band5Enabled** | Pointer to **bool** |  | [optional] 
**Band6Channels** | Pointer to **map[string][]int32** | Property key is the channel width | [optional] 
**Band6Enabled** | Pointer to **bool** |  | [optional] 
**Certified** | Pointer to **bool** |  | [optional] 
**Code** | Pointer to **int32** | country code, ISO 3166-1 numeric | [optional] 
**DfsOk** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** | country code, in two-character | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Uses** | Pointer to **string** |  | [optional] 

## Methods

### NewConstApChannel

`func NewConstApChannel() *ConstApChannel`

NewConstApChannel instantiates a new ConstApChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstApChannelWithDefaults

`func NewConstApChannelWithDefaults() *ConstApChannel`

NewConstApChannelWithDefaults instantiates a new ConstApChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand2440mhzAllowed

`func (o *ConstApChannel) GetBand2440mhzAllowed() bool`

GetBand2440mhzAllowed returns the Band2440mhzAllowed field if non-nil, zero value otherwise.

### GetBand2440mhzAllowedOk

`func (o *ConstApChannel) GetBand2440mhzAllowedOk() (*bool, bool)`

GetBand2440mhzAllowedOk returns a tuple with the Band2440mhzAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand2440mhzAllowed

`func (o *ConstApChannel) SetBand2440mhzAllowed(v bool)`

SetBand2440mhzAllowed sets Band2440mhzAllowed field to given value.

### HasBand2440mhzAllowed

`func (o *ConstApChannel) HasBand2440mhzAllowed() bool`

HasBand2440mhzAllowed returns a boolean if a field has been set.

### GetBand24Channels

`func (o *ConstApChannel) GetBand24Channels() map[string][]int32`

GetBand24Channels returns the Band24Channels field if non-nil, zero value otherwise.

### GetBand24ChannelsOk

`func (o *ConstApChannel) GetBand24ChannelsOk() (*map[string][]int32, bool)`

GetBand24ChannelsOk returns a tuple with the Band24Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Channels

`func (o *ConstApChannel) SetBand24Channels(v map[string][]int32)`

SetBand24Channels sets Band24Channels field to given value.

### HasBand24Channels

`func (o *ConstApChannel) HasBand24Channels() bool`

HasBand24Channels returns a boolean if a field has been set.

### GetBand24Enabled

`func (o *ConstApChannel) GetBand24Enabled() bool`

GetBand24Enabled returns the Band24Enabled field if non-nil, zero value otherwise.

### GetBand24EnabledOk

`func (o *ConstApChannel) GetBand24EnabledOk() (*bool, bool)`

GetBand24EnabledOk returns a tuple with the Band24Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Enabled

`func (o *ConstApChannel) SetBand24Enabled(v bool)`

SetBand24Enabled sets Band24Enabled field to given value.

### HasBand24Enabled

`func (o *ConstApChannel) HasBand24Enabled() bool`

HasBand24Enabled returns a boolean if a field has been set.

### GetBand5Channels

`func (o *ConstApChannel) GetBand5Channels() map[string][]int32`

GetBand5Channels returns the Band5Channels field if non-nil, zero value otherwise.

### GetBand5ChannelsOk

`func (o *ConstApChannel) GetBand5ChannelsOk() (*map[string][]int32, bool)`

GetBand5ChannelsOk returns a tuple with the Band5Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5Channels

`func (o *ConstApChannel) SetBand5Channels(v map[string][]int32)`

SetBand5Channels sets Band5Channels field to given value.

### HasBand5Channels

`func (o *ConstApChannel) HasBand5Channels() bool`

HasBand5Channels returns a boolean if a field has been set.

### GetBand5Enabled

`func (o *ConstApChannel) GetBand5Enabled() bool`

GetBand5Enabled returns the Band5Enabled field if non-nil, zero value otherwise.

### GetBand5EnabledOk

`func (o *ConstApChannel) GetBand5EnabledOk() (*bool, bool)`

GetBand5EnabledOk returns a tuple with the Band5Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5Enabled

`func (o *ConstApChannel) SetBand5Enabled(v bool)`

SetBand5Enabled sets Band5Enabled field to given value.

### HasBand5Enabled

`func (o *ConstApChannel) HasBand5Enabled() bool`

HasBand5Enabled returns a boolean if a field has been set.

### GetBand6Channels

`func (o *ConstApChannel) GetBand6Channels() map[string][]int32`

GetBand6Channels returns the Band6Channels field if non-nil, zero value otherwise.

### GetBand6ChannelsOk

`func (o *ConstApChannel) GetBand6ChannelsOk() (*map[string][]int32, bool)`

GetBand6ChannelsOk returns a tuple with the Band6Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6Channels

`func (o *ConstApChannel) SetBand6Channels(v map[string][]int32)`

SetBand6Channels sets Band6Channels field to given value.

### HasBand6Channels

`func (o *ConstApChannel) HasBand6Channels() bool`

HasBand6Channels returns a boolean if a field has been set.

### GetBand6Enabled

`func (o *ConstApChannel) GetBand6Enabled() bool`

GetBand6Enabled returns the Band6Enabled field if non-nil, zero value otherwise.

### GetBand6EnabledOk

`func (o *ConstApChannel) GetBand6EnabledOk() (*bool, bool)`

GetBand6EnabledOk returns a tuple with the Band6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6Enabled

`func (o *ConstApChannel) SetBand6Enabled(v bool)`

SetBand6Enabled sets Band6Enabled field to given value.

### HasBand6Enabled

`func (o *ConstApChannel) HasBand6Enabled() bool`

HasBand6Enabled returns a boolean if a field has been set.

### GetCertified

`func (o *ConstApChannel) GetCertified() bool`

GetCertified returns the Certified field if non-nil, zero value otherwise.

### GetCertifiedOk

`func (o *ConstApChannel) GetCertifiedOk() (*bool, bool)`

GetCertifiedOk returns a tuple with the Certified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertified

`func (o *ConstApChannel) SetCertified(v bool)`

SetCertified sets Certified field to given value.

### HasCertified

`func (o *ConstApChannel) HasCertified() bool`

HasCertified returns a boolean if a field has been set.

### GetCode

`func (o *ConstApChannel) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ConstApChannel) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ConstApChannel) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ConstApChannel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDfsOk

`func (o *ConstApChannel) GetDfsOk() bool`

GetDfsOk returns the DfsOk field if non-nil, zero value otherwise.

### GetDfsOkOk

`func (o *ConstApChannel) GetDfsOkOk() (*bool, bool)`

GetDfsOkOk returns a tuple with the DfsOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsOk

`func (o *ConstApChannel) SetDfsOk(v bool)`

SetDfsOk sets DfsOk field to given value.

### HasDfsOk

`func (o *ConstApChannel) HasDfsOk() bool`

HasDfsOk returns a boolean if a field has been set.

### GetKey

`func (o *ConstApChannel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConstApChannel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConstApChannel) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ConstApChannel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *ConstApChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConstApChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConstApChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConstApChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUses

`func (o *ConstApChannel) GetUses() string`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *ConstApChannel) GetUsesOk() (*string, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *ConstApChannel) SetUses(v string)`

SetUses sets Uses field to given value.

### HasUses

`func (o *ConstApChannel) HasUses() bool`

HasUses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


