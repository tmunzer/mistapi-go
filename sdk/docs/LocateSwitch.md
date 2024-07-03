# LocateSwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | minutes the leds should keep flashing | [optional] [default to 5]
**Mac** | Pointer to **string** | for virtual chassis, the MAC of the member | [optional] 

## Methods

### NewLocateSwitch

`func NewLocateSwitch() *LocateSwitch`

NewLocateSwitch instantiates a new LocateSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocateSwitchWithDefaults

`func NewLocateSwitchWithDefaults() *LocateSwitch`

NewLocateSwitchWithDefaults instantiates a new LocateSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *LocateSwitch) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *LocateSwitch) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *LocateSwitch) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *LocateSwitch) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetMac

`func (o *LocateSwitch) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *LocateSwitch) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *LocateSwitch) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *LocateSwitch) HasMac() bool`

HasMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


