# ApIotInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | whether to enable a pin | [optional] [default to false]
**Name** | Pointer to **string** | optional; descriptive pin name | [optional] 
**Pullup** | Pointer to [**ApIotInputPullup**](ApIotInputPullup.md) |  | [optional] 

## Methods

### NewApIotInput

`func NewApIotInput() *ApIotInput`

NewApIotInput instantiates a new ApIotInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApIotInputWithDefaults

`func NewApIotInputWithDefaults() *ApIotInput`

NewApIotInputWithDefaults instantiates a new ApIotInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApIotInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApIotInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApIotInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApIotInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *ApIotInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApIotInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApIotInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApIotInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPullup

`func (o *ApIotInput) GetPullup() ApIotInputPullup`

GetPullup returns the Pullup field if non-nil, zero value otherwise.

### GetPullupOk

`func (o *ApIotInput) GetPullupOk() (*ApIotInputPullup, bool)`

GetPullupOk returns a tuple with the Pullup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullup

`func (o *ApIotInput) SetPullup(v ApIotInputPullup)`

SetPullup sets Pullup field to given value.

### HasPullup

`func (o *ApIotInput) HasPullup() bool`

HasPullup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


