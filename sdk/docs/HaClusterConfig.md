# HaClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableAutoConfig** | Pointer to **bool** | if the device is claimed | [optional] 
**Managed** | Pointer to **bool** | if the device is adopted | [optional] 
**Nodes** | Pointer to [**[]HaClusterConfigNode**](HaClusterConfigNode.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 

## Methods

### NewHaClusterConfig

`func NewHaClusterConfig() *HaClusterConfig`

NewHaClusterConfig instantiates a new HaClusterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHaClusterConfigWithDefaults

`func NewHaClusterConfigWithDefaults() *HaClusterConfig`

NewHaClusterConfigWithDefaults instantiates a new HaClusterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableAutoConfig

`func (o *HaClusterConfig) GetDisableAutoConfig() bool`

GetDisableAutoConfig returns the DisableAutoConfig field if non-nil, zero value otherwise.

### GetDisableAutoConfigOk

`func (o *HaClusterConfig) GetDisableAutoConfigOk() (*bool, bool)`

GetDisableAutoConfigOk returns a tuple with the DisableAutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoConfig

`func (o *HaClusterConfig) SetDisableAutoConfig(v bool)`

SetDisableAutoConfig sets DisableAutoConfig field to given value.

### HasDisableAutoConfig

`func (o *HaClusterConfig) HasDisableAutoConfig() bool`

HasDisableAutoConfig returns a boolean if a field has been set.

### GetManaged

`func (o *HaClusterConfig) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *HaClusterConfig) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *HaClusterConfig) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *HaClusterConfig) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetNodes

`func (o *HaClusterConfig) GetNodes() []HaClusterConfigNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *HaClusterConfig) GetNodesOk() (*[]HaClusterConfigNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *HaClusterConfig) SetNodes(v []HaClusterConfigNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *HaClusterConfig) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetSiteId

`func (o *HaClusterConfig) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *HaClusterConfig) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *HaClusterConfig) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *HaClusterConfig) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


