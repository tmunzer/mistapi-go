# BleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeaconEnabled** | Pointer to **bool** | whether Mist beacons is enabled | [optional] [default to false]
**BeaconRate** | Pointer to **int32** | required if &#x60;beacon_rate_mode&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, 1-10, in number-beacons-per-second | [optional] [default to 0]
**BeaconRateMode** | Pointer to [**BleConfigBeaconRateMode**](BleConfigBeaconRateMode.md) |  | [optional] [default to BLECONFIGBEACONRATEMODE_DEFAULT]
**BeamDisabled** | Pointer to **[]int32** | list of AP BLE location beam numbers (1-8) which should be disabled at the AP and not transmit location information (where beam 1 is oriented at the top the AP, growing counter-clock-wise, with 9 being the omni BLE beam) | [optional] 
**CustomBlePacketEnabled** | Pointer to **bool** | can be enabled if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;true&#x60;, whether to send custom packet | [optional] [default to false]
**CustomBlePacketFrame** | Pointer to **string** | The custom frame to be sent out in this beacon. The frame must be a hexstring | [optional] [default to ""]
**CustomBlePacketFreqMsec** | Pointer to **int32** | Frequency (msec) of data emitted by custom ble beacon | [optional] [default to 0]
**EddystoneUidAdvPower** | Pointer to **int32** | advertised TX Power, -100 to 20 (dBm), omit this attribute to use default | [optional] [default to 0]
**EddystoneUidBeams** | Pointer to **string** |  | [optional] [default to ""]
**EddystoneUidEnabled** | Pointer to **bool** | only if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;false&#x60;, Whether Eddystone-UID beacon is enabled | [optional] [default to false]
**EddystoneUidFreqMsec** | Pointer to **int32** | Frequency (msec) of data emmit by Eddystone-UID beacon | [optional] [default to 0]
**EddystoneUidInstance** | Pointer to **string** | Eddystone-UID instance for the device | [optional] [default to ""]
**EddystoneUidNamespace** | Pointer to **string** | Eddystone-UID namespace | [optional] [default to ""]
**EddystoneUrlAdvPower** | Pointer to **int32** | advertised TX Power, -100 to 20 (dBm), omit this attribute to use default | [optional] [default to 0]
**EddystoneUrlBeams** | Pointer to **string** |  | [optional] [default to ""]
**EddystoneUrlEnabled** | Pointer to **bool** | only if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;false&#x60;, Whether Eddystone-URL beacon is enabled | [optional] [default to false]
**EddystoneUrlFreqMsec** | Pointer to **int32** | Frequency (msec) of data emit by Eddystone-UID beacon | [optional] [default to 0]
**EddystoneUrlUrl** | Pointer to **string** | URL pointed by Eddystone-URL beacon | [optional] [default to ""]
**IbeaconAdvPower** | Pointer to **int32** | advertised TX Power, -100 to 20 (dBm), omit this attribute to use default | [optional] [default to 0]
**IbeaconBeams** | Pointer to **string** |  | [optional] [default to ""]
**IbeaconEnabled** | Pointer to **bool** | can be enabled if &#x60;beacon_enabled&#x60;&#x3D;&#x3D;&#x60;true&#x60;, whether to send iBeacon | [optional] [default to false]
**IbeaconFreqMsec** | Pointer to **int32** | Frequency (msec) of data emmit for iBeacon | [optional] [default to 0]
**IbeaconMajor** | Pointer to **int32** | Major number for iBeacon | [optional] [default to 0]
**IbeaconMinor** | Pointer to **int32** | Minor number for iBeacon | [optional] [default to 0]
**IbeaconUuid** | Pointer to **string** | optional, if not specified, the same UUID as the beacon will be used | [optional] 
**Power** | Pointer to **int32** | required if &#x60;power_mode&#x60;&#x3D;&#x3D;&#x60;custom&#x60; | [optional] [default to 9]
**PowerMode** | Pointer to [**BleConfigPowerMode**](BleConfigPowerMode.md) |  | [optional] [default to BLECONFIGPOWERMODE_DEFAULT]

## Methods

### NewBleConfig

`func NewBleConfig() *BleConfig`

NewBleConfig instantiates a new BleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBleConfigWithDefaults

`func NewBleConfigWithDefaults() *BleConfig`

NewBleConfigWithDefaults instantiates a new BleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeaconEnabled

`func (o *BleConfig) GetBeaconEnabled() bool`

GetBeaconEnabled returns the BeaconEnabled field if non-nil, zero value otherwise.

### GetBeaconEnabledOk

`func (o *BleConfig) GetBeaconEnabledOk() (*bool, bool)`

GetBeaconEnabledOk returns a tuple with the BeaconEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconEnabled

`func (o *BleConfig) SetBeaconEnabled(v bool)`

SetBeaconEnabled sets BeaconEnabled field to given value.

### HasBeaconEnabled

`func (o *BleConfig) HasBeaconEnabled() bool`

HasBeaconEnabled returns a boolean if a field has been set.

### GetBeaconRate

`func (o *BleConfig) GetBeaconRate() int32`

GetBeaconRate returns the BeaconRate field if non-nil, zero value otherwise.

### GetBeaconRateOk

`func (o *BleConfig) GetBeaconRateOk() (*int32, bool)`

GetBeaconRateOk returns a tuple with the BeaconRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconRate

`func (o *BleConfig) SetBeaconRate(v int32)`

SetBeaconRate sets BeaconRate field to given value.

### HasBeaconRate

`func (o *BleConfig) HasBeaconRate() bool`

HasBeaconRate returns a boolean if a field has been set.

### GetBeaconRateMode

`func (o *BleConfig) GetBeaconRateMode() BleConfigBeaconRateMode`

GetBeaconRateMode returns the BeaconRateMode field if non-nil, zero value otherwise.

### GetBeaconRateModeOk

`func (o *BleConfig) GetBeaconRateModeOk() (*BleConfigBeaconRateMode, bool)`

GetBeaconRateModeOk returns a tuple with the BeaconRateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconRateMode

`func (o *BleConfig) SetBeaconRateMode(v BleConfigBeaconRateMode)`

SetBeaconRateMode sets BeaconRateMode field to given value.

### HasBeaconRateMode

`func (o *BleConfig) HasBeaconRateMode() bool`

HasBeaconRateMode returns a boolean if a field has been set.

### GetBeamDisabled

`func (o *BleConfig) GetBeamDisabled() []int32`

GetBeamDisabled returns the BeamDisabled field if non-nil, zero value otherwise.

### GetBeamDisabledOk

`func (o *BleConfig) GetBeamDisabledOk() (*[]int32, bool)`

GetBeamDisabledOk returns a tuple with the BeamDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeamDisabled

`func (o *BleConfig) SetBeamDisabled(v []int32)`

SetBeamDisabled sets BeamDisabled field to given value.

### HasBeamDisabled

`func (o *BleConfig) HasBeamDisabled() bool`

HasBeamDisabled returns a boolean if a field has been set.

### GetCustomBlePacketEnabled

`func (o *BleConfig) GetCustomBlePacketEnabled() bool`

GetCustomBlePacketEnabled returns the CustomBlePacketEnabled field if non-nil, zero value otherwise.

### GetCustomBlePacketEnabledOk

`func (o *BleConfig) GetCustomBlePacketEnabledOk() (*bool, bool)`

GetCustomBlePacketEnabledOk returns a tuple with the CustomBlePacketEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBlePacketEnabled

`func (o *BleConfig) SetCustomBlePacketEnabled(v bool)`

SetCustomBlePacketEnabled sets CustomBlePacketEnabled field to given value.

### HasCustomBlePacketEnabled

`func (o *BleConfig) HasCustomBlePacketEnabled() bool`

HasCustomBlePacketEnabled returns a boolean if a field has been set.

### GetCustomBlePacketFrame

`func (o *BleConfig) GetCustomBlePacketFrame() string`

GetCustomBlePacketFrame returns the CustomBlePacketFrame field if non-nil, zero value otherwise.

### GetCustomBlePacketFrameOk

`func (o *BleConfig) GetCustomBlePacketFrameOk() (*string, bool)`

GetCustomBlePacketFrameOk returns a tuple with the CustomBlePacketFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBlePacketFrame

`func (o *BleConfig) SetCustomBlePacketFrame(v string)`

SetCustomBlePacketFrame sets CustomBlePacketFrame field to given value.

### HasCustomBlePacketFrame

`func (o *BleConfig) HasCustomBlePacketFrame() bool`

HasCustomBlePacketFrame returns a boolean if a field has been set.

### GetCustomBlePacketFreqMsec

`func (o *BleConfig) GetCustomBlePacketFreqMsec() int32`

GetCustomBlePacketFreqMsec returns the CustomBlePacketFreqMsec field if non-nil, zero value otherwise.

### GetCustomBlePacketFreqMsecOk

`func (o *BleConfig) GetCustomBlePacketFreqMsecOk() (*int32, bool)`

GetCustomBlePacketFreqMsecOk returns a tuple with the CustomBlePacketFreqMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBlePacketFreqMsec

`func (o *BleConfig) SetCustomBlePacketFreqMsec(v int32)`

SetCustomBlePacketFreqMsec sets CustomBlePacketFreqMsec field to given value.

### HasCustomBlePacketFreqMsec

`func (o *BleConfig) HasCustomBlePacketFreqMsec() bool`

HasCustomBlePacketFreqMsec returns a boolean if a field has been set.

### GetEddystoneUidAdvPower

`func (o *BleConfig) GetEddystoneUidAdvPower() int32`

GetEddystoneUidAdvPower returns the EddystoneUidAdvPower field if non-nil, zero value otherwise.

### GetEddystoneUidAdvPowerOk

`func (o *BleConfig) GetEddystoneUidAdvPowerOk() (*int32, bool)`

GetEddystoneUidAdvPowerOk returns a tuple with the EddystoneUidAdvPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidAdvPower

`func (o *BleConfig) SetEddystoneUidAdvPower(v int32)`

SetEddystoneUidAdvPower sets EddystoneUidAdvPower field to given value.

### HasEddystoneUidAdvPower

`func (o *BleConfig) HasEddystoneUidAdvPower() bool`

HasEddystoneUidAdvPower returns a boolean if a field has been set.

### GetEddystoneUidBeams

`func (o *BleConfig) GetEddystoneUidBeams() string`

GetEddystoneUidBeams returns the EddystoneUidBeams field if non-nil, zero value otherwise.

### GetEddystoneUidBeamsOk

`func (o *BleConfig) GetEddystoneUidBeamsOk() (*string, bool)`

GetEddystoneUidBeamsOk returns a tuple with the EddystoneUidBeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidBeams

`func (o *BleConfig) SetEddystoneUidBeams(v string)`

SetEddystoneUidBeams sets EddystoneUidBeams field to given value.

### HasEddystoneUidBeams

`func (o *BleConfig) HasEddystoneUidBeams() bool`

HasEddystoneUidBeams returns a boolean if a field has been set.

### GetEddystoneUidEnabled

`func (o *BleConfig) GetEddystoneUidEnabled() bool`

GetEddystoneUidEnabled returns the EddystoneUidEnabled field if non-nil, zero value otherwise.

### GetEddystoneUidEnabledOk

`func (o *BleConfig) GetEddystoneUidEnabledOk() (*bool, bool)`

GetEddystoneUidEnabledOk returns a tuple with the EddystoneUidEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidEnabled

`func (o *BleConfig) SetEddystoneUidEnabled(v bool)`

SetEddystoneUidEnabled sets EddystoneUidEnabled field to given value.

### HasEddystoneUidEnabled

`func (o *BleConfig) HasEddystoneUidEnabled() bool`

HasEddystoneUidEnabled returns a boolean if a field has been set.

### GetEddystoneUidFreqMsec

`func (o *BleConfig) GetEddystoneUidFreqMsec() int32`

GetEddystoneUidFreqMsec returns the EddystoneUidFreqMsec field if non-nil, zero value otherwise.

### GetEddystoneUidFreqMsecOk

`func (o *BleConfig) GetEddystoneUidFreqMsecOk() (*int32, bool)`

GetEddystoneUidFreqMsecOk returns a tuple with the EddystoneUidFreqMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidFreqMsec

`func (o *BleConfig) SetEddystoneUidFreqMsec(v int32)`

SetEddystoneUidFreqMsec sets EddystoneUidFreqMsec field to given value.

### HasEddystoneUidFreqMsec

`func (o *BleConfig) HasEddystoneUidFreqMsec() bool`

HasEddystoneUidFreqMsec returns a boolean if a field has been set.

### GetEddystoneUidInstance

`func (o *BleConfig) GetEddystoneUidInstance() string`

GetEddystoneUidInstance returns the EddystoneUidInstance field if non-nil, zero value otherwise.

### GetEddystoneUidInstanceOk

`func (o *BleConfig) GetEddystoneUidInstanceOk() (*string, bool)`

GetEddystoneUidInstanceOk returns a tuple with the EddystoneUidInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidInstance

`func (o *BleConfig) SetEddystoneUidInstance(v string)`

SetEddystoneUidInstance sets EddystoneUidInstance field to given value.

### HasEddystoneUidInstance

`func (o *BleConfig) HasEddystoneUidInstance() bool`

HasEddystoneUidInstance returns a boolean if a field has been set.

### GetEddystoneUidNamespace

`func (o *BleConfig) GetEddystoneUidNamespace() string`

GetEddystoneUidNamespace returns the EddystoneUidNamespace field if non-nil, zero value otherwise.

### GetEddystoneUidNamespaceOk

`func (o *BleConfig) GetEddystoneUidNamespaceOk() (*string, bool)`

GetEddystoneUidNamespaceOk returns a tuple with the EddystoneUidNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidNamespace

`func (o *BleConfig) SetEddystoneUidNamespace(v string)`

SetEddystoneUidNamespace sets EddystoneUidNamespace field to given value.

### HasEddystoneUidNamespace

`func (o *BleConfig) HasEddystoneUidNamespace() bool`

HasEddystoneUidNamespace returns a boolean if a field has been set.

### GetEddystoneUrlAdvPower

`func (o *BleConfig) GetEddystoneUrlAdvPower() int32`

GetEddystoneUrlAdvPower returns the EddystoneUrlAdvPower field if non-nil, zero value otherwise.

### GetEddystoneUrlAdvPowerOk

`func (o *BleConfig) GetEddystoneUrlAdvPowerOk() (*int32, bool)`

GetEddystoneUrlAdvPowerOk returns a tuple with the EddystoneUrlAdvPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlAdvPower

`func (o *BleConfig) SetEddystoneUrlAdvPower(v int32)`

SetEddystoneUrlAdvPower sets EddystoneUrlAdvPower field to given value.

### HasEddystoneUrlAdvPower

`func (o *BleConfig) HasEddystoneUrlAdvPower() bool`

HasEddystoneUrlAdvPower returns a boolean if a field has been set.

### GetEddystoneUrlBeams

`func (o *BleConfig) GetEddystoneUrlBeams() string`

GetEddystoneUrlBeams returns the EddystoneUrlBeams field if non-nil, zero value otherwise.

### GetEddystoneUrlBeamsOk

`func (o *BleConfig) GetEddystoneUrlBeamsOk() (*string, bool)`

GetEddystoneUrlBeamsOk returns a tuple with the EddystoneUrlBeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlBeams

`func (o *BleConfig) SetEddystoneUrlBeams(v string)`

SetEddystoneUrlBeams sets EddystoneUrlBeams field to given value.

### HasEddystoneUrlBeams

`func (o *BleConfig) HasEddystoneUrlBeams() bool`

HasEddystoneUrlBeams returns a boolean if a field has been set.

### GetEddystoneUrlEnabled

`func (o *BleConfig) GetEddystoneUrlEnabled() bool`

GetEddystoneUrlEnabled returns the EddystoneUrlEnabled field if non-nil, zero value otherwise.

### GetEddystoneUrlEnabledOk

`func (o *BleConfig) GetEddystoneUrlEnabledOk() (*bool, bool)`

GetEddystoneUrlEnabledOk returns a tuple with the EddystoneUrlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlEnabled

`func (o *BleConfig) SetEddystoneUrlEnabled(v bool)`

SetEddystoneUrlEnabled sets EddystoneUrlEnabled field to given value.

### HasEddystoneUrlEnabled

`func (o *BleConfig) HasEddystoneUrlEnabled() bool`

HasEddystoneUrlEnabled returns a boolean if a field has been set.

### GetEddystoneUrlFreqMsec

`func (o *BleConfig) GetEddystoneUrlFreqMsec() int32`

GetEddystoneUrlFreqMsec returns the EddystoneUrlFreqMsec field if non-nil, zero value otherwise.

### GetEddystoneUrlFreqMsecOk

`func (o *BleConfig) GetEddystoneUrlFreqMsecOk() (*int32, bool)`

GetEddystoneUrlFreqMsecOk returns a tuple with the EddystoneUrlFreqMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlFreqMsec

`func (o *BleConfig) SetEddystoneUrlFreqMsec(v int32)`

SetEddystoneUrlFreqMsec sets EddystoneUrlFreqMsec field to given value.

### HasEddystoneUrlFreqMsec

`func (o *BleConfig) HasEddystoneUrlFreqMsec() bool`

HasEddystoneUrlFreqMsec returns a boolean if a field has been set.

### GetEddystoneUrlUrl

`func (o *BleConfig) GetEddystoneUrlUrl() string`

GetEddystoneUrlUrl returns the EddystoneUrlUrl field if non-nil, zero value otherwise.

### GetEddystoneUrlUrlOk

`func (o *BleConfig) GetEddystoneUrlUrlOk() (*string, bool)`

GetEddystoneUrlUrlOk returns a tuple with the EddystoneUrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlUrl

`func (o *BleConfig) SetEddystoneUrlUrl(v string)`

SetEddystoneUrlUrl sets EddystoneUrlUrl field to given value.

### HasEddystoneUrlUrl

`func (o *BleConfig) HasEddystoneUrlUrl() bool`

HasEddystoneUrlUrl returns a boolean if a field has been set.

### GetIbeaconAdvPower

`func (o *BleConfig) GetIbeaconAdvPower() int32`

GetIbeaconAdvPower returns the IbeaconAdvPower field if non-nil, zero value otherwise.

### GetIbeaconAdvPowerOk

`func (o *BleConfig) GetIbeaconAdvPowerOk() (*int32, bool)`

GetIbeaconAdvPowerOk returns a tuple with the IbeaconAdvPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconAdvPower

`func (o *BleConfig) SetIbeaconAdvPower(v int32)`

SetIbeaconAdvPower sets IbeaconAdvPower field to given value.

### HasIbeaconAdvPower

`func (o *BleConfig) HasIbeaconAdvPower() bool`

HasIbeaconAdvPower returns a boolean if a field has been set.

### GetIbeaconBeams

`func (o *BleConfig) GetIbeaconBeams() string`

GetIbeaconBeams returns the IbeaconBeams field if non-nil, zero value otherwise.

### GetIbeaconBeamsOk

`func (o *BleConfig) GetIbeaconBeamsOk() (*string, bool)`

GetIbeaconBeamsOk returns a tuple with the IbeaconBeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconBeams

`func (o *BleConfig) SetIbeaconBeams(v string)`

SetIbeaconBeams sets IbeaconBeams field to given value.

### HasIbeaconBeams

`func (o *BleConfig) HasIbeaconBeams() bool`

HasIbeaconBeams returns a boolean if a field has been set.

### GetIbeaconEnabled

`func (o *BleConfig) GetIbeaconEnabled() bool`

GetIbeaconEnabled returns the IbeaconEnabled field if non-nil, zero value otherwise.

### GetIbeaconEnabledOk

`func (o *BleConfig) GetIbeaconEnabledOk() (*bool, bool)`

GetIbeaconEnabledOk returns a tuple with the IbeaconEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconEnabled

`func (o *BleConfig) SetIbeaconEnabled(v bool)`

SetIbeaconEnabled sets IbeaconEnabled field to given value.

### HasIbeaconEnabled

`func (o *BleConfig) HasIbeaconEnabled() bool`

HasIbeaconEnabled returns a boolean if a field has been set.

### GetIbeaconFreqMsec

`func (o *BleConfig) GetIbeaconFreqMsec() int32`

GetIbeaconFreqMsec returns the IbeaconFreqMsec field if non-nil, zero value otherwise.

### GetIbeaconFreqMsecOk

`func (o *BleConfig) GetIbeaconFreqMsecOk() (*int32, bool)`

GetIbeaconFreqMsecOk returns a tuple with the IbeaconFreqMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconFreqMsec

`func (o *BleConfig) SetIbeaconFreqMsec(v int32)`

SetIbeaconFreqMsec sets IbeaconFreqMsec field to given value.

### HasIbeaconFreqMsec

`func (o *BleConfig) HasIbeaconFreqMsec() bool`

HasIbeaconFreqMsec returns a boolean if a field has been set.

### GetIbeaconMajor

`func (o *BleConfig) GetIbeaconMajor() int32`

GetIbeaconMajor returns the IbeaconMajor field if non-nil, zero value otherwise.

### GetIbeaconMajorOk

`func (o *BleConfig) GetIbeaconMajorOk() (*int32, bool)`

GetIbeaconMajorOk returns a tuple with the IbeaconMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMajor

`func (o *BleConfig) SetIbeaconMajor(v int32)`

SetIbeaconMajor sets IbeaconMajor field to given value.

### HasIbeaconMajor

`func (o *BleConfig) HasIbeaconMajor() bool`

HasIbeaconMajor returns a boolean if a field has been set.

### GetIbeaconMinor

`func (o *BleConfig) GetIbeaconMinor() int32`

GetIbeaconMinor returns the IbeaconMinor field if non-nil, zero value otherwise.

### GetIbeaconMinorOk

`func (o *BleConfig) GetIbeaconMinorOk() (*int32, bool)`

GetIbeaconMinorOk returns a tuple with the IbeaconMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMinor

`func (o *BleConfig) SetIbeaconMinor(v int32)`

SetIbeaconMinor sets IbeaconMinor field to given value.

### HasIbeaconMinor

`func (o *BleConfig) HasIbeaconMinor() bool`

HasIbeaconMinor returns a boolean if a field has been set.

### GetIbeaconUuid

`func (o *BleConfig) GetIbeaconUuid() string`

GetIbeaconUuid returns the IbeaconUuid field if non-nil, zero value otherwise.

### GetIbeaconUuidOk

`func (o *BleConfig) GetIbeaconUuidOk() (*string, bool)`

GetIbeaconUuidOk returns a tuple with the IbeaconUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconUuid

`func (o *BleConfig) SetIbeaconUuid(v string)`

SetIbeaconUuid sets IbeaconUuid field to given value.

### HasIbeaconUuid

`func (o *BleConfig) HasIbeaconUuid() bool`

HasIbeaconUuid returns a boolean if a field has been set.

### GetPower

`func (o *BleConfig) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *BleConfig) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *BleConfig) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *BleConfig) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPowerMode

`func (o *BleConfig) GetPowerMode() BleConfigPowerMode`

GetPowerMode returns the PowerMode field if non-nil, zero value otherwise.

### GetPowerModeOk

`func (o *BleConfig) GetPowerModeOk() (*BleConfigPowerMode, bool)`

GetPowerModeOk returns a tuple with the PowerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMode

`func (o *BleConfig) SetPowerMode(v BleConfigPowerMode)`

SetPowerMode sets PowerMode field to given value.

### HasPowerMode

`func (o *BleConfig) HasPowerMode() bool`

HasPowerMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


