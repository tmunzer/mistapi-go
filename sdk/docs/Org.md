# Org

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmtemplateId** | Pointer to **NullableString** |  | [optional] 
**AllowMist** | Pointer to **bool** |  | [optional] [default to true]
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**MspId** | Pointer to **string** |  | [optional] 
**MspLogoUrl** | Pointer to **string** | logo uploaded by the MSP with advanced tier, only present if provided | [optional] [readonly] 
**MspName** | Pointer to **string** | name of the msp the org belongs to | [optional] [readonly] 
**Name** | **string** |  | 
**OrggroupIds** | Pointer to **[]string** |  | [optional] 
**SessionExpiry** | Pointer to **float32** |  | [optional] [default to 1440]

## Methods

### NewOrg

`func NewOrg(name string, ) *Org`

NewOrg instantiates a new Org object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWithDefaults

`func NewOrgWithDefaults() *Org`

NewOrgWithDefaults instantiates a new Org object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmtemplateId

`func (o *Org) GetAlarmtemplateId() string`

GetAlarmtemplateId returns the AlarmtemplateId field if non-nil, zero value otherwise.

### GetAlarmtemplateIdOk

`func (o *Org) GetAlarmtemplateIdOk() (*string, bool)`

GetAlarmtemplateIdOk returns a tuple with the AlarmtemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmtemplateId

`func (o *Org) SetAlarmtemplateId(v string)`

SetAlarmtemplateId sets AlarmtemplateId field to given value.

### HasAlarmtemplateId

`func (o *Org) HasAlarmtemplateId() bool`

HasAlarmtemplateId returns a boolean if a field has been set.

### SetAlarmtemplateIdNil

`func (o *Org) SetAlarmtemplateIdNil(b bool)`

 SetAlarmtemplateIdNil sets the value for AlarmtemplateId to be an explicit nil

### UnsetAlarmtemplateId
`func (o *Org) UnsetAlarmtemplateId()`

UnsetAlarmtemplateId ensures that no value is present for AlarmtemplateId, not even an explicit nil
### GetAllowMist

`func (o *Org) GetAllowMist() bool`

GetAllowMist returns the AllowMist field if non-nil, zero value otherwise.

### GetAllowMistOk

`func (o *Org) GetAllowMistOk() (*bool, bool)`

GetAllowMistOk returns a tuple with the AllowMist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMist

`func (o *Org) SetAllowMist(v bool)`

SetAllowMist sets AllowMist field to given value.

### HasAllowMist

`func (o *Org) HasAllowMist() bool`

HasAllowMist returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Org) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Org) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Org) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Org) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetId

`func (o *Org) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Org) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Org) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Org) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Org) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Org) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Org) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Org) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMspId

`func (o *Org) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *Org) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *Org) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *Org) HasMspId() bool`

HasMspId returns a boolean if a field has been set.

### GetMspLogoUrl

`func (o *Org) GetMspLogoUrl() string`

GetMspLogoUrl returns the MspLogoUrl field if non-nil, zero value otherwise.

### GetMspLogoUrlOk

`func (o *Org) GetMspLogoUrlOk() (*string, bool)`

GetMspLogoUrlOk returns a tuple with the MspLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspLogoUrl

`func (o *Org) SetMspLogoUrl(v string)`

SetMspLogoUrl sets MspLogoUrl field to given value.

### HasMspLogoUrl

`func (o *Org) HasMspLogoUrl() bool`

HasMspLogoUrl returns a boolean if a field has been set.

### GetMspName

`func (o *Org) GetMspName() string`

GetMspName returns the MspName field if non-nil, zero value otherwise.

### GetMspNameOk

`func (o *Org) GetMspNameOk() (*string, bool)`

GetMspNameOk returns a tuple with the MspName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspName

`func (o *Org) SetMspName(v string)`

SetMspName sets MspName field to given value.

### HasMspName

`func (o *Org) HasMspName() bool`

HasMspName returns a boolean if a field has been set.

### GetName

`func (o *Org) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Org) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Org) SetName(v string)`

SetName sets Name field to given value.


### GetOrggroupIds

`func (o *Org) GetOrggroupIds() []string`

GetOrggroupIds returns the OrggroupIds field if non-nil, zero value otherwise.

### GetOrggroupIdsOk

`func (o *Org) GetOrggroupIdsOk() (*[]string, bool)`

GetOrggroupIdsOk returns a tuple with the OrggroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrggroupIds

`func (o *Org) SetOrggroupIds(v []string)`

SetOrggroupIds sets OrggroupIds field to given value.

### HasOrggroupIds

`func (o *Org) HasOrggroupIds() bool`

HasOrggroupIds returns a boolean if a field has been set.

### GetSessionExpiry

`func (o *Org) GetSessionExpiry() float32`

GetSessionExpiry returns the SessionExpiry field if non-nil, zero value otherwise.

### GetSessionExpiryOk

`func (o *Org) GetSessionExpiryOk() (*float32, bool)`

GetSessionExpiryOk returns a tuple with the SessionExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionExpiry

`func (o *Org) SetSessionExpiry(v float32)`

SetSessionExpiry sets SessionExpiry field to given value.

### HasSessionExpiry

`func (o *Org) HasSessionExpiry() bool`

HasSessionExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


