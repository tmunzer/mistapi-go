# WxlanRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**WxlanRuleAction**](WxlanRuleAction.md) |  | [optional] 
**ApplyTags** | Pointer to **[]string** |  | [optional] 
**BlockedApps** | Pointer to **[]string** | blocked apps (always blocking, ignoring action), the key of Get Application List | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DstAllowWxtags** | Pointer to **[]string** | tag list to indicate these tags are allowed access | [optional] 
**DstDenyWxtags** | Pointer to **[]string** | tag list to indicate these tags are blocked access | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Order** | **int32** | the order how rules would be looked up, &gt; 0 and bigger order got matched first, -1 means LAST, uniqueness not checked | 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SrcWxtags** | **[]string** | tag list to determine if this rule would match | 
**TemplateId** | Pointer to **string** | Only for Org Level WxRule | [optional] 

## Methods

### NewWxlanRule

`func NewWxlanRule(order int32, srcWxtags []string, ) *WxlanRule`

NewWxlanRule instantiates a new WxlanRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWxlanRuleWithDefaults

`func NewWxlanRuleWithDefaults() *WxlanRule`

NewWxlanRuleWithDefaults instantiates a new WxlanRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WxlanRule) GetAction() WxlanRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WxlanRule) GetActionOk() (*WxlanRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WxlanRule) SetAction(v WxlanRuleAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *WxlanRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetApplyTags

`func (o *WxlanRule) GetApplyTags() []string`

GetApplyTags returns the ApplyTags field if non-nil, zero value otherwise.

### GetApplyTagsOk

`func (o *WxlanRule) GetApplyTagsOk() (*[]string, bool)`

GetApplyTagsOk returns a tuple with the ApplyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTags

`func (o *WxlanRule) SetApplyTags(v []string)`

SetApplyTags sets ApplyTags field to given value.

### HasApplyTags

`func (o *WxlanRule) HasApplyTags() bool`

HasApplyTags returns a boolean if a field has been set.

### GetBlockedApps

`func (o *WxlanRule) GetBlockedApps() []string`

GetBlockedApps returns the BlockedApps field if non-nil, zero value otherwise.

### GetBlockedAppsOk

`func (o *WxlanRule) GetBlockedAppsOk() (*[]string, bool)`

GetBlockedAppsOk returns a tuple with the BlockedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedApps

`func (o *WxlanRule) SetBlockedApps(v []string)`

SetBlockedApps sets BlockedApps field to given value.

### HasBlockedApps

`func (o *WxlanRule) HasBlockedApps() bool`

HasBlockedApps returns a boolean if a field has been set.

### GetCreatedTime

`func (o *WxlanRule) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *WxlanRule) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *WxlanRule) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *WxlanRule) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDstAllowWxtags

`func (o *WxlanRule) GetDstAllowWxtags() []string`

GetDstAllowWxtags returns the DstAllowWxtags field if non-nil, zero value otherwise.

### GetDstAllowWxtagsOk

`func (o *WxlanRule) GetDstAllowWxtagsOk() (*[]string, bool)`

GetDstAllowWxtagsOk returns a tuple with the DstAllowWxtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstAllowWxtags

`func (o *WxlanRule) SetDstAllowWxtags(v []string)`

SetDstAllowWxtags sets DstAllowWxtags field to given value.

### HasDstAllowWxtags

`func (o *WxlanRule) HasDstAllowWxtags() bool`

HasDstAllowWxtags returns a boolean if a field has been set.

### GetDstDenyWxtags

`func (o *WxlanRule) GetDstDenyWxtags() []string`

GetDstDenyWxtags returns the DstDenyWxtags field if non-nil, zero value otherwise.

### GetDstDenyWxtagsOk

`func (o *WxlanRule) GetDstDenyWxtagsOk() (*[]string, bool)`

GetDstDenyWxtagsOk returns a tuple with the DstDenyWxtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstDenyWxtags

`func (o *WxlanRule) SetDstDenyWxtags(v []string)`

SetDstDenyWxtags sets DstDenyWxtags field to given value.

### HasDstDenyWxtags

`func (o *WxlanRule) HasDstDenyWxtags() bool`

HasDstDenyWxtags returns a boolean if a field has been set.

### GetEnabled

`func (o *WxlanRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WxlanRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WxlanRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WxlanRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForSite

`func (o *WxlanRule) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *WxlanRule) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *WxlanRule) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *WxlanRule) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *WxlanRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WxlanRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WxlanRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WxlanRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *WxlanRule) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *WxlanRule) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *WxlanRule) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *WxlanRule) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetOrder

`func (o *WxlanRule) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WxlanRule) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WxlanRule) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetOrgId

`func (o *WxlanRule) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WxlanRule) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WxlanRule) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WxlanRule) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSiteId

`func (o *WxlanRule) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WxlanRule) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WxlanRule) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WxlanRule) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSrcWxtags

`func (o *WxlanRule) GetSrcWxtags() []string`

GetSrcWxtags returns the SrcWxtags field if non-nil, zero value otherwise.

### GetSrcWxtagsOk

`func (o *WxlanRule) GetSrcWxtagsOk() (*[]string, bool)`

GetSrcWxtagsOk returns a tuple with the SrcWxtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcWxtags

`func (o *WxlanRule) SetSrcWxtags(v []string)`

SetSrcWxtags sets SrcWxtags field to given value.


### GetTemplateId

`func (o *WxlanRule) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *WxlanRule) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *WxlanRule) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *WxlanRule) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


