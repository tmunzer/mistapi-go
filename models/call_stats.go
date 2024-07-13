package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// CallStats represents a CallStats struct.
type CallStats struct {
    App                  *string        `json:"app,omitempty"`
    AudioQuality         *int           `json:"audio_quality,omitempty"`
    EndTime              *int           `json:"end_time,omitempty"`
    Mac                  *string        `json:"mac,omitempty"`
    MeetingId            *string        `json:"meeting_id,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    Rating               *int           `json:"rating,omitempty"`
    ScreenShareQuality   *int           `json:"screen_share_quality,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    StartTime            *int           `json:"start_time,omitempty"`
    VideoQuality         *int           `json:"video_quality,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CallStats.
// It customizes the JSON marshaling process for CallStats objects.
func (c CallStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CallStats object to a map representation for JSON marshaling.
func (c CallStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.App != nil {
        structMap["app"] = c.App
    }
    if c.AudioQuality != nil {
        structMap["audio_quality"] = c.AudioQuality
    }
    if c.EndTime != nil {
        structMap["end_time"] = c.EndTime
    }
    if c.Mac != nil {
        structMap["mac"] = c.Mac
    }
    if c.MeetingId != nil {
        structMap["meeting_id"] = c.MeetingId
    }
    if c.OrgId != nil {
        structMap["org_id"] = c.OrgId
    }
    if c.Rating != nil {
        structMap["rating"] = c.Rating
    }
    if c.ScreenShareQuality != nil {
        structMap["screen_share_quality"] = c.ScreenShareQuality
    }
    if c.SiteId != nil {
        structMap["site_id"] = c.SiteId
    }
    if c.StartTime != nil {
        structMap["start_time"] = c.StartTime
    }
    if c.VideoQuality != nil {
        structMap["video_quality"] = c.VideoQuality
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CallStats.
// It customizes the JSON unmarshaling process for CallStats objects.
func (c *CallStats) UnmarshalJSON(input []byte) error {
    var temp callStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "app", "audio_quality", "end_time", "mac", "meeting_id", "org_id", "rating", "screen_share_quality", "site_id", "start_time", "video_quality")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.App = temp.App
    c.AudioQuality = temp.AudioQuality
    c.EndTime = temp.EndTime
    c.Mac = temp.Mac
    c.MeetingId = temp.MeetingId
    c.OrgId = temp.OrgId
    c.Rating = temp.Rating
    c.ScreenShareQuality = temp.ScreenShareQuality
    c.SiteId = temp.SiteId
    c.StartTime = temp.StartTime
    c.VideoQuality = temp.VideoQuality
    return nil
}

// callStats is a temporary struct used for validating the fields of CallStats.
type callStats  struct {
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
