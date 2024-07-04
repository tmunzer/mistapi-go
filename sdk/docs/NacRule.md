# NacRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**NacRuleAction**](NacRuleAction.md) |  | 
**ApplyTags** | Pointer to **[]string** | all optional, this goes into Access-Accept | [optional] 
**CreatedTime** | Pointer to **int32** |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** | enabled or not | [optional] [default to true]
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Matching** | Pointer to [**NacRuleMatching**](NacRuleMatching.md) |  | [optional] 
**ModifiedTime** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**NotMatching** | Pointer to [**NacRuleMatching**](NacRuleMatching.md) |  | [optional] 
**Order** | Pointer to **int32** | the order of the rule, lower value implies higher priority | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewNacRule

`func NewNacRule(action NacRuleAction, name string, ) *NacRule`

NewNacRule instantiates a new NacRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNacRuleWithDefaults

`func NewNacRuleWithDefaults() *NacRule`

NewNacRuleWithDefaults instantiates a new NacRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *NacRule) GetAction() NacRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NacRule) GetActionOk() (*NacRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NacRule) SetAction(v NacRuleAction)`

SetAction sets Action field to given value.


### GetApplyTags

`func (o *NacRule) GetApplyTags() []string`

GetApplyTags returns the ApplyTags field if non-nil, zero value otherwise.

### GetApplyTagsOk

`func (o *NacRule) GetApplyTagsOk() (*[]string, bool)`

GetApplyTagsOk returns a tuple with the ApplyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTags

`func (o *NacRule) SetApplyTags(v []string)`

SetApplyTags sets ApplyTags field to given value.

### HasApplyTags

`func (o *NacRule) HasApplyTags() bool`

HasApplyTags returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NacRule) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NacRule) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NacRule) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NacRule) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEnabled

`func (o *NacRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NacRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NacRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NacRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *NacRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NacRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NacRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NacRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMatching

`func (o *NacRule) GetMatching() NacRuleMatching`

GetMatching returns the Matching field if non-nil, zero value otherwise.

### GetMatchingOk

`func (o *NacRule) GetMatchingOk() (*NacRuleMatching, bool)`

GetMatchingOk returns a tuple with the Matching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatching

`func (o *NacRule) SetMatching(v NacRuleMatching)`

SetMatching sets Matching field to given value.

### HasMatching

`func (o *NacRule) HasMatching() bool`

HasMatching returns a boolean if a field has been set.

### GetModifiedTime

`func (o *NacRule) GetModifiedTime() int32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *NacRule) GetModifiedTimeOk() (*int32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *NacRule) SetModifiedTime(v int32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *NacRule) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *NacRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NacRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NacRule) SetName(v string)`

SetName sets Name field to given value.


### GetNotMatching

`func (o *NacRule) GetNotMatching() NacRuleMatching`

GetNotMatching returns the NotMatching field if non-nil, zero value otherwise.

### GetNotMatchingOk

`func (o *NacRule) GetNotMatchingOk() (*NacRuleMatching, bool)`

GetNotMatchingOk returns a tuple with the NotMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotMatching

`func (o *NacRule) SetNotMatching(v NacRuleMatching)`

SetNotMatching sets NotMatching field to given value.

### HasNotMatching

`func (o *NacRule) HasNotMatching() bool`

HasNotMatching returns a boolean if a field has been set.

### GetOrder

`func (o *NacRule) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *NacRule) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *NacRule) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *NacRule) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrgId

`func (o *NacRule) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NacRule) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NacRule) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NacRule) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


