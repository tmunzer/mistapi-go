# UtilsDevicesRestart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to **string** | optional for VC member | [optional] 
**Node** | Pointer to **string** | only for SSR: if node is not present, both nodes are restarted for other devices: node should not be present | [optional] 

## Methods

### NewUtilsDevicesRestart

`func NewUtilsDevicesRestart() *UtilsDevicesRestart`

NewUtilsDevicesRestart instantiates a new UtilsDevicesRestart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsDevicesRestartWithDefaults

`func NewUtilsDevicesRestartWithDefaults() *UtilsDevicesRestart`

NewUtilsDevicesRestartWithDefaults instantiates a new UtilsDevicesRestart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *UtilsDevicesRestart) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *UtilsDevicesRestart) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *UtilsDevicesRestart) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *UtilsDevicesRestart) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetNode

`func (o *UtilsDevicesRestart) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *UtilsDevicesRestart) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *UtilsDevicesRestart) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *UtilsDevicesRestart) HasNode() bool`

HasNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


