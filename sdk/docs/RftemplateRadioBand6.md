# RftemplateRadioBand6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRrmDisable** | Pointer to **bool** |  | [optional] [default to false]
**AntGain** | Pointer to **NullableInt32** |  | [optional] [default to 0]
**AntennaMode** | Pointer to [**RadioBandAntennaMode**](RadioBandAntennaMode.md) |  | [optional] [default to RADIOBANDANTENNAMODE_DEFAULT]
**Bandwidth** | Pointer to [**Dot11Bandwidth6**](Dot11Bandwidth6.md) |  | [optional] [default to DOT11BANDWIDTH6__80]
**Channels** | Pointer to **[]int32** | For RFTemplates. List of channels, null or empty array means auto | [optional] [default to []]
**Disabled** | Pointer to **bool** | whether to disable the radio | [optional] [default to false]
**Power** | Pointer to **NullableInt32** | TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / … | [optional] 
**PowerMax** | Pointer to **NullableInt32** | when power&#x3D;0, max tx power to use, HW-specific values will be used if not set | [optional] [default to 18]
**PowerMin** | Pointer to **NullableInt32** | when power&#x3D;0, min tx power to use, HW-specific values will be used if not set | [optional] [default to 8]
**Preamble** | Pointer to [**RadioBandPreamble**](RadioBandPreamble.md) |  | [optional] [default to RADIOBANDPREAMBLE_SHORT]
**StandardPower** | Pointer to **bool** | for 6GHz Only, standard-power operation, AFC (Automatic Frequency Coordination) will be performed and we&#39;ll fallback to Low Power Indoor if AFC failed | [optional] [default to false]

## Methods

### NewRftemplateRadioBand6

`func NewRftemplateRadioBand6() *RftemplateRadioBand6`

NewRftemplateRadioBand6 instantiates a new RftemplateRadioBand6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRftemplateRadioBand6WithDefaults

`func NewRftemplateRadioBand6WithDefaults() *RftemplateRadioBand6`

NewRftemplateRadioBand6WithDefaults instantiates a new RftemplateRadioBand6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowRrmDisable

`func (o *RftemplateRadioBand6) GetAllowRrmDisable() bool`

GetAllowRrmDisable returns the AllowRrmDisable field if non-nil, zero value otherwise.

### GetAllowRrmDisableOk

`func (o *RftemplateRadioBand6) GetAllowRrmDisableOk() (*bool, bool)`

GetAllowRrmDisableOk returns a tuple with the AllowRrmDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRrmDisable

`func (o *RftemplateRadioBand6) SetAllowRrmDisable(v bool)`

SetAllowRrmDisable sets AllowRrmDisable field to given value.

### HasAllowRrmDisable

`func (o *RftemplateRadioBand6) HasAllowRrmDisable() bool`

HasAllowRrmDisable returns a boolean if a field has been set.

### GetAntGain

`func (o *RftemplateRadioBand6) GetAntGain() int32`

GetAntGain returns the AntGain field if non-nil, zero value otherwise.

### GetAntGainOk

`func (o *RftemplateRadioBand6) GetAntGainOk() (*int32, bool)`

GetAntGainOk returns a tuple with the AntGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntGain

`func (o *RftemplateRadioBand6) SetAntGain(v int32)`

SetAntGain sets AntGain field to given value.

### HasAntGain

`func (o *RftemplateRadioBand6) HasAntGain() bool`

HasAntGain returns a boolean if a field has been set.

### SetAntGainNil

`func (o *RftemplateRadioBand6) SetAntGainNil(b bool)`

 SetAntGainNil sets the value for AntGain to be an explicit nil

### UnsetAntGain
`func (o *RftemplateRadioBand6) UnsetAntGain()`

UnsetAntGain ensures that no value is present for AntGain, not even an explicit nil
### GetAntennaMode

`func (o *RftemplateRadioBand6) GetAntennaMode() RadioBandAntennaMode`

GetAntennaMode returns the AntennaMode field if non-nil, zero value otherwise.

### GetAntennaModeOk

`func (o *RftemplateRadioBand6) GetAntennaModeOk() (*RadioBandAntennaMode, bool)`

GetAntennaModeOk returns a tuple with the AntennaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntennaMode

`func (o *RftemplateRadioBand6) SetAntennaMode(v RadioBandAntennaMode)`

SetAntennaMode sets AntennaMode field to given value.

### HasAntennaMode

`func (o *RftemplateRadioBand6) HasAntennaMode() bool`

HasAntennaMode returns a boolean if a field has been set.

### GetBandwidth

`func (o *RftemplateRadioBand6) GetBandwidth() Dot11Bandwidth6`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *RftemplateRadioBand6) GetBandwidthOk() (*Dot11Bandwidth6, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *RftemplateRadioBand6) SetBandwidth(v Dot11Bandwidth6)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *RftemplateRadioBand6) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetChannels

`func (o *RftemplateRadioBand6) GetChannels() []int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *RftemplateRadioBand6) GetChannelsOk() (*[]int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *RftemplateRadioBand6) SetChannels(v []int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *RftemplateRadioBand6) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *RftemplateRadioBand6) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *RftemplateRadioBand6) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetDisabled

`func (o *RftemplateRadioBand6) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RftemplateRadioBand6) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RftemplateRadioBand6) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RftemplateRadioBand6) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetPower

`func (o *RftemplateRadioBand6) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *RftemplateRadioBand6) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *RftemplateRadioBand6) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *RftemplateRadioBand6) HasPower() bool`

HasPower returns a boolean if a field has been set.

### SetPowerNil

`func (o *RftemplateRadioBand6) SetPowerNil(b bool)`

 SetPowerNil sets the value for Power to be an explicit nil

### UnsetPower
`func (o *RftemplateRadioBand6) UnsetPower()`

UnsetPower ensures that no value is present for Power, not even an explicit nil
### GetPowerMax

`func (o *RftemplateRadioBand6) GetPowerMax() int32`

GetPowerMax returns the PowerMax field if non-nil, zero value otherwise.

### GetPowerMaxOk

`func (o *RftemplateRadioBand6) GetPowerMaxOk() (*int32, bool)`

GetPowerMaxOk returns a tuple with the PowerMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMax

`func (o *RftemplateRadioBand6) SetPowerMax(v int32)`

SetPowerMax sets PowerMax field to given value.

### HasPowerMax

`func (o *RftemplateRadioBand6) HasPowerMax() bool`

HasPowerMax returns a boolean if a field has been set.

### SetPowerMaxNil

`func (o *RftemplateRadioBand6) SetPowerMaxNil(b bool)`

 SetPowerMaxNil sets the value for PowerMax to be an explicit nil

### UnsetPowerMax
`func (o *RftemplateRadioBand6) UnsetPowerMax()`

UnsetPowerMax ensures that no value is present for PowerMax, not even an explicit nil
### GetPowerMin

`func (o *RftemplateRadioBand6) GetPowerMin() int32`

GetPowerMin returns the PowerMin field if non-nil, zero value otherwise.

### GetPowerMinOk

`func (o *RftemplateRadioBand6) GetPowerMinOk() (*int32, bool)`

GetPowerMinOk returns a tuple with the PowerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMin

`func (o *RftemplateRadioBand6) SetPowerMin(v int32)`

SetPowerMin sets PowerMin field to given value.

### HasPowerMin

`func (o *RftemplateRadioBand6) HasPowerMin() bool`

HasPowerMin returns a boolean if a field has been set.

### SetPowerMinNil

`func (o *RftemplateRadioBand6) SetPowerMinNil(b bool)`

 SetPowerMinNil sets the value for PowerMin to be an explicit nil

### UnsetPowerMin
`func (o *RftemplateRadioBand6) UnsetPowerMin()`

UnsetPowerMin ensures that no value is present for PowerMin, not even an explicit nil
### GetPreamble

`func (o *RftemplateRadioBand6) GetPreamble() RadioBandPreamble`

GetPreamble returns the Preamble field if non-nil, zero value otherwise.

### GetPreambleOk

`func (o *RftemplateRadioBand6) GetPreambleOk() (*RadioBandPreamble, bool)`

GetPreambleOk returns a tuple with the Preamble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreamble

`func (o *RftemplateRadioBand6) SetPreamble(v RadioBandPreamble)`

SetPreamble sets Preamble field to given value.

### HasPreamble

`func (o *RftemplateRadioBand6) HasPreamble() bool`

HasPreamble returns a boolean if a field has been set.

### GetStandardPower

`func (o *RftemplateRadioBand6) GetStandardPower() bool`

GetStandardPower returns the StandardPower field if non-nil, zero value otherwise.

### GetStandardPowerOk

`func (o *RftemplateRadioBand6) GetStandardPowerOk() (*bool, bool)`

GetStandardPowerOk returns a tuple with the StandardPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPower

`func (o *RftemplateRadioBand6) SetStandardPower(v bool)`

SetStandardPower sets StandardPower field to given value.

### HasStandardPower

`func (o *RftemplateRadioBand6) HasStandardPower() bool`

HasStandardPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


