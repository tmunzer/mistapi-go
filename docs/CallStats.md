# CallStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** |  | [optional] 
**AudioQuality** | Pointer to **int32** |  | [optional] 
**EndTime** | Pointer to **int32** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**MeetingId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Rating** | Pointer to **int32** |  | [optional] 
**ScreenShareQuality** | Pointer to **int32** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**StartTime** | Pointer to **int32** |  | [optional] 
**VideoQuality** | Pointer to **int32** |  | [optional] 

## Methods

### NewCallStats

`func NewCallStats() *CallStats`

NewCallStats instantiates a new CallStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallStatsWithDefaults

`func NewCallStatsWithDefaults() *CallStats`

NewCallStatsWithDefaults instantiates a new CallStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *CallStats) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CallStats) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CallStats) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *CallStats) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAudioQuality

`func (o *CallStats) GetAudioQuality() int32`

GetAudioQuality returns the AudioQuality field if non-nil, zero value otherwise.

### GetAudioQualityOk

`func (o *CallStats) GetAudioQualityOk() (*int32, bool)`

GetAudioQualityOk returns a tuple with the AudioQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioQuality

`func (o *CallStats) SetAudioQuality(v int32)`

SetAudioQuality sets AudioQuality field to given value.

### HasAudioQuality

`func (o *CallStats) HasAudioQuality() bool`

HasAudioQuality returns a boolean if a field has been set.

### GetEndTime

`func (o *CallStats) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CallStats) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CallStats) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CallStats) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMac

`func (o *CallStats) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *CallStats) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *CallStats) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *CallStats) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMeetingId

`func (o *CallStats) GetMeetingId() string`

GetMeetingId returns the MeetingId field if non-nil, zero value otherwise.

### GetMeetingIdOk

`func (o *CallStats) GetMeetingIdOk() (*string, bool)`

GetMeetingIdOk returns a tuple with the MeetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingId

`func (o *CallStats) SetMeetingId(v string)`

SetMeetingId sets MeetingId field to given value.

### HasMeetingId

`func (o *CallStats) HasMeetingId() bool`

HasMeetingId returns a boolean if a field has been set.

### GetOrgId

`func (o *CallStats) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CallStats) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CallStats) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *CallStats) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRating

`func (o *CallStats) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CallStats) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CallStats) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *CallStats) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetScreenShareQuality

`func (o *CallStats) GetScreenShareQuality() int32`

GetScreenShareQuality returns the ScreenShareQuality field if non-nil, zero value otherwise.

### GetScreenShareQualityOk

`func (o *CallStats) GetScreenShareQualityOk() (*int32, bool)`

GetScreenShareQualityOk returns a tuple with the ScreenShareQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenShareQuality

`func (o *CallStats) SetScreenShareQuality(v int32)`

SetScreenShareQuality sets ScreenShareQuality field to given value.

### HasScreenShareQuality

`func (o *CallStats) HasScreenShareQuality() bool`

HasScreenShareQuality returns a boolean if a field has been set.

### GetSiteId

`func (o *CallStats) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CallStats) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CallStats) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CallStats) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStartTime

`func (o *CallStats) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CallStats) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CallStats) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CallStats) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetVideoQuality

`func (o *CallStats) GetVideoQuality() int32`

GetVideoQuality returns the VideoQuality field if non-nil, zero value otherwise.

### GetVideoQualityOk

`func (o *CallStats) GetVideoQualityOk() (*int32, bool)`

GetVideoQualityOk returns a tuple with the VideoQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuality

`func (o *CallStats) SetVideoQuality(v int32)`

SetVideoQuality sets VideoQuality field to given value.

### HasVideoQuality

`func (o *CallStats) HasVideoQuality() bool`

HasVideoQuality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


