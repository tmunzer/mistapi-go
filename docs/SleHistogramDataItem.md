# SleHistogramDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to **[]float32** |  | [optional] 
**Value** | **float32** |  | 

## Methods

### NewSleHistogramDataItem

`func NewSleHistogramDataItem(value float32, ) *SleHistogramDataItem`

NewSleHistogramDataItem instantiates a new SleHistogramDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSleHistogramDataItemWithDefaults

`func NewSleHistogramDataItemWithDefaults() *SleHistogramDataItem`

NewSleHistogramDataItemWithDefaults instantiates a new SleHistogramDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *SleHistogramDataItem) GetRange() []*float32`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *SleHistogramDataItem) GetRangeOk() (*[]*float32, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *SleHistogramDataItem) SetRange(v []*float32)`

SetRange sets Range field to given value.

### HasRange

`func (o *SleHistogramDataItem) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetValue

`func (o *SleHistogramDataItem) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SleHistogramDataItem) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SleHistogramDataItem) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


