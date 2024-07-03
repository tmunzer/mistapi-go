# UtilsServicePing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] [default to 10]
**Host** | **string** |  | 
**Service** | **string** | ping packet takes the same path as the service | 
**Size** | Pointer to **int32** |  | [optional] [default to 56]
**Tenant** | Pointer to **string** | tenant context in which the packet is sent | [optional] 

## Methods

### NewUtilsServicePing

`func NewUtilsServicePing(host string, service string, ) *UtilsServicePing`

NewUtilsServicePing instantiates a new UtilsServicePing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsServicePingWithDefaults

`func NewUtilsServicePingWithDefaults() *UtilsServicePing`

NewUtilsServicePingWithDefaults instantiates a new UtilsServicePing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UtilsServicePing) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UtilsServicePing) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UtilsServicePing) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UtilsServicePing) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHost

`func (o *UtilsServicePing) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UtilsServicePing) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UtilsServicePing) SetHost(v string)`

SetHost sets Host field to given value.


### GetService

`func (o *UtilsServicePing) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *UtilsServicePing) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *UtilsServicePing) SetService(v string)`

SetService sets Service field to given value.


### GetSize

`func (o *UtilsServicePing) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UtilsServicePing) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UtilsServicePing) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *UtilsServicePing) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTenant

`func (o *UtilsServicePing) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *UtilsServicePing) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *UtilsServicePing) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *UtilsServicePing) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


