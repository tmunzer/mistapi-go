# ResponseLoginFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** |  | 
**ForwardUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseLoginFailure

`func NewResponseLoginFailure(detail string, ) *ResponseLoginFailure`

NewResponseLoginFailure instantiates a new ResponseLoginFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseLoginFailureWithDefaults

`func NewResponseLoginFailureWithDefaults() *ResponseLoginFailure`

NewResponseLoginFailureWithDefaults instantiates a new ResponseLoginFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *ResponseLoginFailure) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseLoginFailure) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseLoginFailure) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetForwardUrl

`func (o *ResponseLoginFailure) GetForwardUrl() string`

GetForwardUrl returns the ForwardUrl field if non-nil, zero value otherwise.

### GetForwardUrlOk

`func (o *ResponseLoginFailure) GetForwardUrlOk() (*string, bool)`

GetForwardUrlOk returns a tuple with the ForwardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardUrl

`func (o *ResponseLoginFailure) SetForwardUrl(v string)`

SetForwardUrl sets ForwardUrl field to given value.

### HasForwardUrl

`func (o *ResponseLoginFailure) HasForwardUrl() bool`

HasForwardUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


