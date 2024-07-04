# VrfInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraRoutes** | Pointer to [**map[string]VrfExtraRoute**](VrfExtraRoute.md) | Property key is the destination CIDR (e.g. \&quot;10.0.0.0/8\&quot;) | [optional] 
**Networks** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVrfInstance

`func NewVrfInstance() *VrfInstance`

NewVrfInstance instantiates a new VrfInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfInstanceWithDefaults

`func NewVrfInstanceWithDefaults() *VrfInstance`

NewVrfInstanceWithDefaults instantiates a new VrfInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraRoutes

`func (o *VrfInstance) GetExtraRoutes() map[string]VrfExtraRoute`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *VrfInstance) GetExtraRoutesOk() (*map[string]VrfExtraRoute, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *VrfInstance) SetExtraRoutes(v map[string]VrfExtraRoute)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *VrfInstance) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetNetworks

`func (o *VrfInstance) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VrfInstance) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VrfInstance) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VrfInstance) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


