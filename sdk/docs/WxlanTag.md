# WxlanTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastIps** | Pointer to **[]string** |  | [optional] 
**Mac** | Pointer to **NullableString** |  | [optional] 
**Match** | Pointer to [**WxlanTagMatch**](WxlanTagMatch.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | **string** | The name | 
**Op** | Pointer to [**WxlanTagOperation**](WxlanTagOperation.md) |  | [optional] [default to WXLANTAGOPERATION_IN]
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**ResourceMac** | Pointer to **NullableString** |  | [optional] 
**Services** | Pointer to **[]string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Specs** | Pointer to [**[]WxlanTagSpec**](WxlanTagSpec.md) | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;specs&#x60; | [optional] 
**Subnet** | Pointer to **string** |  | [optional] 
**Type** | [**WxlanTagType**](WxlanTagType.md) |  | 
**Values** | Pointer to **[]string** | if &#x60;type&#x60;!&#x3D;&#x60;vlan_id&#x60; and &#x60;type&#x60;!&#x3D;&#x60;specs&#x60;, list of values to match | [optional] 
**VlanId** | Pointer to **int32** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;vlan_id&#x60; | [optional] 

## Methods

### NewWxlanTag

`func NewWxlanTag(name string, type_ WxlanTagType, ) *WxlanTag`

NewWxlanTag instantiates a new WxlanTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWxlanTagWithDefaults

`func NewWxlanTagWithDefaults() *WxlanTag`

NewWxlanTagWithDefaults instantiates a new WxlanTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *WxlanTag) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *WxlanTag) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *WxlanTag) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *WxlanTag) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetForSite

`func (o *WxlanTag) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *WxlanTag) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *WxlanTag) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *WxlanTag) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *WxlanTag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WxlanTag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WxlanTag) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WxlanTag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastIps

`func (o *WxlanTag) GetLastIps() []string`

GetLastIps returns the LastIps field if non-nil, zero value otherwise.

### GetLastIpsOk

`func (o *WxlanTag) GetLastIpsOk() (*[]string, bool)`

GetLastIpsOk returns a tuple with the LastIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIps

`func (o *WxlanTag) SetLastIps(v []string)`

SetLastIps sets LastIps field to given value.

### HasLastIps

`func (o *WxlanTag) HasLastIps() bool`

HasLastIps returns a boolean if a field has been set.

### GetMac

`func (o *WxlanTag) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WxlanTag) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WxlanTag) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *WxlanTag) HasMac() bool`

HasMac returns a boolean if a field has been set.

### SetMacNil

`func (o *WxlanTag) SetMacNil(b bool)`

 SetMacNil sets the value for Mac to be an explicit nil

### UnsetMac
`func (o *WxlanTag) UnsetMac()`

UnsetMac ensures that no value is present for Mac, not even an explicit nil
### GetMatch

`func (o *WxlanTag) GetMatch() WxlanTagMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *WxlanTag) GetMatchOk() (*WxlanTagMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *WxlanTag) SetMatch(v WxlanTagMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *WxlanTag) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetModifiedTime

`func (o *WxlanTag) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *WxlanTag) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *WxlanTag) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *WxlanTag) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *WxlanTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WxlanTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WxlanTag) SetName(v string)`

SetName sets Name field to given value.


### GetOp

`func (o *WxlanTag) GetOp() WxlanTagOperation`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *WxlanTag) GetOpOk() (*WxlanTagOperation, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *WxlanTag) SetOp(v WxlanTagOperation)`

SetOp sets Op field to given value.

### HasOp

`func (o *WxlanTag) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetOrgId

`func (o *WxlanTag) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WxlanTag) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WxlanTag) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WxlanTag) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetResourceMac

`func (o *WxlanTag) GetResourceMac() string`

GetResourceMac returns the ResourceMac field if non-nil, zero value otherwise.

### GetResourceMacOk

`func (o *WxlanTag) GetResourceMacOk() (*string, bool)`

GetResourceMacOk returns a tuple with the ResourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMac

`func (o *WxlanTag) SetResourceMac(v string)`

SetResourceMac sets ResourceMac field to given value.

### HasResourceMac

`func (o *WxlanTag) HasResourceMac() bool`

HasResourceMac returns a boolean if a field has been set.

### SetResourceMacNil

`func (o *WxlanTag) SetResourceMacNil(b bool)`

 SetResourceMacNil sets the value for ResourceMac to be an explicit nil

### UnsetResourceMac
`func (o *WxlanTag) UnsetResourceMac()`

UnsetResourceMac ensures that no value is present for ResourceMac, not even an explicit nil
### GetServices

`func (o *WxlanTag) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *WxlanTag) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *WxlanTag) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *WxlanTag) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSiteId

`func (o *WxlanTag) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WxlanTag) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WxlanTag) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WxlanTag) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSpecs

`func (o *WxlanTag) GetSpecs() []WxlanTagSpec`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *WxlanTag) GetSpecsOk() (*[]WxlanTagSpec, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *WxlanTag) SetSpecs(v []WxlanTagSpec)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *WxlanTag) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetSubnet

`func (o *WxlanTag) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *WxlanTag) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *WxlanTag) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *WxlanTag) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetType

`func (o *WxlanTag) GetType() WxlanTagType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WxlanTag) GetTypeOk() (*WxlanTagType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WxlanTag) SetType(v WxlanTagType)`

SetType sets Type field to given value.


### GetValues

`func (o *WxlanTag) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WxlanTag) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *WxlanTag) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *WxlanTag) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetVlanId

`func (o *WxlanTag) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *WxlanTag) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *WxlanTag) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *WxlanTag) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


