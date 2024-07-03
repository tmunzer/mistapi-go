# RadsecServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 

## Methods

### NewRadsecServer

`func NewRadsecServer() *RadsecServer`

NewRadsecServer instantiates a new RadsecServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadsecServerWithDefaults

`func NewRadsecServerWithDefaults() *RadsecServer`

NewRadsecServerWithDefaults instantiates a new RadsecServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RadsecServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RadsecServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RadsecServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RadsecServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *RadsecServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RadsecServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RadsecServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RadsecServer) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


