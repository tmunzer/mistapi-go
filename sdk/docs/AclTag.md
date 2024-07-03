# AclTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GbpTag** | Pointer to **int32** | required if - &#x60;type&#x60;&#x3D;&#x3D;&#x60;dynamic_gbp&#x60; (gbp_tag received from RADIUS) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; (applying gbp tag against matching conditions) | [optional] 
**Macs** | Pointer to **[]string** | required if  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;mac&#x60; - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching mac | [optional] 
**Network** | Pointer to **string** | if: - &#x60;type&#x60;&#x3D;&#x3D;&#x60;mac&#x60; (optional. default is &#x60;any&#x60;) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;subnet&#x60; (optional. default is &#x60;any&#x60;) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;network&#x60; - &#x60;type&#x60;&#x3D;&#x3D;&#x60;resource&#x60; (optional. default is &#x60;any&#x60;) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching network (vlan) | [optional] 
**RadiusGroup** | Pointer to **string** | required if  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;radius_group&#x60;  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching radius_group | [optional] 
**Specs** | Pointer to [**[]AclTagSpec**](AclTagSpec.md) | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;resource&#x60; empty means unrestricted, i.e. any | [optional] 
**Subnets** | Pointer to **[]string** | if  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;subnet&#x60;  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;resource&#x60; (optional. default is &#x60;any&#x60;) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching subnet | [optional] 
**Type** | [**AclTagType**](AclTagType.md) |  | 

## Methods

### NewAclTag

`func NewAclTag(type_ AclTagType, ) *AclTag`

NewAclTag instantiates a new AclTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclTagWithDefaults

`func NewAclTagWithDefaults() *AclTag`

NewAclTagWithDefaults instantiates a new AclTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGbpTag

`func (o *AclTag) GetGbpTag() int32`

GetGbpTag returns the GbpTag field if non-nil, zero value otherwise.

### GetGbpTagOk

`func (o *AclTag) GetGbpTagOk() (*int32, bool)`

GetGbpTagOk returns a tuple with the GbpTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbpTag

`func (o *AclTag) SetGbpTag(v int32)`

SetGbpTag sets GbpTag field to given value.

### HasGbpTag

`func (o *AclTag) HasGbpTag() bool`

HasGbpTag returns a boolean if a field has been set.

### GetMacs

`func (o *AclTag) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *AclTag) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *AclTag) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *AclTag) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetNetwork

`func (o *AclTag) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AclTag) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AclTag) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *AclTag) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRadiusGroup

`func (o *AclTag) GetRadiusGroup() string`

GetRadiusGroup returns the RadiusGroup field if non-nil, zero value otherwise.

### GetRadiusGroupOk

`func (o *AclTag) GetRadiusGroupOk() (*string, bool)`

GetRadiusGroupOk returns a tuple with the RadiusGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroup

`func (o *AclTag) SetRadiusGroup(v string)`

SetRadiusGroup sets RadiusGroup field to given value.

### HasRadiusGroup

`func (o *AclTag) HasRadiusGroup() bool`

HasRadiusGroup returns a boolean if a field has been set.

### GetSpecs

`func (o *AclTag) GetSpecs() []AclTagSpec`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *AclTag) GetSpecsOk() (*[]AclTagSpec, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *AclTag) SetSpecs(v []AclTagSpec)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *AclTag) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetSubnets

`func (o *AclTag) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *AclTag) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *AclTag) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *AclTag) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetType

`func (o *AclTag) GetType() AclTagType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AclTag) GetTypeOk() (*AclTagType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AclTag) SetType(v AclTagType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


