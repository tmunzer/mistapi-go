# SnmpVacmSecurityToGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]SnmpVacmSecurityToGroupContentItem**](SnmpVacmSecurityToGroupContentItem.md) |  | [optional] 
**SecurityModel** | Pointer to [**SnmpVacmSecurityModel**](SnmpVacmSecurityModel.md) |  | [optional] 

## Methods

### NewSnmpVacmSecurityToGroup

`func NewSnmpVacmSecurityToGroup() *SnmpVacmSecurityToGroup`

NewSnmpVacmSecurityToGroup instantiates a new SnmpVacmSecurityToGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpVacmSecurityToGroupWithDefaults

`func NewSnmpVacmSecurityToGroupWithDefaults() *SnmpVacmSecurityToGroup`

NewSnmpVacmSecurityToGroupWithDefaults instantiates a new SnmpVacmSecurityToGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SnmpVacmSecurityToGroup) GetContent() []SnmpVacmSecurityToGroupContentItem`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SnmpVacmSecurityToGroup) GetContentOk() (*[]SnmpVacmSecurityToGroupContentItem, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SnmpVacmSecurityToGroup) SetContent(v []SnmpVacmSecurityToGroupContentItem)`

SetContent sets Content field to given value.

### HasContent

`func (o *SnmpVacmSecurityToGroup) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetSecurityModel

`func (o *SnmpVacmSecurityToGroup) GetSecurityModel() SnmpVacmSecurityModel`

GetSecurityModel returns the SecurityModel field if non-nil, zero value otherwise.

### GetSecurityModelOk

`func (o *SnmpVacmSecurityToGroup) GetSecurityModelOk() (*SnmpVacmSecurityModel, bool)`

GetSecurityModelOk returns a tuple with the SecurityModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityModel

`func (o *SnmpVacmSecurityToGroup) SetSecurityModel(v SnmpVacmSecurityModel)`

SetSecurityModel sets SecurityModel field to given value.

### HasSecurityModel

`func (o *SnmpVacmSecurityToGroup) HasSecurityModel() bool`

HasSecurityModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


