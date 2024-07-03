# ResponseOrgSuppressAlarmItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | duration, in seconds. Maximum duration is 86400 * 14 (14 days). 0 is to un-suppress alarms. | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewResponseOrgSuppressAlarmItem

`func NewResponseOrgSuppressAlarmItem() *ResponseOrgSuppressAlarmItem`

NewResponseOrgSuppressAlarmItem instantiates a new ResponseOrgSuppressAlarmItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrgSuppressAlarmItemWithDefaults

`func NewResponseOrgSuppressAlarmItemWithDefaults() *ResponseOrgSuppressAlarmItem`

NewResponseOrgSuppressAlarmItemWithDefaults instantiates a new ResponseOrgSuppressAlarmItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *ResponseOrgSuppressAlarmItem) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ResponseOrgSuppressAlarmItem) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ResponseOrgSuppressAlarmItem) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ResponseOrgSuppressAlarmItem) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetSiteId

`func (o *ResponseOrgSuppressAlarmItem) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ResponseOrgSuppressAlarmItem) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ResponseOrgSuppressAlarmItem) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ResponseOrgSuppressAlarmItem) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


