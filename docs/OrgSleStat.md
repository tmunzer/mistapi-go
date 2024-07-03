# OrgSleStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**UserMinutes** | Pointer to [**OrgSleStatUserMinutes**](OrgSleStatUserMinutes.md) |  | [optional] 

## Methods

### NewOrgSleStat

`func NewOrgSleStat(path string, ) *OrgSleStat`

NewOrgSleStat instantiates a new OrgSleStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSleStatWithDefaults

`func NewOrgSleStatWithDefaults() *OrgSleStat`

NewOrgSleStatWithDefaults instantiates a new OrgSleStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *OrgSleStat) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OrgSleStat) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OrgSleStat) SetPath(v string)`

SetPath sets Path field to given value.


### GetUserMinutes

`func (o *OrgSleStat) GetUserMinutes() OrgSleStatUserMinutes`

GetUserMinutes returns the UserMinutes field if non-nil, zero value otherwise.

### GetUserMinutesOk

`func (o *OrgSleStat) GetUserMinutesOk() (*OrgSleStatUserMinutes, bool)`

GetUserMinutesOk returns a tuple with the UserMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMinutes

`func (o *OrgSleStat) SetUserMinutes(v OrgSleStatUserMinutes)`

SetUserMinutes sets UserMinutes field to given value.

### HasUserMinutes

`func (o *OrgSleStat) HasUserMinutes() bool`

HasUserMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


