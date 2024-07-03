# IdpProfileOverwrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**IdpProfileAction**](IdpProfileAction.md) |  | [optional] [default to IDPPROFILEACTION_ALERT]
**Matching** | Pointer to [**IdpProfileMatching**](IdpProfileMatching.md) |  | [optional] 

## Methods

### NewIdpProfileOverwrite

`func NewIdpProfileOverwrite() *IdpProfileOverwrite`

NewIdpProfileOverwrite instantiates a new IdpProfileOverwrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpProfileOverwriteWithDefaults

`func NewIdpProfileOverwriteWithDefaults() *IdpProfileOverwrite`

NewIdpProfileOverwriteWithDefaults instantiates a new IdpProfileOverwrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IdpProfileOverwrite) GetAction() IdpProfileAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IdpProfileOverwrite) GetActionOk() (*IdpProfileAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IdpProfileOverwrite) SetAction(v IdpProfileAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *IdpProfileOverwrite) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMatching

`func (o *IdpProfileOverwrite) GetMatching() IdpProfileMatching`

GetMatching returns the Matching field if non-nil, zero value otherwise.

### GetMatchingOk

`func (o *IdpProfileOverwrite) GetMatchingOk() (*IdpProfileMatching, bool)`

GetMatchingOk returns a tuple with the Matching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatching

`func (o *IdpProfileOverwrite) SetMatching(v IdpProfileMatching)`

SetMatching sets Matching field to given value.

### HasMatching

`func (o *IdpProfileOverwrite) HasMatching() bool`

HasMatching returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


