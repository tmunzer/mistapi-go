# GatewayPortVpnPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BfdProfile** | Pointer to [**GatewayPortVpnPathBfdProfile**](GatewayPortVpnPathBfdProfile.md) |  | [optional] [default to GATEWAYPORTVPNPATHBFDPROFILE_BROADBAND]
**BfdUseTunnelMode** | Pointer to **bool** | whether to use tunnel mode. SSR only | [optional] [default to false]
**Preference** | Pointer to **int32** | for a given VPN, when &#x60;path_selection.strategy&#x60;&#x3D;&#x3D;&#x60;simple&#x60;, the preference for a path (lower is preferred) | [optional] 
**Role** | Pointer to [**GatewayPortVpnPathRole**](GatewayPortVpnPathRole.md) |  | [optional] [default to GATEWAYPORTVPNPATHROLE_SPOKE]
**TrafficShaping** | Pointer to [**GatewayTrafficShaping**](GatewayTrafficShaping.md) |  | [optional] 

## Methods

### NewGatewayPortVpnPath

`func NewGatewayPortVpnPath() *GatewayPortVpnPath`

NewGatewayPortVpnPath instantiates a new GatewayPortVpnPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPortVpnPathWithDefaults

`func NewGatewayPortVpnPathWithDefaults() *GatewayPortVpnPath`

NewGatewayPortVpnPathWithDefaults instantiates a new GatewayPortVpnPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBfdProfile

`func (o *GatewayPortVpnPath) GetBfdProfile() GatewayPortVpnPathBfdProfile`

GetBfdProfile returns the BfdProfile field if non-nil, zero value otherwise.

### GetBfdProfileOk

`func (o *GatewayPortVpnPath) GetBfdProfileOk() (*GatewayPortVpnPathBfdProfile, bool)`

GetBfdProfileOk returns a tuple with the BfdProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdProfile

`func (o *GatewayPortVpnPath) SetBfdProfile(v GatewayPortVpnPathBfdProfile)`

SetBfdProfile sets BfdProfile field to given value.

### HasBfdProfile

`func (o *GatewayPortVpnPath) HasBfdProfile() bool`

HasBfdProfile returns a boolean if a field has been set.

### GetBfdUseTunnelMode

`func (o *GatewayPortVpnPath) GetBfdUseTunnelMode() bool`

GetBfdUseTunnelMode returns the BfdUseTunnelMode field if non-nil, zero value otherwise.

### GetBfdUseTunnelModeOk

`func (o *GatewayPortVpnPath) GetBfdUseTunnelModeOk() (*bool, bool)`

GetBfdUseTunnelModeOk returns a tuple with the BfdUseTunnelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdUseTunnelMode

`func (o *GatewayPortVpnPath) SetBfdUseTunnelMode(v bool)`

SetBfdUseTunnelMode sets BfdUseTunnelMode field to given value.

### HasBfdUseTunnelMode

`func (o *GatewayPortVpnPath) HasBfdUseTunnelMode() bool`

HasBfdUseTunnelMode returns a boolean if a field has been set.

### GetPreference

`func (o *GatewayPortVpnPath) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *GatewayPortVpnPath) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *GatewayPortVpnPath) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *GatewayPortVpnPath) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetRole

`func (o *GatewayPortVpnPath) GetRole() GatewayPortVpnPathRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GatewayPortVpnPath) GetRoleOk() (*GatewayPortVpnPathRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GatewayPortVpnPath) SetRole(v GatewayPortVpnPathRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *GatewayPortVpnPath) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTrafficShaping

`func (o *GatewayPortVpnPath) GetTrafficShaping() GatewayTrafficShaping`

GetTrafficShaping returns the TrafficShaping field if non-nil, zero value otherwise.

### GetTrafficShapingOk

`func (o *GatewayPortVpnPath) GetTrafficShapingOk() (*GatewayTrafficShaping, bool)`

GetTrafficShapingOk returns a tuple with the TrafficShaping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficShaping

`func (o *GatewayPortVpnPath) SetTrafficShaping(v GatewayTrafficShaping)`

SetTrafficShaping sets TrafficShaping field to given value.

### HasTrafficShaping

`func (o *GatewayPortVpnPath) HasTrafficShaping() bool`

HasTrafficShaping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


