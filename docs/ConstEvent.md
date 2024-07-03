# ConstEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Display** | **string** |  | 
**Example** | Pointer to **map[string]interface{}** |  | [optional] 
**Key** | **string** |  | 

## Methods

### NewConstEvent

`func NewConstEvent(display string, key string, ) *ConstEvent`

NewConstEvent instantiates a new ConstEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstEventWithDefaults

`func NewConstEventWithDefaults() *ConstEvent`

NewConstEventWithDefaults instantiates a new ConstEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ConstEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConstEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConstEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConstEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplay

`func (o *ConstEvent) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConstEvent) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConstEvent) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetExample

`func (o *ConstEvent) GetExample() map[string]interface{}`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *ConstEvent) GetExampleOk() (*map[string]interface{}, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *ConstEvent) SetExample(v map[string]interface{})`

SetExample sets Example field to given value.

### HasExample

`func (o *ConstEvent) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetKey

`func (o *ConstEvent) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConstEvent) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConstEvent) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


