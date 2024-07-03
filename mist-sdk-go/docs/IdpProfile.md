# IdpProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseProfile** | Pointer to [**IdpProfileBaseProfile**](IdpProfileBaseProfile.md) |  | [optional] 
**CreatedTime** | Pointer to **int32** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedTime** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Overwrites** | Pointer to [**[]IdpProfileOverwrite**](IdpProfileOverwrite.md) |  | [optional] 

## Methods

### NewIdpProfile

`func NewIdpProfile() *IdpProfile`

NewIdpProfile instantiates a new IdpProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpProfileWithDefaults

`func NewIdpProfileWithDefaults() *IdpProfile`

NewIdpProfileWithDefaults instantiates a new IdpProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseProfile

`func (o *IdpProfile) GetBaseProfile() IdpProfileBaseProfile`

GetBaseProfile returns the BaseProfile field if non-nil, zero value otherwise.

### GetBaseProfileOk

`func (o *IdpProfile) GetBaseProfileOk() (*IdpProfileBaseProfile, bool)`

GetBaseProfileOk returns a tuple with the BaseProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseProfile

`func (o *IdpProfile) SetBaseProfile(v IdpProfileBaseProfile)`

SetBaseProfile sets BaseProfile field to given value.

### HasBaseProfile

`func (o *IdpProfile) HasBaseProfile() bool`

HasBaseProfile returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IdpProfile) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IdpProfile) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IdpProfile) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IdpProfile) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetId

`func (o *IdpProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdpProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdpProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdpProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *IdpProfile) GetModifiedTime() int32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *IdpProfile) GetModifiedTimeOk() (*int32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *IdpProfile) SetModifiedTime(v int32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *IdpProfile) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *IdpProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdpProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverwrites

`func (o *IdpProfile) GetOverwrites() []IdpProfileOverwrite`

GetOverwrites returns the Overwrites field if non-nil, zero value otherwise.

### GetOverwritesOk

`func (o *IdpProfile) GetOverwritesOk() (*[]IdpProfileOverwrite, bool)`

GetOverwritesOk returns a tuple with the Overwrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrites

`func (o *IdpProfile) SetOverwrites(v []IdpProfileOverwrite)`

SetOverwrites sets Overwrites field to given value.

### HasOverwrites

`func (o *IdpProfile) HasOverwrites() bool`

HasOverwrites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


