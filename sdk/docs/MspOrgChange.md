# MspOrgChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**MspOrgChangeOperation**](MspOrgChangeOperation.md) |  | 
**OrgIds** | **[]string** | list of org_id | 

## Methods

### NewMspOrgChange

`func NewMspOrgChange(op MspOrgChangeOperation, orgIds []string, ) *MspOrgChange`

NewMspOrgChange instantiates a new MspOrgChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspOrgChangeWithDefaults

`func NewMspOrgChangeWithDefaults() *MspOrgChange`

NewMspOrgChangeWithDefaults instantiates a new MspOrgChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *MspOrgChange) GetOp() MspOrgChangeOperation`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *MspOrgChange) GetOpOk() (*MspOrgChangeOperation, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *MspOrgChange) SetOp(v MspOrgChangeOperation)`

SetOp sets Op field to given value.


### GetOrgIds

`func (o *MspOrgChange) GetOrgIds() []string`

GetOrgIds returns the OrgIds field if non-nil, zero value otherwise.

### GetOrgIdsOk

`func (o *MspOrgChange) GetOrgIdsOk() (*[]string, bool)`

GetOrgIdsOk returns a tuple with the OrgIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgIds

`func (o *MspOrgChange) SetOrgIds(v []string)`

SetOrgIds sets OrgIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


