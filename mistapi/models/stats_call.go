package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// StatsCall represents a StatsCall struct.
type StatsCall struct {
    App                  *string                `json:"app,omitempty"`
    AudioQuality         *int                   `json:"audio_quality,omitempty"`
    EndTime              *int                   `json:"end_time,omitempty"`
    Mac                  *string                `json:"mac,omitempty"`
    MeetingId            *string                `json:"meeting_id,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    Rating               *int                   `json:"rating,omitempty"`
    ScreenShareQuality   *int                   `json:"screen_share_quality,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    StartTime            *int                   `json:"start_time,omitempty"`
    VideoQuality         *int                   `json:"video_quality,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsCall,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsCall) String() string {
    return fmt.Sprintf(
    	"StatsCall[App=%v, AudioQuality=%v, EndTime=%v, Mac=%v, MeetingId=%v, OrgId=%v, Rating=%v, ScreenShareQuality=%v, SiteId=%v, StartTime=%v, VideoQuality=%v, AdditionalProperties=%v]",
    	s.App, s.AudioQuality, s.EndTime, s.Mac, s.MeetingId, s.OrgId, s.Rating, s.ScreenShareQuality, s.SiteId, s.StartTime, s.VideoQuality, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsCall.
// It customizes the JSON marshaling process for StatsCall objects.
func (s StatsCall) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "app", "audio_quality", "end_time", "mac", "meeting_id", "org_id", "rating", "screen_share_quality", "site_id", "start_time", "video_quality"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsCall object to a map representation for JSON marshaling.
func (s StatsCall) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.App != nil {
        structMap["app"] = s.App
    }
    if s.AudioQuality != nil {
        structMap["audio_quality"] = s.AudioQuality
    }
    if s.EndTime != nil {
        structMap["end_time"] = s.EndTime
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    }
    if s.MeetingId != nil {
        structMap["meeting_id"] = s.MeetingId
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.Rating != nil {
        structMap["rating"] = s.Rating
    }
    if s.ScreenShareQuality != nil {
        structMap["screen_share_quality"] = s.ScreenShareQuality
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.StartTime != nil {
        structMap["start_time"] = s.StartTime
    }
    if s.VideoQuality != nil {
        structMap["video_quality"] = s.VideoQuality
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsCall.
// It customizes the JSON unmarshaling process for StatsCall objects.
func (s *StatsCall) UnmarshalJSON(input []byte) error {
    var temp tempStatsCall
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "app", "audio_quality", "end_time", "mac", "meeting_id", "org_id", "rating", "screen_share_quality", "site_id", "start_time", "video_quality")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.App = temp.App
    s.AudioQuality = temp.AudioQuality
    s.EndTime = temp.EndTime
    s.Mac = temp.Mac
    s.MeetingId = temp.MeetingId
    s.OrgId = temp.OrgId
    s.Rating = temp.Rating
    s.ScreenShareQuality = temp.ScreenShareQuality
    s.SiteId = temp.SiteId
    s.StartTime = temp.StartTime
    s.VideoQuality = temp.VideoQuality
    return nil
}

// tempStatsCall is a temporary struct used for validating the fields of StatsCall.
type tempStatsCall  struct {
    App                *string    `json:"app,omitempty"`
    AudioQuality       *int       `json:"audio_quality,omitempty"`
    EndTime            *int       `json:"end_time,omitempty"`
    Mac                *string    `json:"mac,omitempty"`
    MeetingId          *string    `json:"meeting_id,omitempty"`
    OrgId              *uuid.UUID `json:"org_id,omitempty"`
    Rating             *int       `json:"rating,omitempty"`
    ScreenShareQuality *int       `json:"screen_share_quality,omitempty"`
    SiteId             *uuid.UUID `json:"site_id,omitempty"`
    StartTime          *int       `json:"start_time,omitempty"`
    VideoQuality       *int       `json:"video_quality,omitempty"`
}
