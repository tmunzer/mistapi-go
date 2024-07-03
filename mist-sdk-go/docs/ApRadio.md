# ApRadio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRrmDisable** | Pointer to **bool** |  | [optional] [default to false]
**AntGain24** | Pointer to **int32** | antenna gain for 2.4G - for models with external antenna only | [optional] 
**AntGain5** | Pointer to **int32** | antenna gain for 5G - for models with external antenna only | [optional] 
**AntGain6** | Pointer to **int32** | antenna gain for 6G - for models with external antenna only | [optional] 
**AntennaMode** | Pointer to [**ApRadioAntennaMode**](ApRadioAntennaMode.md) |  | [optional] [default to APRADIOANTENNAMODE_DEFAULT]
**Band24** | Pointer to [**ApRadioBand24**](ApRadioBand24.md) |  | [optional] 
**Band24Usage** | Pointer to [**RadioBand24Usage**](RadioBand24Usage.md) |  | [optional] 
**Band5** | Pointer to [**ApRadioBand5**](ApRadioBand5.md) |  | [optional] 
**Band5On24Radio** | Pointer to [**ApRadioBand5**](ApRadioBand5.md) |  | [optional] 
**Band6** | Pointer to [**ApRadioBand6**](ApRadioBand6.md) |  | [optional] 
**IndoorUse** | Pointer to **bool** | to make an outdoor operate indoor. for an outdoor-ap, some channels are disallowed by default, this allows the user to use it as an indoor-ap | [optional] [default to false]
**ScanningEnabled** | Pointer to **bool** | whether scanning radio is enabled | [optional] 

## Methods

### NewApRadio

`func NewApRadio() *ApRadio`

NewApRadio instantiates a new ApRadio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApRadioWithDefaults

`func NewApRadioWithDefaults() *ApRadio`

NewApRadioWithDefaults instantiates a new ApRadio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowRrmDisable

`func (o *ApRadio) GetAllowRrmDisable() bool`

GetAllowRrmDisable returns the AllowRrmDisable field if non-nil, zero value otherwise.

### GetAllowRrmDisableOk

`func (o *ApRadio) GetAllowRrmDisableOk() (*bool, bool)`

GetAllowRrmDisableOk returns a tuple with the AllowRrmDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRrmDisable

`func (o *ApRadio) SetAllowRrmDisable(v bool)`

SetAllowRrmDisable sets AllowRrmDisable field to given value.

### HasAllowRrmDisable

`func (o *ApRadio) HasAllowRrmDisable() bool`

HasAllowRrmDisable returns a boolean if a field has been set.

### GetAntGain24

`func (o *ApRadio) GetAntGain24() int32`

GetAntGain24 returns the AntGain24 field if non-nil, zero value otherwise.

### GetAntGain24Ok

`func (o *ApRadio) GetAntGain24Ok() (*int32, bool)`

GetAntGain24Ok returns a tuple with the AntGain24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntGain24

`func (o *ApRadio) SetAntGain24(v int32)`

SetAntGain24 sets AntGain24 field to given value.

### HasAntGain24

`func (o *ApRadio) HasAntGain24() bool`

HasAntGain24 returns a boolean if a field has been set.

### GetAntGain5

`func (o *ApRadio) GetAntGain5() int32`

GetAntGain5 returns the AntGain5 field if non-nil, zero value otherwise.

### GetAntGain5Ok

`func (o *ApRadio) GetAntGain5Ok() (*int32, bool)`

GetAntGain5Ok returns a tuple with the AntGain5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntGain5

`func (o *ApRadio) SetAntGain5(v int32)`

SetAntGain5 sets AntGain5 field to given value.

### HasAntGain5

`func (o *ApRadio) HasAntGain5() bool`

HasAntGain5 returns a boolean if a field has been set.

### GetAntGain6

`func (o *ApRadio) GetAntGain6() int32`

GetAntGain6 returns the AntGain6 field if non-nil, zero value otherwise.

### GetAntGain6Ok

`func (o *ApRadio) GetAntGain6Ok() (*int32, bool)`

GetAntGain6Ok returns a tuple with the AntGain6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntGain6

`func (o *ApRadio) SetAntGain6(v int32)`

SetAntGain6 sets AntGain6 field to given value.

### HasAntGain6

`func (o *ApRadio) HasAntGain6() bool`

HasAntGain6 returns a boolean if a field has been set.

### GetAntennaMode

`func (o *ApRadio) GetAntennaMode() ApRadioAntennaMode`

GetAntennaMode returns the AntennaMode field if non-nil, zero value otherwise.

### GetAntennaModeOk

`func (o *ApRadio) GetAntennaModeOk() (*ApRadioAntennaMode, bool)`

GetAntennaModeOk returns a tuple with the AntennaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntennaMode

`func (o *ApRadio) SetAntennaMode(v ApRadioAntennaMode)`

SetAntennaMode sets AntennaMode field to given value.

### HasAntennaMode

`func (o *ApRadio) HasAntennaMode() bool`

HasAntennaMode returns a boolean if a field has been set.

### GetBand24

`func (o *ApRadio) GetBand24() ApRadioBand24`

GetBand24 returns the Band24 field if non-nil, zero value otherwise.

### GetBand24Ok

`func (o *ApRadio) GetBand24Ok() (*ApRadioBand24, bool)`

GetBand24Ok returns a tuple with the Band24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24

`func (o *ApRadio) SetBand24(v ApRadioBand24)`

SetBand24 sets Band24 field to given value.

### HasBand24

`func (o *ApRadio) HasBand24() bool`

HasBand24 returns a boolean if a field has been set.

### GetBand24Usage

`func (o *ApRadio) GetBand24Usage() RadioBand24Usage`

GetBand24Usage returns the Band24Usage field if non-nil, zero value otherwise.

### GetBand24UsageOk

`func (o *ApRadio) GetBand24UsageOk() (*RadioBand24Usage, bool)`

GetBand24UsageOk returns a tuple with the Band24Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Usage

`func (o *ApRadio) SetBand24Usage(v RadioBand24Usage)`

SetBand24Usage sets Band24Usage field to given value.

### HasBand24Usage

`func (o *ApRadio) HasBand24Usage() bool`

HasBand24Usage returns a boolean if a field has been set.

### GetBand5

`func (o *ApRadio) GetBand5() ApRadioBand5`

GetBand5 returns the Band5 field if non-nil, zero value otherwise.

### GetBand5Ok

`func (o *ApRadio) GetBand5Ok() (*ApRadioBand5, bool)`

GetBand5Ok returns a tuple with the Band5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5

`func (o *ApRadio) SetBand5(v ApRadioBand5)`

SetBand5 sets Band5 field to given value.

### HasBand5

`func (o *ApRadio) HasBand5() bool`

HasBand5 returns a boolean if a field has been set.

### GetBand5On24Radio

`func (o *ApRadio) GetBand5On24Radio() ApRadioBand5`

GetBand5On24Radio returns the Band5On24Radio field if non-nil, zero value otherwise.

### GetBand5On24RadioOk

`func (o *ApRadio) GetBand5On24RadioOk() (*ApRadioBand5, bool)`

GetBand5On24RadioOk returns a tuple with the Band5On24Radio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5On24Radio

`func (o *ApRadio) SetBand5On24Radio(v ApRadioBand5)`

SetBand5On24Radio sets Band5On24Radio field to given value.

### HasBand5On24Radio

`func (o *ApRadio) HasBand5On24Radio() bool`

HasBand5On24Radio returns a boolean if a field has been set.

### GetBand6

`func (o *ApRadio) GetBand6() ApRadioBand6`

GetBand6 returns the Band6 field if non-nil, zero value otherwise.

### GetBand6Ok

`func (o *ApRadio) GetBand6Ok() (*ApRadioBand6, bool)`

GetBand6Ok returns a tuple with the Band6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6

`func (o *ApRadio) SetBand6(v ApRadioBand6)`

SetBand6 sets Band6 field to given value.

### HasBand6

`func (o *ApRadio) HasBand6() bool`

HasBand6 returns a boolean if a field has been set.

### GetIndoorUse

`func (o *ApRadio) GetIndoorUse() bool`

GetIndoorUse returns the IndoorUse field if non-nil, zero value otherwise.

### GetIndoorUseOk

`func (o *ApRadio) GetIndoorUseOk() (*bool, bool)`

GetIndoorUseOk returns a tuple with the IndoorUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorUse

`func (o *ApRadio) SetIndoorUse(v bool)`

SetIndoorUse sets IndoorUse field to given value.

### HasIndoorUse

`func (o *ApRadio) HasIndoorUse() bool`

HasIndoorUse returns a boolean if a field has been set.

### GetScanningEnabled

`func (o *ApRadio) GetScanningEnabled() bool`

GetScanningEnabled returns the ScanningEnabled field if non-nil, zero value otherwise.

### GetScanningEnabledOk

`func (o *ApRadio) GetScanningEnabledOk() (*bool, bool)`

GetScanningEnabledOk returns a tuple with the ScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanningEnabled

`func (o *ApRadio) SetScanningEnabled(v bool)`

SetScanningEnabled sets ScanningEnabled field to given value.

### HasScanningEnabled

`func (o *ApRadio) HasScanningEnabled() bool`

HasScanningEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


