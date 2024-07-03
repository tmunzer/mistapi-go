# VpnPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BfdProfile** | Pointer to [**VpnPathBfdProfile**](VpnPathBfdProfile.md) |  | [optional] [default to VPNPATHBFDPROFILE_BROADBAND]
**Ip** | Pointer to **string** | if different from the wan port | [optional] 
**Pod** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewVpnPath

`func NewVpnPath() *VpnPath`

NewVpnPath instantiates a new VpnPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPathWithDefaults

`func NewVpnPathWithDefaults() *VpnPath`

NewVpnPathWithDefaults instantiates a new VpnPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBfdProfile

`func (o *VpnPath) GetBfdProfile() VpnPathBfdProfile`

GetBfdProfile returns the BfdProfile field if non-nil, zero value otherwise.

### GetBfdProfileOk

`func (o *VpnPath) GetBfdProfileOk() (*VpnPathBfdProfile, bool)`

GetBfdProfileOk returns a tuple with the BfdProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdProfile

`func (o *VpnPath) SetBfdProfile(v VpnPathBfdProfile)`

SetBfdProfile sets BfdProfile field to given value.

### HasBfdProfile

`func (o *VpnPath) HasBfdProfile() bool`

HasBfdProfile returns a boolean if a field has been set.

### GetIp

`func (o *VpnPath) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *VpnPath) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *VpnPath) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *VpnPath) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPod

`func (o *VpnPath) GetPod() int32`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *VpnPath) GetPodOk() (*int32, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *VpnPath) SetPod(v int32)`

SetPod sets Pod field to given value.

### HasPod

`func (o *VpnPath) HasPod() bool`

HasPod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


