# NacTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsermacOverride** | Pointer to **bool** | can be set to true to allow the override by usermac result | [optional] [default to false]
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**EgressVlanNames** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;egress_vlan_names&#x60;, list of egress vlans to return | [optional] 
**GbpTag** | Pointer to **int32** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;gbp_tag&#x60; | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Match** | Pointer to [**NacTagMatch**](NacTagMatch.md) |  | [optional] 
**MatchAll** | Pointer to **bool** | This field is applicable only when &#x60;type&#x60;&#x3D;&#x3D;&#x60;match&#x60; * &#x60;false&#x60;: means it is sufficient to match any of the values (i.e., match-any behavior) * &#x60;true&#x60;: means all values should be matched (i.e., match-all behavior)   Currently it makes sense to set this field to &#x60;true&#x60; only if the &#x60;match&#x60;&#x3D;&#x3D;&#x60;idp_role&#x60; or &#x60;match&#x60;&#x3D;&#x3D;&#x60;usermac_label&#x60; | [optional] [default to false]
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**RadiusAttrs** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;radius_attrs&#x60;, user can specify a list of one or more standard attributes in the field \&quot;radius_attrs\&quot;.  It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected. Note that it is allowed to have more than one radius_attrs in the result of a given rule. | [optional] 
**RadiusGroup** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;radius_group&#x60; | [optional] 
**RadiusVendorAttrs** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;radius_vendor_attrs&#x60;, user can specify a list of one or more vendor-specific attributes in the field \&quot;radius_vendor_attrs\&quot;.  It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected. Note that it is allowed to have more than one radius_vendor_attrs in the result of a given rule. | [optional] 
**SessionTimeout** | Pointer to **int32** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;session_timeout, in seconds | [optional] 
**Type** | [**NacTagType**](NacTagType.md) |  | 
**Values** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;match&#x60; | [optional] 
**Vlan** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;vlan&#x60; | [optional] 

## Methods

### NewNacTag

`func NewNacTag(name string, type_ NacTagType, ) *NacTag`

NewNacTag instantiates a new NacTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNacTagWithDefaults

`func NewNacTagWithDefaults() *NacTag`

NewNacTagWithDefaults instantiates a new NacTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUsermacOverride

`func (o *NacTag) GetAllowUsermacOverride() bool`

GetAllowUsermacOverride returns the AllowUsermacOverride field if non-nil, zero value otherwise.

### GetAllowUsermacOverrideOk

`func (o *NacTag) GetAllowUsermacOverrideOk() (*bool, bool)`

GetAllowUsermacOverrideOk returns a tuple with the AllowUsermacOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsermacOverride

`func (o *NacTag) SetAllowUsermacOverride(v bool)`

SetAllowUsermacOverride sets AllowUsermacOverride field to given value.

### HasAllowUsermacOverride

`func (o *NacTag) HasAllowUsermacOverride() bool`

HasAllowUsermacOverride returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NacTag) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NacTag) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NacTag) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NacTag) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEgressVlanNames

`func (o *NacTag) GetEgressVlanNames() []string`

GetEgressVlanNames returns the EgressVlanNames field if non-nil, zero value otherwise.

### GetEgressVlanNamesOk

`func (o *NacTag) GetEgressVlanNamesOk() (*[]string, bool)`

GetEgressVlanNamesOk returns a tuple with the EgressVlanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressVlanNames

`func (o *NacTag) SetEgressVlanNames(v []string)`

SetEgressVlanNames sets EgressVlanNames field to given value.

### HasEgressVlanNames

`func (o *NacTag) HasEgressVlanNames() bool`

HasEgressVlanNames returns a boolean if a field has been set.

### GetGbpTag

`func (o *NacTag) GetGbpTag() int32`

GetGbpTag returns the GbpTag field if non-nil, zero value otherwise.

### GetGbpTagOk

`func (o *NacTag) GetGbpTagOk() (*int32, bool)`

GetGbpTagOk returns a tuple with the GbpTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbpTag

`func (o *NacTag) SetGbpTag(v int32)`

SetGbpTag sets GbpTag field to given value.

### HasGbpTag

`func (o *NacTag) HasGbpTag() bool`

HasGbpTag returns a boolean if a field has been set.

### GetId

`func (o *NacTag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NacTag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NacTag) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NacTag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMatch

`func (o *NacTag) GetMatch() NacTagMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *NacTag) GetMatchOk() (*NacTagMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *NacTag) SetMatch(v NacTagMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *NacTag) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetMatchAll

`func (o *NacTag) GetMatchAll() bool`

GetMatchAll returns the MatchAll field if non-nil, zero value otherwise.

### GetMatchAllOk

`func (o *NacTag) GetMatchAllOk() (*bool, bool)`

GetMatchAllOk returns a tuple with the MatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAll

`func (o *NacTag) SetMatchAll(v bool)`

SetMatchAll sets MatchAll field to given value.

### HasMatchAll

`func (o *NacTag) HasMatchAll() bool`

HasMatchAll returns a boolean if a field has been set.

### GetModifiedTime

`func (o *NacTag) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *NacTag) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *NacTag) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *NacTag) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *NacTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NacTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NacTag) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *NacTag) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NacTag) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NacTag) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NacTag) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRadiusAttrs

`func (o *NacTag) GetRadiusAttrs() []string`

GetRadiusAttrs returns the RadiusAttrs field if non-nil, zero value otherwise.

### GetRadiusAttrsOk

`func (o *NacTag) GetRadiusAttrsOk() (*[]string, bool)`

GetRadiusAttrsOk returns a tuple with the RadiusAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAttrs

`func (o *NacTag) SetRadiusAttrs(v []string)`

SetRadiusAttrs sets RadiusAttrs field to given value.

### HasRadiusAttrs

`func (o *NacTag) HasRadiusAttrs() bool`

HasRadiusAttrs returns a boolean if a field has been set.

### GetRadiusGroup

`func (o *NacTag) GetRadiusGroup() string`

GetRadiusGroup returns the RadiusGroup field if non-nil, zero value otherwise.

### GetRadiusGroupOk

`func (o *NacTag) GetRadiusGroupOk() (*string, bool)`

GetRadiusGroupOk returns a tuple with the RadiusGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroup

`func (o *NacTag) SetRadiusGroup(v string)`

SetRadiusGroup sets RadiusGroup field to given value.

### HasRadiusGroup

`func (o *NacTag) HasRadiusGroup() bool`

HasRadiusGroup returns a boolean if a field has been set.

### GetRadiusVendorAttrs

`func (o *NacTag) GetRadiusVendorAttrs() []string`

GetRadiusVendorAttrs returns the RadiusVendorAttrs field if non-nil, zero value otherwise.

### GetRadiusVendorAttrsOk

`func (o *NacTag) GetRadiusVendorAttrsOk() (*[]string, bool)`

GetRadiusVendorAttrsOk returns a tuple with the RadiusVendorAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusVendorAttrs

`func (o *NacTag) SetRadiusVendorAttrs(v []string)`

SetRadiusVendorAttrs sets RadiusVendorAttrs field to given value.

### HasRadiusVendorAttrs

`func (o *NacTag) HasRadiusVendorAttrs() bool`

HasRadiusVendorAttrs returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *NacTag) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *NacTag) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *NacTag) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *NacTag) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetType

`func (o *NacTag) GetType() NacTagType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NacTag) GetTypeOk() (*NacTagType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NacTag) SetType(v NacTagType)`

SetType sets Type field to given value.


### GetValues

`func (o *NacTag) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *NacTag) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *NacTag) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *NacTag) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetVlan

`func (o *NacTag) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NacTag) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NacTag) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *NacTag) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


