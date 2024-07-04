# LastTrouble

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code definitions list at /api/v1/consts/ap_led_status | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 

## Methods

### NewLastTrouble

`func NewLastTrouble() *LastTrouble`

NewLastTrouble instantiates a new LastTrouble object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastTroubleWithDefaults

`func NewLastTroubleWithDefaults() *LastTrouble`

NewLastTroubleWithDefaults instantiates a new LastTrouble object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *LastTrouble) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LastTrouble) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LastTrouble) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LastTrouble) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTimestamp

`func (o *LastTrouble) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LastTrouble) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LastTrouble) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LastTrouble) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


