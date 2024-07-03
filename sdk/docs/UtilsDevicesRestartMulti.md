# UtilsDevicesRestartMulti

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | **[]string** |  | 
**Node** | Pointer to **string** | only for SSR: if node is not present, both nodes are restarted for other devices: node should not be present | [optional] 

## Methods

### NewUtilsDevicesRestartMulti

`func NewUtilsDevicesRestartMulti(deviceIds []string, ) *UtilsDevicesRestartMulti`

NewUtilsDevicesRestartMulti instantiates a new UtilsDevicesRestartMulti object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsDevicesRestartMultiWithDefaults

`func NewUtilsDevicesRestartMultiWithDefaults() *UtilsDevicesRestartMulti`

NewUtilsDevicesRestartMultiWithDefaults instantiates a new UtilsDevicesRestartMulti object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *UtilsDevicesRestartMulti) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *UtilsDevicesRestartMulti) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *UtilsDevicesRestartMulti) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.


### GetNode

`func (o *UtilsDevicesRestartMulti) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *UtilsDevicesRestartMulti) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *UtilsDevicesRestartMulti) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *UtilsDevicesRestartMulti) HasNode() bool`

HasNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


