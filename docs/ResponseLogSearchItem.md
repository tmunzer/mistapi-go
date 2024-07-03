# ResponseLogSearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminId** | **NullableString** | admin id | [readonly] 
**AdminName** | **NullableString** | name of the admin that performs the action | [readonly] 
**After** | Pointer to **map[string]interface{}** | field values after the change | [optional] [readonly] 
**Before** | Pointer to **map[string]interface{}** | field values prior to the change | [optional] [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Message** | **string** | log message | [readonly] 
**OrgId** | **string** | org id | [readonly] 
**SiteId** | **string** |  | [readonly] 
**Timestamp** | **float32** | start time, in epoch | [readonly] 

## Methods

### NewResponseLogSearchItem

`func NewResponseLogSearchItem(adminId NullableString, adminName NullableString, message string, orgId string, siteId string, timestamp float32, ) *ResponseLogSearchItem`

NewResponseLogSearchItem instantiates a new ResponseLogSearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseLogSearchItemWithDefaults

`func NewResponseLogSearchItemWithDefaults() *ResponseLogSearchItem`

NewResponseLogSearchItemWithDefaults instantiates a new ResponseLogSearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminId

`func (o *ResponseLogSearchItem) GetAdminId() string`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *ResponseLogSearchItem) GetAdminIdOk() (*string, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *ResponseLogSearchItem) SetAdminId(v string)`

SetAdminId sets AdminId field to given value.


### SetAdminIdNil

`func (o *ResponseLogSearchItem) SetAdminIdNil(b bool)`

 SetAdminIdNil sets the value for AdminId to be an explicit nil

### UnsetAdminId
`func (o *ResponseLogSearchItem) UnsetAdminId()`

UnsetAdminId ensures that no value is present for AdminId, not even an explicit nil
### GetAdminName

`func (o *ResponseLogSearchItem) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *ResponseLogSearchItem) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *ResponseLogSearchItem) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.


### SetAdminNameNil

`func (o *ResponseLogSearchItem) SetAdminNameNil(b bool)`

 SetAdminNameNil sets the value for AdminName to be an explicit nil

### UnsetAdminName
`func (o *ResponseLogSearchItem) UnsetAdminName()`

UnsetAdminName ensures that no value is present for AdminName, not even an explicit nil
### GetAfter

`func (o *ResponseLogSearchItem) GetAfter() map[string]interface{}`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ResponseLogSearchItem) GetAfterOk() (*map[string]interface{}, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ResponseLogSearchItem) SetAfter(v map[string]interface{})`

SetAfter sets After field to given value.

### HasAfter

`func (o *ResponseLogSearchItem) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *ResponseLogSearchItem) GetBefore() map[string]interface{}`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ResponseLogSearchItem) GetBeforeOk() (*map[string]interface{}, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ResponseLogSearchItem) SetBefore(v map[string]interface{})`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ResponseLogSearchItem) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetForSite

`func (o *ResponseLogSearchItem) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *ResponseLogSearchItem) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *ResponseLogSearchItem) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *ResponseLogSearchItem) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *ResponseLogSearchItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseLogSearchItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseLogSearchItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponseLogSearchItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseLogSearchItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseLogSearchItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseLogSearchItem) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrgId

`func (o *ResponseLogSearchItem) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResponseLogSearchItem) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResponseLogSearchItem) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetSiteId

`func (o *ResponseLogSearchItem) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ResponseLogSearchItem) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ResponseLogSearchItem) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetTimestamp

`func (o *ResponseLogSearchItem) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponseLogSearchItem) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponseLogSearchItem) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


