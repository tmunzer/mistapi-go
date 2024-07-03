# SiteEngagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DwellTagNames** | Pointer to [**SiteEngagementDwellTagNames**](SiteEngagementDwellTagNames.md) |  | [optional] 
**DwellTags** | Pointer to [**SiteEngagementDwellTags**](SiteEngagementDwellTags.md) |  | [optional] 
**Hours** | Pointer to [**Hours**](Hours.md) |  | [optional] 
**MaxDwell** | Pointer to **int32** | max time, default is 43200(12h), max is 68400 (18h) | [optional] [default to 43200]
**MinDwell** | Pointer to **int32** | min time | [optional] 

## Methods

### NewSiteEngagement

`func NewSiteEngagement() *SiteEngagement`

NewSiteEngagement instantiates a new SiteEngagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteEngagementWithDefaults

`func NewSiteEngagementWithDefaults() *SiteEngagement`

NewSiteEngagementWithDefaults instantiates a new SiteEngagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDwellTagNames

`func (o *SiteEngagement) GetDwellTagNames() SiteEngagementDwellTagNames`

GetDwellTagNames returns the DwellTagNames field if non-nil, zero value otherwise.

### GetDwellTagNamesOk

`func (o *SiteEngagement) GetDwellTagNamesOk() (*SiteEngagementDwellTagNames, bool)`

GetDwellTagNamesOk returns a tuple with the DwellTagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDwellTagNames

`func (o *SiteEngagement) SetDwellTagNames(v SiteEngagementDwellTagNames)`

SetDwellTagNames sets DwellTagNames field to given value.

### HasDwellTagNames

`func (o *SiteEngagement) HasDwellTagNames() bool`

HasDwellTagNames returns a boolean if a field has been set.

### GetDwellTags

`func (o *SiteEngagement) GetDwellTags() SiteEngagementDwellTags`

GetDwellTags returns the DwellTags field if non-nil, zero value otherwise.

### GetDwellTagsOk

`func (o *SiteEngagement) GetDwellTagsOk() (*SiteEngagementDwellTags, bool)`

GetDwellTagsOk returns a tuple with the DwellTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDwellTags

`func (o *SiteEngagement) SetDwellTags(v SiteEngagementDwellTags)`

SetDwellTags sets DwellTags field to given value.

### HasDwellTags

`func (o *SiteEngagement) HasDwellTags() bool`

HasDwellTags returns a boolean if a field has been set.

### GetHours

`func (o *SiteEngagement) GetHours() Hours`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *SiteEngagement) GetHoursOk() (*Hours, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *SiteEngagement) SetHours(v Hours)`

SetHours sets Hours field to given value.

### HasHours

`func (o *SiteEngagement) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetMaxDwell

`func (o *SiteEngagement) GetMaxDwell() int32`

GetMaxDwell returns the MaxDwell field if non-nil, zero value otherwise.

### GetMaxDwellOk

`func (o *SiteEngagement) GetMaxDwellOk() (*int32, bool)`

GetMaxDwellOk returns a tuple with the MaxDwell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDwell

`func (o *SiteEngagement) SetMaxDwell(v int32)`

SetMaxDwell sets MaxDwell field to given value.

### HasMaxDwell

`func (o *SiteEngagement) HasMaxDwell() bool`

HasMaxDwell returns a boolean if a field has been set.

### GetMinDwell

`func (o *SiteEngagement) GetMinDwell() int32`

GetMinDwell returns the MinDwell field if non-nil, zero value otherwise.

### GetMinDwellOk

`func (o *SiteEngagement) GetMinDwellOk() (*int32, bool)`

GetMinDwellOk returns a tuple with the MinDwell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDwell

`func (o *SiteEngagement) SetMinDwell(v int32)`

SetMinDwell sets MinDwell field to given value.

### HasMinDwell

`func (o *SiteEngagement) HasMinDwell() bool`

HasMinDwell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


