# AclPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]AclPolicyAction**](AclPolicyAction.md) | - for GBP-based policy, all src_tags and dst_tags have to be gbp-based - for ACL-based policy, &#x60;network&#x60; is required in either the source or destination so that we know where to attach the policy to | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SrcTags** | Pointer to **[]string** | - for GBP-based policy, all src_tags and dst_tags have to be gbp-based - for ACL-based policy, &#x60;network&#x60; is required in either the source or destination so that we know where to attach the policy to | [optional] 

## Methods

### NewAclPolicy

`func NewAclPolicy() *AclPolicy`

NewAclPolicy instantiates a new AclPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclPolicyWithDefaults

`func NewAclPolicyWithDefaults() *AclPolicy`

NewAclPolicyWithDefaults instantiates a new AclPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AclPolicy) GetActions() []AclPolicyAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AclPolicy) GetActionsOk() (*[]AclPolicyAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AclPolicy) SetActions(v []AclPolicyAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AclPolicy) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetName

`func (o *AclPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AclPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AclPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AclPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSrcTags

`func (o *AclPolicy) GetSrcTags() []string`

GetSrcTags returns the SrcTags field if non-nil, zero value otherwise.

### GetSrcTagsOk

`func (o *AclPolicy) GetSrcTagsOk() (*[]string, bool)`

GetSrcTagsOk returns a tuple with the SrcTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTags

`func (o *AclPolicy) SetSrcTags(v []string)`

SetSrcTags sets SrcTags field to given value.

### HasSrcTags

`func (o *AclPolicy) HasSrcTags() bool`

HasSrcTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


