# ApIotOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | whether to enable a pin | [optional] [default to false]
**Name** | Pointer to **string** | optional; descriptive pin name | [optional] 
**Output** | Pointer to **bool** | whether the pin is configured as an output. DO and A1-A4 can be repurposed by changing | [optional] 
**Pullup** | Pointer to [**ApIotOutputPullup**](ApIotOutputPullup.md) |  | [optional] 
**Value** | Pointer to **int32** | output pin signal level, default 0 | [optional] 

## Methods

### NewApIotOutput

`func NewApIotOutput() *ApIotOutput`

NewApIotOutput instantiates a new ApIotOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApIotOutputWithDefaults

`func NewApIotOutputWithDefaults() *ApIotOutput`

NewApIotOutputWithDefaults instantiates a new ApIotOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApIotOutput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApIotOutput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApIotOutput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApIotOutput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *ApIotOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApIotOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApIotOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApIotOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *ApIotOutput) GetOutput() bool`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ApIotOutput) GetOutputOk() (*bool, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ApIotOutput) SetOutput(v bool)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ApIotOutput) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetPullup

`func (o *ApIotOutput) GetPullup() ApIotOutputPullup`

GetPullup returns the Pullup field if non-nil, zero value otherwise.

### GetPullupOk

`func (o *ApIotOutput) GetPullupOk() (*ApIotOutputPullup, bool)`

GetPullupOk returns a tuple with the Pullup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullup

`func (o *ApIotOutput) SetPullup(v ApIotOutputPullup)`

SetPullup sets Pullup field to given value.

### HasPullup

`func (o *ApIotOutput) HasPullup() bool`

HasPullup returns a boolean if a field has been set.

### GetValue

`func (o *ApIotOutput) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApIotOutput) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApIotOutput) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApIotOutput) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


