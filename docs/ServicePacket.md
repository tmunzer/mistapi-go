# ServicePacket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceData** | Pointer to **string** | ata from service data | [optional] 
**ServiceUuid** | Pointer to **string** | UUID from service data | [optional] 

## Methods

### NewServicePacket

`func NewServicePacket() *ServicePacket`

NewServicePacket instantiates a new ServicePacket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePacketWithDefaults

`func NewServicePacketWithDefaults() *ServicePacket`

NewServicePacketWithDefaults instantiates a new ServicePacket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceData

`func (o *ServicePacket) GetServiceData() string`

GetServiceData returns the ServiceData field if non-nil, zero value otherwise.

### GetServiceDataOk

`func (o *ServicePacket) GetServiceDataOk() (*string, bool)`

GetServiceDataOk returns a tuple with the ServiceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceData

`func (o *ServicePacket) SetServiceData(v string)`

SetServiceData sets ServiceData field to given value.

### HasServiceData

`func (o *ServicePacket) HasServiceData() bool`

HasServiceData returns a boolean if a field has been set.

### GetServiceUuid

`func (o *ServicePacket) GetServiceUuid() string`

GetServiceUuid returns the ServiceUuid field if non-nil, zero value otherwise.

### GetServiceUuidOk

`func (o *ServicePacket) GetServiceUuidOk() (*string, bool)`

GetServiceUuidOk returns a tuple with the ServiceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUuid

`func (o *ServicePacket) SetServiceUuid(v string)`

SetServiceUuid sets ServiceUuid field to given value.

### HasServiceUuid

`func (o *ServicePacket) HasServiceUuid() bool`

HasServiceUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


