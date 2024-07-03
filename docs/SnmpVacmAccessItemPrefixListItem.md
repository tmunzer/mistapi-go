# SnmpVacmAccessItemPrefixListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextPrefix** | Pointer to **string** | only required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;context_prefix&#x60; | [optional] 
**NotifyView** | Pointer to **string** | refer to view name | [optional] 
**ReadView** | Pointer to **string** | refer to view name | [optional] 
**SecurityLevel** | Pointer to [**SnmpVacmAccessItemPrefixListItemLevel**](SnmpVacmAccessItemPrefixListItemLevel.md) |  | [optional] 
**SecurityModel** | Pointer to [**SnmpVacmAccessItemPrefixListItemModel**](SnmpVacmAccessItemPrefixListItemModel.md) |  | [optional] 
**Type** | Pointer to [**SnmpVacmAccessItemType**](SnmpVacmAccessItemType.md) |  | [optional] 
**WriteView** | Pointer to **string** | refer to view name | [optional] 

## Methods

### NewSnmpVacmAccessItemPrefixListItem

`func NewSnmpVacmAccessItemPrefixListItem() *SnmpVacmAccessItemPrefixListItem`

NewSnmpVacmAccessItemPrefixListItem instantiates a new SnmpVacmAccessItemPrefixListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpVacmAccessItemPrefixListItemWithDefaults

`func NewSnmpVacmAccessItemPrefixListItemWithDefaults() *SnmpVacmAccessItemPrefixListItem`

NewSnmpVacmAccessItemPrefixListItemWithDefaults instantiates a new SnmpVacmAccessItemPrefixListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextPrefix

`func (o *SnmpVacmAccessItemPrefixListItem) GetContextPrefix() string`

GetContextPrefix returns the ContextPrefix field if non-nil, zero value otherwise.

### GetContextPrefixOk

`func (o *SnmpVacmAccessItemPrefixListItem) GetContextPrefixOk() (*string, bool)`

GetContextPrefixOk returns a tuple with the ContextPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextPrefix

`func (o *SnmpVacmAccessItemPrefixListItem) SetContextPrefix(v string)`

SetContextPrefix sets ContextPrefix field to given value.

### HasContextPrefix

`func (o *SnmpVacmAccessItemPrefixListItem) HasContextPrefix() bool`

HasContextPrefix returns a boolean if a field has been set.

### GetNotifyView

`func (o *SnmpVacmAccessItemPrefixListItem) GetNotifyView() string`

GetNotifyView returns the NotifyView field if non-nil, zero value otherwise.

### GetNotifyViewOk

`func (o *SnmpVacmAccessItemPrefixListItem) GetNotifyViewOk() (*string, bool)`

GetNotifyViewOk returns a tuple with the NotifyView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyView

`func (o *SnmpVacmAccessItemPrefixListItem) SetNotifyView(v string)`

SetNotifyView sets NotifyView field to given value.

### HasNotifyView

`func (o *SnmpVacmAccessItemPrefixListItem) HasNotifyView() bool`

HasNotifyView returns a boolean if a field has been set.

### GetReadView

`func (o *SnmpVacmAccessItemPrefixListItem) GetReadView() string`

GetReadView returns the ReadView field if non-nil, zero value otherwise.

### GetReadViewOk

`func (o *SnmpVacmAccessItemPrefixListItem) GetReadViewOk() (*string, bool)`

GetReadViewOk returns a tuple with the ReadView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadView

`func (o *SnmpVacmAccessItemPrefixListItem) SetReadView(v string)`

SetReadView sets ReadView field to given value.

### HasReadView

`func (o *SnmpVacmAccessItemPrefixListItem) HasReadView() bool`

HasReadView returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *SnmpVacmAccessItemPrefixListItem) GetSecurityLevel() SnmpVacmAccessItemPrefixListItemLevel`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *SnmpVacmAccessItemPrefixListItem) GetSecurityLevelOk() (*SnmpVacmAccessItemPrefixListItemLevel, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *SnmpVacmAccessItemPrefixListItem) SetSecurityLevel(v SnmpVacmAccessItemPrefixListItemLevel)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *SnmpVacmAccessItemPrefixListItem) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetSecurityModel

`func (o *SnmpVacmAccessItemPrefixListItem) GetSecurityModel() SnmpVacmAccessItemPrefixListItemModel`

GetSecurityModel returns the SecurityModel field if non-nil, zero value otherwise.

### GetSecurityModelOk

`func (o *SnmpVacmAccessItemPrefixListItem) GetSecurityModelOk() (*SnmpVacmAccessItemPrefixListItemModel, bool)`

GetSecurityModelOk returns a tuple with the SecurityModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityModel

`func (o *SnmpVacmAccessItemPrefixListItem) SetSecurityModel(v SnmpVacmAccessItemPrefixListItemModel)`

SetSecurityModel sets SecurityModel field to given value.

### HasSecurityModel

`func (o *SnmpVacmAccessItemPrefixListItem) HasSecurityModel() bool`

HasSecurityModel returns a boolean if a field has been set.

### GetType

`func (o *SnmpVacmAccessItemPrefixListItem) GetType() SnmpVacmAccessItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnmpVacmAccessItemPrefixListItem) GetTypeOk() (*SnmpVacmAccessItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnmpVacmAccessItemPrefixListItem) SetType(v SnmpVacmAccessItemType)`

SetType sets Type field to given value.

### HasType

`func (o *SnmpVacmAccessItemPrefixListItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWriteView

`func (o *SnmpVacmAccessItemPrefixListItem) GetWriteView() string`

GetWriteView returns the WriteView field if non-nil, zero value otherwise.

### GetWriteViewOk

`func (o *SnmpVacmAccessItemPrefixListItem) GetWriteViewOk() (*string, bool)`

GetWriteViewOk returns a tuple with the WriteView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteView

`func (o *SnmpVacmAccessItemPrefixListItem) SetWriteView(v string)`

SetWriteView sets WriteView field to given value.

### HasWriteView

`func (o *SnmpVacmAccessItemPrefixListItem) HasWriteView() bool`

HasWriteView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


