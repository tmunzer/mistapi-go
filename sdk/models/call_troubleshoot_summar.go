package models

import (
    "encoding/json"
)

// CallTroubleshootSummar represents a CallTroubleshootSummar struct.
type CallTroubleshootSummar struct {
    AudioIn              *CallTroubleshootSummaryData `json:"audio_in,omitempty"`
    AudioOut             *CallTroubleshootSummaryData `json:"audio_out,omitempty"`
    Timestamp            *int                         `json:"timestamp,omitempty"`
    VideoIn              *CallTroubleshootSummaryData `json:"video_in,omitempty"`
    VideoOut             *CallTroubleshootSummaryData `json:"video_out,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CallTroubleshootSummar.
// It customizes the JSON marshaling process for CallTroubleshootSummar objects.
func (c CallTroubleshootSummar) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CallTroubleshootSummar object to a map representation for JSON marshaling.
func (c CallTroubleshootSummar) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for CallTroubleshootSummar.
// It customizes the JSON unmarshaling process for CallTroubleshootSummar objects.
func (c *CallTroubleshootSummar) UnmarshalJSON(input []byte) error {
    var temp callTroubleshootSummar
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "audio_in", "audio_out", "timestamp", "video_in", "video_out")
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

// callTroubleshootSummar is a temporary struct used for validating the fields of CallTroubleshootSummar.
type callTroubleshootSummar  struct {
    AudioIn   *CallTroubleshootSummaryData `json:"audio_in,omitempty"`
    AudioOut  *CallTroubleshootSummaryData `json:"audio_out,omitempty"`
    Timestamp *int                         `json:"timestamp,omitempty"`
    VideoIn   *CallTroubleshootSummaryData `json:"video_in,omitempty"`
    VideoOut  *CallTroubleshootSummaryData `json:"video_out,omitempty"`
}
