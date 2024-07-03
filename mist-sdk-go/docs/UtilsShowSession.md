# UtilsShowSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | Pointer to [**HaClusterNodeEnum**](HaClusterNodeEnum.md) |  | [optional] 
**ServiceName** | Pointer to **string** | The exact service name for which to display the active sessions | [optional] 

## Methods

### NewUtilsShowSession

`func NewUtilsShowSession() *UtilsShowSession`

NewUtilsShowSession instantiates a new UtilsShowSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsShowSessionWithDefaults

`func NewUtilsShowSessionWithDefaults() *UtilsShowSession`

NewUtilsShowSessionWithDefaults instantiates a new UtilsShowSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *UtilsShowSession) GetNode() HaClusterNodeEnum`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *UtilsShowSession) GetNodeOk() (*HaClusterNodeEnum, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *UtilsShowSession) SetNode(v HaClusterNodeEnum)`

SetNode sets Node field to given value.

### HasNode

`func (o *UtilsShowSession) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetServiceName

`func (o *UtilsShowSession) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *UtilsShowSession) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *UtilsShowSession) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *UtilsShowSession) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


