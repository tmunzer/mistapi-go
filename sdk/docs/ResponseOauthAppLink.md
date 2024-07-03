# ResponseOauthAppLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]ResponseOauthAppLinkItem**](ResponseOauthAppLinkItem.md) | List of linked account details | [optional] 
**Linked** | Pointer to **bool** | Basic Auth application linked status in mist portal enabled for VMware | [optional] 

## Methods

### NewResponseOauthAppLink

`func NewResponseOauthAppLink() *ResponseOauthAppLink`

NewResponseOauthAppLink instantiates a new ResponseOauthAppLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOauthAppLinkWithDefaults

`func NewResponseOauthAppLinkWithDefaults() *ResponseOauthAppLink`

NewResponseOauthAppLinkWithDefaults instantiates a new ResponseOauthAppLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ResponseOauthAppLink) GetAccounts() []ResponseOauthAppLinkItem`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ResponseOauthAppLink) GetAccountsOk() (*[]ResponseOauthAppLinkItem, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ResponseOauthAppLink) SetAccounts(v []ResponseOauthAppLinkItem)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ResponseOauthAppLink) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetLinked

`func (o *ResponseOauthAppLink) GetLinked() bool`

GetLinked returns the Linked field if non-nil, zero value otherwise.

### GetLinkedOk

`func (o *ResponseOauthAppLink) GetLinkedOk() (*bool, bool)`

GetLinkedOk returns a tuple with the Linked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinked

`func (o *ResponseOauthAppLink) SetLinked(v bool)`

SetLinked sets Linked field to given value.

### HasLinked

`func (o *ResponseOauthAppLink) HasLinked() bool`

HasLinked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


