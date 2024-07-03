# ResponseOrgInventoryChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **[]string** |  | 
**Op** | [**ResponseOrgInventoryChangeOp**](ResponseOrgInventoryChangeOp.md) |  | 
**Reason** | **[]string** |  | 
**Success** | **[]string** |  | 

## Methods

### NewResponseOrgInventoryChange

`func NewResponseOrgInventoryChange(error_ []string, op ResponseOrgInventoryChangeOp, reason []string, success []string, ) *ResponseOrgInventoryChange`

NewResponseOrgInventoryChange instantiates a new ResponseOrgInventoryChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrgInventoryChangeWithDefaults

`func NewResponseOrgInventoryChangeWithDefaults() *ResponseOrgInventoryChange`

NewResponseOrgInventoryChangeWithDefaults instantiates a new ResponseOrgInventoryChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ResponseOrgInventoryChange) GetError() []string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseOrgInventoryChange) GetErrorOk() (*[]string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseOrgInventoryChange) SetError(v []string)`

SetError sets Error field to given value.


### GetOp

`func (o *ResponseOrgInventoryChange) GetOp() ResponseOrgInventoryChangeOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ResponseOrgInventoryChange) GetOpOk() (*ResponseOrgInventoryChangeOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ResponseOrgInventoryChange) SetOp(v ResponseOrgInventoryChangeOp)`

SetOp sets Op field to given value.


### GetReason

`func (o *ResponseOrgInventoryChange) GetReason() []string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ResponseOrgInventoryChange) GetReasonOk() (*[]string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ResponseOrgInventoryChange) SetReason(v []string)`

SetReason sets Reason field to given value.


### GetSuccess

`func (o *ResponseOrgInventoryChange) GetSuccess() []string`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseOrgInventoryChange) GetSuccessOk() (*[]string, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseOrgInventoryChange) SetSuccess(v []string)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


