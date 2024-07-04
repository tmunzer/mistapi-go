# Msp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMist** | Pointer to **bool** |  | [optional] 
**CreatedTime** | Pointer to **int32** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** | For advanced tier (uMSPs) only | [optional] 
**ModifiedTime** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to [**MspTier**](MspTier.md) |  | [optional] [default to MSPTIER_BASE]
**Url** | Pointer to **string** | For advanced tier (uMSPs) only | [optional] 

## Methods

### NewMsp

`func NewMsp() *Msp`

NewMsp instantiates a new Msp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspWithDefaults

`func NewMspWithDefaults() *Msp`

NewMspWithDefaults instantiates a new Msp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMist

`func (o *Msp) GetAllowMist() bool`

GetAllowMist returns the AllowMist field if non-nil, zero value otherwise.

### GetAllowMistOk

`func (o *Msp) GetAllowMistOk() (*bool, bool)`

GetAllowMistOk returns a tuple with the AllowMist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMist

`func (o *Msp) SetAllowMist(v bool)`

SetAllowMist sets AllowMist field to given value.

### HasAllowMist

`func (o *Msp) HasAllowMist() bool`

HasAllowMist returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Msp) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Msp) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Msp) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Msp) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetId

`func (o *Msp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Msp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Msp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Msp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogoUrl

`func (o *Msp) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Msp) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Msp) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *Msp) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Msp) GetModifiedTime() int32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Msp) GetModifiedTimeOk() (*int32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Msp) SetModifiedTime(v int32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Msp) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Msp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Msp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Msp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Msp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTier

`func (o *Msp) GetTier() MspTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Msp) GetTierOk() (*MspTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Msp) SetTier(v MspTier)`

SetTier sets Tier field to given value.

### HasTier

`func (o *Msp) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUrl

`func (o *Msp) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Msp) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Msp) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Msp) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


