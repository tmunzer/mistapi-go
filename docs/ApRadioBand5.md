# ApRadioBand5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRrmDisable** | Pointer to **bool** |  | [optional] [default to false]
**AntGain** | Pointer to **NullableInt32** |  | [optional] [default to 0]
**AntennaMode** | Pointer to [**RadioBandAntennaMode**](RadioBandAntennaMode.md) |  | [optional] [default to RADIOBANDANTENNAMODE_DEFAULT]
**Bandwidth** | Pointer to [**Dot11Bandwidth5**](Dot11Bandwidth5.md) |  | [optional] 
**Channel** | Pointer to **NullableInt32** | For Device. (primary) channel for the band, 0 means using the Site Setting | [optional] 
**Channels** | Pointer to **[]int32** | For RFTemplates. List of channels, null or empty array means auto | [optional] [default to []]
**Disabled** | Pointer to **bool** | whether to disable the radio | [optional] [default to false]
**Power** | Pointer to **NullableInt32** | TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / … | [optional] 
**PowerMax** | Pointer to **NullableInt32** | when power&#x3D;0, max tx power to use, HW-specific values will be used if not set | [optional] [default to 17]
**PowerMin** | Pointer to **NullableInt32** | when power&#x3D;0, min tx power to use, HW-specific values will be used if not set | [optional] [default to 8]
**Preamble** | Pointer to [**RadioBandPreamble**](RadioBandPreamble.md) |  | [optional] [default to RADIOBANDPREAMBLE_SHORT]

## Methods

### NewApRadioBand5

`func NewApRadioBand5() *ApRadioBand5`

NewApRadioBand5 instantiates a new ApRadioBand5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApRadioBand5WithDefaults

`func NewApRadioBand5WithDefaults() *ApRadioBand5`

NewApRadioBand5WithDefaults instantiates a new ApRadioBand5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowRrmDisable

`func (o *ApRadioBand5) GetAllowRrmDisable() bool`

GetAllowRrmDisable returns the AllowRrmDisable field if non-nil, zero value otherwise.

### GetAllowRrmDisableOk

`func (o *ApRadioBand5) GetAllowRrmDisableOk() (*bool, bool)`

GetAllowRrmDisableOk returns a tuple with the AllowRrmDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRrmDisable

`func (o *ApRadioBand5) SetAllowRrmDisable(v bool)`

SetAllowRrmDisable sets AllowRrmDisable field to given value.

### HasAllowRrmDisable

`func (o *ApRadioBand5) HasAllowRrmDisable() bool`

HasAllowRrmDisable returns a boolean if a field has been set.

### GetAntGain

`func (o *ApRadioBand5) GetAntGain() int32`

GetAntGain returns the AntGain field if non-nil, zero value otherwise.

### GetAntGainOk

`func (o *ApRadioBand5) GetAntGainOk() (*int32, bool)`

GetAntGainOk returns a tuple with the AntGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntGain

`func (o *ApRadioBand5) SetAntGain(v int32)`

SetAntGain sets AntGain field to given value.

### HasAntGain

`func (o *ApRadioBand5) HasAntGain() bool`

HasAntGain returns a boolean if a field has been set.

### SetAntGainNil

`func (o *ApRadioBand5) SetAntGainNil(b bool)`

 SetAntGainNil sets the value for AntGain to be an explicit nil

### UnsetAntGain
`func (o *ApRadioBand5) UnsetAntGain()`

UnsetAntGain ensures that no value is present for AntGain, not even an explicit nil
### GetAntennaMode

`func (o *ApRadioBand5) GetAntennaMode() RadioBandAntennaMode`

GetAntennaMode returns the AntennaMode field if non-nil, zero value otherwise.

### GetAntennaModeOk

`func (o *ApRadioBand5) GetAntennaModeOk() (*RadioBandAntennaMode, bool)`

GetAntennaModeOk returns a tuple with the AntennaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntennaMode

`func (o *ApRadioBand5) SetAntennaMode(v RadioBandAntennaMode)`

SetAntennaMode sets AntennaMode field to given value.

### HasAntennaMode

`func (o *ApRadioBand5) HasAntennaMode() bool`

HasAntennaMode returns a boolean if a field has been set.

### GetBandwidth

`func (o *ApRadioBand5) GetBandwidth() Dot11Bandwidth5`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ApRadioBand5) GetBandwidthOk() (*Dot11Bandwidth5, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ApRadioBand5) SetBandwidth(v Dot11Bandwidth5)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *ApRadioBand5) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetChannel

`func (o *ApRadioBand5) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApRadioBand5) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApRadioBand5) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApRadioBand5) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *ApRadioBand5) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *ApRadioBand5) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetChannels

`func (o *ApRadioBand5) GetChannels() []int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ApRadioBand5) GetChannelsOk() (*[]int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ApRadioBand5) SetChannels(v []int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ApRadioBand5) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *ApRadioBand5) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *ApRadioBand5) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetDisabled

`func (o *ApRadioBand5) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ApRadioBand5) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ApRadioBand5) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ApRadioBand5) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetPower

`func (o *ApRadioBand5) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *ApRadioBand5) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *ApRadioBand5) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *ApRadioBand5) HasPower() bool`

HasPower returns a boolean if a field has been set.

### SetPowerNil

`func (o *ApRadioBand5) SetPowerNil(b bool)`

 SetPowerNil sets the value for Power to be an explicit nil

### UnsetPower
`func (o *ApRadioBand5) UnsetPower()`

UnsetPower ensures that no value is present for Power, not even an explicit nil
### GetPowerMax

`func (o *ApRadioBand5) GetPowerMax() int32`

GetPowerMax returns the PowerMax field if non-nil, zero value otherwise.

### GetPowerMaxOk

`func (o *ApRadioBand5) GetPowerMaxOk() (*int32, bool)`

GetPowerMaxOk returns a tuple with the PowerMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMax

`func (o *ApRadioBand5) SetPowerMax(v int32)`

SetPowerMax sets PowerMax field to given value.

### HasPowerMax

`func (o *ApRadioBand5) HasPowerMax() bool`

HasPowerMax returns a boolean if a field has been set.

### SetPowerMaxNil

`func (o *ApRadioBand5) SetPowerMaxNil(b bool)`

 SetPowerMaxNil sets the value for PowerMax to be an explicit nil

### UnsetPowerMax
`func (o *ApRadioBand5) UnsetPowerMax()`

UnsetPowerMax ensures that no value is present for PowerMax, not even an explicit nil
### GetPowerMin

`func (o *ApRadioBand5) GetPowerMin() int32`

GetPowerMin returns the PowerMin field if non-nil, zero value otherwise.

### GetPowerMinOk

`func (o *ApRadioBand5) GetPowerMinOk() (*int32, bool)`

GetPowerMinOk returns a tuple with the PowerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMin

`func (o *ApRadioBand5) SetPowerMin(v int32)`

SetPowerMin sets PowerMin field to given value.

### HasPowerMin

`func (o *ApRadioBand5) HasPowerMin() bool`

HasPowerMin returns a boolean if a field has been set.

### SetPowerMinNil

`func (o *ApRadioBand5) SetPowerMinNil(b bool)`

 SetPowerMinNil sets the value for PowerMin to be an explicit nil

### UnsetPowerMin
`func (o *ApRadioBand5) UnsetPowerMin()`

UnsetPowerMin ensures that no value is present for PowerMin, not even an explicit nil
### GetPreamble

`func (o *ApRadioBand5) GetPreamble() RadioBandPreamble`

GetPreamble returns the Preamble field if non-nil, zero value otherwise.

### GetPreambleOk

`func (o *ApRadioBand5) GetPreambleOk() (*RadioBandPreamble, bool)`

GetPreambleOk returns a tuple with the Preamble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreamble

`func (o *ApRadioBand5) SetPreamble(v RadioBandPreamble)`

SetPreamble sets Preamble field to given value.

### HasPreamble

`func (o *ApRadioBand5) HasPreamble() bool`

HasPreamble returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


