# UtilsShowForwardingTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | Pointer to [**HaClusterNodeEnum**](HaClusterNodeEnum.md) |  | [optional] 
**Prefix** | Pointer to **string** | IP Prefix | [optional] 
**ServiceIp** | Pointer to **string** | only supported with SSR | [optional] 
**ServiceName** | Pointer to **string** | only supported with SSR | [optional] 
**ServicePort** | Pointer to **int32** | only supported with SSR | [optional] 
**ServiceProtocol** | Pointer to **string** | only supported with SSR | [optional] 
**ServiceTenant** | Pointer to **string** | only supported with SSR | [optional] 
**Vrf** | Pointer to **string** | VRF Name | [optional] 

## Methods

### NewUtilsShowForwardingTable

`func NewUtilsShowForwardingTable() *UtilsShowForwardingTable`

NewUtilsShowForwardingTable instantiates a new UtilsShowForwardingTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsShowForwardingTableWithDefaults

`func NewUtilsShowForwardingTableWithDefaults() *UtilsShowForwardingTable`

NewUtilsShowForwardingTableWithDefaults instantiates a new UtilsShowForwardingTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *UtilsShowForwardingTable) GetNode() HaClusterNodeEnum`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *UtilsShowForwardingTable) GetNodeOk() (*HaClusterNodeEnum, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *UtilsShowForwardingTable) SetNode(v HaClusterNodeEnum)`

SetNode sets Node field to given value.

### HasNode

`func (o *UtilsShowForwardingTable) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPrefix

`func (o *UtilsShowForwardingTable) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *UtilsShowForwardingTable) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *UtilsShowForwardingTable) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *UtilsShowForwardingTable) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetServiceIp

`func (o *UtilsShowForwardingTable) GetServiceIp() string`

GetServiceIp returns the ServiceIp field if non-nil, zero value otherwise.

### GetServiceIpOk

`func (o *UtilsShowForwardingTable) GetServiceIpOk() (*string, bool)`

GetServiceIpOk returns a tuple with the ServiceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIp

`func (o *UtilsShowForwardingTable) SetServiceIp(v string)`

SetServiceIp sets ServiceIp field to given value.

### HasServiceIp

`func (o *UtilsShowForwardingTable) HasServiceIp() bool`

HasServiceIp returns a boolean if a field has been set.

### GetServiceName

`func (o *UtilsShowForwardingTable) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *UtilsShowForwardingTable) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *UtilsShowForwardingTable) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *UtilsShowForwardingTable) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServicePort

`func (o *UtilsShowForwardingTable) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *UtilsShowForwardingTable) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *UtilsShowForwardingTable) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *UtilsShowForwardingTable) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceProtocol

`func (o *UtilsShowForwardingTable) GetServiceProtocol() string`

GetServiceProtocol returns the ServiceProtocol field if non-nil, zero value otherwise.

### GetServiceProtocolOk

`func (o *UtilsShowForwardingTable) GetServiceProtocolOk() (*string, bool)`

GetServiceProtocolOk returns a tuple with the ServiceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProtocol

`func (o *UtilsShowForwardingTable) SetServiceProtocol(v string)`

SetServiceProtocol sets ServiceProtocol field to given value.

### HasServiceProtocol

`func (o *UtilsShowForwardingTable) HasServiceProtocol() bool`

HasServiceProtocol returns a boolean if a field has been set.

### GetServiceTenant

`func (o *UtilsShowForwardingTable) GetServiceTenant() string`

GetServiceTenant returns the ServiceTenant field if non-nil, zero value otherwise.

### GetServiceTenantOk

`func (o *UtilsShowForwardingTable) GetServiceTenantOk() (*string, bool)`

GetServiceTenantOk returns a tuple with the ServiceTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTenant

`func (o *UtilsShowForwardingTable) SetServiceTenant(v string)`

SetServiceTenant sets ServiceTenant field to given value.

### HasServiceTenant

`func (o *UtilsShowForwardingTable) HasServiceTenant() bool`

HasServiceTenant returns a boolean if a field has been set.

### GetVrf

`func (o *UtilsShowForwardingTable) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *UtilsShowForwardingTable) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *UtilsShowForwardingTable) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *UtilsShowForwardingTable) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


