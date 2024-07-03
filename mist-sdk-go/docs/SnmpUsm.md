# SnmpUsm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineId** | Pointer to **string** | required only if &#x60;engine_type&#x60;&#x3D;&#x3D;&#x60;remote_engine&#x60; | [optional] 
**EngineType** | Pointer to [**SnmpUsmEngineType**](SnmpUsmEngineType.md) |  | [optional] 
**Users** | Pointer to [**[]SnmpUsmpUser**](SnmpUsmpUser.md) |  | [optional] 

## Methods

### NewSnmpUsm

`func NewSnmpUsm() *SnmpUsm`

NewSnmpUsm instantiates a new SnmpUsm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpUsmWithDefaults

`func NewSnmpUsmWithDefaults() *SnmpUsm`

NewSnmpUsmWithDefaults instantiates a new SnmpUsm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineId

`func (o *SnmpUsm) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *SnmpUsm) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *SnmpUsm) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *SnmpUsm) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetEngineType

`func (o *SnmpUsm) GetEngineType() SnmpUsmEngineType`

GetEngineType returns the EngineType field if non-nil, zero value otherwise.

### GetEngineTypeOk

`func (o *SnmpUsm) GetEngineTypeOk() (*SnmpUsmEngineType, bool)`

GetEngineTypeOk returns a tuple with the EngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineType

`func (o *SnmpUsm) SetEngineType(v SnmpUsmEngineType)`

SetEngineType sets EngineType field to given value.

### HasEngineType

`func (o *SnmpUsm) HasEngineType() bool`

HasEngineType returns a boolean if a field has been set.

### GetUsers

`func (o *SnmpUsm) GetUsers() []SnmpUsmpUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SnmpUsm) GetUsersOk() (*[]SnmpUsmpUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SnmpUsm) SetUsers(v []SnmpUsmpUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *SnmpUsm) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


