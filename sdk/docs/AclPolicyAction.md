# AclPolicyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**AllowDeny**](AllowDeny.md) |  | [optional] [default to ALLOWDENY_ALLOW]
**DstTag** | Pointer to **string** |  | [optional] 

## Methods

### NewAclPolicyAction

`func NewAclPolicyAction() *AclPolicyAction`

NewAclPolicyAction instantiates a new AclPolicyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclPolicyActionWithDefaults

`func NewAclPolicyActionWithDefaults() *AclPolicyAction`

NewAclPolicyActionWithDefaults instantiates a new AclPolicyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AclPolicyAction) GetAction() AllowDeny`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AclPolicyAction) GetActionOk() (*AllowDeny, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AclPolicyAction) SetAction(v AllowDeny)`

SetAction sets Action field to given value.

### HasAction

`func (o *AclPolicyAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDstTag

`func (o *AclPolicyAction) GetDstTag() string`

GetDstTag returns the DstTag field if non-nil, zero value otherwise.

### GetDstTagOk

`func (o *AclPolicyAction) GetDstTagOk() (*string, bool)`

GetDstTagOk returns a tuple with the DstTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstTag

`func (o *AclPolicyAction) SetDstTag(v string)`

SetDstTag sets DstTag field to given value.

### HasDstTag

`func (o *AclPolicyAction) HasDstTag() bool`

HasDstTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


