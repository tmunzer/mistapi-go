# NetworkInternetAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateSimpleServicePolicy** | Pointer to **bool** |  | [optional] [default to false]
**DestinationNat** | Pointer to [**map[string]NetworkDestinationNatProperty**](NetworkDestinationNatProperty.md) | Property key may be an IP/Port (i.e. \&quot;63.16.0.3:443\&quot;), or a port (i.e. \&quot;:2222\&quot;) | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Restricted** | Pointer to **bool** | by default, all access is allowed, to only allow certain traffic, make &#x60;restricted&#x60;&#x3D;&#x60;true&#x60; and define service_policies | [optional] [default to false]
**StaticNat** | Pointer to [**map[string]NetworkStaticNatProperty**](NetworkStaticNatProperty.md) | Property key may be an IP Address (i.e. \&quot;172.16.0.1\&quot;), and IP Address and Port (i.e. \&quot;172.16.0.1:8443\&quot;) or a CIDR (i.e. \&quot;172.16.0.12/20\&quot;) | [optional] 

## Methods

### NewNetworkInternetAccess

`func NewNetworkInternetAccess() *NetworkInternetAccess`

NewNetworkInternetAccess instantiates a new NetworkInternetAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInternetAccessWithDefaults

`func NewNetworkInternetAccessWithDefaults() *NetworkInternetAccess`

NewNetworkInternetAccessWithDefaults instantiates a new NetworkInternetAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateSimpleServicePolicy

`func (o *NetworkInternetAccess) GetCreateSimpleServicePolicy() bool`

GetCreateSimpleServicePolicy returns the CreateSimpleServicePolicy field if non-nil, zero value otherwise.

### GetCreateSimpleServicePolicyOk

`func (o *NetworkInternetAccess) GetCreateSimpleServicePolicyOk() (*bool, bool)`

GetCreateSimpleServicePolicyOk returns a tuple with the CreateSimpleServicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSimpleServicePolicy

`func (o *NetworkInternetAccess) SetCreateSimpleServicePolicy(v bool)`

SetCreateSimpleServicePolicy sets CreateSimpleServicePolicy field to given value.

### HasCreateSimpleServicePolicy

`func (o *NetworkInternetAccess) HasCreateSimpleServicePolicy() bool`

HasCreateSimpleServicePolicy returns a boolean if a field has been set.

### GetDestinationNat

`func (o *NetworkInternetAccess) GetDestinationNat() map[string]NetworkDestinationNatProperty`

GetDestinationNat returns the DestinationNat field if non-nil, zero value otherwise.

### GetDestinationNatOk

`func (o *NetworkInternetAccess) GetDestinationNatOk() (*map[string]NetworkDestinationNatProperty, bool)`

GetDestinationNatOk returns a tuple with the DestinationNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNat

`func (o *NetworkInternetAccess) SetDestinationNat(v map[string]NetworkDestinationNatProperty)`

SetDestinationNat sets DestinationNat field to given value.

### HasDestinationNat

`func (o *NetworkInternetAccess) HasDestinationNat() bool`

HasDestinationNat returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkInternetAccess) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkInternetAccess) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkInternetAccess) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkInternetAccess) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRestricted

`func (o *NetworkInternetAccess) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *NetworkInternetAccess) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *NetworkInternetAccess) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *NetworkInternetAccess) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetStaticNat

`func (o *NetworkInternetAccess) GetStaticNat() map[string]NetworkStaticNatProperty`

GetStaticNat returns the StaticNat field if non-nil, zero value otherwise.

### GetStaticNatOk

`func (o *NetworkInternetAccess) GetStaticNatOk() (*map[string]NetworkStaticNatProperty, bool)`

GetStaticNatOk returns a tuple with the StaticNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticNat

`func (o *NetworkInternetAccess) SetStaticNat(v map[string]NetworkStaticNatProperty)`

SetStaticNat sets StaticNat field to given value.

### HasStaticNat

`func (o *NetworkInternetAccess) HasStaticNat() bool`

HasStaticNat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


