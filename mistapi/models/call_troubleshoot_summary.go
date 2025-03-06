package models

import (
    "encoding/json"
    "fmt"
)

// CallTroubleshootSummary represents a CallTroubleshootSummary struct.
type CallTroubleshootSummary struct {
    AudioIn              *CallTroubleshootSummaryData `json:"audio_in,omitempty"`
    AudioOut             *CallTroubleshootSummaryData `json:"audio_out,omitempty"`
    // Epoch (seconds)
    Timestamp            *float64                     `json:"timestamp,omitempty"`
    VideoIn              *CallTroubleshootSummaryData `json:"video_in,omitempty"`
    VideoOut             *CallTroubleshootSummaryData `json:"video_out,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for CallTroubleshootSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CallTroubleshootSummary) String() string {
    return fmt.Sprintf(
    	"CallTroubleshootSummary[AudioIn=%v, AudioOut=%v, Timestamp=%v, VideoIn=%v, VideoOut=%v, AdditionalProperties=%v]",
    	c.AudioIn, c.AudioOut, c.Timestamp, c.VideoIn, c.VideoOut, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CallTroubleshootSummary.
// It customizes the JSON marshaling process for CallTroubleshootSummary objects.
func (c CallTroubleshootSummary) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "audio_in", "audio_out", "timestamp", "video_in", "video_out"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CallTroubleshootSummary object to a map representation for JSON marshaling.
func (c CallTroubleshootSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.AudioIn != nil {
        structMap["audio_in"] = c.AudioIn.toMap()
    }
    if c.AudioOut != nil {
        structMap["audio_out"] = c.AudioOut.toMap()
    }
    if c.Timestamp != nil {
        structMap["timestamp"] = c.Timestamp
    }
    if c.VideoIn != nil {
        structMap["video_in"] = c.VideoIn.toMap()
    }
    if c.VideoOut != nil {
        structMap["video_out"] = c.VideoOut.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CallTroubleshootSummary.
// It customizes the JSON unmarshaling process for CallTroubleshootSummary objects.
func (c *CallTroubleshootSummary) UnmarshalJSON(input []byte) error {
    var temp tempCallTroubleshootSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "audio_in", "audio_out", "timestamp", "video_in", "video_out")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.AudioIn = temp.AudioIn
    c.AudioOut = temp.AudioOut
    c.Timestamp = temp.Timestamp
    c.VideoIn = temp.VideoIn
    c.VideoOut = temp.VideoOut
    return nil
}

// tempCallTroubleshootSummary is a temporary struct used for validating the fields of CallTroubleshootSummary.
type tempCallTroubleshootSummary  struct {
    AudioIn   *CallTroubleshootSummaryData `json:"audio_in,omitempty"`
    AudioOut  *CallTroubleshootSummaryData `json:"audio_out,omitempty"`
    Timestamp *float64                     `json:"timestamp,omitempty"`
    VideoIn   *CallTroubleshootSummaryData `json:"video_in,omitempty"`
    VideoOut  *CallTroubleshootSummaryData `json:"video_out,omitempty"`
}
