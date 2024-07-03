# AuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminId** | **string** |  | 
**AdminName** | **string** |  | 
**After** | Pointer to **map[string]interface{}** | field values after the change | [optional] 
**Before** | Pointer to **map[string]interface{}** | field values prior to the change | [optional] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | **string** |  | 
**Message** | **string** |  | 
**OrgId** | **string** |  | 
**SiteId** | **string** |  | [readonly] 
**Timestamp** | **float32** |  | 

## Methods

### NewAuditLog

`func NewAuditLog(adminId string, adminName string, id string, message string, orgId string, siteId string, timestamp float32, ) *AuditLog`

NewAuditLog instantiates a new AuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogWithDefaults

`func NewAuditLogWithDefaults() *AuditLog`

NewAuditLogWithDefaults instantiates a new AuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminId

`func (o *AuditLog) GetAdminId() string`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *AuditLog) GetAdminIdOk() (*string, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *AuditLog) SetAdminId(v string)`

SetAdminId sets AdminId field to given value.


### GetAdminName

`func (o *AuditLog) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *AuditLog) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *AuditLog) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.


### GetAfter

`func (o *AuditLog) GetAfter() map[string]interface{}`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *AuditLog) GetAfterOk() (*map[string]interface{}, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *AuditLog) SetAfter(v map[string]interface{})`

SetAfter sets After field to given value.

### HasAfter

`func (o *AuditLog) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *AuditLog) GetBefore() map[string]interface{}`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *AuditLog) GetBeforeOk() (*map[string]interface{}, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *AuditLog) SetBefore(v map[string]interface{})`

SetBefore sets Before field to given value.

### HasBefore

`func (o *AuditLog) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetForSite

`func (o *AuditLog) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *AuditLog) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *AuditLog) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *AuditLog) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *AuditLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLog) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *AuditLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuditLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuditLog) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrgId

`func (o *AuditLog) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AuditLog) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AuditLog) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetSiteId

`func (o *AuditLog) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AuditLog) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AuditLog) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetTimestamp

`func (o *AuditLog) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditLog) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditLog) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


