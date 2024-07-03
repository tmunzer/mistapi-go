# ApLed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brightness** | Pointer to **int32** |  | [optional] [default to 255]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewApLed

`func NewApLed() *ApLed`

NewApLed instantiates a new ApLed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApLedWithDefaults

`func NewApLedWithDefaults() *ApLed`

NewApLedWithDefaults instantiates a new ApLed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrightness

`func (o *ApLed) GetBrightness() int32`

GetBrightness returns the Brightness field if non-nil, zero value otherwise.

### GetBrightnessOk

`func (o *ApLed) GetBrightnessOk() (*int32, bool)`

GetBrightnessOk returns a tuple with the Brightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrightness

`func (o *ApLed) SetBrightness(v int32)`

SetBrightness sets Brightness field to given value.

### HasBrightness

`func (o *ApLed) HasBrightness() bool`

HasBrightness returns a boolean if a field has been set.

### GetEnabled

`func (o *ApLed) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApLed) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApLed) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApLed) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


