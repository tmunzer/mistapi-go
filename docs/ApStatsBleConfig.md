# ApStatsBleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeaconRate** | Pointer to **int32** |  | [optional] 
**BeaconRateModel** | Pointer to **string** |  | [optional] 
**BeamDisabled** | Pointer to **[]int32** |  | [optional] [default to []]
**Power** | Pointer to **int32** |  | [optional] 
**PowerMode** | Pointer to **string** |  | [optional] 

## Methods

### NewApStatsBleConfig

`func NewApStatsBleConfig() *ApStatsBleConfig`

NewApStatsBleConfig instantiates a new ApStatsBleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApStatsBleConfigWithDefaults

`func NewApStatsBleConfigWithDefaults() *ApStatsBleConfig`

NewApStatsBleConfigWithDefaults instantiates a new ApStatsBleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeaconRate

`func (o *ApStatsBleConfig) GetBeaconRate() int32`

GetBeaconRate returns the BeaconRate field if non-nil, zero value otherwise.

### GetBeaconRateOk

`func (o *ApStatsBleConfig) GetBeaconRateOk() (*int32, bool)`

GetBeaconRateOk returns a tuple with the BeaconRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconRate

`func (o *ApStatsBleConfig) SetBeaconRate(v int32)`

SetBeaconRate sets BeaconRate field to given value.

### HasBeaconRate

`func (o *ApStatsBleConfig) HasBeaconRate() bool`

HasBeaconRate returns a boolean if a field has been set.

### GetBeaconRateModel

`func (o *ApStatsBleConfig) GetBeaconRateModel() string`

GetBeaconRateModel returns the BeaconRateModel field if non-nil, zero value otherwise.

### GetBeaconRateModelOk

`func (o *ApStatsBleConfig) GetBeaconRateModelOk() (*string, bool)`

GetBeaconRateModelOk returns a tuple with the BeaconRateModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconRateModel

`func (o *ApStatsBleConfig) SetBeaconRateModel(v string)`

SetBeaconRateModel sets BeaconRateModel field to given value.

### HasBeaconRateModel

`func (o *ApStatsBleConfig) HasBeaconRateModel() bool`

HasBeaconRateModel returns a boolean if a field has been set.

### GetBeamDisabled

`func (o *ApStatsBleConfig) GetBeamDisabled() []int32`

GetBeamDisabled returns the BeamDisabled field if non-nil, zero value otherwise.

### GetBeamDisabledOk

`func (o *ApStatsBleConfig) GetBeamDisabledOk() (*[]int32, bool)`

GetBeamDisabledOk returns a tuple with the BeamDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeamDisabled

`func (o *ApStatsBleConfig) SetBeamDisabled(v []int32)`

SetBeamDisabled sets BeamDisabled field to given value.

### HasBeamDisabled

`func (o *ApStatsBleConfig) HasBeamDisabled() bool`

HasBeamDisabled returns a boolean if a field has been set.

### GetPower

`func (o *ApStatsBleConfig) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *ApStatsBleConfig) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *ApStatsBleConfig) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *ApStatsBleConfig) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPowerMode

`func (o *ApStatsBleConfig) GetPowerMode() string`

GetPowerMode returns the PowerMode field if non-nil, zero value otherwise.

### GetPowerModeOk

`func (o *ApStatsBleConfig) GetPowerModeOk() (*string, bool)`

GetPowerModeOk returns a tuple with the PowerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerMode

`func (o *ApStatsBleConfig) SetPowerMode(v string)`

SetPowerMode sets PowerMode field to given value.

### HasPowerMode

`func (o *ApStatsBleConfig) HasPowerMode() bool`

HasPowerMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


