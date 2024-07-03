# VrrpGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKey** | Pointer to **string** | if &#x60;auth_type&#x60;&#x3D;&#x3D;&#x60;md5&#x60; | [optional] 
**AuthPassword** | Pointer to **string** | if &#x60;auth_type&#x60;&#x3D;&#x3D;&#x60;simple&#x60; | [optional] 
**AuthType** | Pointer to [**VrrpGroupAuthType**](VrrpGroupAuthType.md) |  | [optional] [default to VRRPGROUPAUTHTYPE_MD5]
**Networks** | Pointer to [**map[string]VrrpGroupNetwork**](VrrpGroupNetwork.md) | Property key is the network name | [optional] 

## Methods

### NewVrrpGroup

`func NewVrrpGroup() *VrrpGroup`

NewVrrpGroup instantiates a new VrrpGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrrpGroupWithDefaults

`func NewVrrpGroupWithDefaults() *VrrpGroup`

NewVrrpGroupWithDefaults instantiates a new VrrpGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthKey

`func (o *VrrpGroup) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *VrrpGroup) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *VrrpGroup) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *VrrpGroup) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetAuthPassword

`func (o *VrrpGroup) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *VrrpGroup) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *VrrpGroup) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *VrrpGroup) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthType

`func (o *VrrpGroup) GetAuthType() VrrpGroupAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *VrrpGroup) GetAuthTypeOk() (*VrrpGroupAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *VrrpGroup) SetAuthType(v VrrpGroupAuthType)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *VrrpGroup) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetNetworks

`func (o *VrrpGroup) GetNetworks() map[string]VrrpGroupNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VrrpGroup) GetNetworksOk() (*map[string]VrrpGroupNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VrrpGroup) SetNetworks(v map[string]VrrpGroupNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VrrpGroup) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


