/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
)

// checks if the CallStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallStats{}

// CallStats struct for CallStats
type CallStats struct {
	App *string `json:"app,omitempty"`
	AudioQuality *int32 `json:"audio_quality,omitempty"`
	EndTime *int32 `json:"end_time,omitempty"`
	Mac *string `json:"mac,omitempty"`
	MeetingId *string `json:"meeting_id,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	Rating *int32 `json:"rating,omitempty"`
	ScreenShareQuality *int32 `json:"screen_share_quality,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	StartTime *int32 `json:"start_time,omitempty"`
	VideoQuality *int32 `json:"video_quality,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CallStats CallStats

// NewCallStats instantiates a new CallStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallStats() *CallStats {
	this := CallStats{}
	return &this
}

// NewCallStatsWithDefaults instantiates a new CallStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallStatsWithDefaults() *CallStats {
	this := CallStats{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *CallStats) GetApp() string {
	if o == nil || IsNil(o.App) {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallStats) GetAppOk() (*string, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *CallStats) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *CallStats) SetApp(v string) {
	o.App = &v
}

// GetAudioQuality returns the AudioQuality field value if set, zero value otherwise.
func (o *CallStats) GetAudioQuality() int32 {
	if o == nil || IsNil(o.AudioQuality) {
		var ret int32
		return ret
	}
	return *o.AudioQuality
}

// GetAudioQualityOk returns a tuple with the AudioQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallStats) GetAudioQualityOk() (*int32, bool) {
	if o == nil || IsNil(o.AudioQuality) {
		return nil, false
	}
	return o.AudioQuality, true
}

// HasAudioQuality returns a boolean if a field has been set.
func (o *CallStats) HasAudioQuality() bool {
	if o != nil && !IsNil(o.AudioQuality) {
		return true
	}

	return false
}

// SetAudioQuality gets a reference to the given int32 and assigns it to the AudioQuality field.
func (o *CallStats) SetAudioQuality(v int32) {
	o.AudioQuality = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *CallStats) GetEndTime() int32 {
	if o == nil || IsNil(o.EndTime) {
		var ret int32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallStats) GetEndTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *CallStats) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int32 and assigns it to the EndTime field.
func (o *CallStats) SetEndTime(v int32) {
	o.EndTime = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *CallStats) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallStats) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *CallStats) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *CallStats) SetMac(v string) {
	o.Mac = &v
}

// GetMeetingId returns the MeetingId field value if set, zero value otherwise.
func (o *CallStats) GetMeetingId() string {
	if o == nil || IsNil(o.MeetingId) {
		var ret string
		return ret
	}
	return *o.MeetingId
}

// GetMeetingIdOk returns a tuple with the MeetingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallStats) GetMeetingIdOk() (*string, bool) {
	if o == nil || IsNil(o.MeetingId) {
		return nil, false
	}
	return o.MeetingId, true
}

// HasMeetingId returns a boolean if a field has been set.
func (o *CallStats) HasMeetingId() bool {
	if o != nil && !IsNil(o.MeetingId) {
		return true
	}

	return false
}

// SetMeetingId gets a reference to the given string and assigns it to the MeetingId field.
func (o *CallStats) SetMeetingId(v string) {
	o.MeetingId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *CallStats) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallStats) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *CallStats) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *CallStats) SetOrgId(v string) {
	o.OrgId = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *CallStats) GetRating() int32 {
	if o == nil || IsNil(o.Rating) {
		var ret int32
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallStats) GetRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *CallStats) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given int32 and assigns it to the Rating field.
func (o *CallStats) SetRating(v int32) {
	o.Rating = &v
}

// GetScreenShareQuality returns the ScreenShareQuality field value if set, zero value otherwise.
func (o *CallStats) GetScreenShareQuality() int32 {
	if o == nil || IsNil(o.ScreenShareQuality) {
		var ret int32
		return ret
	}
	return *o.ScreenShareQuality
}

// GetScreenShareQualityOk returns a tuple with the ScreenShareQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallStats) GetScreenShareQualityOk() (*int32, bool) {
	if o == nil || IsNil(o.ScreenShareQuality) {
		return nil, false
	}
	return o.ScreenShareQuality, true
}

// HasScreenShareQuality returns a boolean if a field has been set.
func (o *CallStats) HasScreenShareQuality() bool {
	if o != nil && !IsNil(o.ScreenShareQuality) {
		return true
	}

	return false
}

// SetScreenShareQuality gets a reference to the given int32 and assigns it to the ScreenShareQuality field.
func (o *CallStats) SetScreenShareQuality(v int32) {
	o.ScreenShareQuality = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *CallStats) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallStats) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *CallStats) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *CallStats) SetSiteId(v string) {
	o.SiteId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CallStats) GetStartTime() int32 {
	if o == nil || IsNil(o.StartTime) {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallStats) GetStartTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CallStats) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *CallStats) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetVideoQuality returns the VideoQuality field value if set, zero value otherwise.
func (o *CallStats) GetVideoQuality() int32 {
	if o == nil || IsNil(o.VideoQuality) {
		var ret int32
		return ret
	}
	return *o.VideoQuality
}

// GetVideoQualityOk returns a tuple with the VideoQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallStats) GetVideoQualityOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoQuality) {
		return nil, false
	}
	return o.VideoQuality, true
}

// HasVideoQuality returns a boolean if a field has been set.
func (o *CallStats) HasVideoQuality() bool {
	if o != nil && !IsNil(o.VideoQuality) {
		return true
	}

	return false
}

// SetVideoQuality gets a reference to the given int32 and assigns it to the VideoQuality field.
func (o *CallStats) SetVideoQuality(v int32) {
	o.VideoQuality = &v
}

func (o CallStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.AudioQuality) {
		toSerialize["audio_quality"] = o.AudioQuality
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.MeetingId) {
		toSerialize["meeting_id"] = o.MeetingId
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.ScreenShareQuality) {
		toSerialize["screen_share_quality"] = o.ScreenShareQuality
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.VideoQuality) {
		toSerialize["video_quality"] = o.VideoQuality
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CallStats) UnmarshalJSON(data []byte) (err error) {
	varCallStats := _CallStats{}

	err = json.Unmarshal(data, &varCallStats)

	if err != nil {
		return err
	}

	*o = CallStats(varCallStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "app")
		delete(additionalProperties, "audio_quality")
		delete(additionalProperties, "end_time")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "meeting_id")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "rating")
		delete(additionalProperties, "screen_share_quality")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "video_quality")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallStats struct {
	value *CallStats
	isSet bool
}

func (v NullableCallStats) Get() *CallStats {
	return v.value
}

func (v *NullableCallStats) Set(val *CallStats) {
	v.value = val
	v.isSet = true
}

func (v NullableCallStats) IsSet() bool {
	return v.isSet
}

func (v *NullableCallStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallStats(val *CallStats) *NullableCallStats {
	return &NullableCallStats{value: val, isSet: true}
}

func (v NullableCallStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

