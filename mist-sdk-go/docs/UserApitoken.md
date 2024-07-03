# UserApitoken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **int32** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Key** | Pointer to **string** |  | [optional] [readonly] 
**LastUsed** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**Name** | Pointer to **string** | name of the token | [optional] 

## Methods

### NewUserApitoken

`func NewUserApitoken() *UserApitoken`

NewUserApitoken instantiates a new UserApitoken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserApitokenWithDefaults

`func NewUserApitokenWithDefaults() *UserApitoken`

NewUserApitokenWithDefaults instantiates a new UserApitoken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *UserApitoken) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *UserApitoken) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *UserApitoken) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *UserApitoken) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetId

`func (o *UserApitoken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserApitoken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserApitoken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserApitoken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *UserApitoken) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserApitoken) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserApitoken) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UserApitoken) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastUsed

`func (o *UserApitoken) GetLastUsed() int32`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *UserApitoken) GetLastUsedOk() (*int32, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *UserApitoken) SetLastUsed(v int32)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *UserApitoken) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsedNil

`func (o *UserApitoken) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *UserApitoken) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetName

`func (o *UserApitoken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserApitoken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserApitoken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserApitoken) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


