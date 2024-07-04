# LicenseSub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] 
**EndTime** | Pointer to **int32** | end date of the license term | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] [readonly] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Quantity** | Pointer to **int32** | number of devices entitled for this license | [optional] [readonly] 
**RemainingQuantity** | Pointer to **int32** | Number of licenses left in this subscription | [optional] 
**StartTime** | Pointer to **int32** | start date of the license term | [optional] [readonly] 
**SubscriptionId** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to [**LicenseType**](LicenseType.md) |  | [optional] 

## Methods

### NewLicenseSub

`func NewLicenseSub() *LicenseSub`

NewLicenseSub instantiates a new LicenseSub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseSubWithDefaults

`func NewLicenseSubWithDefaults() *LicenseSub`

NewLicenseSubWithDefaults instantiates a new LicenseSub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *LicenseSub) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *LicenseSub) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *LicenseSub) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *LicenseSub) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEndTime

`func (o *LicenseSub) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *LicenseSub) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *LicenseSub) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *LicenseSub) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *LicenseSub) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicenseSub) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicenseSub) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicenseSub) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *LicenseSub) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *LicenseSub) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *LicenseSub) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *LicenseSub) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetOrderId

`func (o *LicenseSub) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *LicenseSub) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *LicenseSub) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *LicenseSub) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrgId

`func (o *LicenseSub) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *LicenseSub) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *LicenseSub) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *LicenseSub) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetQuantity

`func (o *LicenseSub) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LicenseSub) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LicenseSub) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LicenseSub) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *LicenseSub) GetRemainingQuantity() int32`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *LicenseSub) GetRemainingQuantityOk() (*int32, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *LicenseSub) SetRemainingQuantity(v int32)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *LicenseSub) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetStartTime

`func (o *LicenseSub) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LicenseSub) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LicenseSub) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *LicenseSub) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *LicenseSub) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *LicenseSub) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *LicenseSub) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *LicenseSub) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetType

`func (o *LicenseSub) GetType() LicenseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicenseSub) GetTypeOk() (*LicenseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicenseSub) SetType(v LicenseType)`

SetType sets Type field to given value.

### HasType

`func (o *LicenseSub) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


