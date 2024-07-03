# UtilsReleaseDhcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | Pointer to [**HaClusterNodeEnum**](HaClusterNodeEnum.md) |  | [optional] 
**Port** | **string** | The nework interface on which to release the current DHCP release | 

## Methods

### NewUtilsReleaseDhcp

`func NewUtilsReleaseDhcp(port string, ) *UtilsReleaseDhcp`

NewUtilsReleaseDhcp instantiates a new UtilsReleaseDhcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsReleaseDhcpWithDefaults

`func NewUtilsReleaseDhcpWithDefaults() *UtilsReleaseDhcp`

NewUtilsReleaseDhcpWithDefaults instantiates a new UtilsReleaseDhcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *UtilsReleaseDhcp) GetNode() HaClusterNodeEnum`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *UtilsReleaseDhcp) GetNodeOk() (*HaClusterNodeEnum, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *UtilsReleaseDhcp) SetNode(v HaClusterNodeEnum)`

SetNode sets Node field to given value.

### HasNode

`func (o *UtilsReleaseDhcp) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPort

`func (o *UtilsReleaseDhcp) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UtilsReleaseDhcp) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UtilsReleaseDhcp) SetPort(v string)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


