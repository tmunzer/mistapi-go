# OrgLicenseAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmendmentId** | Pointer to **string** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;unamend&#x60;, the ID of the operation to cancel | [optional] 
**DstOrgId** | Pointer to **string** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;amend&#x60;, the id of the org where the license is moved | [optional] 
**Notes** | Pointer to **string** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;annotate&#x60; | [optional] 
**Op** | [**OrgLicenseActionOperation**](OrgLicenseActionOperation.md) |  | 
**Quantity** | Pointer to **int32** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;amend&#x60;, the number of licenses to move | [optional] 
**SubscriptionId** | Pointer to **string** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;amend&#x60; or &#x60;op&#x60;&#x3D;&#x3D;&#x60;delete&#x60;, the ID of the subscription to use | [optional] 

## Methods

### NewOrgLicenseAction

`func NewOrgLicenseAction(op OrgLicenseActionOperation, ) *OrgLicenseAction`

NewOrgLicenseAction instantiates a new OrgLicenseAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgLicenseActionWithDefaults

`func NewOrgLicenseActionWithDefaults() *OrgLicenseAction`

NewOrgLicenseActionWithDefaults instantiates a new OrgLicenseAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmendmentId

`func (o *OrgLicenseAction) GetAmendmentId() string`

GetAmendmentId returns the AmendmentId field if non-nil, zero value otherwise.

### GetAmendmentIdOk

`func (o *OrgLicenseAction) GetAmendmentIdOk() (*string, bool)`

GetAmendmentIdOk returns a tuple with the AmendmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentId

`func (o *OrgLicenseAction) SetAmendmentId(v string)`

SetAmendmentId sets AmendmentId field to given value.

### HasAmendmentId

`func (o *OrgLicenseAction) HasAmendmentId() bool`

HasAmendmentId returns a boolean if a field has been set.

### GetDstOrgId

`func (o *OrgLicenseAction) GetDstOrgId() string`

GetDstOrgId returns the DstOrgId field if non-nil, zero value otherwise.

### GetDstOrgIdOk

`func (o *OrgLicenseAction) GetDstOrgIdOk() (*string, bool)`

GetDstOrgIdOk returns a tuple with the DstOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstOrgId

`func (o *OrgLicenseAction) SetDstOrgId(v string)`

SetDstOrgId sets DstOrgId field to given value.

### HasDstOrgId

`func (o *OrgLicenseAction) HasDstOrgId() bool`

HasDstOrgId returns a boolean if a field has been set.

### GetNotes

`func (o *OrgLicenseAction) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrgLicenseAction) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrgLicenseAction) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrgLicenseAction) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOp

`func (o *OrgLicenseAction) GetOp() OrgLicenseActionOperation`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *OrgLicenseAction) GetOpOk() (*OrgLicenseActionOperation, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *OrgLicenseAction) SetOp(v OrgLicenseActionOperation)`

SetOp sets Op field to given value.


### GetQuantity

`func (o *OrgLicenseAction) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrgLicenseAction) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrgLicenseAction) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrgLicenseAction) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *OrgLicenseAction) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *OrgLicenseAction) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *OrgLicenseAction) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *OrgLicenseAction) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


