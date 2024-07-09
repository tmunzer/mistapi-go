package models

import (
    "encoding/json"
)

// TroubleshootCallItem represents a TroubleshootCallItem struct.
type TroubleshootCallItem struct {
    AudioIn              *CallTroubleshootData `json:"audio_in,omitempty"`
    AudioOut             *CallTroubleshootData `json:"audio_out,omitempty"`
    Timestamp            *int                  `json:"timestamp,omitempty"`
    VideoIn              *CallTroubleshootData `json:"video_in,omitempty"`
    VideoOut             *CallTroubleshootData `json:"video_out,omitempty"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TroubleshootCallItem.
// It customizes the JSON marshaling process for TroubleshootCallItem objects.
func (t TroubleshootCallItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TroubleshootCallItem object to a map representation for JSON marshaling.
func (t TroubleshootCallItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.AudioIn != nil {
        structMap["audio_in"] = t.AudioIn.toMap()
    }
    if t.AudioOut != nil {
        structMap["audio_out"] = t.AudioOut.toMap()
    }
    if t.Timestamp != nil {
        structMap["timestamp"] = t.Timestamp
    }
    if t.VideoIn != nil {
        structMap["video_in"] = t.VideoIn.toMap()
    }
    if t.VideoOut != nil {
        structMap["video_out"] = t.VideoOut.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TroubleshootCallItem.
// It customizes the JSON unmarshaling process for TroubleshootCallItem objects.
func (t *TroubleshootCallItem) UnmarshalJSON(input []byte) error {
    var temp troubleshootCallItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "audio_in", "audio_out", "timestamp", "video_in", "video_out")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.AudioIn = temp.AudioIn
    t.AudioOut = temp.AudioOut
    t.Timestamp = temp.Timestamp
    t.VideoIn = temp.VideoIn
    t.VideoOut = temp.VideoOut
    return nil
}

// troubleshootCallItem is a temporary struct used for validating the fields of TroubleshootCallItem.
type troubleshootCallItem  struct {
    AudioIn   *CallTroubleshootData `json:"audio_in,omitempty"`
    AudioOut  *CallTroubleshootData `json:"audio_out,omitempty"`
    Timestamp *int                  `json:"timestamp,omitempty"`
    VideoIn   *CallTroubleshootData `json:"video_in,omitempty"`
    VideoOut  *CallTroubleshootData `json:"video_out,omitempty"`
}
