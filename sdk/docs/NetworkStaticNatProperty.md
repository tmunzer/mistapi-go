# NetworkStaticNatProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalIp** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**WanName** | Pointer to **string** | If not set, we configure the nat policies against all WAN ports for simplicity | [optional] 

## Methods

### NewNetworkStaticNatProperty

`func NewNetworkStaticNatProperty() *NetworkStaticNatProperty`

NewNetworkStaticNatProperty instantiates a new NetworkStaticNatProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkStaticNatPropertyWithDefaults

`func NewNetworkStaticNatPropertyWithDefaults() *NetworkStaticNatProperty`

NewNetworkStaticNatPropertyWithDefaults instantiates a new NetworkStaticNatProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalIp

`func (o *NetworkStaticNatProperty) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *NetworkStaticNatProperty) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *NetworkStaticNatProperty) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *NetworkStaticNatProperty) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetName

`func (o *NetworkStaticNatProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkStaticNatProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkStaticNatProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkStaticNatProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWanName

`func (o *NetworkStaticNatProperty) GetWanName() string`

GetWanName returns the WanName field if non-nil, zero value otherwise.

### GetWanNameOk

`func (o *NetworkStaticNatProperty) GetWanNameOk() (*string, bool)`

GetWanNameOk returns a tuple with the WanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanName

`func (o *NetworkStaticNatProperty) SetWanName(v string)`

SetWanName sets WanName field to given value.

### HasWanName

`func (o *NetworkStaticNatProperty) HasWanName() bool`

HasWanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


