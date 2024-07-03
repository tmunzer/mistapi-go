# ResponseInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | Pointer to **[]string** |  | [optional] 
**Duplicated** | Pointer to **[]string** |  | [optional] 
**Error** | Pointer to **[]string** |  | [optional] 
**InventoryAdded** | Pointer to [**[]ResponseInventoryInventoryAddedItems**](ResponseInventoryInventoryAddedItems.md) |  | [optional] 
**InventoryDuplicated** | Pointer to [**[]ResponseInventoryInventoryDuplicatedItems**](ResponseInventoryInventoryDuplicatedItems.md) |  | [optional] 

## Methods

### NewResponseInventory

`func NewResponseInventory() *ResponseInventory`

NewResponseInventory instantiates a new ResponseInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseInventoryWithDefaults

`func NewResponseInventoryWithDefaults() *ResponseInventory`

NewResponseInventoryWithDefaults instantiates a new ResponseInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *ResponseInventory) GetAdded() []string`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *ResponseInventory) GetAddedOk() (*[]string, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *ResponseInventory) SetAdded(v []string)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *ResponseInventory) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetDuplicated

`func (o *ResponseInventory) GetDuplicated() []string`

GetDuplicated returns the Duplicated field if non-nil, zero value otherwise.

### GetDuplicatedOk

`func (o *ResponseInventory) GetDuplicatedOk() (*[]string, bool)`

GetDuplicatedOk returns a tuple with the Duplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicated

`func (o *ResponseInventory) SetDuplicated(v []string)`

SetDuplicated sets Duplicated field to given value.

### HasDuplicated

`func (o *ResponseInventory) HasDuplicated() bool`

HasDuplicated returns a boolean if a field has been set.

### GetError

`func (o *ResponseInventory) GetError() []string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseInventory) GetErrorOk() (*[]string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseInventory) SetError(v []string)`

SetError sets Error field to given value.

### HasError

`func (o *ResponseInventory) HasError() bool`

HasError returns a boolean if a field has been set.

### GetInventoryAdded

`func (o *ResponseInventory) GetInventoryAdded() []ResponseInventoryInventoryAddedItems`

GetInventoryAdded returns the InventoryAdded field if non-nil, zero value otherwise.

### GetInventoryAddedOk

`func (o *ResponseInventory) GetInventoryAddedOk() (*[]ResponseInventoryInventoryAddedItems, bool)`

GetInventoryAddedOk returns a tuple with the InventoryAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryAdded

`func (o *ResponseInventory) SetInventoryAdded(v []ResponseInventoryInventoryAddedItems)`

SetInventoryAdded sets InventoryAdded field to given value.

### HasInventoryAdded

`func (o *ResponseInventory) HasInventoryAdded() bool`

HasInventoryAdded returns a boolean if a field has been set.

### GetInventoryDuplicated

`func (o *ResponseInventory) GetInventoryDuplicated() []ResponseInventoryInventoryDuplicatedItems`

GetInventoryDuplicated returns the InventoryDuplicated field if non-nil, zero value otherwise.

### GetInventoryDuplicatedOk

`func (o *ResponseInventory) GetInventoryDuplicatedOk() (*[]ResponseInventoryInventoryDuplicatedItems, bool)`

GetInventoryDuplicatedOk returns a tuple with the InventoryDuplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDuplicated

`func (o *ResponseInventory) SetInventoryDuplicated(v []ResponseInventoryInventoryDuplicatedItems)`

SetInventoryDuplicated sets InventoryDuplicated field to given value.

### HasInventoryDuplicated

`func (o *ResponseInventory) HasInventoryDuplicated() bool`

HasInventoryDuplicated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


