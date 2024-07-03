# MspLicenseAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmendmentId** | Pointer to **string** | required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;unamend&#x60; | [optional] 
**DstOrgId** | Pointer to **string** | required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;amend&#x60;, destination org id | [optional] 
**Notes** | Pointer to **string** | required if &#x60;op&#x60;&#x3D;&#x3D; &#x60;annotate&#x60; | [optional] 
**Op** | [**MspLicenseActionOperation**](MspLicenseActionOperation.md) |  | 
**Quantity** | Pointer to **float32** | required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;amend&#x60; | [optional] 
**SubscriptionId** | Pointer to **string** | required if &#x60;op&#x60;&#x3D;&#x3D; &#x60;annotate&#x60; | [optional] 

## Methods

### NewMspLicenseAction

`func NewMspLicenseAction(op MspLicenseActionOperation, ) *MspLicenseAction`

NewMspLicenseAction instantiates a new MspLicenseAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspLicenseActionWithDefaults

`func NewMspLicenseActionWithDefaults() *MspLicenseAction`

NewMspLicenseActionWithDefaults instantiates a new MspLicenseAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmendmentId

`func (o *MspLicenseAction) GetAmendmentId() string`

GetAmendmentId returns the AmendmentId field if non-nil, zero value otherwise.

### GetAmendmentIdOk

`func (o *MspLicenseAction) GetAmendmentIdOk() (*string, bool)`

GetAmendmentIdOk returns a tuple with the AmendmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentId

`func (o *MspLicenseAction) SetAmendmentId(v string)`

SetAmendmentId sets AmendmentId field to given value.

### HasAmendmentId

`func (o *MspLicenseAction) HasAmendmentId() bool`

HasAmendmentId returns a boolean if a field has been set.

### GetDstOrgId

`func (o *MspLicenseAction) GetDstOrgId() string`

GetDstOrgId returns the DstOrgId field if non-nil, zero value otherwise.

### GetDstOrgIdOk

`func (o *MspLicenseAction) GetDstOrgIdOk() (*string, bool)`

GetDstOrgIdOk returns a tuple with the DstOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstOrgId

`func (o *MspLicenseAction) SetDstOrgId(v string)`

SetDstOrgId sets DstOrgId field to given value.

### HasDstOrgId

`func (o *MspLicenseAction) HasDstOrgId() bool`

HasDstOrgId returns a boolean if a field has been set.

### GetNotes

`func (o *MspLicenseAction) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MspLicenseAction) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MspLicenseAction) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MspLicenseAction) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOp

`func (o *MspLicenseAction) GetOp() MspLicenseActionOperation`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *MspLicenseAction) GetOpOk() (*MspLicenseActionOperation, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *MspLicenseAction) SetOp(v MspLicenseActionOperation)`

SetOp sets Op field to given value.


### GetQuantity

`func (o *MspLicenseAction) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MspLicenseAction) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MspLicenseAction) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *MspLicenseAction) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *MspLicenseAction) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *MspLicenseAction) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *MspLicenseAction) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *MspLicenseAction) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


